package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/coc1961/goinfo/internal/callhierarchy"
)

func main() {
	src := flag.String("s", "", "source go file")
	line := flag.Int("l", 0, "line number")
	col := flag.Int("c", 0, "column number")
	find := flag.String("f", "", "function name with parameters")
	flag.Parse()

	if *src == "" {
		flag.CommandLine.Usage()
		os.Exit(1)
	}
	if (*line == -1 || *col == -1) && *find == "" {
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	if *find != "" {
		b, _ := os.ReadFile(*src)
		arr := strings.Split(string(b), "\n")
		for i, s := range arr {
			idx := strings.Index(s, find)
			if idx >= 0 {
				*line = i + 1
				*col = idx + 1
			}
		}
	}

	a, err := callhierarchy.New().Parse(*src, *line, *col)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%v \n", a)
}
