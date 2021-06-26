package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type YamlConfigProvider struct {
	// 作为名字到配置的映射
	Services map[string]*Config `yaml:"services"`
}

func NewYamlConfigProvider(filepath string) (*YamlConfigProvider, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	ycp := &YamlConfigProvider{}
	err = yaml.Unmarshal(content, ycp)
	return ycp, err
}

func (y *YamlConfigProvider) GetServiceConfig(serviceName string) (*Config, error) {
	cfg, ok := y.Services[serviceName]
	if !ok {
		return nil, ErrServiceNotFound
	}
	return cfg, nil
}
