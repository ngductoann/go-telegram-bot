package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"go-telegram-bot/internal/domain/entity"
	domainService "go-telegram-bot/internal/domain/service"
)

type ipService struct {
	httpClient *http.Client
	ipUrls     []string
}

// NewIPService creates a new instance of ipService.
func NewIPService() domainService.IPService {
	// Create a custom transport with TLS configuration
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false, // Try secure first
		},
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
	}

	return &ipService{
		httpClient: &http.Client{
			Timeout:   15 * time.Second,
			Transport: transport,
		},
		// Multiple fallback URLs in case one fails
		ipUrls: []string{
			"https://ipinfo.io/ip",
			"https://api.ipify.org",
			"https://icanhazip.com",
			"http://ipinfo.io/ip", // HTTP fallback for TLS issues
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
	var lastErr error

	// Try each URL until one succeeds
	for _, url := range s.ipUrls {
		ip, err := s.fetchIPFromURL(ctx, url)
		if err == nil {
			return ip, nil
		}
		lastErr = err
	}

	return "", fmt.Errorf("failed to get public IP from all sources, last error: %w", lastErr)
}

// fetchIPFromURL attempts to fetch the IP from a specific URL
func (s *ipService) fetchIPFromURL(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for %s: %w", url, err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d from %s", resp.StatusCode, url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body from %s: %w", url, err)
	}

	publicIP := strings.TrimSpace(string(body))
	if publicIP == "" {
		return "", fmt.Errorf("empty response from %s", url)
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

// ValidateIP validates if the provided string is a valid IP address
func (s *ipService) ValidateIP(ip string) bool {
	return net.ParseIP(strings.TrimSpace(ip)) != nil
}

// IsPrivateIP checks if the provided IP is a private IP address
func (s *ipService) IsPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(strings.TrimSpace(ip))
	if parsedIP == nil {
		return false
	}

	// Check for IPv4 private ranges
	if parsedIP.To4() != nil {
		// 10.0.0.0/8
		if parsedIP.To4()[0] == 10 {
			return true
		}
		// 172.16.0.0/12
		if parsedIP.To4()[0] == 172 && parsedIP.To4()[1] >= 16 && parsedIP.To4()[1] <= 31 {
			return true
		}
		// 192.168.0.0/16
		if parsedIP.To4()[0] == 192 && parsedIP.To4()[1] == 168 {
			return true
		}
		// 127.0.0.0/8 (localhost)
		if parsedIP.To4()[0] == 127 {
			return true
		}
	}

	// Check for IPv6 private ranges
	if parsedIP.To16() != nil {
		// ::1 (localhost)
		if parsedIP.IsLoopback() {
			return true
		}
		// fc00::/7 (unique local addresses)
		if parsedIP.To16()[0] == 0xfc || parsedIP.To16()[0] == 0xfd {
			return true
		}
		// fe80::/10 (link-local addresses)
		if parsedIP.To16()[0] == 0xfe && (parsedIP.To16()[1]&0xc0) == 0x80 {
			return true
		}
	}

	return false
}

// IsPublicIP checks if the provided IP is a public IP address
func (s *ipService) IsPublicIP(ip string) bool {
	if !s.ValidateIP(ip) {
		return false
	}
	return !s.IsPrivateIP(ip)
}
