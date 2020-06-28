package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

// AppBetaTwo is storage for the sub command beta-two
type AppBetaTwo struct {
	*AppBeta        // embedded parent command storage
	Two      string // storage for flag --two
}

// AppBetaTwo is a method of AppBeta that returns a Cmder instance for the sub command
func (app *AppBeta) AppBetaTwo() cmder.Cmder {
	return &AppBetaTwo{AppBeta: app}
}

// Cmd returns a cobra.Command instance to be added to the parent command
func (app *AppBetaTwo) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "two",
		Short:        "Sub command beta-two",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	cmd.Flags().StringVarP(&app.Two, "two", "2", "", "command specific flag")
	return cmd
}

// RunE is a main routine of the sub command, returning an error
func (app *AppBetaTwo) RunE(cmd *cobra.Command, args []string) error {
	fmt.Println(cmd.Short, "executed")
	fmt.Println(time.Now().In(app.Location))
	app.Args = args
	app.Dump(app)
	return fmt.Errorf("%s error", cmd.Short)
}
