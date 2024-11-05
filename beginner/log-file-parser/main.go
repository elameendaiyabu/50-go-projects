package main

import (
	"fmt"
	"regexp"
)

func main() {
	logExample := `Nov 05 08:47:35 bloodhound kernel: [drm:drm_new_set_master] *ERROR* [nvidia-drm] [GPU ID 0x00000100] Failed to grab modeset ownership`
	logFormat := `(?P<month>\w+) (?P<day>\d{2}) (?P<time>\d{2}:\d{2}:\d{2}) (?P<hostname>\w+) (?P<service>\w+): \[(?P<component>.*?)\] \*(?P<log_level>.*?)\* \[(?P<driver>.*?)\] \[(?P<gpu_id>.*?)\] (?P<message>.*)`

	re := regexp.MustCompile(logFormat)
	matches := re.FindStringSubmatch(logExample)

	for i, k := range re.SubexpNames() {
		// ignore the first and the $_
		if i == 0 || k == "_" {
			continue
		}

		fmt.Printf("%-15s => %s\n", k, matches[i])
	}
}
