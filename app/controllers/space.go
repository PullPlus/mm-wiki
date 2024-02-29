package controllers

import (
	"strings"
	"time"

	"github.com/chaiyd/mm-wiki/app/models"
)

type SpaceController struct {
	BaseController
}

func (this *SpaceController) Index() {

	// get space tags
	spaceTags := map[string]string{}
	spaces, err := models.SpaceModel.GetSpaces()
	if err == nil {
		for _, space := range spaces {
			tags := space["tags"]
			if tags == "" {
				continue
			}
			tagList := strings.Split(tags, ",")
			for _, tagName := range tagList {
				spaceTags[tagName] = tagName
			}
		}
	}

	this.Data["spaceTags"] = spaceTags
	this.viewLayout("space/index", "space")
}

func (this *SpaceController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)

	limit := (page - 1) * number
	var err error
	var count int64
	var spaces []map[string]string
	if keyword != "" {
		count, err = models.SpaceModel.CountSpacesByKeyword(keyword)
		spaces, err = models.SpaceModel.GetSpacesByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.SpaceModel.CountSpaces()
		spaces, err = models.SpaceModel.GetSpacesByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("<LABEL_205>: " + err.Error())
		this.ViewError("<LABEL_391>", "/main/index")
	}

	collectionSpaces, err := models.CollectionModel.GetCollectionsByUserIdAndType(this.UserId, models.Collection_Type_Space)
	if err != nil {
		this.ErrorLog("<LABEL_205>: " + err.Error())
		this.ViewError("<LABEL_205>", "/main/index")
	}
	for _, space := range spaces {
		space["collection"] = "0"
		space["collection_id"] = "0"
		for _, collectionSpace := range collectionSpaces {
			if collectionSpace["resource_id"] == space["space_id"] {
				space["collection"] = "1"
				space["collection_id"] = collectionSpace["collection_id"]
				break
			}
		}
	}

	this.Data["spaces"] = spaces
	this.Data["keyword"] = keyword
	this.Data["count"] = count
	this.SetPaginator(number, count)
	this.viewLayout("space/list", "default")
}

func (this *SpaceController) Member() {

	page, _ := this.GetInt("page", 1)
	spaceId := strings.TrimSpace(this.GetString("space_id", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)

	if spaceId == "" {
		this.ViewError("<LABEL_720>！")
	}
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1158> " + spaceId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_966>！")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_966>！")
	}

	limit := (page - 1) * number

	count, err := models.SpaceUserModel.CountSpaceUsersBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1159> " + spaceId + " <LABEL_759>: " + err.Error())
		this.ViewError("<LABEL_206>！", "/space/index")
	}
	spaceUsers, err := models.SpaceUserModel.GetSpaceUsersBySpaceIdAndLimit(spaceId, limit, number)
	if err != nil {
		this.ErrorLog("<LABEL_1159> " + spaceId + " <LABEL_759>: " + err.Error())
		this.ViewError("<LABEL_206>！", "/space/index")
	}

	var userIds = []string{}
	for _, spaceUser := range spaceUsers {
		userIds = append(userIds, spaceUser["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("<LABEL_1159> " + spaceId + " <LABEL_759>: " + err.Error())
		this.ViewError("<LABEL_206>！", "/space/index")
	}
	for _, user := range users {
		for _, spaceUser := range spaceUsers {
			if spaceUser["user_id"] == user["user_id"] {
				user["space_privilege"] = spaceUser["privilege"]
				user["space_user_id"] = spaceUser["space_user_id"]
			}
		}
	}
	this.Data["users"] = users
	this.Data["space_id"] = spaceId
	this.SetPaginator(number, count)

	// check user space privilege
	_, _, isManager := this.GetDocumentPrivilege(space)
	if isManager {
		var otherUsers = []map[string]string{}
		if len(userIds) > 0 {
			otherUsers, err = models.UserModel.GetUserByNotUserIds(userIds)
		} else {
			otherUsers, err = models.UserModel.GetUsers()
		}
		if err != nil {
			this.ErrorLog("<LABEL_1159> " + spaceId + " <LABEL_759>: " + err.Error())
			this.ViewError("<LABEL_206>！", "/space/index")
		}
		this.Data["otherUsers"] = otherUsers
		this.viewLayout("space/manager_member", "default")
	} else {
		this.viewLayout("space/member", "default")
	}
}

func (this *SpaceController) AddMember() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/space/index")
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
		this.ErrorLog("<LABEL_1160> " + spaceId + " <LABEL_1161>: " + err.Error())
		this.jsonError("<LABEL_392>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_966>！")
	}

	// check login user space member privilege
	_, _, isManager := this.GetDocumentPrivilege(space)
	if !isManager {
		this.jsonError("<LABEL_86>！", "/space/index")
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
		"user_id":     userId,
		"space_id":    spaceId,
		"privilege":   privilege,
		"create_time": time.Now().Unix(),
		"update_time": time.Now().Unix(),
	}
	_, err = models.SpaceUserModel.Insert(insertValue)
	if err != nil {
		this.ErrorLog("<LABEL_1160> " + spaceId + " <LABEL_1623> " + userId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_761>！")
	}

	this.InfoLog("<LABEL_1624> " + spaceId + " <LABEL_1162> " + userId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_762>！", nil, "/space/member?space_id="+spaceId)
}

func (this *SpaceController) RemoveMember() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/space/index")
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

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1163> " + spaceId + " <LABEL_1161>: " + err.Error())
		this.jsonError("<LABEL_393>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_966>！")
	}

	// check login user space member privilege
	_, _, isManager := this.GetDocumentPrivilege(space)
	if !isManager {
		this.ViewError("<LABEL_87>！", "/space/index")
	}

	err = models.SpaceUserModel.Delete(spaceUserId)
	if err != nil {
		this.ErrorLog("<LABEL_1163> " + spaceId + " <LABEL_1500> " + userId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_763>！")
	}

	this.InfoLog("<LABEL_1163> " + spaceId + " <LABEL_1500> " + userId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_764>！", nil, "/space/member?space_id="+spaceId)
}

func (this *SpaceController) ModifyMember() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/space/list")
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

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1164> " + spaceId + " <LABEL_767>: " + err.Error())
		this.jsonError("<LABEL_209>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_966>！")
	}

	// check login user space member privilege
	_, _, isManager := this.GetDocumentPrivilege(space)
	if !isManager {
		this.ViewError("<LABEL_88>！", "/space/index")
	}

	_, err = models.SpaceUserModel.Update(spaceUserId, map[string]interface{}{
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

func (this *SpaceController) Collection() {

	collectionSpaces, err := models.CollectionModel.GetCollectionsByUserIdAndType(this.UserId, models.Collection_Type_Space)
	if err != nil {
		this.ErrorLog("<LABEL_210>: " + err.Error())
		this.ViewError("<LABEL_210>", "/space/list")
	}

	spaceIds := []string{}
	for _, collectionSpace := range collectionSpaces {
		spaceIds = append(spaceIds, collectionSpace["resource_id"])
	}

	spaces, err := models.SpaceModel.GetSpaceBySpaceIds(spaceIds)
	if err != nil {
		this.ErrorLog("<LABEL_210>: " + err.Error())
		this.ViewError("<LABEL_210>", "/space/list")
	}

	for _, space := range spaces {
		space["collection_id"] = "0"
		for _, collectionSpace := range collectionSpaces {
			if collectionSpace["resource_id"] == space["space_id"] {
				space["collection_id"] = collectionSpace["collection_id"]
				break
			}
		}
	}
	this.Data["spaces"] = spaces
	this.Data["count"] = len(spaces)
	this.viewLayout("space/collection", "default")
}

func (this *SpaceController) Document() {

	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.ViewError("<LABEL_966>")
	}
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1158> " + spaceId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_966>！")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_966>！")
	}

	// check space visit_level
	isVisit, _, _ := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("<LABEL_199>！")
	}

	spaceDocument, err := models.DocumentModel.GetSpaceDefaultDocument(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1158> " + spaceId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_394>！")
	}
	if len(spaceDocument) == 0 {
		this.ViewError("<LABEL_561>！")
	}

	documentId := spaceDocument["document_id"]

	this.Redirect("/document/index?document_id="+documentId, 302)
}

func (this *SpaceController) Search() {

	tagName := strings.TrimSpace(this.GetString("tag", ""))

	spaces, err := models.SpaceModel.GetSpacesByTags(tagName)
	if err != nil {
		this.ErrorLog("<LABEL_211>: " + err.Error())
		this.ViewError("<LABEL_391>", "/main/index")
	}

	collectionSpaces, err := models.CollectionModel.GetCollectionsByUserIdAndType(this.UserId, models.Collection_Type_Space)
	if err != nil {
		this.ErrorLog("<LABEL_205>: " + err.Error())
		this.ViewError("<LABEL_205>", "/main/index")
	}
	for _, space := range spaces {
		space["collection"] = "0"
		space["collection_id"] = "0"
		for _, collectionSpace := range collectionSpaces {
			if collectionSpace["resource_id"] == space["space_id"] {
				space["collection"] = "1"
				space["collection_id"] = collectionSpace["collection_id"]
				break
			}
		}
	}

	this.Data["tag"] = tagName
	this.Data["spaces"] = spaces
	this.Data["count"] = len(spaces)
	this.viewLayout("space/search", "default")
}
