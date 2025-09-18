package service

import (
	"context"
	"testing"
	"time"
)

func TestIPService_GetPublicIP(t *testing.T) {
	service := NewIPService()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	publicIP, err := service.GetPublicIP(ctx)
	if err != nil {
		t.Fatalf("Failed to get public IP: %v", err)
	}

	if publicIP == "" {
		t.Fatal("Public IP should not be empty")
	}

	t.Logf("Public IP: %s", publicIP)
}

func TestIPService_GetLocalIP(t *testing.T) {
	service := NewIPService()
	ctx := context.Background()

	localIP, err := service.GetLocalIP(ctx)
	if err != nil {
		t.Fatalf("Failed to get local IP: %v", err)
	}

	if localIP == "" {
		t.Fatal("Local IP should not be empty")
	}

	t.Logf("Local IP: %s", localIP)
}

func TestIPService_GetIPInfo(t *testing.T) {
	service := NewIPService()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ipInfo, err := service.GetIPInfo(ctx)
	if err != nil {
		t.Fatalf("Failed to get IP info: %v", err)
	}

	if ipInfo.LocalIP == "" {
		t.Fatal("Local IP should not be empty")
	}

	if ipInfo.PublicIP == "" {
		t.Fatal("Public IP should not be empty")
	}

	t.Logf("IP Info - Local: %s, Public: %s", ipInfo.LocalIP, ipInfo.PublicIP)
}
