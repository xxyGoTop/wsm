package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	app2 "github.com/xxyGoTop/wsm/internal/app"
	"github.com/xxyGoTop/wsm/internal/lib/daemon"
	"github.com/xxyGoTop/wsm/internal/lib/util"
)

func main() {
	app := cli.NewApp()
	app.Usage = "user server"
	app.Version = "0.1.0"
	app.Authors = []*cli.Author {
		{
			Name: "xxy",
			Email: "xxy@qq.com",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name: "start",
			Usage: "start server",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name: "daemon",
					Aliases: []string{"d"},
					Usage: "running in daemon mode",
				},
			},
			Action: func(c *cli.Context) error {
				return daemon.Start(app2.Serve, c.Bool("daemon"))
			},
		},
		{
			Name: "stop",
			Usage: "stop user server",
			Action: func(c *cli.Context) error {
				return daemon.Stop()
			},
		},
		{
			Name: "env",
			Usage: "print runtime environment",
			Action: func(c *cli.Context) error {
				util.PrintEnv()
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
