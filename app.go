package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron/v3"
	"github.com/syndtr/goleveldb/leveldb"
	"html/template"
	"jyksServer/common"
	"jyksServer/model"
	"jyksServer/router"
	"log"
	"os"
	"strconv"
)

type MyJob struct{}

func (j MyJob) Run() {
	fmt.Println("执行")
}

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func main() {

	c := newWithSeconds()
	spec := "*/5 * * * * ?"
	//_, e := c.AddFunc(spec, func() {
	//	fmt.Println("定时任务")
	//})
	_, e := c.AddJob(spec, MyJob{})
	if e != nil {
		fmt.Println("error:", e)
	}
	c.Start()

	//select {}	// 强制阻塞

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Redis
	err := common.InitRedisClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	// 初始化 MySQL
	db, err := model.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// 初始化 LevelDB
	levelDb, err := model.InitLevelDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(levelDb *leveldb.DB) {
		err := levelDb.Close()
		if err != nil {

		}
	}(levelDb)

	// Initialize HTTP server
	server := gin.Default()
	server.SetHTMLTemplate(loadTemplate())
	router.SetRouter(server)
	var realPort = os.Getenv("PORT")
	if realPort == "" { // 环境变量没有，就读constants.go文件
		realPort = strconv.Itoa(*common.Port)
	}
	if *common.Host == "localhost" {
		ip := common.GetIp()
		if ip != "" {
			*common.Host = ip
		}
	}
	serverUrl := "http://" + *common.Host + ":" + realPort + "/"
	common.OpenBrowser(serverUrl)
	err = server.Run(":" + realPort)
	if err != nil {
		log.Println(err)
	}

}

func loadTemplate() *template.Template {
	t := template.Must(template.New("").ParseFS(common.FS, "public/*.html"))
	return t
}
