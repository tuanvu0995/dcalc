package dev

import (
	"fmt"
	"strconv"
)

type Helpers struct {
	Config *Config
}

// FormatFloat formats a float64 value to a string
func (helpers *Helpers) FormatFloat(value float64) string {
	format := "%." + strconv.Itoa(helpers.Config.GetFormatDecimals()) + "f"
	return fmt.Sprintf(format, value)
}
