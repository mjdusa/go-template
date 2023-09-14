package version

import (
	"fmt"
	"os"
	"runtime"
)

var (
	// AppVersion contains the current version in SemVer format.
	AppVersion string //nolint:gochecknoglobals  // only used for usage

	// BuildTime is the compiled build time.
	BuildTime string //nolint:gochecknoglobals  // only used for usage

	// GitBranch is the name of the branch referenced by HEAD.
	GitBranch string //nolint:gochecknoglobals  // only used for usage

	// GitCommit contains the hash of the latest commit on Branch.
	GitCommit string //nolint:gochecknoglobals  // only used for usage

	// GoVersion contains the the version of the go that performed the build.
	GoVersion string = runtime.Version() //nolint:gochecknoglobals  // only used for usage

	// OsArch contains the OS and Arch used to build the binary
	OsArch string = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH) //nolint:gochecknoglobals  // only used for usage
)

func GetVersion() string {
	msg := fmt.Sprintf("%s version: [%s]\n", os.Args[0], AppVersion)
	msg += fmt.Sprintf("- Build Time:      [%s]\n", BuildTime)
	msg += fmt.Sprintf("- Git Branch:      [%s]\n", GitBranch)
	msg += fmt.Sprintf("- Git Commit:      [%s]\n", GitCommit)
	msg += fmt.Sprintf("- Go Version:      [%s]\n", GoVersion)
	msg += fmt.Sprintf("- OS/Architecture: [%s]\n", OsArch)

	return msg
}
