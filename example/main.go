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