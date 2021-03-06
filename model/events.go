package model

// Event ...
type Event struct {
	ID         uint   `gorm:"primaryKey"`
	QueryImage string `gorm:"Column:query_image"`
	Timestamp  string `gorm:"Column:time_stamp"`

	CameraID  string `gorm:"Column:camera_id"`
	VectorIDs string `gorm:"Column:vector_ids"`
}

// ImageInfo ...
type ImageInfo struct {
	ImageId     string `json:"vector_ids"`
	ImageBase64 string `json:"image_base64"`
	PersonId    string `json:"person_id"`
}

// ImagePath ...
type ImagePath struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Path  string `json:"path"`
}
