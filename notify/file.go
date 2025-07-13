package notify

import (
	"fmt"
	"os"
	"time"
)

func CreateAndAddContents(content string) {
	
	file, err := os.OpenFile("weatherReport.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()


	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s\n\n", timestamp, content)

	if _, err := file.WriteString(logEntry); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
