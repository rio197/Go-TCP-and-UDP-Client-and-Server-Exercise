package handleError

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
