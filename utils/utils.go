package utils

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	/// BasePath is used to retrieve absolute path
	_, bp, _, _ = runtime.Caller(0)
	BasePath    = filepath.Dir(bp)

	/// used to run a function only once during program execution
	configFile = "configuration_data = x"
	RunOnce    = sync.Once{}
)

func ScrollToTop(ctx app.Context, e app.Event) {
	app.Window().ScrollToID("app-header")
}
