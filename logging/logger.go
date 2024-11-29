
package logging
import (
	"time"
	"os"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/lipgloss"
)
//DefineLogger ...
func DefineLogger(level string) log.Logger {
	switch {
	case level == "WARN":
		log.SetLevel(log.WarnLevel)
	case level == "INFO":
		log.SetLevel(log.InfoLevel)
	case level == "DEBUG":
		log.SetLevel(log.DebugLevel)
	case level == "ERROR":
		log.SetLevel(log.ErrorLevel)
	}
	
	logger := log.NewWithOptions(os.Stderr, log.Options{
    ReportCaller: true,
    ReportTimestamp: true,
    TimeFormat: time.Kitchen,
    Prefix: "🪲 ",
	Level: log.GetLevel(),
	})
	styles := log.DefaultStyles()

	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DEBUG").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("252")).
		Foreground(lipgloss.Color("0"))

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("FATAL").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("160")).
		Foreground(lipgloss.Color("0"))

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("214")).
		Foreground(lipgloss.Color("0"))

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("301")).
		Foreground(lipgloss.Color("0"))

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("201")).
		Foreground(lipgloss.Color("0"))


	logger.SetReportTimestamp(false)
	logger.SetStyles(styles)
	return *logger
}

//MyLipglossStyle ...
func MyLipglossStyle(word string) lipgloss.Style {
	width := 40
	var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    PaddingTop(1).
	PaddingBottom(1).
    PaddingLeft(4).
	PaddingRight(4).
	Align(lipgloss.Center).
    Width(width).
	Margin(1)

	return style

}