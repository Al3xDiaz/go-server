package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/gorilla/mux"
)

type Config struct {
	Port string
	JWTSecret string
	Database string
}
type Server interface {
	Config() *Config
}
type Broker struct {
	config *Config
	router *mux.Router
}
func (b *Broker) Config() *Config {
	return b.Config
}
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nill, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.Database == "" {
		return nil, errors.New("database is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker,nil
}