package service

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
)

type IPService interface {
	GetLocalIP(ctx context.Context) (string, error)
	GetPublicIP(ctx context.Context) (string, error)
	GetIPInfo(ctx context.Context) (*entity.IPInfo, error)
}
