// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/toomanysource/atreus/app/publish/service/internal/biz"
	"github.com/toomanysource/atreus/app/publish/service/internal/conf"
	"github.com/toomanysource/atreus/app/publish/service/internal/data"
	"github.com/toomanysource/atreus/app/publish/service/internal/server"
	"github.com/toomanysource/atreus/app/publish/service/internal/service"

	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, client *conf.Client, minio *conf.Minio, jwt *conf.JWT, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewMysqlConn(confData, logger)
	extraConn := data.NewMinioExtraConn(minio, logger)
	intraConn := data.NewMinioIntraConn(minio, logger)
	minioXClient := data.NewMinioConn(minio, extraConn, intraConn, logger)
	writer := data.NewKafkaWriter(confData, logger)
	kfkReader := data.NewKafkaReader(confData, logger)
	dataData, cleanup, err := data.NewData(db, minioXClient, writer, kfkReader, logger)
	if err != nil {
		return nil, nil, err
	}
	discovery := server.NewDiscovery(registry)
	userServiceClient := server.NewUserClient(discovery, logger)
	favoriteServiceClient := server.NewFavoriteClient(discovery, logger)
	publishRepo := data.NewPublishRepo(dataData, userServiceClient, favoriteServiceClient, logger)
	publishUseCase := biz.NewPublishUseCase(publishRepo, logger)
	publishService := service.NewPublishService(publishUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, publishService, logger)
	httpServer := server.NewHTTPServer(confServer, jwt, publishService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
