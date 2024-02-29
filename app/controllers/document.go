package controllers

import (
	"fmt"
	"github.com/chaiyd/mm-wiki/app/services"
	"regexp"
	"strings"

	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
)

type DocumentController struct {
	BaseController
}

// document index
func (this *DocumentController) Index() {

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
	isVisit, _, _ := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("<LABEL_37>！")
	}

	// get default space document
	spaceDocument, err := models.DocumentModel.GetSpaceDefaultDocument(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_718>！")
	}
	if len(spaceDocument) == 0 {
		this.ViewError(" <LABEL_255>！")
	}

	// get space all document
	documents, err := models.DocumentModel.GetAllSpaceDocuments(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_719>：" + err.Error())
		this.ViewError("<LABEL_718>！")
	}

	// get space privilege
	_, isEditor, isDelete := this.GetDocumentPrivilege(space)

	this.Data["is_editor"] = isEditor
	this.Data["is_delete"] = isDelete
	this.Data["documents"] = documents
	this.Data["default_document_id"] = documentId
	this.Data["space"] = space
	this.Data["space_document"] = spaceDocument
	this.viewLayout("document/index", "document")
}

// add document
func (this *DocumentController) Add() {

	spaceId := this.GetString("space_id", "0")
	parentId := this.GetString("parent_id", "0")

	if spaceId == "0" {
		this.ViewError("<LABEL_720>！")
	}
	if parentId == "0" {
		this.ViewError("<LABEL_721>！")
	}
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_722>：" + err.Error())
		this.ViewError("<LABEL_722>！")
	}
	if len(space) == 0 {
		this.ViewError("<LABEL_966>！")
	}

	// check space document privilege
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.ViewError("<LABEL_38>！")
	}

	parentDocument, err := models.DocumentModel.GetDocumentByDocumentId(parentId)
	if err != nil {
		this.ErrorLog("<LABEL_1133> " + parentId + " <LABEL_1618>：" + err.Error())
		this.ViewError("<LABEL_722>！")
	}
	if len(parentDocument) == 0 {
		this.ViewError("<LABEL_723>！")
	}
	path := parentDocument["path"] + "," + parentId
	// get parent documents by path
	parentDocuments, err := models.DocumentModel.GetParentDocumentsByPath(path)
	if err != nil {
		this.ErrorLog("<LABEL_550>：" + err.Error())
		this.ViewError("<LABEL_550>！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("<LABEL_723>！")
	}

	this.Data["parent_documents"] = parentDocuments
	this.Data["parent_id"] = parentId
	this.Data["space_id"] = spaceId
	this.viewLayout("document/form", "default")

}

// save document
func (this *DocumentController) Save() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/main/index")
	}
	spaceId := strings.TrimSpace(this.GetString("space_id", "0"))
	parentId := strings.TrimSpace(this.GetString("parent_id", "0"))
	docType, _ := this.GetInt("type", models.Document_Type_Page)
	name := strings.TrimSpace(this.GetString("name", ""))

	if spaceId == "0" {
		this.jsonError("<LABEL_720>！")
	}
	if parentId == "0" {
		this.jsonError("<LABEL_551>！")
	}
	if name == "" {
		this.jsonError("<LABEL_366>！")
	}
	match, err := regexp.MatchString(`[\\\\/:*?\"<>、|]`, name)
	if err != nil {
		this.jsonError("<LABEL_256>！")
	}
	if match {
		this.jsonError("<LABEL_256>！")
	}
	if name == utils.Document_Default_FileName {
		this.jsonError("<LABEL_552> " + utils.Document_Default_FileName + " ！")
	}
	if docType != models.Document_Type_Page &&
		docType != models.Document_Type_Dir {
		this.jsonError("<LABEL_724>！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_367>：" + err.Error())
		this.jsonError("<LABEL_725>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_966>！")
	}

	// check space document privilege
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.jsonError("<LABEL_38>！")
	}

	parentDocument, err := models.DocumentModel.GetDocumentByDocumentId(parentId)
	if err != nil {
		this.ErrorLog("<LABEL_367>：" + err.Error())
		this.jsonError("<LABEL_725>！")
	}
	if len(parentDocument) == 0 {
		this.jsonError("<LABEL_723>！")
	}
	if parentDocument["type"] != fmt.Sprintf("%d", models.Document_Type_Dir) {
		this.jsonError("<LABEL_553>！")
	}

	// check document name
	document, err := models.DocumentModel.GetDocumentByNameParentIdAndSpaceId(name, parentId, spaceId, docType)
	if err != nil {
		this.ErrorLog("<LABEL_367>：" + err.Error())
		this.jsonError("<LABEL_725>！")
	}
	if len(document) != 0 {
		this.jsonError("<LABEL_257>！")
	}

	insertDocument := map[string]interface{}{
		"parent_id":      parentId,
		"space_id":       spaceId,
		"name":           name,
		"type":           docType,
		"path":           parentDocument["path"] + "," + parentId,
		"create_user_id": this.UserId,
		"edit_user_id":   this.UserId,
	}
	documentId, err := models.DocumentModel.Insert(insertDocument)
	if err != nil {
		this.ErrorLog("<LABEL_725>：" + err.Error())
		this.jsonError("<LABEL_725>")
	}
	this.InfoLog("<LABEL_1134> " + utils.Convert.IntToString(documentId, 10) + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_726>", nil, "/document/index?document_id="+utils.Convert.IntToString(documentId, 10))
}

// document history
func (this *DocumentController) History() {

	page, _ := this.GetInt("page", 1)
	documentId := this.GetString("document_id", "0")
	number, _ := this.GetRangeInt("number", 10, 10, 100)
	limit := (page - 1) * number

	if documentId == "0" {
		this.ViewError("<LABEL_368>！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_1135> " + documentId + " <LABEL_727>：" + err.Error())
		this.ViewError("<LABEL_192>！")
	}
	if len(document) == 0 {
		this.jsonError("<LABEL_965>！")
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
	isVisit, _, _ := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("<LABEL_39>！")
	}

	logDocuments, err := models.LogDocumentModel.GetLogDocumentsByDocumentIdAndLimit(documentId, limit, number)
	if err != nil {
		this.ErrorLog("<LABEL_1135> " + documentId + " <LABEL_727>：" + err.Error())
		this.ViewError("<LABEL_192>！")
	}
	count, err := models.LogDocumentModel.CountLogDocumentsByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_1135> " + documentId + " <LABEL_727>：" + err.Error())
		this.ViewError("<LABEL_192>！")
	}

	userIds := []string{}
	for _, logDocument := range logDocuments {
		userIds = append(userIds, logDocument["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("<LABEL_1135> " + documentId + " <LABEL_727>：" + err.Error())
		this.ViewError("<LABEL_192>！")
	}
	for _, logDocument := range logDocuments {
		logDocument["username"] = ""
		for _, user := range users {
			if logDocument["user_id"] == user["user_id"] {
				logDocument["username"] = user["username"]
				break
			}
		}
	}

	this.Data["logDocuments"] = logDocuments
	this.SetPaginator(number, count)
	this.viewLayout("document/history", "default")
}

// move document
func (this *DocumentController) Move() {

	documentId := this.GetString("document_id", "0")
	targetId := this.GetString("target_id", "0")
	moveType := this.GetString("move_type", "") // <LABEL_728>

	if documentId == "0" {
		this.jsonError("<LABEL_369>！")
	}
	if targetId == "0" {
		this.jsonError("<LABEL_193>！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_370>：" + err.Error())
		this.jsonError("<LABEL_729>！")
	}
	if len(document) == 0 {
		this.jsonError("<LABEL_965>！")
	}
	if moveType != "next" && moveType != "prev" {
		if document["type"] == fmt.Sprintf("%d", models.Document_Type_Dir) {
			this.jsonError("<LABEL_371>！")
		}
	}

	targetDocument, err := models.DocumentModel.GetDocumentByDocumentId(targetId)
	if err != nil {
		this.ErrorLog("<LABEL_372>：" + err.Error())
		this.jsonError("<LABEL_729>！")
	}
	if len(targetDocument) == 0 {
		this.jsonError("<LABEL_554>！")
	}
	if document["space_id"] != targetDocument["space_id"] {
		this.jsonError("<LABEL_57>！")
	}
	if moveType != "next" && moveType != "prev" {
		if targetDocument["type"] != fmt.Sprintf("%d", models.Document_Type_Dir) {
			this.jsonError("<LABEL_258>！")
		}
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_729>：" + err.Error())
		this.jsonError("<LABEL_729>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_555>！")
	}
	// check space document privilege
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.jsonError("<LABEL_40>！")
	}

	// <LABEL_1136>：next-<LABEL_259> prev-<LABEL_260>
	if moveType == "next" || moveType == "prev" {
		this.updateDocSequence(moveType, document, targetDocument)
		return
	}

	_, oldPageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("<LABEL_1137> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_729>！")
	}
	newDocument := map[string]string{
		"space_id":  document["space_id"],
		"parent_id": targetId,
		"name":      document["name"],
		"type":      document["type"],
		"path":      targetDocument["path"] + "," + targetId,
	}
	_, newPageFile, err := models.DocumentModel.GetParentDocumentsByDocument(newDocument)
	if err != nil {
		this.ErrorLog("<LABEL_1137> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_729>！")
	}

	// update database and move document file
	updateValue := map[string]interface{}{
		"parent_id":    targetId,
		"path":         targetDocument["path"] + "," + targetId,
		"edit_user_id": this.UserId,
	}
	_, err = models.DocumentModel.MoveDBAndFile(documentId, spaceId, updateValue,
		oldPageFile, newPageFile, document["type"], "<LABEL_967> "+targetDocument["name"])
	if err != nil {
		this.ErrorLog("<LABEL_1137> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_729>！")
	}

	this.InfoLog("<LABEL_1137> " + documentId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_730>", nil, "/document/index?document_id="+documentId)
}

// <LABEL_731>
func (this *DocumentController) updateDocSequence(moveType string, document map[string]string, targetDocument map[string]string) {

	sequence := utils.Convert.StringToInt(targetDocument["sequence"])
	spaceId := targetDocument["space_id"]
	targetDocumentId := targetDocument["document_id"]
	movedDocumentId := document["document_id"]

	updateSequence := sequence
	if moveType == "next" {
		updateSequence = sequence + 1
	}

	// <LABEL_732>
	_, err := models.DocumentModel.MoveSequenceBySpaceIdAndGtSequence(spaceId, updateSequence, 1)
	if err != nil {
		this.ErrorLog("<LABEL_1137> " + movedDocumentId + "<LABEL_968> " + targetDocumentId + " " + moveType + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_729>！")
	}

	// <LABEL_261>
	updateValue := map[string]interface{}{
		"sequence":     updateSequence,
		"edit_user_id": this.UserId,
	}
	_, err = models.DocumentModel.Update(movedDocumentId, updateValue, fmt.Sprintf("<LABEL_1137>"), spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1137> " + movedDocumentId + "<LABEL_968> " + targetDocumentId + " " + moveType + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_729>！")
	}
	this.jsonSuccess("<LABEL_730>", "", "/document/index?document_id="+movedDocumentId)
}

// delete document
func (this *DocumentController) Delete() {

	documentId := this.GetString("document_id", "0")

	if documentId == "0" {
		this.jsonError("<LABEL_733>！")
	}
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_734>：" + err.Error())
		this.jsonError("<LABEL_734>！")
	}
	if len(document) == 0 {
		this.jsonError("<LABEL_965>！")
	}
	if document["type"] == fmt.Sprintf("%d", models.Document_Type_Dir) {
		childDocs, err := models.DocumentModel.GetDocumentsByParentId(document["document_id"])
		if err != nil {
			this.ErrorLog("<LABEL_734>：" + err.Error())
			this.jsonError("<LABEL_734>！")
		}
		if len(childDocs) > 0 {
			this.jsonError("<LABEL_41>！")
		}
	}
	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(document["space_id"])
	if err != nil {
		this.ErrorLog("<LABEL_734>：" + err.Error())
		this.jsonError("<LABEL_734>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_555>！")
	}
	// check space document privilege
	_, _, isManager := this.GetDocumentPrivilege(space)
	if !isManager {
		this.jsonError("<LABEL_42>！")
	}

	_, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("<LABEL_1138> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_734>！")
	}

	err = models.DocumentModel.DeleteDBAndFile(documentId, spaceId, this.UserId, pageFile, document["type"])
	if err != nil {
		this.ErrorLog("<LABEL_1138> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_734>！")
	}

	// delete attachment
	err = models.AttachmentModel.DeleteAttachmentsDBFileByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_1138> " + documentId + " <LABEL_1139>：" + err.Error())
	}

	// <LABEL_735>
	go func(documentId string) {
		services.DocIndexService.ForceDelDocIdIndex(documentId)
	}(documentId)

	this.InfoLog("<LABEL_1138> " + documentId + " <LABEL_1617>")
	this.jsonSuccess("<LABEL_736>", "", "/document/index?document_id="+document["parent_id"])
}
