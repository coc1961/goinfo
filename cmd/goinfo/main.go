package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coc1961/goinfo/internal/callhierarchy"
)

func main() {
	src := flag.String("s", "", "source go file")
	line := flag.Int("l", 0, "line number")
	col := flag.Int("c", 0, "column number")
	flag.Parse()

	if *src == "" || *line == -1 || *col == -1 {
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	a, err := callhierarchy.New().Parse(*src, *line, *col)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%v\n", a)
}
