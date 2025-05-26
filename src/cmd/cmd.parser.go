package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	Method int

	MantisFlag int

	CliCommand struct {
		Program    string
		Method     Method
		TargetFile string
		Output     string
		Subset     int8
	}
)

const (
	Run Method = iota
	Build
	Help
)

func ParseInput() (ret CliCommand) {
	ret.Program = os.Args[0]
	if len(os.Args) == 1 {
		log.Fatalf("error: not enough arguments;")
	}

	switch os.Args[1] {
	case "run":
		ret.Method = Run
		ret.parseRun()
		break
	case "build":
		ret.Method = Build
		ret.parseBuild()
		break
	case "help":
		ret.Method = Help
		sb := strings.Builder{}
		sb.WriteString("\n\nMantis is a programming language created by Arthur for educational purposes.\n\n")
		sb.WriteString("Usage:\n")
		sb.WriteString("\tmnts <command> [arguments]\n")
		sb.WriteString("\nThe commands are:\n")
		sb.WriteString("\tbuild\t\tbuild a program from file. To pass the output name, use '-o <OUTPUT>', otherwise it'll use the file name as output\n")
		//sb.WriteString("\trun: run a mantis file\n")
		sb.WriteString("\thelp\t\tget help panel\n\n")
		fmt.Print(sb.String())
		os.Exit(0)
	default:
		log.Fatalf("error: unknown command '%s'", os.Args[1])
	}
	return ret
}

func (this *CliCommand) parseBuild() {
	if len(os.Args) <= 2 {
		log.Fatalf("error: not enough arguments, missing target file for build")
	}
	this.TargetFile = os.Args[2]
	this.parseFlag()
}

func (this *CliCommand) parseRun() {
	if len(os.Args) <= 2 {
		log.Fatalf("error: not enough arguments, missing target file for run")
	}
	this.TargetFile = os.Args[2]
	this.parseFlag()
}

var flags = map[string]MantisFlag{
	"-o": Output,
}

var options = map[string]MantisFlag{}

const (
	None MantisFlag = iota
	Output
)

func (this *CliCommand) parseFlag() {
	list := os.Args
	for i := 3; i < len(list); i++ {
		arg := list[i]
		switch flags[arg] {
		case Output:
			i++
			if i >= len(list) {
				log.Fatalf("error: not enough arguments, missing target file for run")
			}
			if this.Output != "" {
				log.Fatalf("error: duplicate flag, output must be unique")
			}
			this.Output = list[i]
			break
		case None:
			fallthrough
		default:
			this.parseOption(arg)
		}
	}
}

func (this *CliCommand) parseOption(option string) {
	split := strings.Split(option, "=")
	if len(split) != 2 {
		log.Fatalf("error: unknown flag '%s'", option)
	}

	switch options[split[0]] {
	case None:
		fallthrough
	default:
		log.Fatalf("error: unknown option '%s'", split[0])
	}
}
