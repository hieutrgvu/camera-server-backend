package model

import "mime/multipart"

// Event ...
type Event struct {
	QueryImage *multipart.FileHeader `form:"query_image" binding:"required"`
	Timestamp  string                `form:"timestamp" binding:"required"`
	CameraID   string                `form:"camera_id" binding:"required"`
	VectorIds  []string              `form:"vector_ids" binding:"required"`
}

// EventRecord ...
type EventRecord struct {
	QueryImagePath string
	Timestamp      string
	CameraID       string
	VectorIds      []string
}
