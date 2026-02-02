package main

import "testing"

func TestGetStatusMessage(t *testing.T) {
	// Test Case 1: Healthy
	status, _ := getStatusMessage(10.0, 20.0)
	if status != "HEALTHY - All systems nominal." {
		t.Errorf("Expected Healthy status, got %s", status)
	}

	// Test Case 2: Warning
	status, _ = getStatusMessage(60.0, 30.0)
	if status != "WARNING - Moderate resource usage." {
		t.Errorf("Expected Warning status, got %s", status)
	}

	// Test Case 3: Critical
	status, _ = getStatusMessage(10.0, 95.0)
	if status != "CRITICAL - System is under heavy load!" {
		t.Errorf("Expected Critical status, got %s", status)
	}
}
