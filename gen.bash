#!/bin/bash
echo Enter serivce name:
read serviceName
mkdir $serviceName
echo Enter yaml file name:
read yamlFile
./go-gen $yamlFile $serviceName
dep init 
dep ensure
go fmt $serviceName/*.go