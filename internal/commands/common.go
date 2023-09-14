package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/mjdusa/go-template/internal/config"
	"github.com/mjdusa/go-template/internal/log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	OsExistCode int = 1
)

var (
	panicOnExit bool = false // Set to true to tell Exit() to Panic rather than call os.Exit() - should ONLY be used for testing
)

func SetPanicOnExitTestMode(val bool) {
	panicOnExit = val
}

func CheckError(err error) {
	if err != nil {
		Error(err)
	}
}

func Error(err error) {
	log.Error(err)

	Exit(OsExistCode)
}

func Fatal(err error) {
	log.Fatal(err)

	Exit(OsExistCode)
}

func Exit(code int) {
	if panicOnExit {
		panic(fmt.Sprintf("PanicOnExit is true, code=%d", code))
	}

	os.Exit(code)
}

func InitConfig(appName string, prefix string, level logrus.Level) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if cmd.Name() != "help" {
			initConfig(appName, prefix, level)
		}
	}
}

func initConfig(appName string, prefix string, level logrus.Level) {
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	log.DefaultLogger.SetLevel(level)

	if viper.GetBool("debug") {
		log.DefaultLogger.SetLevel(logrus.DebugLevel)
	}

	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		Error(err)
	}

	if cfg := viper.GetString("config"); cfg != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfg)
	} else {
		// Search config in home directory with name ".pstore" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(fmt.Sprintf(".%s", prefix))
	}

	usingConfig := false
	if err := viper.ReadInConfig(); err == nil {
		usingConfig = true
	}

	provider := config.LoadConfigProvider(appName)

	log.DefaultLogger = log.NewLogger(provider)

	if usingConfig {
		log.DefaultLogger.Info()
	}
}
