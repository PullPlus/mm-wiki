package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/chaiyd/mm-wiki/app"
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/utils"
	"os"
	"path"
)

type UploadResponse struct {
	Success int    `json:"success"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

type ImageController struct {
	BaseController
}

func (this *ImageController) Upload() {

	if !this.IsPost() {
		this.ViewError("<LABEL_705>！", "/space/index")
	}
	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.jsonError("<LABEL_1144>！")
	}

	// handle document
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("<LABEL_717> " + documentId + " <LABEL_1618>：" + err.Error())
		this.jsonError("<LABEL_718>！")
	}
	if len(document) == 0 {
		this.jsonError("<LABEL_965>！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("<LABEL_1132> " + documentId + " <LABEL_719>：" + err.Error())
		this.jsonError("<LABEL_718>！")
	}
	if len(space) == 0 {
		this.jsonError("<LABEL_254>！")
	}
	// check space visit_level
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.jsonError("<LABEL_43>！")
	}

	// handle upload
	f, h, err := this.GetFile("editormd-image-file")
	if err != nil {
		this.ErrorLog("<LABEL_380>: " + err.Error())
		this.jsonError("<LABEL_380>")
		return
	}
	if h == nil || f == nil {
		this.ErrorLog("<LABEL_747>")
		this.jsonError("<LABEL_747>")
		return
	}
	_ = f.Close()

	// file save dir
	saveDir := fmt.Sprintf("%s/%s/%s", app.ImageAbsDir, spaceId, documentId)
	ok, _ := utils.File.PathIsExists(saveDir)
	if !ok {
		err := os.MkdirAll(saveDir, 0777)
		if err != nil {
			this.ErrorLog("<LABEL_747>: " + err.Error())
			this.jsonError("<LABEL_748>")
			return
		}
	}
	// check file is exists
	imageFile := path.Join(saveDir, h.Filename)
	ok, _ = utils.File.PathIsExists(imageFile)
	if ok {
		this.jsonError("<LABEL_381>！")
	}
	// save file
	err = this.SaveToFile("editormd-image-file", imageFile)
	if err != nil {
		this.ErrorLog("<LABEL_749>: " + err.Error())
		this.jsonError("<LABEL_749>")
	}

	// insert db
	attachment := map[string]interface{}{
		"user_id":     this.UserId,
		"document_id": documentId,
		"name":        h.Filename,
		"path":        fmt.Sprintf("images/%s/%s/%s", spaceId, documentId, h.Filename),
		"source":      models.Attachment_Source_Image,
	}
	_, err = models.AttachmentModel.Insert(attachment, spaceId)
	if err != nil {
		_ = os.Remove(imageFile)
		this.ErrorLog("<LABEL_198>: " + err.Error())
		this.jsonError("<LABEL_382>")
	}

	this.InfoLog(fmt.Sprintf("<LABEL_1619> %s <LABEL_1153> %s <LABEL_1617>", documentId, h.Filename))
	this.jsonSuccess("<LABEL_1154>", fmt.Sprintf("/%s", attachment["path"]))
}

func (this *ImageController) jsonError(message string) {

	uploadRes := UploadResponse{
		Success: 0,
		Message: message,
		Url:     "",
	}

	j, err := json.Marshal(uploadRes)
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Abort(string(j))
	}
}

func (this *ImageController) jsonSuccess(message string, url string) {

	uploadRes := UploadResponse{
		Success: 1,
		Message: message,
		Url:     url,
	}

	j, err := json.Marshal(uploadRes)
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Abort(string(j))
	}
}
