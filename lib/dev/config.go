package dev

type Config struct {
	formatDecimals int
	nameStartAt    int
}

func NewConfig() *Config {
	config := &Config{
		formatDecimals: 2,
		nameStartAt:    0,
	}
	return config
}

// Setter
func (config *Config) SetFormatDecimals(formatDecimals int) {
	config.formatDecimals = formatDecimals
}

func (config *Config) SetNameStartAt(nameStartAt int) {
	config.nameStartAt = nameStartAt
}

// Getter
func (config *Config) GetFormatDecimals() int {
	return config.formatDecimals
}

func (config *Config) GetNameStartAt() int {
	return config.nameStartAt
}
