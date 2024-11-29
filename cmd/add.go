// Package cmd ...
package cmd

import (
	"strconv"
	"time"

	"drbaz.com/timesheet/logging"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Timesheet ...
type Timesheet struct {
	gorm.Model
	TimesheetDate   string
	TimesheetNumber string
	WeekNumber      string
	WorkDate      string
	StartTime  string
	EndTime    string
	HoursWorked       string
	OvertimeHours    string
	TotalHours       string
	Approved     	bool
	Pdf           string
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new entry to the timesheet",
	Long:  `Adds a new entry to the timesheet. Saves the entry to a database.`,
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Info("Add timesheet\n")
		// get input , week number, commencing, w1daydate ,w1start, w1end, w1over
		// calculate w1hours and w1total(w1hours+w1over)
		Logger.Debugf("Hours = %s", getHours("08:30", "12:30"))
		Logger.Debugf("Total = %s", calcTotalHours("8.5", "0.5"))

		//TODO add pterm
		myStyle := logging.MyLipglossStyle("Add Timesheet")
		Logger.Print(myStyle.Render("Add Timesheet"))

		// //* Timesheet Date
		timesheetDate := pterm.DefaultInteractiveTextInput
		timesheetDate.DefaultText = "Please enter the timesheet date (DDth MMMMMMMM YYYY)"
		result0, _ := timesheetDate.Show()
		pterm.Println()

		//* Timesheet Number
		timesheetNumber := pterm.DefaultInteractiveTextInput
		timesheetNumber.DefaultText = "Please enter the timesheet number as a digit e.g 5"
		result1, _ := timesheetNumber.Show()
		pterm.Println()

		// //* Week Number
		weekNumber := pterm.DefaultInteractiveTextInput
		weekNumber.DefaultText = "Please enter the week number of this timesheet as a digit between 1 and 5"
		result2, _ := weekNumber.Show()
		pterm.Println()

		// //* WorkDate
		workDate := pterm.DefaultInteractiveTextInput
		workDate.DefaultText = "Please enter the work date as DD-MMM-YYYY"
		result3, _ := workDate.Show()
		pterm.Println()

		// //* Start Time
		startTime := pterm.DefaultInteractiveTextInput
		startTime.DefaultText = "Please enter the start time as hh:mm"
		result4, _ := startTime.Show()
		pterm.Println()

		// //* End Time
		endTime := pterm.DefaultInteractiveTextInput
		endTime.DefaultText = "Please enter the end time as hh:mm"
		result5, _ := endTime.Show()
		pterm.Println()

		// //* Hours Worked (claculated)
		hoursWorked := getHours(result4, result5)
		// hoursWorked.DefaultText = "Please enter the hours worked this period"
		// result5, _ := hoursWorked.Show()
		// pterm.Println()

		// //* Overtime Hours
		overtimeHours := pterm.DefaultInteractiveTextInput
		overtimeHours.DefaultText = "Please enter the overtime hours worked"
		result6,_ := overtimeHours.Show()
		pterm.Println()

		Logger.Debug("Hours Worked", "is", hoursWorked)
		Logger.Debug("Overtime Hours", "is", result6)

		// //* Total Hours
		totalHours := calcTotalHours(hoursWorked, result6)


		result, _ := pterm.DefaultInteractiveConfirm.Show()
		// Print a blank line for better readability.
		pterm.Println()
		// Print the user's answer in a formatted way.
		Logger.Debug("you answered: ", "answer", result)
		pterm.Printfln("You answered: %s", BoolToText(result))
		if result {
			db, _ := gorm.Open(sqlite.Open("timesheet.db"))
			db.AutoMigrate(&Timesheet{})
			db.Create(&Timesheet{
				TimesheetDate: result0,
		 		TimesheetNumber:      result1,
				WeekNumber:    result2,
				WorkDate:      result3,
		 		StartTime:  result4,
				EndTime:    result5,
		 		HoursWorked:       hoursWorked,
				OvertimeHours: result6,
		 		TotalHours:    totalHours,
				Approved: false,
					Pdf: "not ready",
		 	})

		 	Logger.Debug("saved invoice to database", "function", "add")
		 } else {
		 	Logger.Debug("invoice aborted", "function", "add")
		 	pterm.Printfln("Invoice aborted")
		}

	},
}

func calcTotalHours(hours string, over string) string {
	totalHours, err := strconv.ParseFloat(hours, 64)
	if err != nil {
		Logger.Error("Error parsing hours")
	}
	overHours, err := strconv.ParseFloat(over, 64)
	if err != nil {
		Logger.Error("Error parsing over")
	}
	totalHours = totalHours + overHours
	return strconv.FormatFloat(totalHours, 'f', 2, 64)
}
func getHours(start string, end string) string {
	startTime, err := time.Parse("15:04", start)
	endTime, err := time.Parse("15:04", end)
	if err != nil {
		Logger.Error("Error parsing time")
	}
	diff := endTime.Sub(startTime)
	hours := diff.Hours()
	return strconv.FormatFloat(hours, 'f', 2, 64)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//BoolToText ...
	func BoolToText(b bool) string {
	if b {
		return pterm.Green("Yes")
	}
	return pterm.Red("No")
}

