package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
)

type Command struct {
	cmdChains []*exec.Cmd
	Out       io.ReadCloser
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

func (c *Command) Done() error {
	for _, cmd := range c.cmdChains {
		if err := cmd.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	cmd := new(Command)
	cmd.Set()

	if err := cmd.Pipe(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	a, _ := ioutil.ReadAll(cmd.Out)
	fmt.Println(string(a))

	if err := cmd.Done(); err != nil {
		log.Fatal(err)
	}
}
