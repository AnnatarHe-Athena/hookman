package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AnnatarHe-Athena/hookman/jobs"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func setupLogFile() {
	filename := fmt.Sprintf("/var/log/hookman/log-%s.log", time.Now().Format(time.RFC3339))
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.Panicln(err)
	}
	logrus.SetOutput(f)
}

func main() {
	setupLogFile()
	app := &cli.App{
		Name:  "hookman",
		Usage: "athena cli tool",
		Commands: []*cli.Command{
			{
				Name:  "updateWeiboUserInfo",
				Usage: "update weibo users info",
				Action: func(c *cli.Context) error {
					return jobs.UpdateUsersFor1week()
				},
			},
			{
				Name:  "tagAll",
				Usage: "set tags to all images",
				Action: func(c *cli.Context) error {
					return jobs.TagAll()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Errorln(err)
	}
}
