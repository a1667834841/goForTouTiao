package api

import (
	"testing"

	config "github.com/goForTouTiao/config"
)

func TestPostWeiTouTiao(t *testing.T) {
	PostWeiTouTiao("你好再见！", config.GetConfig().Cookie, make([]string, 0))
}

func TestPostWeiPic(t *testing.T) {
	PostWeiPic("C:\\Users\\16678\\Pictures\\test.jpg", config.GetConfig().Cookie)
}

// 发布文字和图片
func TestPostWeiWithPic(t *testing.T) {
	img, _ := PostWeiPic("C:\\Users\\16678\\Pictures\\test.jpg", config.GetConfig().Cookie)

	imgs := make([]string, 1)
	imgs[0] = img

	PostWeiTouTiao("你好再见！", config.GetConfig().Cookie, imgs)

}

func TestGetHotQuestion(t *testing.T) {
	GetHotQuestion(config.GetConfig().Cookie)
}

func TestPublishAnswer(t *testing.T) {
	publishAnswer("7173500486790725918", "testest111111111111111111111111111111111111111111111111111111", config.GetConfig().Cookie, "hot_question")
}
