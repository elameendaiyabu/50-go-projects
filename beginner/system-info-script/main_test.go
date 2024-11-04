package main

import (
	"testing"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func TestMemoryInfo(t *testing.T) {
	m := &mem.VirtualMemoryStat{
		Total:       16139497472,
		Available:   6285946880,
		Used:        9026813952,
		Free:        2896203776,
		UsedPercent: 55.929956726722054,
	}

	got := MemoryInfo(*m)
	want := MemInfo{
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

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestHostInformation(t *testing.T) {
	h := &host.InfoStat{
		Hostname: "bloodhound",
		OS:       "linux",
		Platform: "arch",
	}

	got := HostInformation(*h)
	want := HostInfo{
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

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestCpuInformation(t *testing.T) {
	c := cpu.InfoStat{
		ModelName: "AMD Ryzen 9 5900HS with Radeon Graphics",
		Mhz:       4680,
		CacheSize: 512,
	}
	numOfCores := 16

	got := CpuInformation(c, numOfCores)
	want := CpuInfo{
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
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
