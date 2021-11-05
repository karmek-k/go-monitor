package resources

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCpuStats(interval time.Duration) ([]float64, error) {
	return cpu.Percent(interval, true)
}