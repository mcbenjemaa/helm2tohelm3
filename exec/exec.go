package exec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)



type Command struct {
	Command string
	Dryrun bool
	Execute bool
}

var dryRunFlag = "--dry-run"

func (c *Command) Exec() error {
	if !c.Execute {
		return fmt.Errorf("Could not execute script!")
	}
	mydir, err := os.Getwd()
	if err != nil {
		return err
	}

	c.DryRunMode(c.Dryrun)
	cmd := exec.Command("bash", "-c", c.Command)
	cmd.Dir = mydir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

   return err
}

func (c *Command) ExecAndGetOutput() (string, error)  {
	out, err := exec.Command("bash", "-c", c.Command).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}


func (c *Command) DryRunMode(b bool) {
	if b {
		c.Dryrun = true
		c. Command = c.Command+" "+dryRunFlag
	} else {
		c.Dryrun = false
		c.Command = strings.Replace(c.Command, dryRunFlag, "", -1)
	}
}