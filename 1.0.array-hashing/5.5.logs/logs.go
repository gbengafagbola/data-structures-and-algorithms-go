package main

import (
	"strings"
)
 
func processLogs(logs []string, threshold int32) []string {
	m := make(map[string]int) // Store user transaction counts

	// Process each log entry
	for _, log := range logs {
		transaction := strings.Fields(log) // Split log into [senderID, receiverID, amount]

		if len(transaction) < 3 {
			continue // Skip invalid logs
		}

		sender := transaction[0]
		receiver := transaction[1]

		// Count transactions for sender and receiver
		m[sender]++ // Sender always gets counted
		if sender != receiver {
			m[receiver]++ // Receiver only gets counted if different
		}
	}
	return 
}
 