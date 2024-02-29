package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/chaiyd/mm-wiki/app"
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/services"
	"github.com/chaiyd/mm-wiki/app/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/astaxie/beego/logs"
)

type PageController struct {
	BaseController
}

// document page view
func (this *PageController) View() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("<LABEL_972>！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1618>：" + err.Error())
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
		this.ViewError("<LABEL_199>！")
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("<LABEL_550>：" + err.Error())
		this.ViewError("<LABEL_550>！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("<LABEL_723>！")
	}

	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_965>！")
	}

	// get edit user and create user
	users, err := models.UserModel.GetUsersByUserIds([]string{document["create_user_id"], document["edit_user_id"]})
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_718>！")
	}
	if len(users) == 0 {
		this.ViewError("<LABEL_266>！")
	}

	var createUser = map[string]string{}
	var editUser = map[string]string{}
	for _, user := range users {
		if user["user_id"] == document["create_user_id"] {
			createUser = user
		}
		if user["user_id"] == document["edit_user_id"] {
			editUser = user
		}
	}

	collectionId := "0"
	collection, err := models.CollectionModel.GetCollectionByUserIdTypeAndResourceId(this.UserId, models.Collection_Type_Doc, documentId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_750>！")
	}
	if len(collection) > 0 {
		collectionId = collection["collection_id"]
	}

	this.Data["is_editor"] = isEditor
	this.Data["space"] = space
	this.Data["create_user"] = createUser
	this.Data["edit_user"] = editUser
	this.Data["document"] = document
	this.Data["collection_id"] = collectionId
	this.Data["page_content"] = documentContent
	this.Data["parent_documents"] = parentDocuments
	this.viewLayout("page/view", "document_page")
}

// page edit
func (this *PageController) Edit() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("<LABEL_972>！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_1155> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_751>！")
	}
	if len(document) == 0 {
		this.ViewError("<LABEL_965>！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1155> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_751>！")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_254>！")
	}
	// check space visit_level
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.ViewError("<LABEL_59>！")
	}

	// get parent documents by document
	_, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("<LABEL_550>：" + err.Error())
		this.ViewError("<LABEL_550>！")
	}

	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_965>！")
	}

	autoFollowDoc := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyAutoFollowdoc, "0")
	sendEmail := models.ConfigModel.GetConfigValueByKey(models.ConfigKeySendEmail, "0")

	this.Data["sendEmail"] = sendEmail
	this.Data["autoFollowDoc"] = autoFollowDoc
	this.Data["page_content"] = documentContent
	this.Data["document"] = document
	this.viewLayout("page/edit", "document_page")
}

// page modify
func (this *PageController) Modify() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/space/index")
	}
	documentId := this.GetString("document_id", "")
	newName := strings.TrimSpace(this.GetString("name", ""))
	documentContent := this.GetString("document_page_editor-markdown-doc", "")
	comment := strings.TrimSpace(this.GetString("comment", ""))
	isNoticeUser := strings.TrimSpace(this.GetString("is_notice_user", "0"))
	isFollowDoc := strings.TrimSpace(this.GetString("is_follow_doc", "0"))

	// rm document_page_editor-markdown-doc
	this.Ctx.Request.PostForm.Del("document_page_editor-markdown-doc")

	if documentId == "" {
		this.jsonError("<LABEL_559>！")
	}
	if newName == "" {
		this.jsonError("<LABEL_366>！")
	}
	match, err := regexp.MatchString(`[\\\\/:*?\"<>、|]`, newName)
	if err != nil {
		this.jsonError("<LABEL_256>！")
	}
	if match {
		this.jsonError("<LABEL_256>！")
	}
	if newName == utils.Document_Default_FileName {
		this.jsonError("<LABEL_552> " + utils.Document_Default_FileName + " ！")
	}
	//if comment == "" {
	//	this.jsonError("<LABEL_127>！")
	//}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_1155> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_752>！")
	}
	if len(document) == 0 {
		this.jsonError("<LABEL_965>！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1155> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_752>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_254>！")
	}
	// check space document privilege
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.jsonError("<LABEL_59>！")
	}

	// not allow update space document home page name
	if document["parent_id"] == "0" {
		newName = document["name"]
	}
	// check document name
	if newName != document["name"] {
		newDocument, err := models.DocumentModel.GetDocumentByNameParentIdAndSpaceId(newName,
			document["parent_id"], document["space_id"], utils.Convert.StringToInt(document["type"]))
		if err != nil {
			this.ErrorLog("<LABEL_751>：" + err.Error())
			this.jsonError("<LABEL_752>！")
		}
		if len(newDocument) != 0 {
			this.jsonError("<LABEL_257>！")
		}
	}

	// update document and file content
	updateValue := map[string]interface{}{
		"name":         newName,
		"edit_user_id": this.UserId,
	}
	_, err = models.DocumentModel.UpdateDBAndFile(documentId, spaceId, document, documentContent, updateValue, comment)
	if err != nil {
		this.ErrorLog("<LABEL_1155> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_751>！")
	}

	// send email to follow user
	if isNoticeUser == "1" {
		logInfo := this.GetLogInfoByCtx()
		url := fmt.Sprintf("%s:%d/document/index?document_id=%s", this.Ctx.Input.Site(), this.Ctx.Input.Port(), documentId)
		go func(documentId string, username string, comment string, url string) {
			err := sendEmail(documentId, username, comment, url)
			if err != nil {
				logInfo["message"] = "<LABEL_60>：" + err.Error()
				logInfo["level"] = models.Log_Level_Error
				models.LogModel.Insert(logInfo)
				logs.Error("<LABEL_60>：" + err.Error())
			}
		}(documentId, this.User["username"], comment, url)
	}
	// follow doc
	if isFollowDoc == "1" {
		go func(userId string, documentId string) {
			_, _ = models.FollowModel.FollowDocument(userId, documentId)
		}(this.UserId, documentId)
	}
	// <LABEL_753>
	go func(documentId string) {
		_ = services.DocIndexService.ForceUpdateDocIndexByDocId(documentId)
	}(documentId)

	this.InfoLog("<LABEL_1155> " + documentId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_754>！", nil, "/document/index?document_id="+documentId)
}

// document share display
func (this *PageController) Display() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("<LABEL_972>！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_718>！")
	}
	if len(document) == 0 {
		this.ViewError("<LABEL_965>！")
	}

	// get document space
	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1156> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_752>！")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_254>！")
	}

	// check space is allow display
	if space["is_share"] != fmt.Sprintf("%d", models.Space_Share_True) {
		this.ViewError("<LABEL_383>！")
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("<LABEL_550>：" + err.Error())
		this.ViewError("<LABEL_550>！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("<LABEL_723>！")
	}

	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1157>：" + err.Error())
		this.ViewError("<LABEL_965>！")
	}

	// get edit user and create user
	users, err := models.UserModel.GetUsersByUserIds([]string{document["create_user_id"], document["edit_user_id"]})
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_718>！")
	}
	if len(users) == 0 {
		this.ViewError("<LABEL_266>！")
	}

	var createUser = map[string]string{}
	var editUser = map[string]string{}
	for _, user := range users {
		if user["user_id"] == document["create_user_id"] {
			createUser = user
		}
		if user["user_id"] == document["edit_user_id"] {
			editUser = user
		}
	}

	this.Data["create_user"] = createUser
	this.Data["edit_user"] = editUser
	this.Data["document"] = document
	this.Data["page_content"] = documentContent
	this.Data["parent_documents"] = parentDocuments
	this.viewLayout("page/display", "document_share")
}

// export file
func (this *PageController) Export() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("<LABEL_972>！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1618>：" + err.Error())
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

	// check space document privilege
	isVisit, _, _ := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("<LABEL_61>！")
	}

	// check space is allow export
	if space["is_export"] != fmt.Sprintf("%d", models.Space_Download_True) {
		this.ViewError("<LABEL_267>！")
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("<LABEL_550>：" + err.Error())
		this.ViewError("<LABEL_550>！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("<LABEL_723>！")
	}

	packFiles := []*utils.CompressFileInfo{}

	absPageFile := utils.Document.GetAbsPageFileByPageFile(pageFile)
	// pack document file
	packFiles = append(packFiles, &utils.CompressFileInfo{
		File:       absPageFile,
		PrefixPath: "",
	})

	// get document attachments
	attachments, err := models.AttachmentModel.GetAttachmentsByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_374>：" + err.Error())
		this.ViewError("<LABEL_374>！")
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
	var dest = fmt.Sprintf("%s/mm_wiki/%s.zip", os.TempDir(), document["name"])
	err = utils.Zipx.PackFile(packFiles, dest)
	if err != nil {
		this.ErrorLog("<LABEL_384>：" + err.Error())
		this.ViewError("<LABEL_755>！")
	}

	this.Ctx.Output.Download(dest, document["name"]+".zip")
}

func sendEmail(documentId string, username string, comment string, url string) error {

	// get document by documentId
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		return errors.New("<LABEL_85>：" + err.Error())
	}

	// get send email open config
	sendEmailConfig := models.ConfigModel.GetConfigValueByKey(models.ConfigKeySendEmail, "0")
	if sendEmailConfig == "0" {
		return nil
	}

	// get email config
	emailConfig, err := models.EmailModel.GetUsedEmail()
	if err != nil {
		return errors.New("<LABEL_13>：" + err.Error())
	}
	if len(emailConfig) == 0 {
		return nil
	}

	// get follow doc user
	follows, err := models.FollowModel.GetFollowsByObjectIdAndType(documentId, models.Follow_Type_Doc)
	if err != nil {
		return errors.New("<LABEL_44>：" + err.Error())
	}
	if len(follows) == 0 {
		return nil
	}
	userIds := []string{}
	for _, follow := range follows {
		userIds = append(userIds, follow["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		return errors.New("<LABEL_44>：" + err.Error())
	}
	if len(users) == 0 {
		return nil
	}
	emails := []string{}
	for _, user := range users {
		if user["email"] != "" {
			emails = append(emails, user["email"])
		}
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		return errors.New("<LABEL_385>: " + err.Error())
	}
	if len(parentDocuments) == 0 {
		return errors.New("<LABEL_385>")
	}
	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		return errors.New("<LABEL_385>: " + err.Error())
	}

	if len([]byte(documentContent)) > 500 {
		documentContent = string([]byte(documentContent)[:500])
	}

	documentValue := document
	documentValue["content"] = documentContent
	documentValue["username"] = username
	documentValue["comment"] = comment
	documentValue["url"] = url

	emailTemplate := beego.BConfig.WebConfig.ViewsPath + "/system/email/template.html"
	body, err := utils.Email.MakeDocumentHtmlBody(documentValue, emailTemplate)
	if err != nil {
		return errors.New("<LABEL_200>：" + err.Error())
	}
	// start send email
	return utils.Email.Send(emailConfig, emails, "<LABEL_756>", body)
}
