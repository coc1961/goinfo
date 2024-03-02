package run

import (
	"fmt"
	"os/exec"
	"strings"
)

func CallHierarchy(args ...string) (string, error) {
	_pls, err := pls()
	if err != nil {
		return "", err
	}
	_args := append([]string{"call_hierarchy"}, args...)
	cmd := exec.Command(_pls, _args...)

	b, err := cmd.Output()
	if err != nil {
		fmt.Printf("_pls: %v\n", cmd)
		return "", err
	}
	return string(b), nil
}

func pls() (string, error) {
	cmd := exec.Command("which", "gopls")
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	cmd = exec.Command("readlink", "-f", strings.TrimRight(string(b), "\n"))
	b, err = cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(b), "\n"), nil
}
