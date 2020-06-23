// Package cmder contains a command builder library for spf13/cobra.
//
// It helps you to easily build a cobra.Command hierarchy
// in the non-invasive and test-friendly way without any global variables.
package cmder

import (
	"reflect"

	"github.com/spf13/cobra"
)

// Cmder is an interface for objects that return cobra.Command
type Cmder interface {
	Cmd() *cobra.Command
}

// Cmd traverses a Cmder hierarchy and returns a configured cobra.Command.
// It recursively calls all methods that return a Cmder
// to collect and associate cobra.Command instances.
func Cmd(cmder Cmder) *cobra.Command {
	return recCmd(cmder, map[string]bool{})
}

// outTypes is a constant for the array of function output types
var outTypes = []reflect.Type{reflect.TypeOf((*Cmder)(nil)).Elem()}

// recCmd is an actual worker function to visit and collect Cmder instances.
// methodMap is for bookkeeping already visted method names.
func recCmd(cmder Cmder, methodMap map[string]bool) *cobra.Command {
	cmd := cmder.Cmd()
	inV := reflect.ValueOf(cmder)
	inT := reflect.TypeOf(cmder)
	funcT := reflect.FuncOf([]reflect.Type{inT}, outTypes, false)
	methodList := []reflect.Method{}
	for i := 0; i < inT.NumMethod(); i++ {
		m := inT.Method(i)
		if m.Func.Type() != funcT {
			continue
		}
		if methodMap[m.Name] {
			continue
		}
		methodList = append(methodList, m)
	}
	for _, m := range methodList {
		methodMap[m.Name] = true
	}
	for _, m := range methodList {
		subCmder := m.Func.Call([]reflect.Value{inV})[0].Interface().(Cmder)
		subCmd := recCmd(subCmder, methodMap)
		cmd.AddCommand(subCmd)
	}
	for _, m := range methodList {
		methodMap[m.Name] = false
	}
	return cmd
}
