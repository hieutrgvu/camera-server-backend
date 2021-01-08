package helper

import (
	"camera-server-backend/model"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

var dataImagePath = "[{\"id\":\"0\",\"label\":\"0\",\"path\":\"1m_face/0/0.jpg\"},{\"id\":\"1\",\"label\":\"0\",\"path\":\"1m_face/0/1.jpg\"},{\"id\":\"2\",\"label\":\"0\",\"path\":\"1m_face/0/2.jpg\"},{\"id\":\"3\",\"label\":\"0\",\"path\":\"1m_face/0/3.jpg\"},{\"id\":\"4\",\"label\":\"0\",\"path\":\"1m_face/0/4.jpg\"},{\"id\":\"5\",\"label\":\"0\",\"path\":\"1m_face/0/5.jpg\"},{\"id\":\"6\",\"label\":\"0\",\"path\":\"1m_face/0/6.jpg\"},{\"id\":\"7\",\"label\":\"0\",\"path\":\"1m_face/0/7.jpg\"},{\"id\":\"8\",\"label\":\"0\",\"path\":\"1m_face/0/8.jpg\"},{\"id\":\"9\",\"label\":\"0\",\"path\":\"1m_face/0/9.jpg\"},{\"id\":\"10\",\"label\":\"0\",\"path\":\"1m_face/0/10.jpg\"},{\"id\":\"11\",\"label\":\"0\",\"path\":\"1m_face/0/11.jpg\"},{\"id\":\"12\",\"label\":\"0\",\"path\":\"1m_face/0/12.jpg\"},{\"id\":\"13\",\"label\":\"0\",\"path\":\"1m_face/0/13.jpg\"},{\"id\":\"14\",\"label\":\"0\",\"path\":\"1m_face/0/14.jpg\"},{\"id\":\"15\",\"label\":\"0\",\"path\":\"1m_face/0/15.jpg\"},{\"id\":\"16\",\"label\":\"0\",\"path\":\"1m_face/0/16.jpg\"},{\"id\":\"17\",\"label\":\"0\",\"path\":\"1m_face/0/17.jpg\"},{\"id\":\"18\",\"label\":\"0\",\"path\":\"1m_face/0/18.jpg\"},{\"id\":\"19\",\"label\":\"0\",\"path\":\"1m_face/0/19.jpg\"},{\"id\":\"20\",\"label\":\"0\",\"path\":\"1m_face/0/20.jpg\"},{\"id\":\"21\",\"label\":\"0\",\"path\":\"1m_face/0/21.jpg\"},{\"id\":\"22\",\"label\":\"0\",\"path\":\"1m_face/0/22.jpg\"},{\"id\":\"23\",\"label\":\"0\",\"path\":\"1m_face/0/23.jpg\"},{\"id\":\"24\",\"label\":\"0\",\"path\":\"1m_face/0/24.jpg\"},{\"id\":\"25\",\"label\":\"0\",\"path\":\"1m_face/0/25.jpg\"},{\"id\":\"26\",\"label\":\"0\",\"path\":\"1m_face/0/26.jpg\"},{\"id\":\"27\",\"label\":\"0\",\"path\":\"1m_face/0/27.jpg\"},{\"id\":\"28\",\"label\":\"0\",\"path\":\"1m_face/0/28.jpg\"},{\"id\":\"29\",\"label\":\"0\",\"path\":\"1m_face/0/29.jpg\"},{\"id\":\"30\",\"label\":\"0\",\"path\":\"1m_face/0/30.jpg\"},{\"id\":\"31\",\"label\":\"0\",\"path\":\"1m_face/0/31.jpg\"},{\"id\":\"32\",\"label\":\"0\",\"path\":\"1m_face/0/32.jpg\"},{\"id\":\"999999\",\"label\":\"khang\",\"path\":\"1m_face/khang/999999.jpg\"},{\"id\":\"1000000\",\"label\":\"khang\",\"path\":\"1m_face/khang/1000000.jpg\"},{\"id\":\"1000001\",\"label\":\"khang\",\"path\":\"1m_face/khang/1000001.jpg\"},{\"id\":\"1000002\",\"label\":\"dat\",\"path\":\"1m_face/dat/1000002.jpg\"},{\"id\":\"1000003\",\"label\":\"dat\",\"path\":\"1m_face/dat/1000003.jpg\"},{\"id\":\"1000004\",\"label\":\"hieu\",\"path\":\"1m_face/hieu/1000004.jpg\"},{\"id\":\"1000005\",\"label\":\"hieu\",\"path\":\"1m_face/hieu/1000005.jpg\"},{\"id\":\"1000006\",\"label\":\"hieu\",\"path\":\"1m_face/hieu/1000006.jpg\"},{\"id\":\"1000007\",\"label\":\"hoai\",\"path\":\"1m_face/hoai/1000007.jpg\"},{\"id\":\"1000008\",\"label\":\"hoai\",\"path\":\"1m_face/hoai/1000008.jpg\"},{\"id\":\"1000009\",\"label\":\"hoai\",\"path\":\"1m_face/hoai/1000009.jpg\"}]"
var MapImage = make(map[string]*model.ImagePath)

func InitDataImagePath() {
	listImagePath := []*model.ImagePath{}
	err := json.Unmarshal([]byte(dataImagePath), &listImagePath)
	if err != nil {
		log.Error("InitDataImagePath: err unmarshal data = ", err)
		return
	}

	for _, imgPathInfo := range listImagePath {
		imgPathTemp := &model.ImagePath{
			Id:    imgPathInfo.Id,
			Label: imgPathInfo.Label,
			Path:  "./file/image/" + imgPathInfo.Path,
		}

		MapImage[imgPathTemp.Id] = imgPathTemp
	}

	return
}
