package models

// TypeUpload upload type
type TypeUpload int

const (
	// UploadUnknown unknown type
	UploadUnknown TypeUpload = iota
	// UploadImageType image type
	UploadImageType
	// UploadVideoType video type
	UploadVideoType
	// UploadFileType file type
	UploadFileType
)

// File file model
type File struct {
	Model
	Name                string     `json:"name"`
	TypeUpload          TypeUpload `json:"type_upload"`
	IsSecure            bool       `json:"is_secure"`
	FileSize            int64      `json:"file_size"`
	ThumbnailImageURL   string     `gorm:"-" json:"thumbnail_image_url,omitempty"`
	OriginalImageURL    string     `gorm:"-" json:"original_image_url,omitempty"`
	FileURL             string     `gorm:"-" json:"file_url,omitempty"`
	IsFromAnotherSource bool       `jsoN:"is_from_another_source"`
	ThumbnailPath       string     `json:"-"`
	OriginalPath        string     `json:"-"`
	FilePath            string     `json:"-"`
	UUID                string     `json:"-"`
	CreatedByUserID     uint       `json:"created_by_user_id"`
}
