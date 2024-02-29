package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/chaiyd/mm-wiki/app/utils"
	"github.com/chaiyd/mm-wiki/global"
	"github.com/chaiyd/mm-wiki/install/storage"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type InstallController struct {
	BaseController
}

// <LABEL_1113>
func (this *InstallController) Index() {
	this.view("install/index")
}

// <LABEL_1114>
func (this *InstallController) License() {

	if this.isPost() {
		licenseAgree := this.GetString("license_agree", "")
		if licenseAgree == "" || licenseAgree == "0" {
			this.jsonError("<LABEL_184>")
		}
		storage.Data.License = storage.License_Agree
		this.jsonSuccess("", nil, "/install/env")
	} else {
		bytes, _ := ioutil.ReadFile(filepath.Join(storage.RootDir, "./LICENSE"))
		license := string(bytes)
		this.Data["license"] = license
		this.Data["license_agree"] = storage.Data.License

		this.view("install/license")
	}
}

// <LABEL_1115>
func (this *InstallController) Env() {

	if this.isPost() {
		if storage.Data.Env == storage.Env_NotAccess {
			this.jsonError("<LABEL_540>")
		}
		storage.Data.Env = storage.Env_Access
		this.jsonSuccess("", nil, "/install/config")
	}
	storage.Data.Env = storage.Env_Access
	//<LABEL_541>
	host := utils.Misc.GetLocalIp()
	osSys := runtime.GOOS
	server := map[string]string{
		"host":        host,
		"sys":         osSys,
		"install_dir": storage.RootDir,
		"version":     global.SYSTEM_VERSION,
	}

	// <LABEL_1115>
	vm, _ := mem.VirtualMemory()
	vmTotal := vm.Total / 1024 / 1024
	cpuCount, _ := cpu.Counts(true)
	memData := map[string]interface{}{
		"name":    "<LABEL_1607>",
		"require": "400M",
		"value":   strconv.FormatInt(int64(vmTotal), 10) + "M",
		"result":  "1",
	}
	if int(vmTotal) < 400 {
		storage.Data.Env = storage.Env_NotAccess
		memData["result"] = "0"
	}
	cpuData := map[string]interface{}{
		"name":    "CPU",
		"require": "1<LABEL_1828>",
		"value":   strconv.Itoa(cpuCount) + "<LABEL_1828>",
		"result":  "1",
	}
	if cpuCount < 1 {
		storage.Data.Env = storage.Env_NotAccess
		cpuData["result"] = "0"
	}
	envData := []map[string]interface{}{}
	envData = append(envData, memData)
	envData = append(envData, cpuData)

	// <LABEL_698>
	fileTool := utils.NewFile()
	templateConfDir := map[string]string{
		"path":    "conf/template.conf",
		"require": "<LABEL_1829>/<LABEL_1830>",
		"result":  "1",
	}
	err := fileTool.IsWriterReadable(filepath.Join(storage.RootDir, templateConfDir["path"]))
	if err != nil {
		storage.Data.Env = storage.Env_NotAccess
		templateConfDir["result"] = "0"
	}

	databaseTable := map[string]string{
		"path":    "docs/databases/table.sql",
		"require": "<LABEL_1829>/<LABEL_1830>",
		"result":  "1",
	}
	err = fileTool.IsWriterReadable(filepath.Join(storage.RootDir, databaseTable["path"]))
	if err != nil {
		storage.Data.Env = storage.Env_NotAccess
		databaseTable["result"] = "0"
	}

	databaseData := map[string]string{
		"path":    "docs/databases/data.sql",
		"require": "<LABEL_1829>/<LABEL_1830>",
		"result":  "1",
	}
	err = fileTool.IsWriterReadable(filepath.Join(storage.RootDir, databaseData["path"]))
	if err != nil {
		storage.Data.Env = storage.Env_NotAccess
		databaseData["result"] = "0"
	}

	viewsDir := map[string]string{
		"path":    "views",
		"require": "<LABEL_699>",
		"result":  "1",
	}
	isEmpty := utils.File.PathIsEmpty(filepath.Join(storage.RootDir, viewsDir["path"]))
	if isEmpty == true {
		storage.Data.Env = storage.Env_NotAccess
		viewsDir["result"] = "0"
	}

	staticDir := map[string]string{
		"path":    "static",
		"require": "<LABEL_699>",
		"result":  "1",
	}
	isEmpty = utils.File.PathIsEmpty(filepath.Join(storage.RootDir, staticDir["path"]))
	if isEmpty == true {
		storage.Data.Env = storage.Env_NotAccess
		staticDir["result"] = "0"
	}

	dirData := []map[string]string{}
	dirData = append(dirData, templateConfDir)
	dirData = append(dirData, databaseTable)
	dirData = append(dirData, databaseData)
	dirData = append(dirData, viewsDir)
	dirData = append(dirData, staticDir)

	this.Data["server"] = server
	this.Data["envData"] = envData
	this.Data["dirData"] = dirData
	this.view("install/env")
}

// <LABEL_1116>
func (this *InstallController) Config() {

	if this.isPost() {
		addr := strings.TrimSpace(this.GetString("addr", ""))
		documentDir := strings.TrimSpace(this.GetString("document_dir", ""))
		port, _ := this.GetInt32("port", 0)

		if addr == "" {
			this.jsonError("addr <LABEL_1117>，<LABEL_954> 0.0.0.0")
		}
		if port == 0 {
			this.jsonError("<LABEL_353>")
		}
		if port > int32(65535) {
			this.jsonError("<LABEL_700>")
		}
		if documentDir == "" {
			this.jsonError("<LABEL_185>")
		}
		if !filepath.IsAbs(documentDir) {
			this.jsonError("<LABEL_81>")
		}
		docAbsDir, err := filepath.Abs(documentDir)
		if err != nil {
			this.jsonError("<LABEL_249>")
		}
		ok, _ := utils.File.PathIsExists(docAbsDir)
		if !ok {
			this.jsonError("<LABEL_186>")
		}

		storage.Data.SystemConf = map[string]string{
			"addr":         addr,
			"port":         strconv.FormatInt(int64(port), 10),
			"document_dir": documentDir,
		}
		storage.Data.System = storage.Sys_Access
		this.jsonSuccess("", nil, "/install/database")
	}

	sysConf := storage.Data.SystemConf
	this.Data["sysConf"] = sysConf
	this.view("install/config")
}

// <LABEL_955>
func (this *InstallController) Database() {

	if !this.isPost() {
		this.Data["databaseConf"] = storage.Data.DatabaseConf
		this.viewLayoutTitle("mm-wiki-<LABEL_1608>-<LABEL_955>", "install/database", "install")
		return
	}

	host := strings.TrimSpace(this.GetString("host", ""))
	port := strings.TrimSpace(this.GetString("port", ""))
	name := strings.TrimSpace(this.GetString("name", ""))
	user := strings.TrimSpace(this.GetString("user", ""))
	pass := strings.TrimSpace(this.GetString("pass", ""))
	connMaxIdle := strings.TrimSpace(this.GetString("conn_max_idle", "0"))
	connMaxConn := strings.TrimSpace(this.GetString("conn_max_connection", "0"))
	adminName := strings.TrimSpace(this.GetString("admin_name", ""))
	adminPass := strings.TrimSpace(this.GetString("admin_pass", ""))

	if host == "" {
		this.jsonError("<LABEL_1495> host <LABEL_1117>！")
	}
	if port == "" {
		this.jsonError("<LABEL_250>！")
	}
	if name == "" {
		this.jsonError("<LABEL_354>！")
	}
	if user == "" {
		this.jsonError("<LABEL_187>！")
	}
	if pass == "" {
		this.jsonError("<LABEL_251>！")
	}
	if connMaxIdle == "0" {
		this.jsonError("<LABEL_252>0！")
	}
	if connMaxConn == "0" {
		this.jsonError("<LABEL_355>0！")
	}
	if adminName == "" {
		this.jsonError("<LABEL_82>！")
	} else {
		v := validation.Validation{}
		if !v.AlphaNumeric(adminName, "admin_name").Ok {
			this.jsonError("<LABEL_356>！")
		}
	}

	if adminPass == "" {
		this.jsonError("<LABEL_123>！")
	}

	storage.Data.DatabaseConf = map[string]string{
		"host":                host,
		"port":                port,
		"name":                name,
		"user":                user,
		"pass":                pass,
		"conn_max_idle":       connMaxIdle,
		"conn_max_connection": connMaxConn,
		"admin_name":          adminName,
		"admin_pass":          adminPass,
	}
	storage.Data.Database = storage.Database_Access
	this.jsonSuccess("", nil, "/install/ready")
}

// <LABEL_1118>
func (this *InstallController) Ready() {

	if this.isPost() {
		if (storage.Data.License != storage.License_Agree) ||
			(storage.Data.Env != storage.Env_Access) ||
			(storage.Data.System != storage.Sys_Access) ||
			(storage.Data.Database != storage.Database_Access) {
			this.jsonError("<LABEL_357>")
		}
		storage.StartInstall()
		this.jsonSuccess("", nil, "/install/end")
	}

	// <LABEL_1609>
	licenseConf := map[string]interface{}{
		"name":   "<LABEL_1114>",
		"value":  "<LABEL_1610>",
		"result": "1",
		"url":    "/install/license",
	}
	if storage.Data.License != storage.License_Agree {
		licenseConf["value"] = "<LABEL_1496>"
		licenseConf["result"] = "0"
	}
	//<LABEL_1115>
	envConf := map[string]interface{}{
		"name":   "<LABEL_1115>",
		"value":  "<LABEL_1611>",
		"result": "1",
		"url":    "/install/env",
	}
	if storage.Data.Env != storage.Env_Access {
		envConf["value"] = "<LABEL_1497>"
		envConf["result"] = "0"
	}
	//<LABEL_1116>
	sysConf := map[string]interface{}{
		"name":   "<LABEL_1116>",
		"value":  "<LABEL_1612>",
		"result": "1",
		"url":    "/install/config",
	}
	if storage.Data.System != storage.Sys_Access {
		sysConf["value"] = "<LABEL_1498>"
		sysConf["result"] = "0"
	}
	//<LABEL_955>
	databaseConf := map[string]interface{}{
		"name":   "<LABEL_955>",
		"value":  "<LABEL_1612>",
		"result": "1",
		"url":    "/install/database",
	}
	if storage.Data.Database != storage.Database_Access {
		databaseConf["value"] = "<LABEL_1498>"
		databaseConf["result"] = "0"
	}

	readyConf := []map[string]interface{}{}
	readyConf = append(readyConf, licenseConf)
	readyConf = append(readyConf, envConf)
	readyConf = append(readyConf, sysConf)
	readyConf = append(readyConf, databaseConf)

	this.Data["readyConf"] = readyConf
	this.view("install/ready")
}

// <LABEL_1119>
func (this *InstallController) End() {

	if storage.Data.Status == storage.Install_Ready {
		this.Redirect("/install/ready", 302)
	}

	this.view("install/end")
}

// <LABEL_1120>
func (this *InstallController) Status() {

	data := map[string]interface{}{
		"status":     storage.Data.Status,
		"is_success": storage.Data.IsSuccess,
		"result":     storage.Data.Result,
	}

	this.jsonSuccess("ok", data)
}
