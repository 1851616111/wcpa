package api

import (
	"encoding/json"
	"github.com/1851616111/util/http"
	"github.com/1851616111/wcpa/pkg_old/errors"
	"io/ioutil"
	"fmt"
	"bytes"
)

const (
	NewMenuURL = "https://api.weixin.qq.com/cgi-bin/menu/create"
	SelfMenuURL = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info"
)

func ListMenu(access_token string) ([]Button, error) {
	//rsp, err := http.Send(&http.HttpSpec{
	//	URL: SelfMenuURL,
	//	Method:"GET",
	//	ContentType: http.ContentType_JSON,
	//	URLParams:   http.NewParams().Add("access_token", access_token),
	//})
	//
	//var res struct {
	//	IsMenuOpen int `json:"is_menu_open"`
	//	SelfMenuInfo struct {
	//		Button []struct{
	//			Name string `json:"name"`
	//			SubButton struct {
	//				List []struct {
	//					Type string `json:"type"`
	//					Name string `json:"name"`
	//					Url  string `json:"url"`
	//					Value string `json:"value"`
	//				} `json:"list"`
	//			} `json:"sub_button"`
	//		} `json:"button"`
	//	} `json:"selfmenu_info"`
	//}

return nil, nil
}

func CreateMenu(access_token string, bt ...*Button) error {
	req := &http.HttpSpec{
		URL:         NewMenuURL,
		Method:      "POST",
		ContentType: http.ContentType_JSON,
		URLParams:   http.NewParams().Add("access_token", access_token),
		BodyParams:  *http.NewBody().Add("button", bt),
	}

	rsp, err := http.Send(req)
	if err != nil {
		return err
	}

	data ,_ := ioutil.ReadAll(rsp.Body)
	fmt.Printf("%s", string(data))

	tmp := errors.Error{}
	if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&tmp); err != nil {
		return err
	}

	if tmp.Code == errors.CODE_SUCCESS {
		return nil
	} else {
		return tmp.Error()
	}
}

type ButtonType int

const (
	Click              = iota
	View
	ScanCode_Push
	ScanCode_WaitMsg
	Pic_SysPhoto
	Pic_Photo_or_album
	Pic_WeiXin
	Location_Select
	Media_Id
	View_Limited
)

var buttonTypes []string = []string{
	Click:           "click", View: "view", ScanCode_Push: "scancode_push", ScanCode_WaitMsg: "scancode_waitmsg",
	Pic_SysPhoto:    "pic_sysphoto", Pic_Photo_or_album: "pic_photo_or_album", Pic_WeiXin: "pic_weixin",
	Location_Select: "location_select", Media_Id: "media_id", View_Limited: "view_limited"}

type Button struct {
	Type      string       `json:"type,omitempty"`
	Name      string       `json:"name"`
	Key       string       `json:"key,omitempty"`
	SubButton *[]SubButton `json:"sub_button,omitempty"`
}

type SubButton struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	Key     string `json:"key,omitempty"`
	MediaId string `json:"media_id,omitempty"`
	AppId   string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`

}

func NewViewButton(name, url string) SubButton {
	return SubButton{
		Type: buttonTypes[View],
		Name: name,
		URL:  url,
	}
}

func NewViewLimitedButton(name, mediaId string) SubButton {
	return SubButton{
		Type: buttonTypes[View_Limited],
		Name: name,
		MediaId:  mediaId,
	}
}

func NewMediaButton(name, mediaId string) SubButton {
	return SubButton{
		Type:    buttonTypes[Media_Id],
		Name:    name,
		MediaId: mediaId,
	}
}


func NewMiniProgramButton(appId, name, url, pagePath string) SubButton {
	return SubButton{
		Type: "miniprogram",
		Name: name,
		URL: url,
		AppId: appId,
		PagePath: pagePath,
	}
}

func NewTopButton(name string) *Button {
	subs := []SubButton{}
	return &Button{
		Name:      name,
		SubButton: &subs,
	}
}

func (b *Button) AddSub(sub SubButton) *Button {
	*b.SubButton = append(*b.SubButton, sub)
	return b
}
