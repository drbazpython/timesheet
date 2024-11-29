## Timesheet

```
//go:embed templates/timesheet.docx
var wordTemplate []byte
```

go get github.com/glebarez/sqlite
go get gorm.io/gorm

### add

``` go
  "strconv"
	"time"
	"drbaz.com/timesheet/logging"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"github.com/glebarez/sqlite"
```
	