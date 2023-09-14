package version_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mjdusa/go-template/internal/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Setup Suite
type VersionSuite struct {
	suite.Suite
}

func Test_VersionSuite(t *testing.T) {
	suite.Run(t, &VersionSuite{})
}

func (s *VersionSuite) Test_GetVersion_unpopulated() {
	expected := fmt.Sprintf("%s version: []\n- Build Time:      []\n- Git Branch:      []\n- Git Commit:      []\n- Go Version:      []\n- OS/Architecture: []\n", os.Args[0])

	version.AppVersion = ""
	version.BuildTime = ""
	version.GitBranch = ""
	version.GitCommit = ""
	version.GoVersion = ""
	version.OsArch = ""

	actual := version.GetVersion()

	assert.Equal(s.T(), expected, actual, "GetVersion() unpopulated message expected '%s', but got '%s'", expected, actual)
}

func (s *VersionSuite) Test_GetVersion_populated() {
	version.AppVersion = "v1.2.3"
	version.BuildTime = "01/01/1970T00:00:00.0000 GMT"
	version.GitBranch = "main"
	version.GitCommit = "1234567890abcdef"
	version.GoVersion = "1.20.5"
	version.OsArch = "Darwin/amd64"

	expected := fmt.Sprintf("%s version: [%s]\n- Build Time:      [%s]\n- Git Branch:      [%s]\n- Git Commit:      [%s]\n- Go Version:      [%s]\n- OS/Architecture: [%s]\n",
		os.Args[0], version.AppVersion, version.BuildTime, version.GitBranch, version.GitCommit, version.GoVersion, version.OsArch)

	actual := version.GetVersion()

	assert.Equal(s.T(), expected, actual, "GetVersion() populated message expected '%s', but got '%s'", expected, actual)
}
