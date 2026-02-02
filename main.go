package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// ANSI color codes for terminal output accessibility
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31;1m"
	ColorYellow = "\033[33;1m"
	ColorGreen  = "\033[32;1m"
	ColorCyan   = "\033[36m"
)

func main() {
	// Fetch the current time dynamically to ensure the report is always up to date
	now := time.Now().Format("Monday, Jan 02, 2006 | 15:04:05")
	fmt.Printf("%s--- VPS Health Check | %s ---%s\n", ColorCyan, now, ColorReset)

	// Retrieve Virtual Memory statistics from the OS
	v, _ := mem.VirtualMemory()
	memUsed := v.UsedPercent

	// Retrieve CPU usage percentage over a 1-second interval
	c, _ := cpu.Percent(time.Second, false)
	cpuLoad := c[0]

	fmt.Printf("Memory: %vMB Free / %vMB Total (Used: %.2f%%)\n",
		v.Available/1024/1024, v.Total/1024/1024, memUsed)
	fmt.Printf("CPU Load: %.2f%%\n", cpuLoad)

	fmt.Println("----------------------------------------------")
	printStatus(cpuLoad, memUsed)
}

// printStatus evaluates system load and prints a color-coded health message
func printStatus(cpu float64, mem float64) {
	var status string
	var color string

	// Threshold logic: Red for >80% CPU or >90% RAM
	if cpu > 80 || mem > 90 {
		color = ColorRed
		status = "CRITICAL - System is under heavy load!"
	} else if cpu > 50 || mem > 70 {
		color = ColorYellow
		status = "WARNING - Moderate resource usage."
	} else {
		color = ColorGreen
		status = "HEALTHY - All systems nominal."
	}

	fmt.Printf("Status: %s%s%s\n", color, status, ColorReset)
}
