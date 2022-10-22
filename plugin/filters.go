package plugin

import (
	"os/exec"
)

func HasExec(name string) func() bool {
	return func() bool {
		_, err := exec.LookPath(name)
		return err == nil
	}
}
