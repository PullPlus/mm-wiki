package controllers

import (
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
)

type FollowController struct {
	BaseController
}

func (this *FollowController) Add() {

	redirect := this.Ctx.Request.Referer()
	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/main/index")
	}
	objectId := this.GetString("object_id", "")
	followType, _ := this.GetInt("type", 1)
	if objectId == "" {
		this.jsonError("<LABEL_361>！")
	}
	if followType != models.Follow_Type_Doc && followType != models.Follow_Type_User {
		this.jsonError("<LABEL_706>！")
	}
	if followType == models.Follow_Type_User && objectId == this.UserId {
		this.jsonError("<LABEL_707>！")
	}

	follow, err := models.FollowModel.GetFollowByUserIdAndTypeAndObjectId(this.UserId, followType, objectId)
	if err != nil {
		this.ErrorLog("<LABEL_708>：" + err.Error())
		this.jsonError("<LABEL_708>！")
	}
	if len(follow) > 0 {
		this.jsonError("<LABEL_960>，<LABEL_709>！")
	}
	fId, err := models.FollowModel.Insert(this.UserId, followType, objectId)
	if err != nil {
		this.ErrorLog("<LABEL_708>：" + err.Error())
		this.jsonError("<LABEL_708>！")
	}

	this.InfoLog("<LABEL_1126>" + utils.Convert.IntToString(fId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_1127>！", nil, redirect)
}

func (this *FollowController) Cancel() {

	redirect := this.Ctx.Request.Referer()
	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/main/index")
	}
	followId := this.GetString("follow_id", "")
	if followId == "" {
		this.jsonError("<LABEL_361>！")
	}

	follow, err := models.FollowModel.GetFollowByFollowId(followId)
	if err != nil {
		this.ErrorLog("<LABEL_710>：" + err.Error())
		this.jsonError("<LABEL_710>！")
	}
	if len(follow) == 0 {
		this.jsonError("<LABEL_548>！")
	}
	if follow["user_id"] != this.UserId {
		this.jsonError("<LABEL_188>！")
	}

	err = models.FollowModel.Delete(followId)
	if err != nil {
		this.ErrorLog("<LABEL_1128> " + followId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_362>！")
	}

	this.InfoLog("<LABEL_1128> " + followId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_961>！", nil, redirect)
}
