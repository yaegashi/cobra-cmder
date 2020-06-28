package main

import (
	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

// AppBeta is storage for the sub command beta
type AppBeta struct {
	*App        // embedded parent command storage
	Beta string // storage for flag --beta
}

// AppBeta is a method of App that returns a Cmder instance for the sub command
func (app *App) AppBeta() cmder.Cmder {
	return &AppBeta{App: app}
}

// Cmd returns a cobra.Command instance to be added to the parent command
func (app *AppBeta) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "beta",
		Short:        "Sub command beta",
		SilenceUsage: true,
	}
	cmd.PersistentFlags().StringVarP(&app.Beta, "beta", "b", "", "command specific flag")
	return cmd
}
