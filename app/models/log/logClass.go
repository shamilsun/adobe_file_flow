package log

import (
	"fmt"
	"github.com/andlabs/ui"
	stdlog "log"
	"time"
)

var uiLogContainer *ui.MultilineEntry

type IClassLog interface {
	AddLog(container *ui.MultilineEntry)
}

func (l *ILog) AddLog() {
	var t = time.Now()
	var s = fmt.Sprintf("%s \t %s \n", t.Format("2006-01-02 15:04:05"), l.Message)

	if getUiLogContainer() != nil {
		getUiLogContainer().Append(s)
	} else {
		stdlog.Println(s)
	}
}

func AddLog(v string) {
	var logRow = ILog{Message: v}
	logRow.AddLog()
}

func Println(v ...interface{}) {
	stdlog.Println(v)
}

func SetUiLogContainer(v *ui.MultilineEntry) {
	uiLogContainer = v
}

func getUiLogContainer() *ui.MultilineEntry {
	return uiLogContainer
}
