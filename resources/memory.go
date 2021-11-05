package resources

import "github.com/shirou/gopsutil/v3/mem"

type MemoryStats struct {
	Total uint64 `json:"total"`
	Used uint64 `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func GetMemoryStats() (MemoryStats, error) {
	stats, err := mem.VirtualMemory()

	return MemoryStats{
		stats.Total,
		stats.Used,
		stats.UsedPercent,
	}, err
}