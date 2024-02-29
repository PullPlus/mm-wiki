package controllers

import (
	"fmt"
	"github.com/chaiyd/mm-wiki/app/work"
	"strings"

	"github.com/chaiyd/mm-wiki/app/models"
)

type ConfigController struct {
	BaseController
}

func (this *ConfigController) Global() {

	configs, err := models.ConfigModel.GetConfigs()
	if err != nil {
		this.ErrorLog("<LABEL_433>: " + err.Error())
		this.ViewError("<LABEL_418>", "/system/main/index")
	}

	var configValue = map[string]string{}
	for _, config := range configs {
		if config["key"] == models.ConfigKeyAutoFollowdoc && config["value"] != "1" {
			config["value"] = "0"
		}
		if config["key"] == models.ConfigKeySendEmail && config["value"] != "1" {
			config["value"] = "0"
		}
		if config["key"] == models.ConfigKeyAuthLogin && config["value"] != "1" {
			config["value"] = "0"
		}
		configValue[config["key"]] = config["value"]
	}

	this.Data["configValue"] = configValue
	this.viewLayout("config/form", "config")
}

func (this *ConfigController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/email/list")
	}
	mainTitle := this.GetString(models.ConfigKeyMainTitle, "")
	mainDescription := strings.TrimSpace(this.GetString(models.ConfigKeyMainDescription, ""))
	autoFollowDocOpen := strings.TrimSpace(this.GetString(models.ConfigKeyAutoFollowdoc, "0"))
	sendEmailOpen := strings.TrimSpace(this.GetString(models.ConfigKeySendEmail, "0"))
	ssoOpen := strings.TrimSpace(this.GetString(models.ConfigKeyAuthLogin, "0"))
	fulltextSearch := strings.TrimSpace(this.GetString(models.ConfigKeyFulltextSearch, "0"))
	docSearchTimer := strings.TrimSpace(this.GetString(models.ConfigKeyDocSearchTimer, "3600"))
	systemName := strings.TrimSpace(this.GetString(models.ConfigKeySystemName, "Markdown Mini Wiki"))

	if sendEmailOpen == "1" {
		email, err := models.EmailModel.GetUsedEmail()
		if err != nil {
			this.ErrorLog("<LABEL_145>: " + err.Error())
			this.jsonError("<LABEL_1195>！")
		}
		if len(email) == 0 {
			this.jsonError("<LABEL_5>！")
		}
	}

	if ssoOpen == "1" {
		auth, err := models.AuthModel.GetUsedAuth()
		if err != nil {
			this.ErrorLog("<LABEL_146>: " + err.Error())
			this.jsonError("<LABEL_1195>！")
		}
		if len(auth) == 0 {
			this.jsonError("<LABEL_6>！")
		}
	}
	updateValues := map[string]string{
		models.ConfigKeyMainTitle:       mainTitle,
		models.ConfigKeyMainDescription: mainDescription,
		models.ConfigKeyAutoFollowdoc:   autoFollowDocOpen,
		models.ConfigKeySendEmail:       sendEmailOpen,
		models.ConfigKeyAuthLogin:       ssoOpen,
		models.ConfigKeyFulltextSearch:  fulltextSearch,
		models.ConfigKeyDocSearchTimer:  docSearchTimer,
		models.ConfigKeySystemName:      systemName,
	}
	// <LABEL_825>
	configs, err := models.ConfigModel.GetConfigs()
	if err != nil {
		this.ErrorLog("<LABEL_434>: " + err.Error())
		this.jsonError("<LABEL_826>！")
	}
	updateKeys := make(map[string]string)
	for _, config := range configs {
		if len(config) == 0 {
			continue
		}
		name := config["name"]
		key := config["key"]
		value := config["value"]
		updateValue, ok := updateValues[key]
		if !ok {
			continue
		}
		// <LABEL_590>
		if value == updateValue {
			continue
		}
		_, err := models.ConfigModel.UpdateByKey(key, updateValue)
		if err != nil {
			this.ErrorLog(fmt.Sprintf("<LABEL_1196> %s <LABEL_1618>: %s", name, err.Error()))
			this.jsonError(fmt.Sprintf("<LABEL_1196> %s <LABEL_1618>", name))
		}
		updateKeys[key] = updateValue
	}

	// <LABEL_827>
	//this.configUpdateCallback(updateKeys)
	this.InfoLog("<LABEL_435>")
	this.jsonSuccess("<LABEL_435>", nil, "/system/config/global")
}

// <LABEL_436>
func (this *ConfigController) configUpdateCallback(updateKeyMaps map[string]string) {
	fullTextOpenUpdate := false
	updateValue, ok := updateKeyMaps[models.ConfigKeyFulltextSearch]
	if ok {
		fullTextOpenUpdate = true
	}
	docSearchTimerUpdate := false
	_, ok = updateKeyMaps[models.ConfigKeyDocSearchTimer]
	if ok {
		docSearchTimerUpdate = true
	}
	// <LABEL_828>，<LABEL_829>，<LABEL_1197> worker
	if docSearchTimerUpdate && !fullTextOpenUpdate {
		work.DocSearchWorker.Restart()
		return
	}
	// <LABEL_1198>
	if fullTextOpenUpdate {
		if updateValue == "1" {
			work.DocSearchWorker.Start()
			return
		}
		work.DocSearchWorker.Stop()
	}
}
