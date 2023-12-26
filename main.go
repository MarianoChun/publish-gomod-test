package publish_module_test

import "fmt"

func LogErrorTest(err error) {
	fmt.Sprintf("There was an unexpected error: %s", err.Error())
}
