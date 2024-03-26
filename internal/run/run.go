package run

import (
	"os/exec"
	"strings"
)

func Implementation(args ...string) (string, error) {
	_pls, err := pls()
	if err != nil {
		return "", err
	}
	_args := append([]string{"implementation"}, args...)
	cmd := exec.Command(_pls, _args...)

	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(b), nil

}

func CallHierarchy(args ...string) (string, error) {
	_pls, err := pls()
	if err != nil {
		return "", err
	}
	_args := append([]string{"call_hierarchy"}, args...)
	cmd := exec.Command(_pls, _args...)

	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Find gopls program path
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
