package images

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/abroudoux/dk/internal/logs"
)

func checkPortInput(port string) bool {
	portRegex := regexp.MustCompile(`^(\d+):(\d+)$`)
	if !portRegex.MatchString(port) {
		logs.ErrorMsg("Invalid port mapping. Please use the following format: host:container")
		return false
	}

	portParts := strings.Split(port, ":")
	_, errHost := strconv.Atoi(portParts[0])
	_, errContainer := strconv.Atoi(portParts[1])
	if errHost != nil || errContainer != nil {
		logs.ErrorMsg("Invalid port numbers. Please use integers for both host and container ports")
		return false
	}

	return true
}
