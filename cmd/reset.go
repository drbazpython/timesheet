//Package cmd ...
package cmd

import (
	"time"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"github.com/spf13/cobra"
	"github.com/pterm/pterm"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "A command to reset the database",
	Long: `A command to reset the database upon aqccpetance of password`,
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Info("Reset database")
		timesheet := Timesheet{}
		// //* Reset password
		password := pterm.DefaultInteractiveTextInput
		password.DefaultText = "Please enter the password"
		result0, _ := password.Show()
		pterm.Println()
		envpassword := Config.ResetPassword

		if result0 == envpassword {	
			db, _ := gorm.Open(sqlite.Open("timesheet.db"))
			rows  := db.Find(&timesheet).RowsAffected
			if rows > 0 {
				db.Where("id > 0").Delete(&timesheet) // soft delete
				db.Unscoped().Where("deleted_at < ?", time.Now()).Delete(&timesheet)
				Logger.Infof("%d Database rows deleted",rows)
			}else{ Logger.Info("No database rows deleted - no data") }
		}else{
			Logger.Info("No database rows deleted - wrong password")
		}
		
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
