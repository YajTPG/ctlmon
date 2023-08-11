package main

import "github.com/IceWhaleTech/CasaOS-Common/utils/systemctl"

const (
	RUNNING = 1
	STOPPED = 0
	UNKNOWN = 2
)

func CheckService(service string) int {
	logger.Debug("Checking service:", service)
	serviceInfo, err := systemctl.IsServiceRunning(service)
	if err != nil {
		logger.Criticalf("Error while checking service %s: %s", service, err)
		return UNKNOWN
	}
	if serviceInfo {
		logger.Debugf("Service %s is running", service)
		return RUNNING
	}
	logger.Debugf("Service %s is down", service)
	return STOPPED

}
