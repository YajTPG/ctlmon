package main

import (
	"bytes"
	"flag"
	"math/rand"
	"os"
	"runtime"

	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

var Version string = "v0.0.0" // Use ldflags to set this https://pkg.go.dev/cmd/link
var logger = logging.MustGetLogger("ctlmon")
var TempFile = "/tmp/.ctlmon-stopped"

func main() {
	var ConfigOverride string

	logger.Debug("OS:", runtime.GOOS)
	if runtime.GOOS != "linux" {
		logger.Panic("This program only works on Linux :(")
	}
	logger.Info("Starting up...")
	logger.Infof("Version: %s", Version)
	flag.StringVar(&ConfigOverride, "c", "", "Override default configuration location")
	flag.Parse()

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/ctlmon")
	viper.SetConfigType("yaml")

	var ReadConfigErr error
	if ConfigOverride == "" {
		ReadConfigErr = viper.ReadInConfig()
	} else {
		ReadConfigErr = nil
		File, err := os.ReadFile(ConfigOverride)
		if err != nil {
			logger.Panic(err)
		}
		viper.ReadConfig(bytes.NewBuffer(File))
	}
	if ReadConfigErr != nil {
		if _, ok := ReadConfigErr.(viper.ConfigFileNotFoundError); ok {
			logger.Info("No config file found, creating one...")
			viper.SetDefault("Services", []string{"dbus.service"})
			viper.SetDefault("WebhookURL", "")
			viper.SetDefault("WebhookEnabled", false)
			viper.SetDefault("RoleID", "")
			viper.SetDefault("Version", Version)
			hostname, err := os.Hostname()
			if err != nil {
				hostname = "example_node_" + string(rune(rand.Intn(10000)))
			}
			viper.SetDefault("NodeName", hostname)
			MakeDir("/etc/ctlmon")
			WriteConfigErr := viper.SafeWriteConfig()
			if WriteConfigErr != nil {
				logger.Fatal("Error while writing to config:", WriteConfigErr)
			} else {
				logger.Info("Location of config: ", viper.GetViper().ConfigFileUsed())
			}
		} else {
			logger.Fatal("Error while reading config:", ReadConfigErr)
		}
	}
	logger.Debug("Keys in config:", viper.AllKeys())
	LastStopped := LoadTempFile(TempFile)
	CurrentStopped := []string{}
	for _, service := range viper.GetStringSlice("Services") {
		ServiceStatus := CheckService(service)
		if ServiceStatus == STOPPED || ServiceStatus == UNKNOWN {
			if !Contains(LastStopped, service) {
				logger.Debugf("Sending webhook for %s...", service)
				SendWebhook(service)
			}
			CurrentStopped = append(CurrentStopped, service)
		}
	}
	err := AppendTempFile(TempFile, CurrentStopped)
	if err != nil {
		logger.Panic(err)
	}
}

func Contains(arr []string, val string) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}
	return false
}
