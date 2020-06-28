# cobra-cmder

[![go.dev](https://img.shields.io/badge/go.dev-reference-000000?logo=go)](https://pkg.go.dev/github.com/yaegashi/cobra-cmder)

## Introduction

This Go module contains `cmder` library package
which is a useful command builder for
[spf13/cobra](https://github.com/spf13/cobra).

It helps you to easily build a cobra.Command hierarchy
in the non-invasive and test-friendly way without any global variables or init().

## Basic usage

Define a sturct (later implements `cmder.Cmder` interface) for each command:

```go
type App struct {
	Root bool // storage for flag -R
}
type AppAlpha struct {
	*App         // storage for parent Cmder (embedded)
	Alpha string // storage for flag -A
}
type AppAlphaOne struct {
	*AppAlpha     // storage for parent Cmder (embedded)
	One       int // storage for flag -1
}
```

Associate Cmders by defining methods returning child `Cmder` instances:

```go
func (app *App) AppAlpha() cmder.Cmder         { return &AppAlpha{App: app} }
func (app *AppAlpha) AppAlphaOne() cmder.Cmder { return &AppAlphaOne{AppAlpha: app} }
```

Define each Comder's Cmd() method returning `*cobra.Command`:

```go

func (app *App) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "app",
	}
	cmd.PersistentFlags().BoolVarP(&app.Root, "root", "R", false, "Root flag")
	return cmd
}

func (app *AppAlpha) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "alpha",
	}
	cmd.PersistentFlags().StringVarP(&app.Alpha, "alpha", "A", "", "Alpha flag")
	return cmd
}

func (app *AppAlphaOne) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "one",
		Run: func(cmd *cobra.Command, args []string) { fmt.Println(app.Root, app.Alpha, app.One) },
	}
	cmd.Flags().IntVarP(&app.One, "one", "1", 0, "One flag")
	return cmd
}
```

Call `cmder.Cmd()` to collect and associate all `*cobra.Command` instances:

```go
func main() {
	app := &App{}
	cmd := cmder.Cmd(app)
	cmd.SetArgs([]string{"alpha", "one", "-R", "-A", "abc", "-1", "123"})
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
```

Visit https://play.golang.org/p/fXF8z8vMyo8 to see and test the complete source code.

## Examples

See the sample CLI app in [cmd/sample](cmd/sample) for more comprehensive example.

Usage in real applications:

- https://github.com/yaegashi/contest.go
- https://github.com/yaegashi/azbill
