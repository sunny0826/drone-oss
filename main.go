package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sunny0826/drone-oss/oss"
	"github.com/urfave/cli"
	"os"
)

var (
	version = "unknown"
)

func main() {
	// Load env-file if it exists first
	if env := os.Getenv("PLUGIN_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := cli.NewApp()
	app.Name = "oss plugin"
	app.Usage = "oss plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "oss.dist",
			Usage:  "dist package path",
			EnvVar: "PLUGIN_DIST",
		},
		cli.StringFlag{
			Name:   "oss.path",
			Usage:  "oss package path",
			EnvVar: "PLUGIN_PATH",
		},
		cli.StringFlag{
			Name:   "oss.endpoint",
			Usage:  "oss endpoint",
			EnvVar: "PLUGIN_ENDPOINT",
		},
		cli.StringFlag{
			Name:   "access.key",
			Usage:  "AccessKeyID",
			EnvVar: "PLUGIN_ACCESS_KEY_ID",
		},
		cli.StringFlag{
			Name:   "access.secret",
			Usage:  "AccessKeySecret",
			EnvVar: "PLUGIN_ACCESS_KEY_SECRET",
		},
		cli.StringFlag{
			Name:   "modname",
			Usage:  "git module name",
			EnvVar: "PLUGIN_MODNAME",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := oss.Plugin{
		Config: oss.Config{
			Dist: c.String("oss.dist"),
			Path: c.String("oss.path"),
			EndPoint: c.String("oss.endpoint"),
			AccessKeyID: c.String("access.key"),
			AccessKeySecret: c.String("access.secret"),
			ModName: c.String("modname"),
		},
	}

	return plugin.Exec()
}
