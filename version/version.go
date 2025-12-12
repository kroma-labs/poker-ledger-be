package version

import (
	"fmt"
	"runtime"
)

// Version returns the main version number that is being run at the moment.
const Version = "0.1.0"

// GoVersion returns the version of the go runtime used to compile the binary
var GoVersion = runtime.Version()

// OsArch returns the os and arch used to build the binary
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)

// AppName returns the name of the app
var AppName = "poker-ledger"

// ScopeName is the instrumentation scope name.
const ScopeName = "github.com/kroma-labs/poker-ledger-be"

// V is a struct that contains all the version information
type V struct {
	Version   string `json:"version"`
	GoVersion string `json:"go_version"`
	OSArch    string `json:"os_arch"`
	AppName   string `json:"app"`
	ScopeName string `json:"scope_name"`
}

// Get returns a V struct with all the version information
// This is used to populate the version endpoint of the API
func Get() V {
	return V{
		Version:   Version,
		GoVersion: GoVersion,
		OSArch:    OsArch,
		AppName:   AppName,
		ScopeName: ScopeName,
	}
}
