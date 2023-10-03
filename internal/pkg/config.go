package pkg

import (
	"fmt"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
	"os"
	"strings"
)

type ServerRuntime string

const (
	ServerRuntimeOpenApiKey = "openapi"
	ServerRuntimeStdHttpKey = "stdhttp"
)

type AppConfig struct {
	// Middlewares configures middlewares for endpoints
	Middlewares map[string][]MiddlewareConfig
	// Port defines what port to map to
	Port int
	// Server defines what server to use
	Server ServerRuntime
}

type MiddlewareConfig struct {
	// Name of middleware to be activated
	Name string
	// Config stores an arbitrary set of keys and values tree
	Config map[string]any
}

const (
	portKey   = "port"
	serverKey = "server"
)

var DefaultConfig AppConfig = AppConfig{
	Port:   8081,
	Server: ServerRuntimeOpenApiKey,
}

func MakeAppConfig() (AppConfig, error) {
	k := koanf.New(".")
	if err := k.Load(file.Provider("./config.yaml"), yaml.Parser()); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("error loading config from ./config.yaml: %s", err)
		}
		log.Println("no config.yaml detected, skipping")
	}
	if err := k.Load(file.Provider("./config.json"), json.Parser()); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("error loading config from ./config.json: %s", err)
		}
		log.Println("no config.json detected, skipping")
	}

	if err := k.Load(env.Provider("STREAMSAPI_", ".", func(s string) string {
		return strings.Replace(
			strings.ToLower(
				strings.TrimPrefix(s, "STREAMSAPI_"),
			), "_", ".", -1)
	}), nil); err != nil {
		log.Fatalf("error loading config from enviroment: %v", err)
	}

	if k.Exists(portKey) {
		DefaultConfig.Port = k.Int(portKey)
	}
	if k.Exists(serverKey) {
		val, serverRuntimeErr := parseServerRuntime(k.String(serverKey))
		if serverRuntimeErr != nil {
			return AppConfig{}, serverRuntimeErr
		}
		DefaultConfig.Server = val
	}
	return DefaultConfig, nil
}

func parseServerRuntime(s string) (ServerRuntime, error) {
	switch s {
	case ServerRuntimeStdHttpKey, ServerRuntimeOpenApiKey:
		return ServerRuntime(s), nil
	}
	return "", fmt.Errorf("unrecognized server runtime %s", s)
}
