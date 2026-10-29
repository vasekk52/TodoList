package database

import (
	"testing"
	"todoapp/internal/model"
)

func TestNewDB(t *testing.T) {
	cfg := model.Config{
		DB_HOST: "127.0.0.1",
		DB_PORT: "5432",
		DB_USER: "myuser",
		DB_PASS: "mypass",
		DB_NAME: "mydb",
	}
	validStr := "host=127.0.0.1 port=5432 user=myuser password=mypass dbname=mydb sslmode=disable"

	str := FormatStr(cfg)

	if str != validStr {
		t.Errorf("got %v, want %v", str, validStr)
	}
}

func TestNewPostgresDB_ConnectionError(t *testing.T) {
	cfg := model.Config{
		DB_HOST: "test123_",
		DB_PORT: "9999p",
	}

	_, err := NewPostgresDB(cfg)
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}

}