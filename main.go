/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/samxyg/k8s-argocd-sync-forcer/health"
	"github.com/samxyg/k8s-argocd-sync-forcer/k8s"
	"github.com/samxyg/k8s-argocd-sync-forcer/logger"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {

	// load env file
	err := godotenv.Load()
	if err != nil {
		logger.Info("File .env does not exist.")
	}

	// start health checking server
	go health.Start()

	// Get the value of the "ENV_VAR_NAME" environment variable
	envVarValue := os.Getenv("ENV_VAR_NAME")
	logger.Debug("Environment variable value:", envVarValue)

	k8s.Watch()
}
