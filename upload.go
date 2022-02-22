// Wrapper around the 'upload' endpoint

package telegraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

var baseUploadURL = "https://telegra.ph/upload"

type UploadResult struct {
	Source []Source
}

type Source struct {
	Src string `json:"src"`
}

type Error struct {
	Error string `json:"error"`
}

// Upload photo/video to Telegra.ph on the '/upload' endpoint.
// Media type should either be "video" or "photo". "Animation" is considered "video" here.
func Upload(f io.Reader, mediaType string) (string, error) {
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	var name string
	if mediaType == "video" {
		name = "file.mp4"
	} else {
		name = "file.jpg"
	}
	part, err := w.CreateFormFile(mediaType, name)
	if err != nil {
		return "", err
	}
	io.Copy(part, f)
	w.Close()
	r, err := http.NewRequest("POST", baseUploadURL, bytes.NewReader(b.Bytes()))
	if err != nil {
		return "", err
	}
	r.Header.Set("Content-Type", w.FormDataContentType())
	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	var jsonData UploadResult
	json.Unmarshal(content, &jsonData.Source)
	if jsonData.Source == nil {
		var err Error
		json.Unmarshal(content, &err)
		return "", fmt.Errorf(err.Error)
	}
	return jsonData.Source[0].Src, err
}
