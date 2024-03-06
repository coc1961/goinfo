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
	it := getIT()
	if it != nil {
		it.Get()
	}
}

func getIT() Inter {
	return nil
}

type Inter interface {
	Get() string
}

type InterImpl1 struct {
}

func (i InterImpl1) Get() string {
	return "InterImpl1"
}

type InterImpl2 struct {
}

func (i InterImpl2) Get() string {
	return "InterImpl1"
}
