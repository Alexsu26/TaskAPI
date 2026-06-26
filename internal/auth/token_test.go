package auth

import (
	"errors"
	"testing"
	"time"
)

func TestTokenManager_GenerateAndParseToken(t *testing.T) {
	manager := NewTokenManager("test-token", time.Hour)

	token, err := manager.GenerateToken(123)
	if err != nil {
		t.Fatalf("generate failed, err: %v", err)
	}

	c, err := manager.ParseToken(token)
	if err != nil {
		t.Fatalf("parse token err: %v", err)
	}
	if c.UserID != 123 {
		t.Fatalf("expected user id is 123, but got %v", c.UserID)
	}
}

func TestTokenManager_InvalidUserID(t *testing.T) {
	manager := NewTokenManager("test-token", time.Hour)
	_, err := manager.GenerateToken(0)
	if !errors.Is(err, ErrTokenInvalid) {
		t.Fatalf("expected ErrTokenInvalid, got %v", err)
	}
}

func TestTokenManager_InvalidToken(t *testing.T) {
	manager := NewTokenManager("test-token", time.Hour)
	_, err := manager.ParseToken("bad-token")
	if !errors.Is(err, ErrTokenInvalid) {
		t.Fatalf("expected ErrTokenInvalid, got %v", err)
	}
}
