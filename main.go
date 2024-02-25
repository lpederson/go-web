package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go-web/pkg"
)

var port uint32

func main() {
	rootCmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, _ []string) error {

			server := pkg.NewServer(
				cmd.Context(),
				port,
			)
			return server.Run()
		},
		Short: `Runs a basic hello world server`,
	}

	rootCmd.Flags().AddFlagSet(getFlags())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func getFlags() *pflag.FlagSet {
	flags := pflag.NewFlagSet("go-web", pflag.ExitOnError)
	flags.Uint32Var(&port, "port", 8080, "")
	return flags
}
