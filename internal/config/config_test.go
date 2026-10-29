package config

import (
	"os"
	"testing"
)

func TestLoad_Defaults(t *testing.T) {
	envKeys := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "SERVER_PORT"}
	for _, key := range envKeys {
		t.Setenv(key, "")
		os.Unsetenv(key)
	}

	cfg := Load()

	if cfg.DB_HOST != "localhost" {
		t.Errorf("got %v, want localhost", cfg.DB_HOST)
	}
	if cfg.DB_PORT != "5432" {
		t.Errorf("got %v, want 5432", cfg.DB_PORT)
	}
	if cfg.DB_USER != "postgres" {
		t.Errorf("got %v, want postgres", cfg.DB_USER)
	}
	if cfg.DB_PASS != "postgres123" {
		t.Errorf("got %v, want postgres123", cfg.DB_PASS)
	}
	if cfg.DB_NAME != "todo_db" {
		t.Errorf("got %v, want todo_db", cfg.DB_NAME)
	}
	if cfg.ServerPort != "8080" {
		t.Errorf("got %v, want 8080", cfg.ServerPort)
	}
}

func TestLoad_CustomEnv(t *testing.T) {
	t.Setenv("DB_HOST", "test")
	t.Setenv("DB_PORT", "5433")
	t.Setenv("DB_USER", "admin")
	t.Setenv("DB_PASSWORD", "pass")
	t.Setenv("DB_NAME", "test1")
	t.Setenv("SERVER_PORT", "8081")

	cfg := Load()

	if cfg.DB_HOST != "test" {
		t.Errorf("got %v, want test", cfg.DB_HOST)
	}
	if cfg.DB_PORT != "5433" {
		t.Errorf("got %v, want 5433", cfg.DB_PORT)
	}
	if cfg.DB_USER != "admin" {
		t.Errorf("got %v, want admin", cfg.DB_USER)
	}
	if cfg.DB_PASS != "pass" {
		t.Errorf("got %v, want pass", cfg.DB_PASS)
	}
	if cfg.DB_NAME != "test1" {
		t.Errorf("got %v, want test1", cfg.DB_NAME)
	}
	if cfg.ServerPort != "8081" {
		t.Errorf("got %v, want 8081", cfg.ServerPort)
	}
}

func TestGetEnv_EmptyStringFallback(t *testing.T) {
	t.Setenv("EMPTY_VAR", "")

	got := getEnv("EMPTY_VAR", "fallback")
	if got != "fallback" {
		t.Errorf("got %v, want %v", got, "fallback")
	}
}