package main

import (
	"errors"
)

type Config struct {
	Endpoint string
}

type ConfigProvider interface {
	GetServiceConfig(serviceName string) (*Config, error)
}

type InMemoryConfigProvider struct {
	cfg map[string]*Config
}

func NewInMemoryConfigProvider() *InMemoryConfigProvider {
	return &InMemoryConfigProvider{
		cfg: make(map[string]*Config, 4),
	}
}

var ErrServiceNotFound = errors.New("service not found")

func (i *InMemoryConfigProvider) GetServiceConfig(serviceName string) (*Config, error) {
	cfg, ok := i.cfg[serviceName]
	if !ok {
		return nil, ErrServiceNotFound
	}
	return cfg, nil
}
