package testpackage

import "fmt"

func TestPackage(level int) {
	internalTest(level)
}

func internalTest(level int) {
	if level < 10 {
		internalTest(level + 1)
	}
	fmt.Println(level)
}
