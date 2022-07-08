package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hedhyw/jsoncjson"
	"github.com/pkg/errors"
)

type Config struct {
	Environment string      `json:"environment"`
	Log         LogConfig   `json:"log"`
	API         API         `json:"api"`
	DB          DB          `json:"db"`
	TraceConfig TraceConfig `json:"trace_config"`
	GrpcCfg     GrpcCfg     `json:"grpc_cfg"`
}

type LogConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"`
}

type API struct {
	Address         string        `json:"address"`
	GRPCAddress     string        `json:"grpc_address"`
	ReadTimeout     time.Duration `json:"read_timeout"`
	WriteTimeout    time.Duration `json:"write_timeout"`
	ShutdownTimeout time.Duration `json:"shutdown_timeout"`
	MainPath        string        `json:"main_path"`
	TokenTTl        int64         `json:"token_ttl"`
	JwtKey          string        `json:"jwt_key"`
}

type DB struct {
	URL          string `json:"url"`
	SchemaName   string `json:"schema_name"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type TraceConfig struct {
	Disabled      bool    `json:"disabled"`
	ServiceName   string  `json:"service_name"`
	AgentHostPort string  `json:"agent_host_port"`
	SamplerType   string  `json:"sampler_type"`
	SamplerParam  float64 `json:"sampler_param"`
}

type GrpcCfg struct {
	AddressGatewaySrv string `json:"gateway_address"`
	AddressOTPSrv     string `json:"address_otp"`
}

func DefaultConfig() (cfg *Config) {
	return &Config{
		Environment: "development",
		Log: LogConfig{
			Level:  "debug",
			Format: "json",
		},
		API: API{
			Address:         ":5004",
			ReadTimeout:     30 * time.Second,
			WriteTimeout:    30 * time.Second,
			ShutdownTimeout: 5 * time.Second,
			MainPath:        "/api/v1",
		},
		DB: DB{
			URL:          "postgres://rpguser:rpguser@10.10.14.10:5452/rpgDB?sslmode=disable",
			SchemaName:   "RpgDB",
			MaxOpenConns: 2,
			MaxIdleConns: 2,
		},
		TraceConfig: TraceConfig{
			Disabled:      true,
			ServiceName:   "go-aut-registration-user-otp",
			AgentHostPort: "",
			SamplerType:   "remote",
			SamplerParam:  1,
		},
		GrpcCfg: GrpcCfg{
			AddressOTPSrv: ":5014",
		},
	}
}

func LoadConfig(path string) (*Config, error) {
	cfg := DefaultConfig()

	confPath := os.Getenv("CONFIG_PATH")
	if confPath != "" {
		path = confPath
	}
	f, err := os.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("opening: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()
	jsoncReader := jsoncjson.NewReader(f)
	if err = json.NewDecoder(jsoncReader).Decode(cfg); err != nil {
		return nil, errors.WithStack(err)
	}
	return cfg, nil
}
