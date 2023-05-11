package model

type File struct {
	Id              int    `json:"id"`
	Filename        string `json:"filename" gorm:"type:varchar(255)"`
	Description     string `json:"description" gorm:"type:varchar(255)"`
	Uploader        string `json:"uploader" gorm:"type:varchar(255)"`
	Link            string `json:"link" gorm:"type:varchar(255) unique"`
	Time            string `json:"time" gorm:"type:varchar(255)"`
	DownloadCounter int    `json:"download_counter" gorm:"type:int"`
	IsLocalFile     bool   `json:"is_local_file" gorm:"type:int"`
}
