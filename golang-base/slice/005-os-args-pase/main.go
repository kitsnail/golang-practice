package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	debugFlag   = "-D"
	logfileFlag = "--logfile"
	choiceFlag  = "-s"
	helpFlag    = "--help"
	versionFlag = "--version"
	flagPrefix  = "-"
)

func main() {
	var globalFlags []string
	args := os.Args

	for i, arg := range args {
		fmt.Println(i, arg)
	}
	slurpGlobalFlags(&args, &globalFlags)
	fmt.Println("main.args:", args)
	fmt.Println("main.globalFlags:", globalFlags)

}

func slurpGlobalFlags(args *[]string, globalFlags *[]string) {
	slurpNextValue := false
	commandIndex := 0

	fmt.Println("slurpGlobalFlags args:", args)

	for i, arg := range *args {
		if slurpNextValue {
			commandIndex = i + 1
			slurpNextValue = false
		} else if arg == versionFlag || arg == helpFlag || !looksLikeFlag(arg) {
			break
		} else {
			commandIndex = i + 1
			if arg == debugFlag || arg == choiceFlag {
				slurpNextValue = true
			}
		}
	}

	fmt.Println("commandIndex:", commandIndex)
	if commandIndex > 0 {
		aa := *args
		*globalFlags = aa[0:commandIndex]
		*args = aa[commandIndex:]
	}
}

func looksLikeFlag(value string) bool {
	return strings.HasPrefix(value, flagPrefix)
}
