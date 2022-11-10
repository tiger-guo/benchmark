package errf

import (
	"fmt"
	"os"
)

// CheckErr check process error.
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
