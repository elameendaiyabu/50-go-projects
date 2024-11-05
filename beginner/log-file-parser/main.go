package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func GetFileLines() []string {
	f, _ := os.Open("./log.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	return fileLines
}

func ParseLogFileLines(re *regexp.Regexp, fileLines []string) {
	for idx, line := range fileLines {
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			return
		}
		fmt.Printf("%s %d\n", strings.ToUpper("error"), idx+1)
		for i, k := range re.SubexpNames() {
			// ignore the first and the $_
			if i == 0 || k == "_" {
				continue
			}

			fmt.Printf("%-15s => %s\n", k, matches[i])
		}
		fmt.Printf("\n\n")
	}
}

func GenerateLogFile() {
	// Run the command within a shell to enable redirection
	cmd := exec.Command("bash", "-c", "journalctl -p err -b > log.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Log file generated.")
	fmt.Println(string(output)) // Outputs any messages from the command, if needed
}

func main() {
	GenerateLogFile()
	logFormat := `(?P<month>\w+) (?P<day>\d{2}) (?P<time>\d{2}:\d{2}:\d{2}) (?P<hostname>\S+) (?P<service>\S+):(?: \[(?P<component>[^\]]*)\])?(?: \*(?P<log_level>[^\*]+)\*)?(?: \[(?P<driver>[^\]]+)\])?(?: \[(?P<gpu_id>[^\]]+)\])? (?P<message>.+)`
	re := regexp.MustCompile(logFormat)

	fileLines := GetFileLines()

	ParseLogFileLines(re, fileLines)
}
