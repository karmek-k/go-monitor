package resources

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CpuStats struct {
	CoreCount int `json:"coreCount"`
	UsagePercent []float64 `json:"usagePercent"`
}

func GetCpuStats(interval time.Duration, logical bool) (*CpuStats, error) {
	count, err := cpu.Counts(logical)
	if err != nil {
		return nil, err
	}

	usage, err := cpu.Percent(interval, logical)
	if err != nil {
		return nil, err
	}

	return &CpuStats{count, usage}, nil
}