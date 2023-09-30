package dev

import (
	"strconv"
	"strings"
)

type NameFactory struct {
	chars       []string
	maxIndex    int
	pos         int
	rangeCount  int
	count       int
	currentName string
	config      *Config
}

func NewNameFactory(config *Config) *NameFactory {
	factory := &NameFactory{config: config}
	factory.chars = strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	factory.maxIndex = len(factory.chars) - 1
	factory.count = 0
	factory.pos = 0
	factory.rangeCount = config.nameStartAt
	factory.currentName = ""
	return factory
}

func (factory *NameFactory) GetName() string {
	if factory.pos == factory.maxIndex {
		factory.pos = 0
		factory.rangeCount++
	}

	factory.currentName = factory.chars[factory.pos] + strconv.Itoa(factory.rangeCount)
	factory.count += 1
	factory.pos += 1
	return factory.currentName
}

func (factory *NameFactory) Reset() {
	factory.pos = 0
	factory.rangeCount = 0
	factory.currentName = ""
	factory.count = 0
}
