package controllers

import (
	"fmt"
	"github.com/chaiyd/mm-wiki/app"
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
	"os"
	"path"
)

type AttachmentController struct {
	BaseController
}

func (this *AttachmentController) Page() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("<LABEL_716>！", "/space/index")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_717> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_718>！")
	}
	if len(document) == 0 {
		this.ViewError("<LABEL_965>！")
	}
	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_719>：" + err.Error())
		this.ViewError("<LABEL_718>！")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_254>！")
	}
	// check space visit_level
	isVisit, isEditor, isManager := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("<LABEL_37>！")
	}

	// get document attachments
	attachments, err := models.AttachmentModel.GetAttachmentsByDocumentIdAndSource(documentId, models.Attachment_Source_Default)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1139>：" + err.Error())
		this.ViewError("<LABEL_374>！")
	}

	// get username
	userIds := []string{}
	for _, attachment := range attachments {
		userIds = append(userIds, attachment["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1139>：" + err.Error())
		this.ViewError("<LABEL_374>！")
	}
	usernameMap := make(map[string]string)
	for _, user := range users {
		usernameMap[user["user_id"]] = user["username"]
	}
	for _, attachment := range attachments {
		attachment["username"] = usernameMap[attachment["user_id"]]
	}

	this.Data["attachments"] = attachments
	this.Data["document_id"] = documentId
	this.Data["is_upload"] = isEditor
	this.Data["is_delete"] = isManager
	this.viewLayout("attachment/page", "attachment")
}

func (this *AttachmentController) Upload() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/space/index")
	}
	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.uploadJsonError("<LABEL_1144>！", "/space/index")
	}

	// handle document
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_717> " + documentId + " <LABEL_1618>：" + err.Error())
		this.uploadJsonError("<LABEL_718>！")
	}
	if len(document) == 0 {
		this.uploadJsonError("<LABEL_965>！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_719>：" + err.Error())
		this.uploadJsonError("<LABEL_718>！")
	}
	if len(space) == 0 {
		this.uploadJsonError("<LABEL_254>！")
	}
	// check space visit_level
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.uploadJsonError("<LABEL_43>！")
	}

	// handle upload
	f, h, err := this.GetFile("attachment")
	if err != nil {
		this.ErrorLog("<LABEL_375>: " + err.Error())
		this.uploadJsonError("<LABEL_375>")
		return
	}
	if h == nil || f == nil {
		this.ErrorLog("<LABEL_739>")
		this.uploadJsonError("<LABEL_739>")
		return
	}
	_ = f.Close()

	// file save dir
	saveDir := fmt.Sprintf("%s/%s/%s", app.AttachmentAbsDir, spaceId, documentId)
	ok, _ := utils.File.PathIsExists(saveDir)
	if !ok {
		err := os.MkdirAll(saveDir, 0777)
		if err != nil {
			this.ErrorLog("<LABEL_739>: " + err.Error())
			this.uploadJsonError("<LABEL_740>")
			return
		}
	}
	// check file is exists
	attachmentFile := path.Join(saveDir, h.Filename)
	ok, _ = utils.File.PathIsExists(attachmentFile)
	if ok {
		this.uploadJsonError("<LABEL_557>！")
	}
	// save file
	err = this.SaveToFile("attachment", attachmentFile)
	if err != nil {
		this.ErrorLog("<LABEL_741>: " + err.Error())
		this.uploadJsonError("<LABEL_741>")
	}

	// insert db
	attachment := map[string]interface{}{
		"user_id":     this.UserId,
		"document_id": documentId,
		"name":        h.Filename,
		"path":        fmt.Sprintf("attachment/%s/%s/%s", spaceId, documentId, h.Filename),
		"source":      models.Attachment_Source_Default,
	}
	_, err = models.AttachmentModel.Insert(attachment, spaceId)
	if err != nil {
		_ = os.Remove(attachmentFile)
		this.ErrorLog("<LABEL_739>: " + err.Error())
		this.uploadJsonError("<LABEL_376>")
	}

	this.InfoLog(fmt.Sprintf("<LABEL_1619> %s <LABEL_1145> %s <LABEL_1617>", documentId, h.Filename))
	this.jsonSuccess("<LABEL_742>", "", "/attachment/page?document_id="+documentId)
}

func (this *AttachmentController) Delete() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/space/index")
	}
	attachmentId := this.GetString("attachment_id", "")
	if attachmentId == "" {
		this.jsonError("<LABEL_743>！")
	}

	attachment, err := models.AttachmentModel.GetAttachmentByAttachmentId(attachmentId)
	if err != nil {
		this.ErrorLog("<LABEL_1146> " + attachmentId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_744>")
	}
	if len(attachment) == 0 {
		this.jsonError("<LABEL_970>")
	}

	documentId := attachment["document_id"]
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_195> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_196>！")
	}
	if len(document) == 0 {
		this.jsonError("<LABEL_264>！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_377> " + documentId + " <LABEL_719>：" + err.Error())
		this.jsonError("<LABEL_84>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_58>！")
	}
	// check space visit_level
	_, _, isManager := this.GetDocumentPrivilege(space)
	if !isManager {
		this.jsonError("<LABEL_42>！")
	}
	attachmentName := attachment["name"]
	attachmentSource := attachment["source"]

	// delete db
	err = models.AttachmentModel.DeleteAttachmentDBFile(attachmentId)
	if err != nil {
		this.ErrorLog("<LABEL_1146> " + attachmentId + " <LABEL_1618>: " + err.Error())
		this.jsonError("<LABEL_744>")
	}

	// update document log
	go func(userId string, documentId string, attachmentName string, spaceId string) {
		_, _ = models.LogDocumentModel.UpdateAction(userId, documentId, "<LABEL_971> "+attachmentName, spaceId)
	}(this.UserId, documentId, attachmentName, spaceId)

	redirect := fmt.Sprintf("/attachment/page?document_id=%s", documentId)
	if attachmentSource == fmt.Sprintf("%d", models.Attachment_Source_Image) {
		redirect = fmt.Sprintf("/attachment/image?document_id=%s", documentId)
	}

	this.InfoLog("<LABEL_1138> " + documentId + " <LABEL_1620> " + attachmentName + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_1147>", nil, redirect)
}

func (this *AttachmentController) Download() {

	attachmentId := this.GetString("attachment_id", "")
	if attachmentId == "" {
		this.ViewError("<LABEL_743>！")
	}

	attachment, err := models.AttachmentModel.GetAttachmentByAttachmentId(attachmentId)
	if err != nil {
		this.ErrorLog("<LABEL_1148> " + attachmentId + " <LABEL_1618>: " + err.Error())
		this.ViewError("<LABEL_745>")
	}
	if len(attachment) == 0 {
		this.ViewError("<LABEL_970>")
	}

	documentId := attachment["document_id"]
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_195> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_196>！")
	}
	if len(document) == 0 {
		this.ViewError("<LABEL_264>！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_377> " + documentId + " <LABEL_719>：" + err.Error())
		this.ViewError("<LABEL_84>！")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_58>！")
	}
	// check space visit_level
	isVisit, _, _ := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("<LABEL_12>！")
	}
	attachmentFilePath := path.Join(app.DocumentAbsDir, attachment["path"])
	attachmentName := attachment["name"]

	this.Ctx.Output.Download(attachmentFilePath, attachmentName)
}

func (this *AttachmentController) Image() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("<LABEL_716>！", "/space/index")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_717> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_718>！")
	}
	if len(document) == 0 {
		this.ViewError("<LABEL_965>！")
	}
	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_719>：" + err.Error())
		this.ViewError("<LABEL_718>！")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_254>！")
	}
	// check space visit_level
	isVisit, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("<LABEL_37>！")
	}

	// get document attachment images
	attachments, err := models.AttachmentModel.GetAttachmentsByDocumentIdAndSource(documentId, models.Attachment_Source_Image)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1149>：" + err.Error())
		this.ViewError("<LABEL_378>！")
	}

	// get username
	userIds := []string{}
	for _, attachment := range attachments {
		userIds = append(userIds, attachment["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1149>：" + err.Error())
		this.ViewError("<LABEL_378>！")
	}
	usernameMap := make(map[string]string)
	for _, user := range users {
		usernameMap[user["user_id"]] = user["username"]
	}
	for _, attachment := range attachments {
		attachment["username"] = usernameMap[attachment["user_id"]]
	}

	this.Data["attachments"] = attachments
	this.Data["document_id"] = documentId
	this.Data["is_delete"] = isEditor
	this.viewLayout("attachment/image", "attachment")
}
