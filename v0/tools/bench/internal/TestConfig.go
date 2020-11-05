package internal

import "time"

type TestConfig struct {
	Name string
	Func func(uint64, int) time.Duration
	Inputs interface{}
	Competitor func(uint64, int) time.Duration
}
