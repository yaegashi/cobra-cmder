package main

import (
	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

// AppAlpha is storage for the sub command alpha
type AppAlpha struct {
	*App         // embedded parent command storage
	Alpha string // storage for flag --alpha
}

// AppAlpha is a method of App that returns a Cmder instance for the sub command
func (app *App) AppAlpha() cmder.Cmder {
	return &AppAlpha{App: app}
}

// Cmd returns a cobra.Command instance to be added to the parent command
func (app *AppAlpha) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "alpha",
		Short:        "Sub command alpha",
		SilenceUsage: true,
	}
	cmd.PersistentFlags().StringVarP(&app.Alpha, "alpha", "a", "", "command specific flag")
	return cmd
}
