package controllers

import (
	"github.com/chaiyd/mm-wiki/app/models"
	"strings"
	"time"
)

type Space_UserController struct {
	BaseController
}

func (this *Space_UserController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/space/list")
	}
	spaceId := strings.TrimSpace(this.GetString("space_id", ""))
	userId := this.GetString("user_id", "")
	privilege := strings.TrimSpace(this.GetString("privilege", "0"))

	if spaceId == "" {
		this.jsonError("<LABEL_966>！")
	}
	if userId == "" {
		this.jsonError("<LABEL_760>！")
	}
	if privilege == "" {
		this.jsonError("<LABEL_207>！")
	}
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1160> " + spaceId + " <LABEL_1623> " + userId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_392>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_966>！")
	}

	spaceUser, err := models.SpaceUserModel.GetSpaceUserBySpaceIdAndUserId(spaceId, userId)
	if err != nil {
		this.ErrorLog("<LABEL_1160> " + spaceId + " <LABEL_1623> " + userId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_392>！")
	}
	if len(spaceUser) > 0 {
		this.jsonError("<LABEL_208>！")
	}

	insertValue := map[string]interface{}{
		"user_id":   userId,
		"space_id":  spaceId,
		"privilege": privilege,
	}
	_, err = models.SpaceUserModel.Insert(insertValue)
	if err != nil {
		this.ErrorLog("<LABEL_1624> " + spaceId + " <LABEL_1162> " + userId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_761>！")
	}

	this.InfoLog("<LABEL_1624> " + spaceId + " <LABEL_1162> " + userId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_762>！", nil, "/system/space/member?space_id="+spaceId)
}

func (this *Space_UserController) Remove() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/space/list")
	}
	spaceId := this.GetString("space_id", "")
	userId := this.GetString("user_id", "")
	spaceUserId := this.GetString("space_user_id", "")

	if spaceUserId == "" {
		this.jsonError("<LABEL_560>！")
	}
	if spaceId == "" {
		this.jsonError("<LABEL_966>！")
	}
	if userId == "" {
		this.jsonError("<LABEL_962>！")
	}

	err := models.SpaceUserModel.Delete(spaceUserId)
	if err != nil {
		this.ErrorLog("<LABEL_1163> " + spaceId + " <LABEL_1500> " + userId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_763>！")
	}

	this.InfoLog("<LABEL_1163> " + spaceId + " <LABEL_1500> " + userId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_764>！", nil, "/system/space/member?space_id="+spaceId)
}

func (this *Space_UserController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/space/list")
	}
	spaceUserId := this.GetString("space_user_id", "")
	privilege := this.GetString("privilege", "0")
	userId := this.GetString("user_id", "")
	spaceId := this.GetString("space_id", "")

	if spaceUserId == "" {
		this.jsonError("<LABEL_765>！")
	}
	if privilege == "" {
		this.jsonError("<LABEL_766>！")
	}

	_, err := models.SpaceUserModel.Update(spaceUserId, map[string]interface{}{
		"privilege":   privilege,
		"update_time": time.Now().Unix(),
	})
	if err != nil {
		this.ErrorLog("<LABEL_1164> " + spaceId + " <LABEL_1500> " + userId + " <LABEL_1151>：" + err.Error())
		this.jsonError("<LABEL_768>！")
	}

	this.InfoLog("<LABEL_1164> " + spaceId + " <LABEL_1500> " + userId + " <LABEL_1165>")
	this.jsonSuccess("<LABEL_769>！", nil)
}
