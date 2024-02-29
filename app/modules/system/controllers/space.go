package controllers

import (
	"fmt"
	"github.com/chaiyd/mm-wiki/app"
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type SpaceController struct {
	BaseController
}

func (this *SpaceController) Add() {
	this.viewLayout("space/form", "space")
}

func (this *SpaceController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/space/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	description := strings.TrimSpace(this.GetString("description", ""))
	tags := strings.TrimSpace(this.GetString("tags", ""))
	visitLevel := strings.TrimSpace(this.GetString("visit_level", "public"))
	isShare := strings.TrimSpace(this.GetString("is_share", "1"))
	isExport := strings.TrimSpace(this.GetString("is_export", "0"))

	if name == "" {
		this.jsonError("<LABEL_449>！")
	}
	match, err := regexp.MatchString(`[\\\\/:*?\"<>、|]`, name)
	if err != nil {
		this.jsonError("<LABEL_298>！")
	}
	if match {
		this.jsonError("<LABEL_298>！")
	}
	ok, err := models.SpaceModel.HasSpaceName(name)
	if err != nil {
		this.ErrorLog("<LABEL_835>：" + err.Error())
		this.jsonError("<LABEL_835>！")
	}
	if ok {
		this.jsonError("<LABEL_592>！")
	}

	// create space database
	spaceId, err := models.SpaceModel.Insert(map[string]interface{}{
		"name":        name,
		"description": description,
		"tags":        tags,
		"visit_level": strings.ToLower(visitLevel),
		"is_share":    isShare,
		"is_export":   isExport,
	})
	if err != nil {
		this.ErrorLog("<LABEL_835>：" + err.Error())
		this.jsonError("<LABEL_835>")
	}

	// create space document
	spaceDocument := map[string]interface{}{
		"space_id":       fmt.Sprintf("%d", spaceId),
		"parent_id":      "0",
		"name":           name,
		"type":           models.Document_Type_Dir,
		"path":           "0",
		"create_user_id": this.UserId,
		"edit_user_id":   this.UserId,
	}
	_, err = models.DocumentModel.Insert(spaceDocument)
	if err != nil {
		// delete space
		models.SpaceModel.Delete(fmt.Sprintf("%d", spaceId))
		this.ErrorLog("<LABEL_450>：" + err.Error())
		this.jsonError("<LABEL_835>！")
	}

	// add space member
	insertValue := map[string]interface{}{
		"user_id":   this.UserId,
		"space_id":  spaceId,
		"privilege": models.SpaceUser_Privilege_Manager,
	}
	_, err = models.SpaceUserModel.Insert(insertValue)
	if err != nil {
		// delete space
		models.SpaceModel.Delete(fmt.Sprintf("%d", spaceId))
		this.ErrorLog("<LABEL_97>: " + err.Error())
		this.jsonError("<LABEL_835>！")
	}

	this.InfoLog("<LABEL_1160> " + utils.Convert.IntToString(spaceId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_836>", nil, "/system/space/list")
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
		this.ErrorLog("<LABEL_391>: " + err.Error())
		this.ViewError("<LABEL_391>", "/system/main/index")
	}

	this.Data["spaces"] = spaces
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("space/list", "space")
}

func (this *SpaceController) Edit() {

	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.ViewError("<LABEL_966>", "/system/space/list")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_837>: " + err.Error())
		this.ViewError("<LABEL_837>", "/system/space/list")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_966>", "/system/space/list")
	}

	this.Data["space"] = space
	this.viewLayout("space/form", "space")
}

func (this *SpaceController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/space/list")
	}
	spaceId := this.GetString("space_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	description := strings.TrimSpace(this.GetString("description", ""))
	tags := strings.TrimSpace(this.GetString("tags", ""))
	visitLevel := strings.TrimSpace(this.GetString("visit_level", "public"))
	isShare := strings.TrimSpace(this.GetString("is_share", "0"))
	isExport := strings.TrimSpace(this.GetString("is_export", "0"))

	if spaceId == "" {
		this.jsonError("<LABEL_966>！")
	}
	if name == "" {
		this.jsonError("<LABEL_449>！")
	}
	match, err := regexp.MatchString(`[\\\\/:*?\"<>、|]`, name)
	if err != nil {
		this.jsonError("<LABEL_298>！")
	}
	if match {
		this.jsonError("<LABEL_298>！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1201> " + spaceId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_838>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_966>！")
	}

	ok, _ := models.SpaceModel.HasSameName(spaceId, name)
	if ok {
		this.jsonError("<LABEL_592>！")
	}

	spaceValue := map[string]interface{}{
		"name":        name,
		"description": description,
		"tags":        tags,
		"visit_level": visitLevel,
		"is_share":    isShare,
		"is_export":   isExport,
	}
	// update space document dir name if name update
	_, err = models.SpaceModel.UpdateDBAndSpaceFileName(spaceId, spaceValue, space["name"])
	if err != nil {
		this.ErrorLog("<LABEL_1201> " + spaceId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_838>")
	}
	this.InfoLog("<LABEL_1201> " + spaceId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_839>", nil, "/system/space/list")
}

func (this *SpaceController) Member() {

	page, _ := this.GetInt("page", 1)
	spaceId := strings.TrimSpace(this.GetString("space_id", ""))
	number, _ := this.GetRangeInt("number", 15, 10, 100)

	if spaceId == "" {
		this.ViewError("<LABEL_720>！")
	}

	limit := (page - 1) * number

	count, err := models.SpaceUserModel.CountSpaceUsersBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1159> " + spaceId + " <LABEL_759>: " + err.Error())
		this.ViewError("<LABEL_206>！", "/system/space/list")
	}
	spaceUsers, err := models.SpaceUserModel.GetSpaceUsersBySpaceIdAndLimit(spaceId, limit, number)
	if err != nil {
		this.ErrorLog("<LABEL_1159> " + spaceId + " <LABEL_759>: " + err.Error())
		this.ViewError("<LABEL_206>！", "/system/space/list")
	}

	var userIds = []string{}
	for _, spaceUser := range spaceUsers {
		userIds = append(userIds, spaceUser["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("<LABEL_1159> " + spaceId + " <LABEL_759>: " + err.Error())
		this.ViewError("<LABEL_206>！", "/system/main/index")
	}
	for _, user := range users {
		for _, spaceUser := range spaceUsers {
			if spaceUser["user_id"] == user["user_id"] {
				user["space_privilege"] = spaceUser["privilege"]
				user["space_user_id"] = spaceUser["space_user_id"]
			}
		}
	}

	var otherUsers = []map[string]string{}
	if len(userIds) > 0 {
		otherUsers, err = models.UserModel.GetUserByNotUserIds(userIds)
	} else {
		otherUsers, err = models.UserModel.GetUsers()
	}
	if err != nil {
		this.ErrorLog("<LABEL_1159> " + spaceId + " <LABEL_759>: " + err.Error())
		this.ViewError("<LABEL_206>！", "/system/main/index")
	}

	this.Data["users"] = users
	this.Data["space_id"] = spaceId
	this.Data["otherUsers"] = otherUsers
	this.SetPaginator(number, count)
	this.viewLayout("space/member", "space")
}

func (this *SpaceController) Delete() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/system/space/list")
	}
	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.jsonError("<LABEL_720>！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1202> " + spaceId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_840>")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_966>")
	}

	// check space documents
	documents, err := models.DocumentModel.GetDocumentsBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1202> " + spaceId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_840>")
	}
	if len(documents) > 1 {
		this.jsonError("<LABEL_841>，<LABEL_147>")
	} else if len(documents) == 1 {
		if documents[0]["name"] != space["name"] {
			this.jsonError("<LABEL_841>，<LABEL_147>")
		} else {
			// delete space dir and documentId
			_, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(documents[0])
			if err != nil {
				this.ErrorLog("<LABEL_1202> " + spaceId + " <LABEL_451>: " + err.Error())
				this.jsonError("<LABEL_840>")
			}
			err = models.DocumentModel.DeleteDBAndFile(documents[0]["document_id"], spaceId, this.UserId,
				pageFile, fmt.Sprintf("%d", models.Document_Type_Dir))
			// delete space document attachments
			_ = models.AttachmentModel.DeleteAttachmentsDBFileByDocumentId(documents[0]["document_id"])
		}
	} else {
		// delete space dir
		err = utils.Document.DeleteSpace(space["name"])
	}
	if err != nil {
		this.ErrorLog("<LABEL_1202> " + spaceId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_840>")
	}

	// delete space user
	err = models.SpaceUserModel.DeleteBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1202> " + spaceId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_840>")
	}
	// delete space and space document
	err = models.SpaceModel.Delete(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1202> " + spaceId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_840>")
	}

	this.InfoLog("<LABEL_1202> " + spaceId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_842>", nil, "/system/space/list")
}

func (this *SpaceController) Download() {

	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.ViewError("<LABEL_966>", "/system/space/list")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_837>: " + err.Error())
		this.ViewError("<LABEL_837>", "/system/space/list")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_966>", "/system/space/list")
	}

	spaceName := space["name"]
	spacePath := utils.Document.GetAbsPageFileByPageFile(spaceName)

	packFiles := []*utils.CompressFileInfo{}

	// pack space all markdown file
	packFiles = append(packFiles, &utils.CompressFileInfo{
		File:       spacePath,
		PrefixPath: "",
	})

	// get space all document attachments
	attachments, err := models.AttachmentModel.GetAttachmentsBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_223>：" + err.Error())
		this.ViewError("<LABEL_223>！")
	}
	for _, attachment := range attachments {
		if attachment["path"] == "" {
			continue
		}
		path := attachment["path"]
		attachmentFile := filepath.Join(app.DocumentAbsDir, path)
		packFile := &utils.CompressFileInfo{
			File:       attachmentFile,
			PrefixPath: filepath.Dir(path),
		}
		packFiles = append(packFiles, packFile)
	}
	var dest = fmt.Sprintf("%s/mm_wiki/%s.zip", os.TempDir(), spaceName)
	err = utils.Zipx.PackFile(packFiles, dest)
	if err != nil {
		this.ErrorLog("<LABEL_452>：" + err.Error())
		this.ViewError("<LABEL_452>！")
	}

	this.Ctx.Output.Download(dest, spaceName+".zip")
}
