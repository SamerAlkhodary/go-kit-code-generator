#!/bin/bash
echo Enter serivce name:
read serviceName
mkdir $serviceName
echo Enter yaml file name:
read yamlFile
./go-gen $yamlFile $serviceName
echo dep init 
cd $serviceName
go mod init github.com/$serviceName
cd ..
echo getting dependancies
go build
go fmt $serviceName/*.go