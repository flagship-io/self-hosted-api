package main

import (
	"fmt"
	"strings"

	"github.com/flagship-io/self-hosted-api/pkg/fsclient"
	"github.com/flagship-io/self-hosted-api/pkg/log"
	"github.com/flagship-io/self-hosted-api/pkg/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/flagship-io/self-hosted-api/docs"
)

// @title Flagship Decision Host
// @version 2.0
// @BasePath /v2
// @description This is the Flagship Decision Host API documentation

// @contact.name API Support
// @contact.url https://www.abtasty.com/solutions-product-teams/
// @contact.email support@flagship.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	// Init logger with default Warn level
	log.InitLogger(logrus.WarnLevel)

	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.GetLogger().Warnf("Could not find config file: %v", err)
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	viper.SetDefault("port", 8080)
	port := viper.GetInt("port")

	fsClient, err := fsclient.InitFsClient()

	if err != nil {
		log.GetLogger().Panicf("Error when initializing Flagship: %v", err)
	}

	r := router.Init(fsClient)

	err = r.Run(fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		panic(err)
	}
}
