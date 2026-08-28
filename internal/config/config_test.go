package config

import "testing"

func TestConfigLoad(t *testing.T) {
	if !Load().Validate() {
		t.Fatal()
	}
}
