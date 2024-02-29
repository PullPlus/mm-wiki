package controllers

import (
	"encoding/json"
	"github.com/chaiyd/mm-wiki/app/models"
	"strings"
)

type PluginController struct {
	BaseController
}

func (this *PluginController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var plugins []map[string]string
	if keyword != "" {
		count, err = models.PluginModel.CountPluginsByKeyword(keyword)
		plugins, err = models.PluginModel.GetPluginsByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.PluginModel.CountPlugins()
		plugins, err = models.PluginModel.GetPluginsByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("<LABEL_411>: " + err.Error())
		this.ViewError("<LABEL_411>", "/system/main/index")
	}

	this.Data["plugins"] = plugins
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("plugin/list", "plugin")
}

func (this *PluginController) Config() {

	pluginId := this.GetString("plugin_id", "")
	if pluginId == "" {
		this.ViewError("<LABEL_1144>", "/system/plugin/list")
	}

	plugin, err := models.PluginModel.GetPluginByPluginId(pluginId)
	if err != nil {
		this.ErrorLog("<LABEL_804>：" + err.Error())
		this.ViewError("<LABEL_804>", "/system/plugin/list")
	}
	if len(plugin) == 0 {
		this.ViewError("<LABEL_982>", "/system/plugin/list")
	}
	pluginKey, ok := plugin["plugin_key"]
	if !ok || pluginKey == "" {
		this.ViewError("<LABEL_580> Key <LABEL_1506>", "/system/plugin/list")
	}
	configValue := make(map[string]string)
	configValueStr, ok := plugin["conf_value"]
	if ok && configValueStr != "" {
		json.Unmarshal([]byte(configValueStr), &configValue)
	}
	// <LABEL_581>，<LABEL_276>
	this.Data["plugin_config"] = configValue
	this.Data["plugin"] = plugin
	this.viewLayout("plugin/"+pluginKey, "plugin")
}

func (this *PluginController) ConfigModify() {

	pluginId := this.GetString("plugin_id", "")
	confValue := this.GetString("conf_value", "")
	if pluginId == "" {
		this.jsonError("<LABEL_1144>")
	}
	if confValue == "" {
		this.jsonError("<LABEL_805>")
	}

	plugin, err := models.PluginModel.GetPluginByPluginId(pluginId)
	if err != nil {
		this.ErrorLog("<LABEL_804>：" + err.Error())
		this.jsonError("<LABEL_804>")
	}
	if len(plugin) == 0 {
		this.jsonError("<LABEL_982>")
	}

	// <LABEL_1180>
	_, err = models.PluginModel.UpdateConfValueByPluginId(pluginId, confValue)
	if err != nil {
		this.ErrorLog("<LABEL_412>：" + err.Error())
		this.jsonError("<LABEL_413>")
	}
	this.jsonSuccess("<LABEL_414>", nil, "/system/plugin/list")
}
