//go:generate wire
//+build !wireinject

package di

import (
	"configs"
	"internal/dao"
	"internal/service"
	"pkg/email"
)

func NewService(c *configs.DbConfig, m *configs.EMailConfig) (*service.Service, error) {
	mailSender := email.NewMailSender(m)
	sqlDB, err := configs.NewDb(c)
	if err != nil {
		return nil, err
	}
	userRepo := dao.NewUserRepo(sqlDB)
	serviceService := service.NewService(mailSender, userRepo)
	return serviceService, nil
}