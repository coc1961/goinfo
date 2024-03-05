package callhierarchy

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/coc1961/goinfo/internal/run"
)

type callService struct {
}

type call struct {
	caller    string
	level     int
	isMain    bool
	path      string
	callStack []*call
}

func New() *callService {
	return &callService{}
}

func (c *callService) Parse(path string, line, col int) (*call, error) {
	return c.parse(fmt.Sprintf("%s:%d:%d", path, line, col), 0, []string{})
}
func (c *callService) parse(path string, level int, callStack []string) (*call, error) {
	if strings.Contains(path, "@") {
		return nil, errors.New("skip library")
	}
	var goRoot = os.Getenv("GOROOT")

	fmt.Fprintln(os.Stderr, "processing...", path)

	str, err := run.CallHierarchy(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, path, err)
		return nil, err
	}

	arr := strings.Split(strings.TrimRight(str, "\n"), "\n")
	var cl *call

	for _, s := range arr {
		arr1 := strings.Split(s, " ")
		switch len(arr1) {
		case 5:
			cl1 := &call{}
			cl1.caller = arr1[2]
			cl1.path = arr1[4]
			cl1.isMain = true
			cl1.level = level
			cl1.callStack = []*call{}
			cl = cl1
		case 10:
			if cl != nil {
				cl1 := &call{}
				cl1.caller = arr1[7]
				cl1.path = arr1[9]
				cl1.isMain = false
				cl1.level = level
				cl1.callStack = []*call{}
				processFun := true
				for _, cs := range callStack {
					if cs == arr1[7] {
						processFun = false
						break
					}
				}
				if processFun {
					if !strings.Contains(cl1.path, goRoot) {
						_callStack := append([]string{cl1.path}, callStack...)
						tmp, err := c.parse(cl1.path, level+1, _callStack)
						if err == nil && tmp != nil {
							cl1.callStack = tmp.callStack
						} else {
							fmt.Fprintln(os.Stderr, err)
						}
					}
				}
				cl.callStack = append(cl.callStack, cl1)
			}
		}
	}

	if len(arr) > 0 {
		//  Is interface?
		if strings.Index(arr[len(arr)-1], "identifier:") == 0 {
			str, err := run.Implementation(path)
			if err == nil {
				arr := strings.Split(strings.TrimRight(str, "\n"), "\n")
				for _, s := range arr {
					arr1 := strings.Split(s, " ")
					cl1, err := c.parse(arr1[0], level, callStack)
					if err == nil {
						cl.callStack = append(cl.callStack, cl1)
					}
				}
			}
		}
	}

	return cl, nil
}

func (c *call) String() string {
	return print(c, 0)
}

func print(c *call, level int) string {
	currPath, _ := os.Getwd()
	tab := strings.Repeat("   ", level)
	b := bytes.Buffer{}
	pt := c.path
	if strings.Index(pt, currPath) == 0 {
		pt = pt[len(currPath):]
	}
	_, _ = b.WriteString(fmt.Sprintf("%slevel:%d Func:%s Path:%s\n", tab, level, c.caller, pt))
	for _, c1 := range c.callStack {
		b.WriteString(print(c1, level+1))
	}
	return b.String()
}
