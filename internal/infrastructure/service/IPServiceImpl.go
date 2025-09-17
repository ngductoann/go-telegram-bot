package service

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
)

// ipService implements the IPService interface.
type ipService struct {
	httpClient *http.Client
}

// NewIPService creates a new instance of ipService.
func NewIPService() domainService.IPService {
	return &ipService{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetPublicIP fetches the public IP address of the machine using an external service.
func (s *ipService) GetLocalIP(ctx context.Context) (string, error) {
	// Connect to a well-known address to determine the local IP address
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", fmt.Errorf("failed to dial UDP get local ip: %w", err)
	}
	defer conn.Close() // close connection after function returns

	// Extract the local address from the connection
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

// GetPublicIP fetches the public IP address of the machine using an external service.
func (s *ipService) GetPublicIP(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://ipinfo.io/ip", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resq, err := s.httpClient.Do(req) // Perform the HTTP request
	if err != nil {
		return "", fmt.Errorf("failed to perform request: %w", err)
	}
	defer resq.Body.Close() // Ensure the response body is closed after function returns

	if resq.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resq.StatusCode)
	}

	body, err := io.ReadAll(resq.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	publicIP := strings.TrimSpace(string(body))
	if publicIP == "" {
		return "", fmt.Errorf("no result from ipinfo.io")
	}
	return publicIP, nil
}

// GetIPInfo retrieves both the local and public IP addresses and returns them in an IPInfo struct.
func (s *ipService) GetIPInfo(ctx context.Context) (*entity.IPInfo, error) {
	localIP, err := s.GetLocalIP(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get local IP: %w", err)
	}

	publicIP, err := s.GetPublicIP(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get public IP: %w", err)
	}

	return &entity.IPInfo{
		LocalIP:  localIP,
		PublicIP: publicIP,
	}, nil
}
