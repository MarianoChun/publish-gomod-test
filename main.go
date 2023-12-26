package loggerTest

import "fmt"

func LogErrorTest(err error) {
	fmt.Sprintf("There was an unexpected error: %s", err.Error())
}
