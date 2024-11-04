package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v4/host"
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

type HostInfo struct {
	HostName, Os, Platform struct{ name, val string }
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

func HostInformation(h host.InfoStat) HostInfo {
	var hostInfo HostInfo = HostInfo{
		HostName: struct {
			name string
			val  string
		}{
			name: "HostName",
			val:  h.Hostname,
		},
		Os: struct {
			name string
			val  string
		}{
			name: "OS",
			val:  h.OS,
		},
		Platform: struct {
			name string
			val  string
		}{
			name: "Platform",
			val:  h.Platform,
		},
	}

	return hostInfo
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
	fmt.Printf("%s: %d\n\n", memoryInfo.UsedPercentage.name, memoryInfo.UsedPercentage.val)

	// get host Info
	h, err := host.Info()
	if err != nil {
		log.Println(err)
	}
	hostInfo := HostInformation(*h)
	fmt.Printf("%s: %s \n", hostInfo.HostName.name, hostInfo.HostName.val)
	fmt.Printf("%s: %s \n", hostInfo.Os.name, hostInfo.Os.val)
	fmt.Printf("%s: %s \n", hostInfo.Platform.name, hostInfo.Platform.val)
}
