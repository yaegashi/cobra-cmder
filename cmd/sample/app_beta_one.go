package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

// AppBetaOne is storage for the sub command beta-one
type AppBetaOne struct {
	*AppBeta        // embedded parent command storage
	One      string // storage for flag --one
}

// AppBetaOne is a method of AppBeta that returns a Cmder instance for the sub command
func (app *AppBeta) AppBetaOne() cmder.Cmder {
	return &AppBetaOne{AppBeta: app}
}

// Cmd returns a cobra.Command instance to be added to the parent command
func (app *AppBetaOne) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "one",
		Short:        "Sub command beta-one",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	cmd.Flags().StringVarP(&app.One, "one", "1", "", "command specific flag")
	return cmd
}

// RunE is a main routine of the sub command, returning a nil
func (app *AppBetaOne) RunE(cmd *cobra.Command, args []string) error {
	fmt.Println(cmd.Short, "executed")
	fmt.Println(time.Now().In(app.Location))
	app.Args = args
	app.Dump(app)
	return nil
}
