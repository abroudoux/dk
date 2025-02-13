package images

import (
	"fmt"

	"github.com/abroudoux/dk/internal/logs"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

func getEnv() string {
	var key string
	var value string
	huh.NewInput().Title("Key").Prompt("? ").Value(&key).Run()
	huh.NewInput().Title("Value").Prompt("? ").Value(&value).Run()

	return key + "=" + value
}

func isEnvAlreadySaved(newEnv string, envs *[]string) bool {
	for _, env := range *envs {
		if env == newEnv {
			return true
		}
	}
	return false
}

func getEnvs(envs *[]string) {
	newEnv := getEnv()
	envAlreadySaved := isEnvAlreadySaved(newEnv, envs)
	if envAlreadySaved {
		logs.WarnMsg("Environment variable already saved")
		getEnvs(envs)
		return
	}

	*envs = append(*envs, newEnv)
	log.Info(fmt.Sprintf("Environment variable saved: %s", newEnv))

	addNewEnv := utils.GetConfirmation("Do you want to add another environment variable?")
	if addNewEnv {
		getEnvs(envs)
		return
	}
	return
}
