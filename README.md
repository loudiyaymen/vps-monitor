# VPS Monitor

A lightweight system health monitoring tool written in Go. It provides real-time snapshots of CPU and memory usage with color-coded status reports based on resource load.

## Features

- Dynamic timestamping for every report.
- CPU usage percentage monitoring.
- Virtual memory (RAM) statistics.
- Color-coded terminal output (ANSI) for system status:
  - Green: Healthy (Low load)
  - Yellow: Warning (Moderate load)
  - Red: Critical (High load)

## Prerequisites

- Go 1.21 or higher
- Linux or macOS environment (for ANSI color support)

## Installation

1. Clone the repository:
   git clone git@github.com-personal:YourGitHubUsername/vps-monitor.git

2. Install dependencies:
   go mod tidy

3. Run the application:
   go run main.go

## License

MIT
