package main

import (
	"flag"
	"fmt"
	"go/types"
	"os"
	"strings"
)

// AVRConfig for build or flash
type AVRConfig struct {
}

var version string = ""

func usage() {
	fmt.Fprintln(os.Stdout, "Go compiler for AVR microcontrollers family.")
	fmt.Fprintln(os.Stdout, "version: ", version)
	fmt.Fprintf(os.Stdout, "usage: %s command [-target=<target>] [-output=<output-file>] [-input=<input-file>]", os.Args[0])
	fmt.Fprintln(os.Stdout, "\ncommands :")
	fmt.Fprintln(os.Stdout, " [build]: compiling code and dependencies.")
	fmt.Fprintln(os.Stdout, " [flash]: compile and flash into microcontroller chip.")
	fmt.Fprintln(os.Stdout, " [help]: show help.")
	fmt.Fprintln(os.Stderr, "\nflags:")
	flag.PrintDefaults()
}

// Build go code into executable code
func Build(input, output, target string, config *AVRConfig) error {
	fmt.Fprintln(os.Stdout, "Compiling")
	return nil
}

// Flash go code into microcontroller device
func Flash(input, target, port string, config *AVRConfig) error {
	fmt.Fprintln(os.Stdout, "Flashing")
	return nil
}

// BuildErrorHandler ..
func BuildErrorHandler(err error) {
	if err != nil {
		switch err := err.(type) {
		case types.Error:
			{
				fmt.Fprintln(os.Stderr, err)
			}
		default:
			{
				fmt.Fprintln(os.Stderr, "error : ", err)
			}
		}
	}
}

func main() {
	output := flag.String("output", "", "output filename")
	input := flag.String("input", "", "input filename")
	target := flag.String("target", "", "microcontroller chips target")
	port := flag.String("port", "", "flashing port")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stdout, "No command-line arguments supplied.\n\n")
		usage()
		os.Exit(0)
	}

	command := os.Args[1]

	os.Setenv("CC", "clang -target="+*target)

	switch strings.ToLower(command) {
	case "build":
		{
			if *output == "" {
				fmt.Fprintln(os.Stdout, "no output filename specified.")
				usage()
				os.Exit(0)
			}

			target := *target
			err := Build(*input, *output, target, nil)
			BuildErrorHandler(err)
		}
	case "flash":
		{
			if *output != "" {
				fmt.Fprintln(os.Stderr, "output cannot be specified with flash command.")
				usage()
				os.Exit(0)
			}

			err := Flash(*input, *target, *port, nil)
			BuildErrorHandler(err)
		}
	}
}
