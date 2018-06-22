package main

import (
	_ "goflow/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//setup Log
	logs.Async()
	logFiles := beego.AppConfig.String("LogFiles")
	logFilesMaxDays := beego.AppConfig.String("LogFilesMaxDays")
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"`+logFiles+`","separate":["error","info"],"maxdays":`+logFilesMaxDays+`}`)

	logFilesEs := beego.AppConfig.String("LogFilesEs")
	//logFilesEsLevel := beego.AppConfig.String("LogFilesEsLevel")
	if logFilesEs != "" {
		//logs.SetLogger(logs.AdapterEs, `{"dsn":"`+logFilesEs+`","level":`+logFilesEsLevel+`}`)
		logs.SetLogger(logs.AdapterEs, `{"dsn":"`+logFilesEs+`","level": 1 }`)
		//fmt.Println("==>" + logFilesEs)

	}

	beego.Debug("App started.")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else {
		beego.BeeLogger.DelLogger("console")
	}
	beego.Run()

}
