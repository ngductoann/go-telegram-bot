package service

import (
	"context"

	"go-telegram-bot/internal/domain/entity"
)

// IPService defines the interface for IP-related operations
type IPService interface {
	// GetLocalIP retrieves the local IP address of the machine
	GetLocalIP(ctx context.Context) (string, error)

	// GetPublicIP retrieves the public/WAN IP address of the machine
	GetPublicIP(ctx context.Context) (string, error)

	// GetIPInfo retrieves both local and public IP information
	GetIPInfo(ctx context.Context) (*entity.IPInfo, error)

	// ValidateIP validates if the provided string is a valid IP address
	ValidateIP(ip string) bool

	// IsPrivateIP checks if the provided IP is a private IP address
	IsPrivateIP(ip string) bool

	// IsPublicIP checks if the provided IP is a public IP address
	IsPublicIP(ip string) bool
}
