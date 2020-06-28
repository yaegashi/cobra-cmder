package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// App is storage for app's root command
type App struct {
	TimeZone string         // Storage for common flag --tz
	Location *time.Location // Location converted from TimeZone
	Args     []string       // Args set by sub commands
}

// Cmd returns a cobra.Command instance to be added to app's root command
func (app *App) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "sample",
		Short:             "Sample CLI app for cobra-cmder",
		PersistentPreRunE: app.PersistentPreRunE,
		SilenceUsage:      true,
	}
	cmd.PersistentFlags().StringVarP(&app.TimeZone, "tz", "", "Local", "Time zone")
	return cmd
}

// PersistentPreRunE is a global initializer for this app
func (app *App) PersistentPreRunE(cmd *cobra.Command, args []string) error {
	loc, err := time.LoadLocation(app.TimeZone)
	if err != nil {
		return err
	}
	app.Location = loc
	return nil
}

// Dump dumps any object in JSON.
// This method is commonly available among all commands in this app.
func (app *App) Dump(obj interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(obj)
}
