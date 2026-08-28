package config

import (
	"os"
	"strconv"
)

type Config struct {
	Addr, Database string
	ReadTimeout    int
}

func Load() Config {
	c := Config{Addr: ":8080", Database: "inspection.db", ReadTimeout: 10}
	if v := os.Getenv("INSPECTION_ADDR"); v != "" {
		c.Addr = v
	}
	if v := os.Getenv("INSPECTION_DB"); v != "" {
		c.Database = v
	}
	if v, e := strconv.Atoi(os.Getenv("INSPECTION_TIMEOUT")); e == nil && v > 0 {
		c.ReadTimeout = v
	}
	return c
}
func (c Config) Validate() bool { return c.Addr != "" && c.Database != "" && c.ReadTimeout > 0 }
func (c Config) DSN() string    { return c.Database }
