package main

import (
	"MonitorService/log"
	"MonitorService/monitor"
)

func main()  {
	//todo init log
	log.InitMlog()
	//todo parse cmd input
	//todo monitor
	var m_agent = monitor.MonitorTaskInfo{
		MonitorTaskId  : monitor.MonitorAgentTask,
		MonitorTaskFunc: monitor.MonitorAgent,
	}
	var m_keepalived = monitor.MonitorTaskInfo{
		MonitorTaskId  : monitor.MonitorKeepalivedTask,
		MonitorTaskFunc: monitor.MonitorKeepalived,
	}
	var m_mysql = monitor.MonitorTaskInfo{
		MonitorTaskId  : monitor.MonitorMysqlTask,
		MonitorTaskFunc: monitor.MonitorMysql,
	}
	var m_net = monitor.MonitorTaskInfo{
		MonitorTaskId  : monitor.MonitorNetTask,
		MonitorTaskFunc: monitor.MonitorMysql,
	}
	m_tasks := make([]monitor.MonitorTaskInfo, 0)
	m_tasks = append(m_tasks, m_agent)
	m_tasks = append(m_tasks, m_net)
	m_tasks = append(m_tasks, m_keepalived)
	m_tasks = append(m_tasks, m_mysql)
	monitor.RunMonitorTask(m_tasks)
}
