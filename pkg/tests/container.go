package tests

import "runtime"

func SelectMySQLImageByArch() string {
	switch runtime.GOARCH {
	case "arm64":
		return "arm64v8/mysql:8.0.33"
	case "amd64":
		return "mysql:8.0.33"
	default:
		return "mysql:8.0.33"
	}
}
