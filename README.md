# Gologger

## General Information
This is a simple implementation of a logger for Go. It is based on zap and lumberjack. It is written in go and has been tested on version 1.14.6 on Mac OS X and Go version 1.13.7 on Windows 10.

I created this becuse the out of the box log package in go did not do everythiing I wanted and I needed to integrate log file rotation.

## Sources
This code is based off this blog post and code: https://sunitc.dev/2019/05/27/adding-uber-go-zap-logger-to-golang-project/

## To Do:
- Add logrus option

## Installation
`go get github.com/pedeveaux/gologger`

## Use
As currently configured this logger will log to both the console and a file. The console logs will be colorized to easily identify logging levels. The file will be JSON formatted for ease of automated parsing. If the file specified in the config does not exist it will be created. 

examples/main.go:
```go
package main

import(
	logger "github.com/pedeveaux/gologger"
)

var(

	log = logger.InitLogger()
)

func main(){
	log.Debugf("This should be Debug logged")
	log.Infof("This should be Info logged")
	log.Warnf("This should be Warninig logged")
	log.Errorf("This should be error logged")
}
```
examples/config.json:
```js
{
    "logConfig": {
        "EnableConsole": true,
        "ConsoleLevel": "logger.Debug",
        "ConsoleJSONFormat": false,
        "EnableFile": true,
        "FileLevel": "logger.Info",
        "FileJSONFormat": true,
        "FileLocation": "program.log",
        "MaxSize": 10,
        "Compress": true,
        "MaxAge": 90
    }
}
```
Most of the configuration parameters should be self explanatory. The MaxSize and MaxAge attributes apply to log file rotation. The MaxSize is in megabytes. The MaxAge is in days.