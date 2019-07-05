package main

import(
	"github.com/gin-gonic/gin"
	"account/routs"
	"account/db"
	"os"
	"log"
	"io"
	"github.com/sirupsen/logrus"
	"account/utils"
	"github.com/jinzhu/gorm"
)

var Dbs *gorm.DB
func main(){
	config := utils.ReadConfig()
	f, _ := os.OpenFile(config.LogName+".log",os.O_WRONLY,0666)
	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f)
	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logger.SetOutput(gin.DefaultWriter)
	Dbs = db.Open(config.DbURI,logger)
	routs.Dbs=Dbs
	r := gin.Default()
	r.Use(gin.Recovery())
	r.POST("/init", routs.Init)
	r.Run(":"+config.Port)
}