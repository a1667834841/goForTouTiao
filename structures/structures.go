package structures

import (
	"encoding/json"
)

type UploadPictureRes struct {
	Code         int    `json:"code"`
	Height       int    `json:"height"`
	ImageType    int    `json:"image_type"`
	Message      string `json:"message"`
	MimeType     string `json:"mime_type"`
	OriginWebURI string `json:"origin_web_uri"`
	OriginWebURL string `json:"origin_web_url"`
	Original     string `json:"original"`
	State        string `json:"state"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	WebURI       string `json:"web_uri"`
	WebURL       string `json:"web_url"`
	Width        int    `json:"width"`
}

type ToutiaoRes struct {
	Code int `json:"code"`
	Data struct {
		AddThreadLimit int   `json:"add_thread_limit"`
		ThreadID       int64 `json:"thread_id"`
	} `json:"data"`
	Message string `json:"message"`
}

func (res *ToutiaoRes) Success() bool {
	return res.Code == 0
}

func (res *UploadPictureRes) Success() bool {
	return res.Code == 0
}

func DecodeToutiao(body []byte, res ToutiaoRes) ToutiaoRes {
	_ = json.Unmarshal(body, &res)
	return res
}

func DecodeUploadPicture(body []byte, res UploadPictureRes) UploadPictureRes {
	_ = json.Unmarshal(body, &res)
	return res
}
