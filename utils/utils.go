package utils

import (
	"fmt"
	"os"
)

func Debug(args ...interface{}) {
	if os.Getenv("CDB_DEBUG") != "" {
		fmt.Println(args...)
	}
}
