package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemInfo struct {
	Total, Available, Used, Free struct {
		name string
		val  uint64
	}
	UsedPercentage struct {
		name string
		val  int
	}
}

func MemoryInfo(m mem.VirtualMemoryStat) MemInfo {
	var memInfo MemInfo = MemInfo{
		Total: struct {
			name string
			val  uint64
		}{name: "Total", val: m.Total / 1024 / 1024},
		Available: struct {
			name string
			val  uint64
		}{name: "Available", val: m.Available / 1024 / 1024},
		Used: struct {
			name string
			val  uint64
		}{name: "Used", val: m.Used / 1024 / 1024},
		Free: struct {
			name string
			val  uint64
		}{name: "Free", val: m.Free / 1024 / 1024},
		UsedPercentage: struct {
			name string
			val  int
		}{name: "Used Percentage", val: int(m.UsedPercent)},
	}

	return memInfo
}

func main() {
	// get memory info
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
	}
	memoryInfo := MemoryInfo(*m)
	fmt.Printf("%s: %d GiB\n", memoryInfo.Total.name, memoryInfo.Total.val)
	fmt.Printf("%s: %d GiB\n", memoryInfo.Available.name, memoryInfo.Available.val)
	fmt.Printf("%s: %d GiB\n", memoryInfo.Used.name, memoryInfo.Used.val)
	fmt.Printf("%s: %d GiB\n", memoryInfo.Free.name, memoryInfo.Free.val)
	fmt.Printf("%s: %d\n", memoryInfo.UsedPercentage.name, memoryInfo.UsedPercentage.val)
}
