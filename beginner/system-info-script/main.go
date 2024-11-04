package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/shirou/gopsutil/v4/cpu"
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

type CpuInfo struct {
	ModelName struct {
		name string
		val  string
	}
	CacheSize, ProcessorSpeed, NumberOfCores struct {
		name string
		val  int
	}
}

func MemoryInfo(m mem.VirtualMemoryStat) MemInfo {
	memInfo := MemInfo{
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
	hostInfo := HostInfo{
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

func CpuInformation(c cpu.InfoStat, numOfCores int) CpuInfo {
	cpuinfo := CpuInfo{
		ModelName: struct {
			name string
			val  string
		}{
			name: "Model Name",
			val:  c.ModelName,
		},
		CacheSize: struct {
			name string
			val  int
		}{
			name: "Cache Size",
			val:  int(c.CacheSize),
		},
		ProcessorSpeed: struct {
			name string
			val  int
		}{
			name: "Processor Speed",
			val:  int(c.Mhz),
		},
		NumberOfCores: struct {
			name string
			val  int
		}{
			name: "Number of CPU Cores",
			val:  numOfCores,
		},
	}
	return cpuinfo
}

func main() {
	// get memory info
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
	}
	memoryInfo := MemoryInfo(*m)
	fmt.Println(strings.ToUpper("Memory Information:"))
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
	fmt.Println(strings.ToUpper("Host Information:"))
	fmt.Printf("%s: %s \n", hostInfo.HostName.name, hostInfo.HostName.val)
	fmt.Printf("%s: %s \n", hostInfo.Os.name, hostInfo.Os.val)
	fmt.Printf("%s: %s \n\n", hostInfo.Platform.name, hostInfo.Platform.val)

	// get cpu info
	c, err := cpu.Info()
	if err != nil {
		log.Println(err)
	}
	cpuInfo := CpuInformation(c[0], len(c))
	fmt.Println(strings.ToUpper("CPU Information:"))
	fmt.Printf("%s: %s \n", cpuInfo.ModelName.name, cpuInfo.ModelName.val)
	fmt.Printf("%s: %d \n", cpuInfo.CacheSize.name, cpuInfo.CacheSize.val)
	fmt.Printf("%s: %d \n", cpuInfo.ProcessorSpeed.name, cpuInfo.ProcessorSpeed.val)
	fmt.Printf("%s: %d \n", cpuInfo.NumberOfCores.name, cpuInfo.NumberOfCores.val)
}
