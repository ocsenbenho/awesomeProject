package main

import (
	"awesomeProject/db"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"runtime"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "ocsen",
		Password: "ocsenbenho",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()
	logErr("co loi xay ra !")

	e := echo.New()

	e.Logger.Fatal(e.Start(":8000"))
}
func logErr(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Error(errMsg)
}
