package config

import (
	"os"
)

// Loader handles YAML parsing and environment variable overrides.
type Loader struct{}

// NewLoader creates a new Loader instance.
func NewLoader() *Loader {
	return &Loader{}
}

// LoadFromFile reads a YAML file, flattens nested keys to dot-notation,
// and returns the resulting key-value pairs.
func (l *Loader) LoadFromFile(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return flattenYAML(data)
}

// ApplyEnvOverrides overrides values with environment variables.
// Environment variable format: AUTOMATA_<KEY> where <KEY> is the
// dot-notation key uppercased with dots replaced by underscores.
func (l *Loader) ApplyEnvOverrides(values map[string]string) {
	applyEnvOverrides(values)
}

// ParseYAML parses YAML bytes into a flat map with dot-notation keys.
func (l *Loader) ParseYAML(data []byte) (map[string]string, error) {
	return flattenYAML(data)
}
