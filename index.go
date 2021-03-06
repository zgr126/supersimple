package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/muesli/cache2go"
	bolt "go.etcd.io/bbolt"
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
	db     *bolt.DB
	// DB  *badger.DB
	App   *iris.Application
	cache *cache2go.CacheTable
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
	defer db.Close()

	// run App
	App = iris.New()
	newCache()
	setRouter(App)
	App.Listen(":" + strconv.Itoa(config.Port))
}

// test file or dir limit
func testChmod() error {
	// mask := syscall.Umask(0)
	// defer syscall.Umask(mask)
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
	// DB, err = badger.Open(badger.DefaultOptions("./db"))
	db, err = bolt.Open("supersimple.db", 0700, nil)
	if err != nil {
		log.Fatal(err)
	}
	newAdmin(db)
	newApp()
	// defer db.Close()
	return err
}

func newCache() {
	cache = cache2go.Cache("supersimple")

}
