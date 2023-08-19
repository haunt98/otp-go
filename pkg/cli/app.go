package cli

import "github.com/urfave/cli/v2"

const (
	name  = "gotp"
	usage = "handle otp"
)

type App struct {
	cliApp *cli.App
}
