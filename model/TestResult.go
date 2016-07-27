package model

import (
	"github.com/gaia-docker/tugbot-parse"
	"time"
)

type TestResult struct {
	ContainerId string
	ExitCode    int
	FinishedAt  time.Time
	HostName    string
	ImageName   string
	StartedAt   time.Time
	TestCase    *parse.TestSet
}
