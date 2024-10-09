package app

import (
	"github.com/refluxwhw/nas/cmd/nas-server/app/options"
	"github.com/refluxwhw/nas/pkg/conf"
	"github.com/refluxwhw/nas/pkg/db"
	"github.com/refluxwhw/nas/pkg/logger"
	"github.com/spf13/cobra"
)

func NewNasServerCommand() *cobra.Command {
	opt := options.NewNasServerOption()
	cmd := &cobra.Command{
		Use: "nas-server",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO:
		},
	}

	cmd.Flags().StringVarP(&opt.ConfigFile, "config", "f", opt.ConfigFile, "config file name")
	return cmd
}

func run(opt options.NasServerOption) {
	// 1. init config
	if err := conf.LoadConfig(opt.ConfigFile); err != nil {
		logger.Fatal(err)
		return
	}

	// 2. init check

	// 3. setup db conn
	if err := db.Setup(); err != nil {
		logger.Fatal(err)
		return
	}

	// 4. setup file manager

	// 5. setup http server

}
