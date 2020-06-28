package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

// AppAlphaTwo is storage for the sub command alpha-two
type AppAlphaTwo struct {
	*AppAlpha        // embedded parent command storage
	Two       string // storage for flag --two
}

// AppAlphaTwo is a method of AppAlpha that returns a Cmder instance for the sub command
func (app *AppAlpha) AppAlphaTwo() cmder.Cmder {
	return &AppAlphaTwo{AppAlpha: app}
}

// Cmd returns a cobra.Command instance to be added to the parent command
func (app *AppAlphaTwo) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "two",
		Short:        "Sub command alpha-two",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	cmd.Flags().StringVarP(&app.Two, "two", "2", "", "command specific flag")
	return cmd
}

// RunE is a main routine of the sub command, returning an error
func (app *AppAlphaTwo) RunE(cmd *cobra.Command, args []string) error {
	fmt.Println(cmd.Short, "executed")
	fmt.Println(time.Now().In(app.Location))
	app.Args = args
	app.Dump(app)
	return fmt.Errorf("%s error", cmd.Short)
}
