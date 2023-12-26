package loggerTest

import "fmt"

func LogErrorTest(err error) {
	fmt.Printf("There was an unexpected error: %s", err.Error())
}
