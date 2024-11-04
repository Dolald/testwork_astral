package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Dolald/testwork_astral/configs"
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

	document := models.Document{
		User_id:    userId,
		Filename:   file.Filename,
		Url:        fileLink,
		Created_at: time.Now(),
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	documentId, err := h.service.Document.CreateDocument(ctx, userId, document)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	allDocuments, err := h.service.Document.GetAllDocuments(ctx, userId, filteredDocuments)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	document, err := h.service.Document.GetDocumentById(ctx, userId, documentId)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	err = h.service.Document.DeleteDocument(ctx, userId, documentId)
	if err != nil {
		h.logger.Errorf("DeleteDocument failed: %w", err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"asd": "asd",
	})
}
