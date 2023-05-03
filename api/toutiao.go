package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	s "github.com/goForTouTiao/structures"
)

var Hot_Question_url = "https://mp.toutiao.com/wenda/v5/question/hot"
var publish_answer_url = "https://mp.toutiao.com/mp/agw/wenda/publish_answer"

func PostWeiTouTiao(content, cookie string, imgUrls []string) s.ToutiaoRes {

	url := "https://mp.toutiao.com/mp/agw/article/wtt"
	method := "POST"

	imgUrlsJson, err := json.Marshal(imgUrls)
	if err != nil {
		panic(err)
	}

	payload := strings.NewReader(`{"content":"` + content + `","image_list":` + string(imgUrlsJson) + `,"is_fans_article":2,"pre_upload":1,"welfare_card":""}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	req.Header.Add("authority", "mp.toutiao.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", cookie)
	req.Header.Add("origin", "https://mp.toutiao.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://mp.toutiao.com/profile_v4/weitoutiao/publish")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// fmt.Println(string(body))
	var result s.ToutiaoRes
	result = s.DecodeToutiao(body, result)
	if !result.Success() {
		panic("上传失败")
	} else {
		log.Print("上传成功！")
	}
	return result
}

func PostWeiPic(filePath, cookie string) (string, error) {

	url := "https://mp.toutiao.com/mp/agw/article_material/photo/upload_picture?type=ueditor&pgc_watermark=1&action=uploadimage&encode=utf-8&is_private=1"
	method := "POST"

	//打开要上传的文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	defer file.Close()
	fileBuffer := &bytes.Buffer{}
	//创建一个multipart类型的写文件
	writer := multipart.NewWriter(fileBuffer)
	//使用给出的属性名paramName和文件名filePath创建一个新的form-data头
	part, err := writer.CreateFormFile("upfile", filePath)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	//将源复制到目标，将file写入到part   是按默认的缓冲区32k循环操作的，不会将内容一次性全写入内存中,这样就能解决大文件的问题
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		fmt.Println(" post err=", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, fileBuffer)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("authority", "mp.toutiao.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Add("cache-control", "no-cache")
	// req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary6WYpEbI3b55boJls")
	req.Header.Add("cookie", cookie)
	req.Header.Add("origin", "https://mp.toutiao.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://mp.toutiao.com/profile_v4/weitoutiao/publish")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer func() {
		res.Body.Close()
		fmt.Println("finish")
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var result s.UploadPictureRes
	result = s.DecodeUploadPicture(body, result)

	if !result.Success() {
		panic(filePath + "，图片上传失败")
	} else {
		log.Printf("%s，图片上传成功！", filePath)
	}

	// jsonResp, err := json.Marshal(res.Body)
	// if err != nil {
	// 	log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// }
	// fmt.Print(jsonResp["origin_web_url"])

	// fmt.Println(string(body))
	fmt.Println(result.OriginWebURL)

	return result.OriginWebURI, nil
}

func GetHotQuestion(cookie string) {

	client := &http.Client{}
	reader := strings.NewReader("{}")
	req, err := http.NewRequest("GET", Hot_Question_url, reader)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	req.Header.Add("authority", "mp.toutiao.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", cookie)
	req.Header.Add("origin", "https://mp.toutiao.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://mp.toutiao.com/profile_v4/weitoutiao/publish")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(body))

}

func publishAnswer(qid, answer, cookie, source string) {

	method := "POST"

	form := url.Values{
		"content":       {"<p>" + answer + "</p>"},
		"qid":           {qid},
		"praise":        {"0"},
		"origin_source": {"mp_wenda_tab"},
		"source":        {source},
		"is_original":   {"0"},
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, publish_answer_url, strings.NewReader(form.Encode()))

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	req.Header.Add("authority", "mp.toutiao.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cookie", cookie)
	req.Header.Add("origin", "https://mp.toutiao.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://mp.toutiao.com/profile_v4/weitoutiao/publish")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(body))
}

// func main() {
// 	// PostWeiTouTiao("test", config.GetConfig().Cookie)
// 	PostWeiPic("C:\\Users\\16678\\Pictures\\test.jpg")
// }
