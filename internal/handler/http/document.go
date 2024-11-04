package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Dolald/testwork_astral/configs"
	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/Dolald/testwork_astral/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDocument(c *gin.Context) {
	err := c.Request.ParseMultipartForm(configs.MaxByteForInputting << 20)
	if err != nil {
		h.logger.Errorf("ParseMultipartForm failed: %w", err)
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		h.logger.Errorf("getUserId failed: %w", err)
		return
	}

	file, err := c.FormFile(configs.FormName)
	if err != nil {
		h.logger.Errorf("FormFile failed: %w", err)
		return
	}

	fileLink := fmt.Sprintf(configs.Url + file.Filename)

	document := domain.Document{
		User_id:    userId,
		Filename:   file.Filename,
		Url:        fileLink,
		Created_at: time.Now(),
	}

	documentId, err := h.service.Document.CreateDocument(userId, document)
	if err != nil {
		h.logger.Errorf("CreateDocument failed: %w", err)
		return
	}

	if err := c.SaveUploadedFile(file, configs.UploaderFolder+file.Filename); err != nil {
		h.logger.Errorf("SaveUploadedFile failed: %w", err)
		return
	}

	c.String(http.StatusOK, configs.SucsessUpload+fileLink, documentId)
}

func (h *Handler) getAllDocuments(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		h.logger.Errorf("getUserId failed: %w", err)
		return
	}

	var filteredDocuments models.Filters

	err = c.BindJSON(&filteredDocuments)
	if err != nil {
		h.logger.Errorf("BindJSON failed: %w", err)
		return
	}

	allDocuments, err := h.service.Document.GetAllDocuments(userId, filteredDocuments)
	if err != nil {
		h.logger.Errorf("GetAllDocuments failed: %w", err)
		return
	}

	c.JSON(http.StatusOK, allDocuments)
}

func (h *Handler) getDocumentById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		h.logger.Errorf("getUserId failed: %w", err)
		return
	}

	documentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("strconv.Atoi failed: %w", err)
		return
	}

	document, err := h.service.Document.GetDocumentById(userId, documentId)
	if err != nil {
		h.logger.Errorf("GetDocumentById failed: %w", err)
		return
	}

	c.JSON(http.StatusOK, document)
}

func (h *Handler) deleteDocument(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		h.logger.Errorf("getUserId failed: %w", err)
		return
	}

	documentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("strconv.Atoi failed: %w", err)
		return
	}

	err = h.service.Document.DeleteDocument(userId, documentId)
	if err != nil {
		h.logger.Errorf("DeleteDocument failed: %w", err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
