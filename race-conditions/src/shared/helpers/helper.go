package helpers

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/lazyguyid/gacor"
)

// Helper interface
type Helper interface {
	IdentifyPanic(ctx string, rec interface{}) string
}

type hp struct {
	app gacor.App
}

// NewHelper func
func NewHelper(app gacor.App) Helper {
	helper := new(hp)
	helper.app = app

	return helper
}

// IdentifyPanic for identify line code in panic recover
func (h *hp) IdentifyPanic(ctx string, rec interface{}) string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	var source string
	switch {
	case name != "":
		source = fmt.Sprintf("%v:%v", name, line)
	case file != "":
		source = fmt.Sprintf("%v:%v", file, line)
	default:
		source = fmt.Sprintf("pc:%x", pc)
	}

	return fmt.Sprintf("panic: %v. source: %s", rec, source)
}
