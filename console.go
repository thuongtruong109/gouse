package gouse

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func Cmd(defaultCmd string, windowsCmd ...string) {
	var cmd *exec.Cmd

	splitCmd := Split(defaultCmd, " ")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c")
		if len(windowsCmd) > 0 {
			cmd.Args = append(cmd.Args, windowsCmd...)
		} else {
			cmd.Args = append(cmd.Args, splitCmd...)
		}
	} else {
		cmd = exec.Command("sh", "-c")
		cmd.Args = append(cmd.Args, splitCmd...)
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Cls() {
	Cmd("clear", "cls")
}

func OutputColor[T any](color string, output T) {
	fmt.Printf("%s%v%s", color, output, DEFAULT_CONSOLE)
}

func OutputSuccess(output string) {
	println(GREEN_CONSOLE + output + DEFAULT_CONSOLE)
}

func OutputError(output string, err string) {
	println(RED_CONSOLE + err + DEFAULT_CONSOLE)
}

func Banner(font IFontDouble, s string) {
	split := Split(Uppers(s), "")
	for i := range 3 {
		for _, v := range split {
			fmt.Print(font[v][i])
		}
		fmt.Print("\n")
	}
}

func Select(label string, options []string) (string, error) {
	var selected string
	prompt := &survey.Select{
		Message: label,
		Options: options,
	}
	err := survey.AskOne(prompt, &selected)
	return selected, err
}

func Confirm(label string) (bool, error) {
	var confirm bool
	prompt := &survey.Confirm{
		Message: label,
	}
	err := survey.AskOne(prompt, &confirm)
	return confirm, err
}

/* Help menu */

type IHelpOptions struct {
	Opt    string
	Desc   string
	Val    any
	Action any
}

type IHelpActionStr struct {
	Action func(string)
}

type IHelpActionInt struct {
	Action func(int)
}

type IHelpActionBool struct {
	Action func(bool)
}

func autoDetectAction(f any) any {
	var action any

	switch value := f.(type) {
	case func(string):
		action = IHelpActionStr{Action: value}
	case func(int):
		action = IHelpActionInt{Action: value}
	case func(bool):
		action = IHelpActionBool{Action: value}
	default:
		println(DESC_CONSOLE)
		return nil
	}

	return action
}

// func Help(name string, options []*HelpOptions) {
// 	for _, option := range options {
// 		switch autoDetectAction(option.Action).(type) {
// 		case StringAction:
// 			flag.String(option.Opt, option.Val.(string), option.Desc)
// 		case IntAction:
// 			flag.Int(option.Opt, option.Val.(int), option.Desc)
// 		case BoolAction:
// 			flag.Bool(option.Opt, option.Val.(bool), option.Desc)
// 		default:
// 			fmt.Println("Unsupported action type")
// 		}
// 	}

// 	flag.Parse()

// 	if flag.NArg() > 0 && (flag.Arg(0) == "-h" || flag.Arg(0) == "--help") {
// 		fmt.Printf("Usage: %s [options]\n", name)
// 		fmt.Println("Options:")
// 		for _, option := range options {
// 			fmt.Printf("  -%s\n", option.Opt)
// 			fmt.Printf("        %s\n", option.Desc)
// 		}
// 	}

// 	// Run the associated actions for all set flags
// 	for _, option := range options {
// 		flag.Visit(func(f *flag.Flag) {
// 			if f.Name == option.Opt {
// 				switch a := autoDetectAction(option.Action).(type) {
// 				case StringAction:
// 					a.Action(f.Value.String())
// 				case IntAction:
// 					val, _ := strconv.Atoi(f.Value.String())
// 					a.Action(val)
// 				case BoolAction:
// 					val, _ := strconv.ParseBool(f.Value.String())
// 					a.Action(val)
// 				default:
// 					fmt.Println("Unsupported action type")
// 				}
// 			}
// 		})
// 	}
// }

func Help(name string, options []*IHelpOptions) {
	parseFlags(options)

	PrintUsage(name, options)

	runActions(options)
}

func parseFlags(options []*IHelpOptions) {
	for _, option := range options {
		switch autoDetectAction(option.Action).(type) {
		case IHelpActionStr:
			flag.String(option.Opt, option.Val.(string), option.Desc)
		case IHelpActionInt:
			flag.Int(option.Opt, option.Val.(int), option.Desc)
		case IHelpActionBool:
			flag.Bool(option.Opt, option.Val.(bool), option.Desc)
		default:
			println(DESC_CONSOLE)
		}
	}

	flag.Parse()
}

func PrintUsage(name string, options []*IHelpOptions) {
	if flag.NArg() > 0 && (flag.Arg(0) == "-h" || flag.Arg(0) == "--help") {
		fmt.Printf("Usage: %s [options]\n", name)
		fmt.Println("Options:")
		for _, option := range options {
			fmt.Printf("  -%s\n", option.Opt)
			fmt.Printf("        %s\n", option.Desc)
		}
	}
}

func runActions(options []*IHelpOptions) {
	flag.Visit(func(f *flag.Flag) {
		for _, option := range options {
			if f.Name == option.Opt {
				switch a := autoDetectAction(option.Action).(type) {
				case IHelpActionStr:
					a.Action(f.Value.String())
				case IHelpActionInt:
					val, _ := strconv.Atoi(f.Value.String())
					a.Action(val)
				case IHelpActionBool:
					val, _ := strconv.ParseBool(f.Value.String())
					a.Action(val)
				default:
					println(DESC_CONSOLE)
				}
			}
		}
	})
}
