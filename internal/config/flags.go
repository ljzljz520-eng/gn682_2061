package config

import (
	"flag"
	"os"
)

func FromArgs(args []string) Config {
	c := Load()
	fs := flag.NewFlagSet("inspectiond", flag.ContinueOnError)
	addr := fs.String("addr", c.Addr, "listen address")
	db := fs.String("db", c.Database, "database path")
	_ = fs.Parse(args)
	c.Addr = *addr
	c.Database = *db
	return c
}
func EnsureDir(path string) error {
	if path == "" {
		return nil
	}
	return os.MkdirAll(path, 0755)
}
func IsMemory(path string) bool { return path == ":memory:" || path == "file::memory:" }
