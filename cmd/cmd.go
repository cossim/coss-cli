package cmd

import (
	"github.com/urfave/cli/v2"
)

const VERSION = "v1.0.4"

var App = &cli.App{
	Name:    "coss-cli",
	Usage:   "coss-cli is a command line tool for coss",
	Version: VERSION,
	Commands: []*cli.Command{
		{
			Name:  "start",
			Usage: "start coss",
			Flags: []cli.Flag{
				//&cli.StringFlag{
				//	Name:  "path",
				//	Value: "./",
				//	Usage: "config path",
				//},
				&cli.BoolFlag{
					Name:    "direct",
					Value:   true,
					Usage:   "true or false: --direct=false or -d=false",
					Aliases: []string{"d"},
				},
			},
			Action: start,
		},
		{
			Name:  "config",
			Usage: "init consul config",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Value:   "",
					Aliases: []string{"p"},
					Usage:   "config path (multiple paths separated by commas)",
				},
				&cli.StringFlag{
					Name:    "namespace",
					Aliases: []string{"n"},
					Value:   "default",
					Usage:   "config namespace",
				},
				&cli.StringFlag{
					Name:  "host",
					Value: "http://127.0.0.1:8500",
					Usage: "consul host",
				},
				&cli.StringFlag{
					Name:  "token",
					Value: "",
					Usage: "consul acl token",
				},
				&cli.BoolFlag{
					Name:  "ssl",
					Value: false,
					Usage: "consul enable ssl",
				},
			},
			Action: initConfig,
		},
		{
			Name:  "gen",
			Usage: "gen coss config",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "direct",
					Value:   true,
					Usage:   "true or false: --direct=false or -d=false",
					Aliases: []string{"d"},
				},
				&cli.StringFlag{
					Name:  "path",
					Value: "./",
					Usage: "config path",
				},
				&cli.StringFlag{
					Name:  "domain",
					Value: "127.0.0.1",
					Usage: "your domain name",
				},
				&cli.BoolFlag{
					Name:  "ssl",
					Value: false,
					Usage: "consul enable ssl",
				},
			},
			Action: gen,
		},
		{
			Name:  "route",
			Usage: "init gateway route",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "direct",
					Value:   true,
					Usage:   "true or false: --direct=false or -d=false",
					Aliases: []string{"d"},
				},
				&cli.StringFlag{
					Name:  "key",
					Value: "edd1c9f034335f136f87ad84b625c8f1",
					Usage: "apisix api key ",
				},
				&cli.StringFlag{
					Name:  "host",
					Value: "http://127.0.0.1:9180",
					Usage: "apisix host",
				},
				&cli.StringFlag{
					Name:  "domain",
					Value: "",
					Usage: "route domain name",
				},
				&cli.StringFlag{
					Name:  "livekit",
					Value: "",
					Usage: "livekit domain name",
				},
				&cli.StringFlag{
					Name:  "route-host",
					Value: "",
					Usage: "route host",
				},
			},
			Action: initRoute,
		},
		{
			Name:  "ssl",
			Usage: "init consul config",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "cert",
					Value: "./",
					Usage: "cert file path",
				},
				&cli.StringFlag{
					Name:  "private_key",
					Value: "./",
					Usage: "key file path",
				},
				&cli.StringFlag{
					Name:  "domain",
					Value: "",
					Usage: "your domain name",
				},
				&cli.IntFlag{
					Name:  "num",
					Value: 1,
					Usage: "ssl key num",
				},
				&cli.StringFlag{
					Name:  "key",
					Value: "edd1c9f034335f136f87ad84b625c8f1",
					Usage: "apisix api key ",
				},
				&cli.StringFlag{
					Name:  "host",
					Value: "http://127.0.0.1:9180",
					Usage: "apisix host",
				},
			},
			Action: uploadSSL,
		},
	},
}
