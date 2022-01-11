package main

import (
	"github.com/flagship-io/self-hosted-api/internal/log"
	"github.com/flagship-io/self-hosted-api/runner"

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
	options := runner.GetOptionsFromConfig()
	err := runner.RunAPI(options)

	log.GetLogger().Panic(err)
}
