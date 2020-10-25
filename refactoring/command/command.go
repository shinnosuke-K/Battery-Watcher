package command

import (
	"io"
	"io/ioutil"
	"os/exec"
)

type Command struct {
	cmdChains []*exec.Cmd
	Out       io.ReadCloser
}

func New() *Command {
	return &Command{}
}

func (c *Command) Set() {
	c.cmdChains = []*exec.Cmd{
		exec.Command("ioreg", "-l"),
		exec.Command("grep", "-v", "Apple"),
		exec.Command("grep", "-v", "BatteryData"),
		exec.Command("grep", "-e", "MaxCapacity", "-e", "DesignCapacity", "-e", "CurrentCapacity"),
	}
}

func (c *Command) Pipe() error {
	var err error
	for i := 0; i < len(c.cmdChains)-1; i++ {
		thisCmd := c.cmdChains[i]
		nextCmd := c.cmdChains[i+1]

		nextCmd.Stdin, err = thisCmd.StdoutPipe()
		if err != nil {
			return err
		}
	}

	c.Out, err = c.cmdChains[len(c.cmdChains)-1].StdoutPipe()
	if err != nil {
		return err
	}
	return nil
}

func (c *Command) Start() error {
	for _, cmd := range c.cmdChains {
		if err := cmd.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Command) GetStrOut() (string, error) {
	byteOut, err := ioutil.ReadAll(c.Out)
	if err != nil {
		return "", err
	}
	return string(byteOut), err
}

func (c *Command) Done() error {
	for _, cmd := range c.cmdChains {
		if err := cmd.Wait(); err != nil {
			return err
		}
	}
	return nil
}
