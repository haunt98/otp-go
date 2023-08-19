package cli

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/make-go-great/color-go"
)

const (
	name  = "gotp"
	usage = "handle otp"

	commandAddName  = "add"
	commandAddUsage = "add new otp"

	flagVerboseName  = "verbose"
	flagVerboseUsage = "show what is going on"

	flagDryRunName  = "dry-run"
	flagDryRunUsage = "demo run without actually changing anything"
)

var flagVerboseAliases = []string{"v"}

type App struct {
	cliApp *cli.App
}

func NewApp() *App {
	a := &action{}

	cliApp := &cli.App{
		Name:  name,
		Usage: usage,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagVerboseName,
				Aliases: flagVerboseAliases,
				Usage:   flagVerboseUsage,
			},
			&cli.BoolFlag{
				Name:  flagDryRunName,
				Usage: flagDryRunUsage,
			},
		},
		Commands: []*cli.Command{
			{
				Name:   commandAddName,
				Usage:  commandAddUsage,
				Action: a.add,
			},
		},
		Action: a.RunHelp,
	}

	return &App{
		cliApp: cliApp,
	}
}

func (a *App) Run() {
	if err := a.cliApp.Run(os.Args); err != nil {
		color.PrintAppError(name, err.Error())
	}
}
