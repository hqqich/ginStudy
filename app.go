package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/syndtr/goleveldb/leveldb"
	"html/template"
	"jyksServer/common"
	"jyksServer/model"
	"jyksServer/router"
	"log"
	"os"
	"strconv"
)

func main() {

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
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
