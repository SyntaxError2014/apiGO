// Package used for application configuration
package config

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

// Application configuration file path
var appConfigFile = "config/app.json"

// Application descriptive variables
// These are used throughout the entire application for
// configuration purposess
var (
    ApplicationName   string
    ApiInstance       string
    HttpServerAddress string
    RpcServerAddress  string
    Protocol          string
)

// Struct with the sole purpose of easier serialization
// and deserialization of configuration data
type appConfigHolder struct {
    ApplicationName   string `json:"applicationName"`
    ApiInstance       string `json:"apiInstance"`
    HttpServerAddress string `json:"httpServerAddress"`
    RpcServerAddress  string `json:"rpcServerAddress"`
    Protocol          string `json:"protocol"`
}

// Initialize the application by loading all the necessary configuration.
// If the appConfigPath variable is specified, then the initialization is
// done using the file path in that variable, rather than using the
// default file for storing application configurations
func InitApp(appConfigPath string) {
    if len(appConfigPath) != 0 {
        appConfigFile = appConfigPath
    }

    configData := &appConfigHolder{}

    data, err := ioutil.ReadFile(appConfigFile)

    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(data, &configData)

    if err != nil {
        log.Fatal(err)
    }

    ApplicationName = configData.ApplicationName
    ApiInstance = configData.ApiInstance
    HttpServerAddress = configData.HttpServerAddress
    RpcServerAddress = configData.RpcServerAddress
    Protocol = configData.Protocol
}
