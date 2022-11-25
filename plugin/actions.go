package plugin

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

type Action = func() error

func CommandAction(name string, arg ...string) Action {
	return func() error { return exec.Command(name, arg...).Run() }
}

func ConfirmAndThenActions(prompt string, actions ...Action) Action {
	return func() error {
		for {
			p := promptui.Prompt{
				Label:     prompt,
				IsConfirm: true,
			}

			s, err := p.Run()

			if err != nil && err != promptui.ErrAbort {
				fmt.Printf("Prompt failed %v\n", err)
				return err
			}

			if lower := strings.ToLower(s); lower == "y" || lower == "" {
				for _, action := range actions {
					err := action()
					if err != nil {
						return err
					}
				}
			} else if lower == "n" {
				break
			}
		}
		return nil
	}
}
