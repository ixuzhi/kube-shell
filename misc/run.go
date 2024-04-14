package main

import (
	"encoding/json"
	"io/ioutil"
	//"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // kubectl auth providers.
	"k8s.io/kubectl/pkg/cmd"
)

type Option struct {
	Help string `json:"help"`
	Name string `json:"name"`
}

type Command struct {
	Command     string             `json:"command"`
	Help        string             `json:"help"`
	Subcommands map[string]Command `json:"subcommands"`
	Args        []string           `json:"args"`
	Options     map[string]Option  `json:"options"`
}

var (
	kubectl_cmd    = Command{}
	globalFlagsSet = false
)

func main() {
	command := cmd.NewDefaultKubectlCommand()
	buildCmdMap(&kubectl_cmd, command)

	cli := make(map[string]Command)
	cli["kubectl"] = kubectl_cmd
	b, err := json.Marshal(&cli)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("cli.json", b, 0644)
}

func buildCmdMap(rootCmd *Command, cobraCmd *cobra.Command) {

	rootCmd.Command = cobraCmd.Name()
	rootCmd.Help = cobraCmd.Short
	rootCmd.Subcommands = make(map[string]Command)
	rootCmd.Options = make(map[string]Option, 0)
	rootCmd.Args = make([]string, 0)

	for _, subCobraCmd := range cobraCmd.Commands() {
		subCmd := Command{}
		buildCmdMap(&subCmd, subCobraCmd)
		rootCmd.Subcommands[subCobraCmd.Name()] = subCmd
	}

	validArgs := cobraCmd.ValidArgs
	if len(validArgs) > 0 {
		for _, cobraArg := range validArgs {
			rootCmd.Args = append(rootCmd.Args, cobraArg)
		}
	}

	if cobraCmd.HasFlags() {
		flagVisitFuncton := func(f *pflag.Flag) {
			option := Option{}
			option.Name = "--" + f.Name
			option.Help = f.Usage
			rootCmd.Options[option.Name] = option
		}
		flagSet := cobraCmd.Flags()
		flagSet.VisitAll(flagVisitFuncton)
	}

	if globalFlagsSet != true {
		inheritedFlagSet := cobraCmd.InheritedFlags()
		inheritedFlagVisitFuncton := func(f *pflag.Flag) {
			option := Option{}
			option.Name = "--" + f.Name
			option.Help = f.Usage
			kubectl_cmd.Options[option.Name] = option
		}
		inheritedFlagSet.VisitAll(inheritedFlagVisitFuncton)
		globalFlagsSet = true
	}
}
