package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

// AppAlphaOne is storage for the sub command alpha-one
type AppAlphaOne struct {
	*AppAlpha        // embedded parent command storage
	One       string // storage for flag --one
}

// AppAlphaOne is a method of AppAlpha that returns a Cmder instance for the sub command
func (app *AppAlpha) AppAlphaOne() cmder.Cmder {
	return &AppAlphaOne{AppAlpha: app}
}

// Cmd returns a cobra.Command instance to be added to the parent command
func (app *AppAlphaOne) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "one",
		Short:        "Sub command alpha-one",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	cmd.Flags().StringVarP(&app.One, "one", "1", "", "command specific flag")
	return cmd
}

// RunE is a main routine of the sub command, returning a nil
func (app *AppAlphaOne) RunE(cmd *cobra.Command, args []string) error {
	fmt.Println(cmd.Short, "executed")
	fmt.Println(time.Now().In(app.Location))
	app.Args = args
	app.Dump(app)
	return nil
}
