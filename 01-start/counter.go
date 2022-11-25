package counter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type option func(*counter) error
type counter struct {
	input  io.Reader
	output io.Writer
}

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return c, err
		}
	}
	return c, nil
}

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func (c counter) Lines() {
	lines := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		lines++
	}
	fmt.Fprintf(c.output, "%d lines\n", lines)
}

func Lines() {
	c, err := NewCounter()
	// If the only person who will ever call Lines can’t do anything
	// useful with the error except log it and crash, we may as well do that directly:
	//
	// There is no useful information we can give the user,
	// because the user can’t fix our pro‐ gram, and this is
	// definitely an internal program bug. That’s exactly what panic
	// is for: reporting unrecoverable internal program bugs.
	if err != nil {
		panic(err)
	}
	c.Lines()
}
