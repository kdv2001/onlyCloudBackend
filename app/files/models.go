package files

import "time"

type FileType string

type GetFile struct {
	ID        string     `json:"ID"`
	Parent    string     `json:"parent"`
	Type      FileType   `json:"type"`
	Directory *Directory `json:"directory,omitempty"`
	File      *File      `json:"file,omitempty"`
}

type DownloadFile struct {
	FileInfo    *File  `json:"fileInfo"`
	DownloadURL string `json:"downloadURL"`
}

type UploadFile struct {
	UploadURL string `json:"uploadURL"`
}

type Directory struct {
	Name         string    `json:"name"`
	Content      []string  `json:"content"`
	Owner        string    `json:"owner"`
	CreationDate time.Time `json:"creationDate"`
}

type File struct {
	Name         string    `json:"name"`
	Size         uint64    `json:"size"`
	MimeType     string    `json:"mimeType"`
	Extension    string    `json:"extension"`
	Owner        string    `json:"owner"`
	CreationDate time.Time `json:"creationDate"`
	UpdateDate   time.Time `json:"updateDate"`
	AccessDate   time.Time `json:"accessDate"`
}
