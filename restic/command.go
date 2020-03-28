package restic

import (
	"context"
	"io"
	"os"
	"os/exec"

	"github.com/go-logr/logr"
)

// CommandOptions contains options for the command struct.
type CommandOptions struct {
	Path   string    // path where the command is to be executed
	StdIn  io.Reader // set the StdIn for the command
	StdOut io.Writer // set the StdOut for the command
	StdErr io.Writer // set StdErr for the command
	Args   []string
}

// Command can handle a given command.
type Command struct {
	options    CommandOptions
	FatalError error
	Errors     []error
	cmdLogger  logr.Logger
	ctx        context.Context
}

// NewCommand returns a new command
func NewCommand(ctx context.Context, log logr.Logger, commandOptions CommandOptions) *Command {
	return &Command{
		options:   commandOptions,
		Errors:    []error{},
		cmdLogger: log.WithName("command"),
		ctx:       ctx,
	}
}

// Run will run the currently configured command
func (c *Command) Run() {

	cmd := exec.CommandContext(c.ctx, c.options.Path, c.options.Args...)
	cmd.Env = os.Environ()

	if c.options.StdIn != nil {
		cmd.Stdin = c.options.StdIn
	}

	if c.options.StdOut != nil {
		cmd.Stdout = c.options.StdOut
	}

	if c.options.StdErr != nil {
		cmd.Stderr = c.options.StdErr
	}

	err := cmd.Start()
	if err != nil {
		c.FatalError = err
		return
	}

	err = cmd.Wait()
	if err != nil {
		c.FatalError = err
		return
	}
}
