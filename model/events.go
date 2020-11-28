package model

// Event ...
type Event struct {
	QueryImage string   `json:"query_image" binding:"required"`
	Timestamp  string   `json:"timestamp" binding:"required"`
	CameraID   string   `json:"camera_id" binding:"required"`
	VectorIds  []string `json:"vector_ids" binding:"required"`
}
