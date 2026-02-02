package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31;1m"
	ColorYellow = "\033[33;1m"
	ColorGreen  = "\033[32;1m"
	ColorCyan   = "\033[36m"
)

func main() {
	now := time.Now().Format("Monday, Jan 02, 2006 | 15:04:05")
	fmt.Printf("%s--- VPS Health Check | %s ---%s\n", ColorCyan, now, ColorReset)

	v, _ := mem.VirtualMemory()
	memUsed := v.UsedPercent

	c, _ := cpu.Percent(time.Second, false)
	cpuLoad := c[0]

	fmt.Printf("Memory: %vMB Free / %vMB Total (Used: %.2f%%)\n",
		v.Available/1024/1024, v.Total/1024/1024, memUsed)
	fmt.Printf("CPU Load: %.2f%%\n", cpuLoad)

	fmt.Println("----------------------------------------------")

	msg, color := getStatusMessage(cpuLoad, memUsed)

	fmt.Printf("Status: %s%s%s\n", color, msg, ColorReset)
}

func getStatusMessage(cpu float64, mem float64) (string, string) {
	if cpu > 80 || mem > 90 {
		return "CRITICAL - System is under heavy load!", ColorRed
	} else if cpu > 50 || mem > 70 {
		return "WARNING - Moderate resource usage.", ColorYellow
	}
	return "HEALTHY - All systems nominal.", ColorGreen
}
