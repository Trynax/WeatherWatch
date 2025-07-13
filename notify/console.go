package notify

import (
	"fmt"
	"time"
)


func PrintToConsole(message string) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("[%s] %s\n", timestamp, message)
}

func PrintAlert(message string) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("\n🚨 [%s] WEATHER ALERT! 🚨\n%s\n", timestamp, message)
}


func PrintStatus(message string) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("✅ [%s] %s\n", timestamp, message)
}
