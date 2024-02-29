package controllers

import (
	"time"

	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
)

type CollectionController struct {
	BaseController
}

func (this *CollectionController) Add() {

	redirect := this.Ctx.Request.Referer()

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/main/index")
	}
	resourceId := this.GetString("resource_id", "")
	colType, _ := this.GetInt("type", 1)

	if resourceId == "" {
		this.jsonError("<LABEL_364>！")
	}
	if colType != models.Collection_Type_Doc && colType != models.Collection_Type_Space {
		this.jsonError("<LABEL_712>！")
	}

	collect, err := models.CollectionModel.GetCollectionByUserIdTypeAndResourceId(this.UserId, colType, resourceId)
	if err != nil {
		this.ErrorLog("<LABEL_713>：" + err.Error())
		this.jsonError("<LABEL_713>！")
	}
	if len(collect) > 0 {
		this.jsonError("<LABEL_963>，<LABEL_714>！")
	}
	insertCollection := map[string]interface{}{
		"user_id":     this.UserId,
		"resource_id": resourceId,
		"type":        colType,
		"create_time": time.Now().Unix(),
	}
	collectId, err := models.CollectionModel.Insert(insertCollection)
	if err != nil {
		this.ErrorLog("<LABEL_713>：" + err.Error())
		this.jsonError("<LABEL_713>！")
	}

	this.InfoLog("<LABEL_1129> " + utils.Convert.IntToString(collectId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_1130>！", nil, redirect)
}

func (this *CollectionController) Cancel() {

	redirect := this.Ctx.Request.Referer()
	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/space/list")
	}
	collectionId := this.GetString("collection_id", "")

	if collectionId == "" {
		this.jsonError("<LABEL_365>！")
	}

	collection, err := models.CollectionModel.GetCollectionByCollectionId(collectionId)
	if err != nil {
		this.ErrorLog("<LABEL_1131> " + collectionId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_715>！")
	}
	if len(collection) == 0 {
		this.jsonError("<LABEL_549>！")
	}
	if collection["user_id"] != this.UserId {
		this.jsonError("<LABEL_191>！")
	}

	err = models.CollectionModel.Delete(collectionId)
	if err != nil {
		this.ErrorLog("<LABEL_1131> " + collectionId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_715>！")
	}

	this.InfoLog("<LABEL_1131> " + collectionId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_964>！", nil, redirect)
}
