#!/bin/bash
echo Enter serivce name:
read serviceName
mkdir $serviceName
echo Enter yaml file name:
read yamlFile
./go-gen $yamlFile $serviceName
echo mod init
go mod init github.com/$serviceName
echo getting dependancies
go build
go fmt $serviceName/*.go