package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/chaiyd/mm-wiki/app/utils"
	"github.com/chaiyd/mm-wiki/global"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	Data = NewData()

	installChan = make(chan int, 1)

	InstallDir = ""

	RootDir = ""

	CopyRight = global.SYSTEM_COPYRIGHT
)

const License_Disagree = 0 // <LABEL_956>
const License_Agree = 1    // <LABEL_1121>

const Env_NotAccess = 0 // <LABEL_542>
const Env_Access = 1    // <LABEL_701>

const Sys_NotAccess = 0 // <LABEL_543>
const Sys_Access = 1    // <LABEL_702>

const Database_NotAccess = 0 // <LABEL_358>
const Database_Access = 1    // <LABEL_544>

const Install_Ready = 0 // <LABEL_703>
const Install_Start = 1 // <LABEL_1122>
const Install_End = 2   // <LABEL_1119>

const Install_Default = 0 // <LABEL_1613>
const Install_Failed = 1  // <LABEL_1123>
const Install_Success = 2 // <LABEL_1124>

var defaultSystemConf = map[string]string{
	"addr":         "0.0.0.0",
	"port":         "8080",
	"document_dir": "",
}

var defaultDatabaseConf = map[string]string{
	"host":                "127.0.0.1",
	"port":                "3306",
	"name":                "mm_wiki",
	"user":                "",
	"pass":                "",
	"conn_max_idle":       "30",
	"conn_max_connection": "200",
	"admin_name":          "",
	"admin_pass":          "",
}

func NewData() *data {
	return &data{
		License:      License_Disagree,
		Env:          Env_NotAccess,
		System:       Sys_NotAccess,
		Database:     Database_NotAccess,
		SystemConf:   defaultSystemConf,
		DatabaseConf: defaultDatabaseConf,
		Status:       Install_Ready,
		Result:       "",
		IsSuccess:    Install_Default,
	}
}

type data struct {
	License      int
	Env          int
	System       int
	Database     int
	SystemConf   map[string]string
	DatabaseConf map[string]string
	Status       int
	Result       string
	IsSuccess    int
}

// check db
func checkDB() (err error) {

	host := Data.DatabaseConf["host"]
	port := Data.DatabaseConf["port"]
	user := Data.DatabaseConf["user"]
	pass := Data.DatabaseConf["pass"]

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/")
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return
	}
	return
}

// create db
func createDB() (err error) {
	host := Data.DatabaseConf["host"]
	port := Data.DatabaseConf["port"]
	user := Data.DatabaseConf["user"]
	pass := Data.DatabaseConf["pass"]
	name := Data.DatabaseConf["name"]

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/")
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name + " CHARACTER SET utf8")
	if err != nil {
		return
	}
	return nil
}

// create table
func createTable() (err error) {

	host := Data.DatabaseConf["host"]
	port := Data.DatabaseConf["port"]
	user := Data.DatabaseConf["user"]
	pass := Data.DatabaseConf["pass"]
	name := Data.DatabaseConf["name"]

	sqlBytes, err := ioutil.ReadFile(filepath.Join(RootDir, "docs/databases/table.sql"))
	if err != nil {
		return err
	}
	sqlTable := string(sqlBytes)
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+name+"?charset=utf8&multiStatements=true")
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(sqlTable)
	if err != nil {
		return
	}
	return nil
}

// create admin
func createAdmin() (err error) {
	host := Data.DatabaseConf["host"]
	port := Data.DatabaseConf["port"]
	user := Data.DatabaseConf["user"]
	pass := Data.DatabaseConf["pass"]
	name := Data.DatabaseConf["name"]
	adminName := Data.DatabaseConf["admin_name"]
	adminPass := utils.NewEncrypt().Md5Encode(Data.DatabaseConf["admin_pass"])

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+name+"?charset=utf8")
	if err != nil {
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT mw_user SET username=?,password=?,given_name=?,role_id=?, create_time=?,update_time=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(adminName, adminPass, adminName, 1, time.Now().Unix(), time.Now().Unix())
	return
}

// write database install data
func writeInstallData() (err error) {
	host := Data.DatabaseConf["host"]
	port := Data.DatabaseConf["port"]
	user := Data.DatabaseConf["user"]
	pass := Data.DatabaseConf["pass"]
	name := Data.DatabaseConf["name"]

	sqlBytes, err := ioutil.ReadFile(filepath.Join(RootDir, "docs/databases/data.sql"))
	if err != nil {
		return err
	}
	sqlTable := string(sqlBytes)
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+name+"?charset=utf8&multiStatements=true")
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(sqlTable)
	if err != nil {
		return
	}
	// insert version
	stmt, err := db.Prepare("INSERT mw_config SET `name`=?,`key`=?,`value`=?,create_time=?,update_time=?;")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("<LABEL_957>", "system_version", global.SYSTEM_VERSION, time.Now().Unix(), time.Now().Unix())
	return err
}

// write conf
func makeConf() (err error) {

	templateConf, err := utils.NewFile().GetFileContents(filepath.Join(RootDir, "conf/template.conf"))
	if err != nil {
		return
	}
	// replace conf tag
	templateConf = strings.Replace(templateConf, "#httpaddr#", Data.SystemConf["addr"], 1)
	templateConf = strings.Replace(templateConf, "#httpport#", Data.SystemConf["port"], 1)
	templateConf = strings.Replace(templateConf, "#document_dir#", Data.SystemConf["document_dir"], 1)
	templateConf = strings.Replace(templateConf, "#db.host#", Data.DatabaseConf["host"], 1)
	templateConf = strings.Replace(templateConf, "#db.port#", Data.DatabaseConf["port"], 1)
	templateConf = strings.Replace(templateConf, "#db.name#", Data.DatabaseConf["name"], 1)
	templateConf = strings.Replace(templateConf, "#db.user#", Data.DatabaseConf["user"], 1)
	templateConf = strings.Replace(templateConf, "#db.pass#", Data.DatabaseConf["pass"], 1)
	templateConf = strings.Replace(templateConf, "#db.conn_max_idle#", Data.DatabaseConf["conn_max_idle"], 1)
	templateConf = strings.Replace(templateConf, "#db.conn_max_connection#", Data.DatabaseConf["conn_max_connection"], 1)

	logFilename := strings.Replace(filepath.Join(RootDir, "logs/mm-wiki.log"), `\`, `/`, -1)
	templateConf = strings.Replace(templateConf, "#log.filename#", logFilename, 1)

	fileObject, err := os.OpenFile(filepath.Join(RootDir, "conf/mm-wiki.conf"), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer fileObject.Close()

	_, err = fileObject.Write([]byte(templateConf))
	return
}

func installFailed(err string) {
	Data.Result = err
	Data.Status = Install_End
	Data.IsSuccess = Install_Failed
	log.Println(err)
}

func installSuccess() {

	Data.Status = Install_End
	Data.IsSuccess = Install_Success
	result := map[string]string{
		"cmd": "",
		"url": "http://127.0.0.1:" + Data.SystemConf["port"],
	}
	if runtime.GOOS == "windows" {
		result["cmd"] = "mm-wiki.exe --conf conf/mm-wiki.conf"
	} else {
		result["cmd"] = "./mm-wiki --conf conf/mm-wiki.conf"
	}
	resByte, _ := json.Marshal(result)
	Data.Result = string(resByte)

	// create install lock file
	file, _ := os.Create(filepath.Join(RootDir, "./install.lock"))
	file.Close()
}

func StartInstall() {
	installChan <- 1
}

func ListenInstall() {

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(fmt.Sprintf("install crash: %v", err))
			}
		}()
		for {
			select {
			case <-installChan:
				Data.Status = Install_Start
				// <LABEL_1125>
				log.Println("mm-wiki start install")
				// <LABEL_1614>db
				err := checkDB()
				if err != nil {
					installFailed("<LABEL_545>：" + err.Error())
					continue
				}
				log.Println("database connect success")
				// <LABEL_958>
				err = createDB()
				if err != nil {
					installFailed("<LABEL_546>：" + err.Error())
					continue
				}
				log.Println("create database success")
				// <LABEL_1499>
				err = createTable()
				if err != nil {
					installFailed("<LABEL_959>：" + err.Error())
					continue
				}
				log.Println("create table success")
				// <LABEL_547>
				err = createAdmin()
				if err != nil {
					installFailed("<LABEL_253>：" + err.Error())
					continue
				}
				log.Println("create admin user success")
				// <LABEL_704>
				err = writeInstallData()
				if err != nil {
					installFailed("<LABEL_359>：" + err.Error())
					continue
				}
				log.Println("write install data success")
				// <LABEL_1615> conf <LABEL_1616>
				err = makeConf()
				if err != nil {
					installFailed("<LABEL_360>：" + err.Error())
					continue
				}
				log.Println("make conf file success")
				installSuccess()
				return
			}
		}
	}()
}

func init() {
	ListenInstall()
}
