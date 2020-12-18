package di

import (
	"github.com/google/wire"
	"configs"
	"internal/dao"
	"internal/service"
	"pkg/email"
)

// +build wireinject

// NewService 定义injector的函数签名
func NewService(c *configs.DbConfig, m *configs.EMailConfig) (*service.Service, error) {
	wire.Build(service.NewService, email.EMailSet, dao.UserSet, configs.NewDb)
	return &service.Service{}, nil
}