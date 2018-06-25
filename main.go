package main

import (
	"fmt"
	_ "goflow/routers"
	"os"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

func main() {

	//setup Log
	logs.Async()
	logFiles := beego.AppConfig.String("LogFiles")
	logFilesMaxDays := beego.AppConfig.String("LogFilesMaxDays")
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"`+logFiles+`","separate":["error","info"],"maxdays":`+logFilesMaxDays+`}`)

	beego.Debug("App started.")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else {
		beego.BeeLogger.DelLogger("console")
	}

	beego.Run()

}

func init() { //func init in main.go
	//set file config
	args := os.Args
	if len(args) > 1 && args[1] != "" {
		beego.LoadAppConfig("ini", args[1])

	}

	port, err := beego.AppConfig.Int("httpport")
	fmt.Println("port", port, err)
}
