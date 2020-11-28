package controller

import (
	"bufio"
	"camera-server-backend/model"
	"camera-server-backend/pkgs/store"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// Controller ...
type Controller struct {
	store *store.Store
}

// NewController ...
func NewController(s *store.Store) *Controller {
	return &Controller{store: s}
}

const (
	eventImageDir = "./images"
)

// SaveEvents ...
func (c *Controller) SaveEvents(events []model.Event) error {
	c.store.InsertEvents(events)
	log.Info("Save Events Done")
	return nil
}

// GetEvents ...
func (c *Controller) GetEvents() ([]model.Event, error) {
	response, err := c.store.GetAllEvents()
	if err != nil {
		log.Error("Cannot events from database")
		return nil, err
	}
	log.Info("Load Events Done")
	return response, nil
}

// GetImages ...
func (c *Controller) GetImages(vectorIDs []string) ([]*model.ImageInfo, error) {
	if len(vectorIDs) == 0 {
		return nil, fmt.Errorf("Not found")
	}

	mapImageTemp := make(map[string]string)
	mapImageTemp["0"] = "./file_image/943_7.jpg"
	mapImageTemp["1"] = "./file_image/943_42.jpg"
	mapImageTemp["2"] = "./file_image/943_1.jpg"
	mapImageTemp["3"] = "./file_image/962_7.jpg"
	mapImageTemp["4"] = "./file_image/962_1.jpg"
	response := []*model.ImageInfo{}
	for _, id := range vectorIDs {
		if path, ok := mapImageTemp[id]; ok {
			baseCode, err := getImageFromFilePath(path)
			if err != nil {
				log.Error("Controller.GetImages: err get image base 64 = ", err)
			}

			response = append(response, &model.ImageInfo{
				ImageId:     id,
				ImageBase64: baseCode,
			})
		}
	}

	log.Info("Load Images Done")
	return response, nil
}

func getImageFromFilePath(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Error("getImageFromFilePath: read file path err = ", err)
		return "", err
	}

	defer f.Close()
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Error("getImageFromFilePath: read all file err = ", err)
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded, err
}
