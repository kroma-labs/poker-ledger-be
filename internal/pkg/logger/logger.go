package logger

import "github.com/itsLeonB/ezutil/v2"

var Logger ezutil.Logger

func Init() {
	Logger = ezutil.NewSimpleLogger("poker-ledger-be", true, 0)
}
