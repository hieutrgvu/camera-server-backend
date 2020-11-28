package model

// Event ...
type Event struct {
	QueryImage string `gorm:"Column:query_image"`
	Timestamp  string `gorm:"Column:time_stamp"`
	CameraID   string `gorm:"Column:camera_id"`
	VectorIDs  string `gorm:"Column:vector_ids"`
}

// ImageInfo ...
type ImageInfo struct {
	ImageId     string `json:"vector_ids"`
	ImageBase64 string `json:"image_base64"`
}
