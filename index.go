package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/kataras/iris/v12"
)

const (
// maxSize   = 1 * iris.GB
// uploadDir = "./uploads"

)

type configStruct struct {
	Port                 int
	MaxUploadFileSize_kb int
	UploadDir            string
}

var (
	defaultConfig = &configStruct{
		Port:                 9090,
		MaxUploadFileSize_kb: 1000000,
		UploadDir:            "./uploads",
	}
	config = defaultConfig
	DB     *bolt.DB
	App    *iris.Application
)

func run() {
	var err error
	// read configFile
	err = readConfigFile()
	if err != nil {
		log.Print("cannot read config.json !")
	}
	// file/dir Mod test
	err = testChmod()
	if err != nil {
		log.Print("File mod 0700 false!")
		return
	}
	// connect DB
	err = connectDB()
	if err != nil {
		log.Print("cannot connect db!")
		return
	}
	defer DB.Close()

	// run App
	App = iris.New()
	setRouter(App)
	App.Listen(":" + strconv.Itoa(config.Port))
}

// test file or dir limit
func testChmod() error {
	var err error
	err = os.Mkdir(config.UploadDir, 0700)
	if os.IsNotExist(err) {
		return err
	}
	err = os.Chmod("./", 0700)
	if os.IsNotExist(err) {
		return err
	}
	return err
}

// read config file
func readConfigFile() error {
	_bytes, err := os.ReadFile("./config.json")
	if err != nil {
		if os.IsExist(err) {
			return err
		} else {
			return nil
		}
	}

	err = json.Unmarshal(_bytes, config)
	// CopyFields(&config, d)
	config.MaxUploadFileSize_kb = config.MaxUploadFileSize_kb * iris.KB
	return err
}

func connectDB() error {
	var err error
	DB, err = bolt.Open("chaojijian.db", 0700, nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer DB.Close()
	return err
}
