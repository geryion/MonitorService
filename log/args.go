package log

import "go.uber.org/zap"

var Mlog *zap.Logger

const (
	LOGPATH     = "~/MonitorService/log/"
	LOGFILE     = LOGPATH + "monitor.log"
	CONFIGFILE  = "~/MonitorService/log.config"
	MCONFIGPATH = "~/MonitorService/etc/"
	MAXSIZE     = 50
)
