package common

import (
	"embed"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"time"
)

var StartTime = time.Now()
var Version = "v0.0.1"
var OptionMap map[string]string

const (
	RoleGuestUser  = 0
	RoleCommonUser = 1
	RoleAdminUser  = 10
)

var (
	FileUploadPermission    = RoleAdminUser
	FileDownloadPermission  = RoleAdminUser
	ImageUploadPermission   = RoleAdminUser
	ImageDownloadPermission = RoleAdminUser
)

var (
	GlobalApiRateLimit = 20
	GlobalWebRateLimit = 60
	DownloadRateLimit  = 10
	CriticalRateLimit  = 3
)

const (
	UserStatusEnabled  = 1
	UserStatusDisabled = 2 // don't use 0
)

var (
	Port      = flag.Int("port", 80, "specify the server listening port.")
	Host      = flag.String("host", "localhost", "the server's ip address or domain")
	Path      = flag.String("path", "", "specify a local path to public")
	VideoPath = flag.String("video", "", "specify a video folder to public")
)

var UploadPath = "upload"
var LocalFileRoot = UploadPath
var ImageUploadPath = "upload/images"
var VideoServePath = "upload"

//go:embed public
var FS embed.FS

var SessionSecret = uuid.New().String()

func init() {

	initConfig()

	flag.Parse()
	if *Path != "" {
		LocalFileRoot = *Path
	}
	if *VideoPath != "" {
		VideoServePath = *VideoPath
	}

	LocalFileRoot, _ = filepath.Abs(LocalFileRoot)
	VideoServePath, _ = filepath.Abs(VideoServePath)
	ImageUploadPath, _ = filepath.Abs(ImageUploadPath)

	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(UploadPath, 0777)
	}
	if _, err := os.Stat(ImageUploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(ImageUploadPath, 0777)
	}
}

var (
	DataSource string
)

func initConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("无法加载配置文件：%v\n", err)
		return
	}

	// 读取指定 section 的指定 key 的值
	DataSource = cfg.Section("system").Key("dataSource").String()
	//distPath = cfg.Section("system").Key("distPath").String()
	//port, _ = strconv.Atoi(cfg.Section("system").Key("port").String())
}
