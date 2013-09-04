package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matzhouse/gosync/gosync"
	"launchpad.net/goamz/aws"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gosync"
	app.Usage = "CLI for S3"

	const concurrent = 10

	app.Commands = []cli.Command{
		{
			Name:        "sync",
			Usage:       "gosync sync SOURCE TARGET",
			Description: "Sync directories to / from S3 bucket.",
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fmt.Printf("S3 URL and local directory required.")
					os.Exit(1)
				}
				arg0 := c.Args()[0]
				arg1 := c.Args()[1]
				auth, err := aws.EnvAuth()
				if err != nil {
					panic(err)
				}

				fmt.Printf("Syncing %s with %s\n", arg0, arg1)

				sync := gosync.SyncPair{arg0, arg1, auth, concurrent}
				result := sync.Sync()
				if result == true {
					fmt.Printf("Syncing completed succesfully.")
				} else {
					fmt.Printf("Syncing failed.")
					os.Exit(1)
				}
			},
		},
	}
	app.Run(os.Args)
}
