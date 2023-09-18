package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/toomanysource/atreus/pkg/kafkaX"

	"github.com/toomanysource/atreus/middleware"

	"github.com/toomanysource/atreus/app/message/service/internal/biz"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"github.com/segmentio/kafka-go"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	timeFactor    = 20
	RandTimeBegin = 360
	RandTimeEnd   = 720
)

var ErrMsgYourself = errors.New("can't send message to yourself")

type Message struct {
	Id         uint32 `gorm:"column:id;primary_key;auto_increment"`
	FromUserId uint32 `gorm:"column:from_user_id;not null;index:idx_from_user_to_user"`
	ToUserId   uint32 `gorm:"column:to_user_id;not null;index:idx_from_user_to_user"`
	Content    string `gorm:"column:content;not null"`
	CreateTime int64  `gorm:"column:created_at"`
}

func (Message) TableName() string {
	return "message"
}

type messageRepo struct {
	data *Data
	log  *log.Helper
}

func NewMessageRepo(data *Data, logger log.Logger) biz.MessageRepo {
	return &messageRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/message")),
	}
}

// PublishMessage 发送消息
func (r *messageRepo) PublishMessage(ctx context.Context, toUserId uint32, content string) error {
	userId := ctx.Value(middleware.UserIdKey("user_id")).(uint32)
	if userId == toUserId {
		return ErrMsgYourself
	}
	createTime := time.Now().UnixMilli()
	go func() {
		err := r.MessageProducer(userId, toUserId, content, createTime)
		if err != nil {
			r.log.Errorf("message producer error, err: %w", err)
			return
		}
	}()
	return nil
}

// GetMessageList 获取聊天记录列表
func (r *messageRepo) GetMessageList(ctx context.Context, toUserId uint32, preMsgTime int64) ([]*biz.Message, error) {
	// 先在redis缓存中查询是否存在聊天记录列表
	userId := ctx.Value(middleware.UserIdKey("user_id")).(uint32)
	key := setKey(userId, toUserId)
	ok, err := r.CheckCache(ctx, key)
	if err != nil {
		return nil, err
	}
	if ok {
		return r.GetCache(ctx, userId, toUserId, preMsgTime)
	}
	// 加锁防止私聊两用户同时请求导致重复创建
	ok, err = r.AddCacheMutex(ctx)
	if err != nil {
		return nil, err
	}
	if ok {
		ok, err = r.CheckCache(ctx, key)
		if err != nil {
			return nil, err
		}
		if ok {
			return r.GetCache(ctx, userId, toUserId, preMsgTime)
		}
		cl, err := r.GetMessages(ctx, userId, toUserId, preMsgTime)
		if err != nil {
			return nil, err
		}
		go func() {
			if err = r.CreateCacheByTran(ctx, cl, key); err != nil {
				r.log.Error(err)
				return
			}
			if err = r.DelCacheMutex(ctx); err != nil {
				r.log.Error(err)
				return
			}
			r.log.Info("redis transaction success")
		}()
		return cl, nil
	}
	return r.GetMessages(ctx, userId, toUserId, preMsgTime)
}

// AddCacheMutex 缓存加锁
func (r *messageRepo) AddCacheMutex(ctx context.Context) (bool, error) {
	ok, err := r.data.cache.SetNX(ctx, "mutex", "", time.Second*time.Duration(timeFactor)).Result()
	if err != nil {
		return false, errors.Join(ErrRedisSet, err)
	}
	return ok, nil
}

// DelCacheMutex 缓存解锁
func (r *messageRepo) DelCacheMutex(ctx context.Context) error {
	_, err := r.data.cache.Del(ctx, "mutex").Result()
	if err != nil {
		return errors.Join(ErrRedisDelete, err)
	}
	return nil
}

// CheckCache 检查缓存是否存在
func (r *messageRepo) CheckCache(ctx context.Context, key string) (bool, error) {
	// 先在redis缓存中查询是否存在聊天记录列表
	count, err := r.data.cache.Exists(ctx, key).Result()
	if err != nil {
		return false, errors.Join(ErrRedisQuery, err)
	}
	return count > 0, nil
}

// GetCache 从缓存中获取聊天记录列表
func (r *messageRepo) GetCache(ctx context.Context, userId, toUserId uint32, preMsgTime int64) ([]*biz.Message, error) {
	msgList, err := r.data.cache.ZRangeByScore(ctx, setKey(userId, toUserId), &redis.ZRangeBy{
		Min: fmt.Sprintf("(%f", float64(preMsgTime)),
		Max: "+inf",
	}).Result()
	if err != nil {
		return nil, errors.Join(ErrRedisQuery, err)
	}
	if len(msgList) == 0 {
		return nil, nil
	}
	cl := make([]*biz.Message, 0, len(msgList))
	// 如果存在则直接返回
	for _, v := range msgList {
		co := &biz.Message{}
		if err = json.Unmarshal([]byte(v), co); err != nil {
			return nil, errors.Join(ErrJsonMarshal, err)
		}
		if co.FromUserId != userId || preMsgTime == 0 {
			cl = append(cl, co)
		}
	}
	return cl, nil
}

// MessageProducer 生产消息
func (r *messageRepo) MessageProducer(userId, toUserId uint32, content string, time int64) error {
	mg := &Message{
		FromUserId: userId,
		ToUserId:   toUserId,
		Content:    content,
		CreateTime: time,
	}
	byteValue, err := json.Marshal(mg)
	if err != nil {
		return errors.Join(ErrJsonMarshal, err)
	}
	return kafkaX.Update(r.data.kfk.writer, "", string(byteValue))
}

// InitStoreMessageQueue 初始化聊天记录存储队列
func (r *messageRepo) InitStoreMessageQueue() {
	kafkaX.Reader(r.data.kfk.reader, r.log, func(ctx context.Context, reader *kafka.Reader, msg kafka.Message) {
		value := msg.Value
		var mg *Message
		err := json.Unmarshal(value, &mg)
		if err != nil {
			r.log.Error(errors.Join(ErrJsonMarshal, err))
			return
		}
		m, err := r.InsertMessage(ctx, mg.FromUserId, mg.ToUserId, mg.Content, mg.CreateTime)
		if err != nil {
			r.log.Error(err)
			return
		}
		go func() {
			ctx = context.Background()
			key := setKey(mg.FromUserId, mg.ToUserId)
			data, err := json.Marshal(m)
			if err != nil {
				r.log.Errorf("json marshal error %w", err)
				return
			}
			if err = r.data.cache.ZAdd(ctx, key, &redis.Z{
				Score:  float64(mg.CreateTime),
				Member: string(data),
			}).Err(); err != nil {
				r.log.Errorf("redis store error %w", err)
				return
			}
			r.log.Info("redis store success")
		}()
	})
}

// GetMessages 数据库根据最新消息时间查询消息
func (r *messageRepo) GetMessages(ctx context.Context, userId, toUserId uint32, preMsgTime int64) (ml []*biz.Message, err error) {
	var mel []*Message
	err = r.data.db.WithContext(ctx).Where(
		"(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)",
		userId, toUserId, toUserId, userId).Where("created_at >= ?", preMsgTime).
		Order("created_at").Find(&mel).Error
	if err != nil {
		return nil, errors.Join(ErrMysqlQuery, err)
	}
	if err = copier.Copy(&ml, &mel); err != nil {
		return nil, errors.Join(ErrCopy, err)
	}
	return
}

// InsertMessage 数据库插入消息
func (r *messageRepo) InsertMessage(
	ctx context.Context, userId uint32, toUserId uint32, content string, createTime int64,
) (*Message, error) {
	m := &Message{
		FromUserId: userId,
		ToUserId:   toUserId,
		Content:    content,
		CreateTime: createTime,
	}
	err := r.data.db.WithContext(ctx).Create(m).Error
	if err != nil {
		return nil, errors.Join(ErrMysqlInsert, err)
	}
	return m, nil
}

// CreateCacheByTran 缓存创建事务
func (r *messageRepo) CreateCacheByTran(ctx context.Context, ml []*biz.Message, key string) error {
	// 使用事务将列表存入redis缓存
	_, err := r.data.cache.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		insertList := make([]*redis.Z, 0, len(ml))
		for _, u := range ml {
			data, err := json.Marshal(u)
			if err != nil {
				return errors.Join(ErrJsonMarshal, err)
			}
			insertList = append(insertList, &redis.Z{
				Score:  float64(u.CreateTime),
				Member: string(data),
			})
		}
		if len(insertList) == 0 {
			insertList = append(insertList, &redis.Z{
				Score:  0,
				Member: "",
			})
		}
		if err := pipe.ZAdd(ctx, key, insertList...).Err(); err != nil {
			return errors.Join(ErrRedisSet, err)
		}
		// 将评论数量存入redis缓存,使用随机过期时间防止缓存雪崩
		err := pipe.Expire(ctx, key, randomTime(time.Minute, RandTimeBegin, RandTimeEnd)).Err()
		if err != nil {
			return errors.Join(ErrRedisSet, err)
		}
		return nil
	})
	if err != nil {
		return errors.Join(ErrRedisTransaction, err)
	}
	return nil
}

// randomTime 随机生成时间
func randomTime(timeType time.Duration, begin, end int) time.Duration {
	return timeType * time.Duration(rand.Intn(end-begin+1)+begin)
}

// setKey 设置缓存key
func setKey(userId, toUserId uint32) string {
	if userId > toUserId {
		userId, toUserId = toUserId, userId
	}
	return fmt.Sprint(strconv.Itoa(int(userId)), "-", strconv.Itoa(int(toUserId)))
}
