package argo

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/samxyg/k8s-argocd-sync-forcer/logger"
)

func Forcesync(argocdAppName string) (string, error) {
	logger.Info("Start to force sync on ", argocdAppName)
	commandFile, err := makeCommandFileByArgs(argocdAppName)
	if err != nil {
		return "", fmt.Errorf("error executing command on makeCommandFileByArgs: %v", err)
	}

	result, err := runOSFile(commandFile)
	if err != nil {
		return "", fmt.Errorf("error executing command on runOSFile: %v", err)
	}

	logger.Info("End to force sync on ", argocdAppName)
	return result, nil
}

func runOSFile(file string) (string, error) {
	cmd := exec.Command("/bin/sh", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command in runOSFile: %v", err)
	}
	return string(output), nil
}

func makeCommandFileByArgs(argocdAppName string) (string, error) {
	server := os.Getenv("ARGO_SERVER")
	user := os.Getenv("ARGO_USER")
	password := os.Getenv("ARGO_PASSWORD")

	loginCommand := fmt.Sprintf("argocd --grpc-web login %s --name %s --username %s --password %s \n", server, user, user, password)
	syncCommand := fmt.Sprintf("argocd app sync %s --grpc-web --server %s --force\n", argocdAppName, server)

	content := loginCommand + syncCommand
	strFile, err := writeFile([]byte(content), "commands.txt")
	if err != nil {
		logger.Fatal("Write failed.")
		return strFile, err
	}

	return strFile, nil
}

func writeFile(content []byte, fileName string) (string, error) {
	// Content to write to the file
	//content = []byte("Hello, this is some content to write to the file!")

	// Write content to a new file
	err := os.WriteFile("/tmp/"+fileName, content, 0644)
	if err != nil {
		logger.Fatal(err)
		return "", err
	}

	logger.Debug("File " + fileName + " created and content written successfully.")

	return "/tmp/" + fileName, nil
}
