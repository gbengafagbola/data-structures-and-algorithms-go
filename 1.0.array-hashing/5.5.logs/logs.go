package main

import (
	"sort"
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
		m[sender]++
		if sender != receiver {
			m[receiver]++
		}
	}

	// Collect users that meet the threshold
	var result []string
	for user, count := range m {
		if int32(count) >= threshold {
			result = append(result, user)
		}
	}

	// Sort the result lexicographically
	sort.Strings(result)

	return result
}
