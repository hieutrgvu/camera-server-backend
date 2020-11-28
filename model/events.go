package model

// Event ...
type Event struct {
	QueryImage string   `json:"query_image" binding:"required"`
	Timestamp  string   `json:"timestamp" binding:"required"`
	CameraID   string   `json:"camera_id" binding:"required"`
	VectorIds  []string `json:"vector_ids" binding:"required"`
}

// ImageInfo ...
type ImageInfo struct {
	ImageId     string `json:"vector_ids"`
	ImageBase64 string `json:"image_base64"`
}
