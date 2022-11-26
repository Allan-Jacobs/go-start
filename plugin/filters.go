package plugin

import (
	"os/exec"
)

type Predicate func() bool

func (p Predicate) Or(predicate Predicate) Predicate {
	return func() bool {
		return p() || predicate()
	}
}

func (p Predicate) And(predicate Predicate) Predicate {
	return func() bool {
		return p() && predicate()
	}
}

func True() Predicate {
	return func() bool { return true }
}

func False() Predicate {
	return func() bool { return false }
}

func HasExec(name string) Predicate {
	return func() bool {
		_, err := exec.LookPath(name)
		return err == nil
	}
}
