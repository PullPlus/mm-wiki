package controllers

import (
	"strings"

	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"

	valid "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LinkController struct {
	BaseController
}

func (this *LinkController) Add() {
	this.viewLayout("link/form", "link")
}

func (this *LinkController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/link/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	url := strings.TrimSpace(this.GetString("url", ""))
	sequence := strings.TrimSpace(this.GetString("sequence", "0"))
	if name == "" {
		this.jsonError("<LABEL_415>！")
	}
	if url == "" {
		this.jsonError("<LABEL_416>！")
	}
	if valid.Validate(url, is.URL) != nil {
		this.jsonError("<LABEL_277>！")
	}
	ok, err := models.LinkModel.HasLinkName(name)
	if err != nil {
		this.ErrorLog("<LABEL_806>：" + err.Error())
		this.jsonError("<LABEL_806>！")
	}
	if ok {
		this.jsonError("<LABEL_582>！")
	}

	linkId, err := models.LinkModel.Insert(map[string]interface{}{
		"name":     name,
		"url":      url,
		"sequence": sequence,
	})

	if err != nil {
		this.ErrorLog("<LABEL_806>：" + err.Error())
		this.jsonError("<LABEL_806>")
	}
	this.InfoLog("<LABEL_1181> " + utils.Convert.IntToString(linkId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_807>", nil, "/system/link/list")
}

func (this *LinkController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var links []map[string]string
	if keyword != "" {
		count, err = models.LinkModel.CountLinksByKeyword(keyword)
		links, err = models.LinkModel.GetLinksByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.LinkModel.CountLinks()
		links, err = models.LinkModel.GetLinksByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("<LABEL_417>: " + err.Error())
		this.ViewError("<LABEL_417>", "/system/main/index")
	}

	this.Data["links"] = links
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("link/list", "link")
}

func (this *LinkController) Edit() {

	linkId := this.GetString("link_id", "")
	if linkId == "" {
		this.ViewError("<LABEL_983>", "/system/link/list")
	}

	link, err := models.LinkModel.GetLinkByLinkId(linkId)
	if err != nil {
		this.ViewError("<LABEL_983>", "/system/link/list")
	}

	this.Data["link"] = link
	this.viewLayout("link/form", "link")
}

func (this *LinkController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/link/list")
	}
	linkId := this.GetString("link_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	url := strings.TrimSpace(this.GetString("url", ""))
	sequence := strings.TrimSpace(this.GetString("sequence", ""))

	if linkId == "" {
		this.jsonError("<LABEL_983>！")
	}
	if name == "" {
		this.jsonError("<LABEL_415>！")
	}
	if url == "" {
		this.jsonError("<LABEL_416>！")
	}
	if valid.Validate(url, is.URL) != nil {
		this.jsonError("<LABEL_277>！")
	}

	link, err := models.LinkModel.GetLinkByLinkId(linkId)
	if err != nil {
		this.ErrorLog("<LABEL_1182> " + linkId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_808>！")
	}
	if len(link) == 0 {
		this.jsonError("<LABEL_983>！")
	}

	ok, _ := models.LinkModel.HasSameName(linkId, name)
	if ok {
		this.jsonError("<LABEL_582>！")
	}
	_, err = models.LinkModel.Update(linkId, map[string]interface{}{
		"name":     name,
		"url":      url,
		"sequence": sequence,
	})

	if err != nil {
		this.ErrorLog("<LABEL_1182> " + linkId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_808>")
	}
	this.InfoLog("<LABEL_1182> " + linkId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_809>", nil, "/system/link/list")
}

func (this *LinkController) Delete() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/link/list")
	}
	linkId := this.GetString("link_id", "")
	if linkId == "" {
		this.jsonError("<LABEL_810>！")
	}

	link, err := models.LinkModel.GetLinkByLinkId(linkId)
	if err != nil {
		this.ErrorLog("<LABEL_1183> " + linkId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_811>")
	}
	if len(link) == 0 {
		this.jsonError("<LABEL_983>")
	}

	err = models.LinkModel.Delete(linkId)
	if err != nil {
		this.ErrorLog("<LABEL_1183> " + linkId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_811>")
	}

	this.InfoLog("<LABEL_1183> " + linkId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_812>", nil, "/system/link/list")
}
