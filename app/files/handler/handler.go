package handler

import "github.com/gofiber/fiber/v2"

type FilesHandler struct{}

func NewFilesHandler() *FilesHandler {
	return &FilesHandler{}
}

// GetFile godoc
// @Summary Получение информации о файле
// @Tags files
// @Produce  json
// @Param id query string true "ID файла"
// @Success 200 {object} files.GetFile
// @Failure 400 {object} appErrors.AppError
// @Failure 401 {object} appErrors.AppError
// @Failure 500 {object} appErrors.AppError
// @Router /files [get]
func GetFile(c fiber.Ctx) error {
	return nil
}

// GetFileByPath godoc
// @Summary Получение информации о файле по пути
// @Tags files
// @Produce  json
// @Param path query string true "Путь до файла"
// @Success 200 {object} files.GetFile
// @Failure 400 {object} appErrors.AppError
// @Failure 401 {object} appErrors.AppError
// @Failure 500 {object} appErrors.AppError
// @Router /files/path [get]
func GetFileByPath(c fiber.Ctx) error {
	return nil
}

// DownloadFile godoc
// @Summary Скачивание файла
// @Tags files
// @Produce  json
// @Param id query string true "ID файла"
// @Success 200 {object} files.GetFile
// @Failure 400 {object} appErrors.AppError
// @Failure 401 {object} appErrors.AppError
// @Failure 500 {object} appErrors.AppError
// @Router /files/download [get]
func DownloadFile(c fiber.Ctx) error {
	return nil
}

// UploadFile godoc
// @Summary Загрузка файла
// @Tags files
// @Produce  json
// @Accept json
// @Param file body files.File true "Информация о файле"
// @Success 200 {object} files.UploadFile
// @Failure 400 {object} appErrors.AppError
// @Failure 401 {object} appErrors.AppError
// @Failure 500 {object} appErrors.AppError
// @Router /files/upload [post]
func UploadFile(c fiber.Ctx) error {
	return nil
}
