package fortune

import (
	"os/exec"
	"strings"
)

func Fortune() string {
	// exec fortune and return result from shell
	out, err := exec.Command("fortune").Output()
	if err != nil {
		return "Error executing fortune command"
	}
	return strings.TrimSpace(string(out))
}
