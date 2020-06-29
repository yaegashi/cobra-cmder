# cobra-cmder

[![go.dev](https://img.shields.io/badge/go.dev-reference-000000?logo=go)](https://pkg.go.dev/github.com/yaegashi/cobra-cmder)

## Introduction

This Go module contains `cmder` library package
which is a useful command builder for
[spf13/cobra](https://github.com/spf13/cobra).

It helps you to easily build a cobra.Command hierarchy
in the non-invasive and test-friendly way without any global variables or init().

## Basic usage

Define each command's sturct that implements `Cmder` interface:

```go
type App struct {
	Bool bool // storage for flag -b
}
type AppAlpha struct {
	*App          // storage for parent Cmder (embedded)
	String string // storage for flag -s
}
type AppAlphaOne struct {
	*AppAlpha     // storage for parent Cmder (embedded)
	Int       int // storage for flag -i
}
```

Define each Comder's `Cmd()` method that returns `*cobra.Command`:

```go
func (app *App) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "app",
	}
	cmd.PersistentFlags().BoolVarP(&app.Bool, "bool", "b", false, "Bool flag")
	return cmd
}

func (app *AppAlpha) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "alpha",
	}
	cmd.PersistentFlags().StringVarP(&app.String, "string", "s", "", "String flag")
	return cmd
}

func (app *AppAlphaOne) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "one",
		Run: app.Run,
	}
	cmd.Flags().IntVarP(&app.Int, "int", "i", 0, "Int flag")
	return cmd
}

func (app *AppAlphaOne) Run(cmd *cobra.Command, args []string) {
	fmt.Println(app.Bool, app.String, app.Int)
}
```

Associate Cmders each other by defining a method that returns a child Cmder:

```go
func (app *App) AppAlphaCmder() cmder.Cmder         { return &AppAlpha{App: app} }
func (app *AppAlpha) AppAlphaOneCmder() cmder.Cmder { return &AppAlphaOne{AppAlpha: app} }
```

Call `cmder.Cmd()` to collect and associate all `cobra.Command` instances:

```go
func main() {
	app := &App{}
	cmd := cmder.Cmd(app)
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
```

Visit https://play.golang.org/p/zw4arxJfUkt to see and test the complete source code.

## Unit test

You can easily perform the Go standard unit tests on CLI apps with cobra-cmder.

See https://play.golang.org/p/tijvjDzmwqW for another example.

## Examples

See the sample CLI app in [cmd/sample](cmd/sample) for more comprehensive example.

Usage in real applications:

- https://github.com/yaegashi/contest.go
- https://github.com/yaegashi/azbill
