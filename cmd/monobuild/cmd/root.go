// Copyright © 2018 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/monobuild/monobuild/cmd/monobuild/methods"
	"github.com/sirupsen/logrus"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "monobuild",
	Short: "run a build",
	Long: `Called without an argument monobuild will run a mono repository build with the 
default flags or the ones provided in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// set default log level
		methods.SetLogLevel(viper.GetString("log-level"))

		if err := methods.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.monobuild.yaml)")
	rootCmd.Flags().BoolP("no-parallelism", "s", false, "disable parallel execution of steps")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not show header (version info and name)")
	rootCmd.Flags().StringP("log-level", "l", "warn", "Set log level to debug, info or warn (fallback)")
	rootCmd.Flags().StringP("marker", "m", ".MONOBUILD", "name of marker file")
	rootCmd.Flags().StringP("limit", "o", "", "build only build configuration")

	viper.BindPFlag("no-parallelism", rootCmd.Flags().Lookup("no-parallelism"))
	viper.BindPFlag("quiet", rootCmd.Flags().Lookup("quiet"))
	viper.BindPFlag("log-level", rootCmd.Flags().Lookup("log-level"))
	viper.BindPFlag("marker", rootCmd.Flags().Lookup("marker"))
	viper.BindPFlag("limit", rootCmd.Flags().Lookup("limit"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".monobuild" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".monobuild")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
