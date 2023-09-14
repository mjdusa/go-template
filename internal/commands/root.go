package commands

import (
	"fmt"
	"runtime/debug"

	"github.com/mjdusa/go-template/internal/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "generated code example",
		Short: "A brief description of your application",
		Long: `A longer
	multi-line
	description of your application`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
			InitConfig("go-template", "init-prefix", logrus.InfoLevel)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd Run with args: %v\n", args)
			verbose := viper.GetBool("verbose")
			if verbose {
				fmt.Println(version.GetVersion())
			}

			dbg := viper.GetBool("debug")
			if dbg {
				buildInfo, ok := debug.ReadBuildInfo()
				if ok {
					fmt.Println(buildInfo.String())
				}
			}
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
		},
	}
)

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose logging")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "set log level to debug")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		Fatal(err)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v", err)
		Exit(1)
	}
}
