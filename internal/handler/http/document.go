package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"web-cache/configs"
	webСache "web-cache/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createDocument(c *gin.Context) {
	err := c.Request.ParseMultipartForm(configs.MaxByteForInputting << 20)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, err := c.FormFile(configs.FormName)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fileLink := fmt.Sprintf(configs.Url + file.Filename)

	document := webСache.Document{
		User_id:    userId,
		Filename:   file.Filename,
		Url:        fileLink,
		Created_at: time.Now(),
	}

	documentId, err := h.service.Document.CreateDocument(userId, document)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.SaveUploadedFile(file, configs.UploaderFolder+file.Filename); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, configs.SucsessUpload+fileLink, documentId)
}

type getAllListResponse struct {
	Data []webСache.Document `json:"data"`
}

func (h *Handler) getAllDocuments(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	allDocuments, err := h.service.Document.GetAllDocuments(userId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: allDocuments,
	})
}

func (h *Handler) getDocumentById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	documentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	document, err := h.service.Document.GetById(userId, documentId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	c.JSON(http.StatusOK, document)
}

func (h *Handler) deleteDocument(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.Document.DeleteDocument(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
