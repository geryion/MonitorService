package monitor

import "fmt"

type MonitorTaskInfo struct {
	MonitorTaskId 	 	string
	MonitorTaskFunc	 	func()
}

type MonitorTaskFunc func(interface{}, string, *[]MonitorTaskInfo) (string, error)

func RunMonitorTask(monitorTasks []MonitorTaskInfo) error {
	for _, monitorTask := range monitorTasks {
		switch monitorTask.MonitorTaskId {
		case MonitorAgentTask:
			monitorTask.MonitorTaskFunc()
		case MonitorKeepalivedTask:
			monitorTask.MonitorTaskFunc()
		case MonitorMysqlTask:
			monitorTask.MonitorTaskFunc()
		case MonitorNetTask:
			monitorTask.MonitorTaskFunc()
		default:
			return fmt.Errorf("No matched task ")
		}
	}
	return  nil
}

