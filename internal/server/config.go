package server

import "time"

type Config struct {
	BindAddr       string        `toml:"bind_addr"`
	LogLevel       string        `toml:"log_level"`
	MaxHeaderBytes int           `toml:"max_header_bytes"`
	ReadTimeout    time.Duration `toml:"read_timeout"`
	WriteTimeout   time.Duration `toml:"write_timeout"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:       ":3131",
		LogLevel:       "debug",
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}
