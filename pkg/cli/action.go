package cli

import (
	"log"

	"github.com/urfave/cli/v2"
)

type action struct {
	flags struct {
		verbose bool
		dryRun  bool
	}
}

func (a *action) RunHelp(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.verbose = c.Bool(flagVerboseName)
	a.flags.dryRun = c.Bool(flagDryRunName)

	a.log("Flags %+v\n", a.flags)
}

func (a *action) log(format string, v ...interface{}) {
	if a.flags.verbose {
		log.Printf(format, v...)
	}
}

func (a *action) add(c *cli.Context) error {
	a.getFlags(c)

	return nil
}
