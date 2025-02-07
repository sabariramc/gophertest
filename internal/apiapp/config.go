package apiapp

import (
	"context"
	"fmt"
	"gopertest/internal/appbase"
	"gopertest/internal/service/math"

	e "gopertest/internal/env"

	"gitlab.com/engineering/products/api_security/go-common/api/server"
	"gitlab.com/engineering/products/api_security/go-common/environment"
	"gitlab.com/engineering/products/api_security/go-common/log"
	"gitlab.com/engineering/products/api_security/go-common/log/factory"
)

var env = environment.EnvironmentReader{}

type Config struct {
	ServiceName    string
	Base           *appbase.AppBase
	Server         *server.RESTAPIServer
	Log            log.Log
	Port           uint16
	Math           *math.Math
	MetricsEnabled bool
}

var authorizedClients = []string{"client1", "client2"}

func GetDefaultConfig(ctx context.Context) (*Config, error) {
	port := env.ReadInteger(e.EnvAPIServerPort, 8080)
	log := factory.NewLog(ctx, "HTTPServer")
	if port < 1000 || port > 65535 {
		log.Warnf(ctx, "Invalid port number: %d: setting port to 8080", port)
		port = 8080
	}
	srv, err := server.NewRESTAPIServer(uint16(port), server.RouteConfig{}, env.ReadString(e.EnvJWTSURL, ""), authorizedClients, env.ReadBoolean(e.EnvServerAuthEnabled, false))
	if err != nil {
		return nil, err
	}
	return &Config{
		ServiceName:    env.ReadString(e.EnvServiceName, "default"),
		Base:           appbase.New(ctx),
		Server:         srv,
		Log:            log,
		Port:           uint16(port),
		MetricsEnabled: env.ReadBoolean(e.EnvMetricsEnabled, false),
	}, nil
}

type Option func(*Config) error

func (c *Config) Validate() error {
	if c.Math == nil {
		return fmt.Errorf("math cannot be nil")
	}
	return nil
}

func WithServiceName(serviceName string) Option {
	return func(c *Config) error {
		c.ServiceName = serviceName
		return nil
	}
}

func WithBase(base *appbase.AppBase) Option {
	return func(c *Config) error {
		c.Base = base
		return nil
	}
}

func WithServer(server *server.RESTAPIServer) Option {
	return func(c *Config) error {
		c.Server = server
		return nil
	}
}

func WithLog(log log.Log) Option {
	return func(c *Config) error {
		c.Log = log
		return nil
	}
}

func WithMetricsEnabled() Option {
	return func(c *Config) error {
		c.MetricsEnabled = true
		return nil
	}
}

func WithMath(math *math.Math) Option {
	return func(c *Config) error {
		c.Math = math
		return nil
	}
}
