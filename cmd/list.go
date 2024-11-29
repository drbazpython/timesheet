//Package cmd ...
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	//"github.com/pterm/pterm"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows the timesheet data in a table",
	Long: `Shows the timesheet data in a table`,
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Info("List timesheets\n")
		timesheet := Timesheet{}
		db, _ := gorm.Open(sqlite.Open("timesheet.db"))
		rows  := db.Find(&timesheet).RowsAffected
		Logger.Debugf("Database Rows = %d",rows)

		// put database rows into a t
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
