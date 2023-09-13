package runner

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/go-template-org/go-template-app/internal/version"
)

const (
	OsExistCode int = 1
)

var PanicOnExit bool = false // Set to true to tell Exit() to Panic rather than call os.Exit() - should ONLY be used for testing

func Exit(code int) {
	if PanicOnExit {
		panic(fmt.Sprintf("PanicOnExit is true, code=%d", code))
	}

	os.Exit(code)
}

func GetParameters() (bool, bool) {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	flagSet.SetOutput(os.Stderr)

	var dbg bool
	var verbose bool

	// add flags
	flagSet.BoolVar(&dbg, "debug", false, "Log Debug")
	flagSet.BoolVar(&verbose, "verbose", false, "Show Verbose Logging")

	// Parse the flags
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		flagSet.Usage()
		Exit(OsExistCode)
	}

	if verbose {
		fmt.Println(version.GetVersion())
	}

	if dbg {
		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			fmt.Println(buildInfo.String())
		}
	}

	return dbg, verbose
}

func Run() int {
	ctx := context.Background()
	debugFlag, verboseFlag := GetParameters()

	// ToDo:

	return 0
}
