//Package cmd ...
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create --print=true/false",
	Short: "Create a pdf timesheet from the database",
	Long: `Creates a pdf timesheet from the database by replacing text in the template. The template name is defined in the config file. If the print flag is set to true, the PDF is printed.`,
	Run: func(cmd *cobra.Command, args []string) {
		printFlag, _ := cmd.Flags().GetBool("print")
		if printFlag {
			Logger.Info("Create timesheet, save to pdf, and print it\n")
		}else{
			Logger.Info("Create timesheet, save to pdf but NOT print it\n")
	}
},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().Bool("print", false, "Print timesheet if set to true")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
