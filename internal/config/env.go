package config

import "os"

func Env(name, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}
func EnvBool(name string, def bool) bool {
	v := os.Getenv(name)
	if v == "1" || v == "true" {
		return true
	}
	if v == "0" || v == "false" {
		return false
	}
	return def
}
func EnvInt(name string, def int) int {
	v := os.Getenv(name)
	n := 0
	for _, r := range v {
		if r < '0' || r > '9' {
			return def
		}
		n = n*10 + int(r-'0')
	}
	if v == "" {
		return def
	}
	return n
}
func (c Config) Address() string      { return c.Addr }
func (c Config) DatabasePath() string { return c.Database }
func (c Config) TimeoutSeconds() int  { return c.ReadTimeout }
func (c Config) IsProduction() bool   { return c.Addr != ":8080" }
func (c Config) IsTest() bool         { return c.Database == ":memory:" }
func (c Config) Clone() Config        { return c }
func (c Config) WithAddress(v string) Config {
	if v != "" {
		c.Addr = v
	}
	return c
}
func (c Config) WithDatabase(v string) Config {
	if v != "" {
		c.Database = v
	}
	return c
}
func (c Config) WithTimeout(v int) Config {
	if v > 0 {
		c.ReadTimeout = v
	}
	return c
}
func (c Config) Values() map[string]string {
	return map[string]string{"addr": c.Addr, "database": c.Database}
}
func (c Config) HasDatabase() bool  { return c.Database != "" }
func (c Config) HasAddress() bool   { return c.Addr != "" }
func (c Config) TimeoutValid() bool { return c.ReadTimeout > 0 }
func (c Config) PortHint() string {
	if c.Addr == ":8080" {
		return "8080"
	}
	return c.Addr
}
func (c Config) StorageHint() string {
	if IsMemory(c.Database) {
		return "memory"
	}
	return "sqlite"
}
func (c Config) StringValues() []string { return []string{c.Addr, c.Database} }
func MergeConfig(a, b Config) Config {
	c := a
	if b.Addr != "" {
		c.Addr = b.Addr
	}
	if b.Database != "" {
		c.Database = b.Database
	}
	if b.ReadTimeout > 0 {
		c.ReadTimeout = b.ReadTimeout
	}
	return c
}
func DefaultAddress() string        { return ":8080" }
func DefaultDatabase() string       { return "inspection.db" }
func DefaultTimeout() int           { return 10 }
func IsValidAddress(v string) bool  { return len(v) > 1 && v[0] == ':' || len(v) > 3 }
func IsValidDatabase(v string) bool { return v != "" }
func NormalizeAddress(v string) string {
	if v == "" {
		return DefaultAddress()
	}
	return v
}
func NormalizeDatabase(v string) string {
	if v == "" {
		return DefaultDatabase()
	}
	return v
}
