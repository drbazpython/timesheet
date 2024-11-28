//Package cmd ...
package cmd

import (
	"os"
	"drbaz.com/timesheet/configs"
	"drbaz.com/timesheet/logging"
	"github.com/spf13/cobra"
)

//Config ... 
var Config = configs.LoadEnvViper()
//Logger ...
var Logger = logging.DefineLogger(Config.LogLevel)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "timesheet",
	Short: "Application to generate a timesheet",
	Long: `The timesheet generation is in two parts. 
	add - from the command line, enter data for the timesheet, which is stored in a database. The CLI is defined using bubbletea	
	create - generate the timesheet in the PDF format and print it.
	The create command has a flag "print", which is set to false by default. If the flag is set to true, the PDF is printed.
	list - shows the timesheets in a table`,
	
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.timesheet.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


