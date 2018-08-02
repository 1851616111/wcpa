package api

import (
	"io"
	"github.com/1851616111/util/http/file"
	"fmt"
	"encoding/json"
)

type UploadResponse struct {
	Type     string `json:"type"`
	MediaId  string `json:"media_id"`
	CreateAt int    `json:"created_at"`
	Error
}

func UploadMediaReader(access_token, tp, fileName string, reader io.Reader) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", access_token, tp)
	rsp, err := file.PostFile(url, "media", fileName, reader)
	if err != nil {
		return "", err
	}

	var res UploadResponse
	if err := json.NewDecoder(rsp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res.MediaId, res.Error.Error()
}
