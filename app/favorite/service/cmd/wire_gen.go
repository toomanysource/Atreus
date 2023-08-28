// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/toomanysource/atreus/app/favorite/service/internal/biz"
	"github.com/toomanysource/atreus/app/favorite/service/internal/conf"
	"github.com/toomanysource/atreus/app/favorite/service/internal/data"
	"github.com/toomanysource/atreus/app/favorite/service/internal/server"
	"github.com/toomanysource/atreus/app/favorite/service/internal/service"

	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, client *conf.Client, confData *conf.Data, jwt *conf.JWT, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewMysqlConn(confData, logger)
	redisClient := data.NewRedisConn(confData, logger)
	kfkWriter := data.NewKafkaWriter(confData)
	dataData, cleanup, err := data.NewData(db, redisClient, kfkWriter, logger)
	if err != nil {
		return nil, nil, err
	}
	publishConn := server.NewPublishClient(client, logger)
	favoriteRepo := data.NewFavoriteRepo(dataData, publishConn, logger)
	favoriteUseCase := biz.NewFavoriteUseCase(jwt, favoriteRepo, logger)
	favoriteService := service.NewFavoriteService(favoriteUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, favoriteService, logger)
	httpServer := server.NewHTTPServer(confServer, favoriteService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
