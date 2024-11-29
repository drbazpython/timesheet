//Package main ...
package main

import (
	//_ "embed"
	//docprocessing "drbaz.com/timesheet/cmd/docporocessing"
	"drbaz.com/timesheet/configs"
	"drbaz.com/timesheet/logging"
	"drbaz.com/timesheet/cmd"
)

//Config ...
var Config = configs.LoadEnvViper()
//Logger ...
var Logger = logging.DefineLogger(Config.LogLevel)

// //go:embed templates/timesheet.docx
// //EmWordTemplate The embedded word template 
// var EmWordTemplate []byte

func main(){		
	Logger.Info("Started Timesheet App\n")
	//Logger.Infof("timesheet template: %v",EmWordTemplate)
	//newInvoice := docprocessing.ReplaceDocument(EmWordTemplate,"26Nov2024")
	//Logger.Info(newInvoice)
	cmd.Execute()
}