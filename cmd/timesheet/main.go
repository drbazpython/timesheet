package main

import (
	//docprocessing "drbaz.com/timesheet/cmd/docporocessing"
	"drbaz.com/timesheet/configs"
	"drbaz.com/timesheet/logging"
	"drbaz.com/timesheet/cmd"
)
//Config ...
var Config = configs.LoadEnvViper()
//Logger ...
var Logger = logging.DefineLogger(Config.LogLevel)
func main(){		
	Logger.Info("Started Timesheet App\n")
	//newInvoice := docprocessing.ReplaceDocument("26Nov2024")
	//Logger.Info(newInvoice)
	cmd.Execute()
}