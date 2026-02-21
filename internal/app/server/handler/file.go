package handler

import (
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
	"github.com/sssjuing/media-vault/internal/pkg/oss"
)

func (h *Handler) ListFiles(c *echo.Context) error {
	folderName := c.Param("folder_name")
	if folderName != "" {
		folderName = folderName + "/"
	}
	result, err := oss.GetFiles(folderName)
	if err != nil {
		return c.String(http.StatusBadGateway, "Error")
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) UploadImage(c *echo.Context) error {
	file, _ := c.FormFile("file")
	fileObj, err := file.Open()
	contentType := file.Header.Get("content-type")
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	objectPath := filepath.Join("pictures/covers", file.Filename)
	info, err := oss.UploadFile(&fileObj, objectPath, contentType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, info)
}

type uploadLocalFileRequest struct {
	TargetDir string `json:"target_dir" validate:"required"`
	FilePath  string `json:"file_path" validate:"required"`
}

func (h *Handler) UploadLocalFile(c *echo.Context) error {
	req := &uploadLocalFileRequest{}
	if code, err := validateRequest(c, req); err != nil {
		return c.JSON(code, utils.NewError(err))
	}
	filename := filepath.Base(req.FilePath)
	objectPath := filepath.Join(req.TargetDir, filename)
	info, err := oss.UploadLocalFile(c.Request().Context(), objectPath, req.FilePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, info)
}
