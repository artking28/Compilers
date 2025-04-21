package main

import (
	"compilers/langs/Mantis/lexer"
	"log"
	"os"
	"strconv"
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
	"-s": Subset,
}

var options = map[string]MantisFlag{
	"--subset": Subset,
}

const (
	None MantisFlag = iota
	Output
	Subset
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
		case Subset:
			i++
			if i >= len(list) {
				log.Fatalf("error: not enough arguments, missing subset code")
			}
			parseInt, err := strconv.ParseInt(list[i], 0, 8)
			if err != nil {
				log.Fatalf("error: not valid argument, subset must be a number argument")
			}
			if this.Subset != 0 {
				log.Fatalf("error: duplicate flag, subset must be unique")
			}
			if lexer.Subsets[int(parseInt)] == false {
				log.Fatalf("error: invalid subset value, subset must be between 0 and %d", lexer.SUBSET_MAX)
			}
			this.Subset = int8(parseInt)
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
	case Subset:
		if split[1] == "" {
			log.Fatalf("error: not valid argument, subset code must be a number argument")
		}
		parseInt, err := strconv.ParseInt(split[1], 0, 8)
		if err != nil {
			log.Fatalf("error: not valid argument, subset code must be a number argument")
		}
		if this.Subset != 0 {
			log.Fatalf("error: duplicate flag, subset code must be unique")
		}
		if lexer.Subsets[int(parseInt)] == false {
			log.Fatalf("error: invalid subset value, subset code must be between 0 and %d", lexer.SUBSET_MAX)
		}
		this.Subset = int8(parseInt)
		break
	case None:
		fallthrough
	default:
		log.Fatalf("error: unknown option '%s'", split[0])
	}
}
