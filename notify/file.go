package notify

import (

	"fmt"
	"os"
	
)





func CreateAndAddContents (content string){

	err:= os.WriteFile("weatherReport.txt", []byte(content), 0644)

	if err != nil{
		fmt.Println("Error creating file: ", err)
	}
}