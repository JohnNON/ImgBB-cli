package uploader

import (
	"strconv"

	imgBB "github.com/JohnNON/ImgBB"
)

type outData struct {
	Source     string  `json:"source" yaml:"source"`
	ID         string  `json:"id" yaml:"id"`
	Title      string  `json:"title" yaml:"title"`
	URLViewer  string  `json:"url_viewer" yaml:"url_viewer"`
	URL        string  `json:"url" yaml:"url"`
	DisplayURL string  `json:"display_url" yaml:"display_url"`
	Width      int     `json:"width" yaml:"width"`
	Height     int     `json:"height" yaml:"height"`
	Size       int     `json:"size" yaml:"size"`
	Time       int64   `json:"time" yaml:"time"`
	TTL        int64   `json:"expiration" yaml:"expiration"`
	Image      outInfo `json:"image" yaml:"image"`
	Thumb      outInfo `json:"thumb" yaml:"thumb"`
	Medium     outInfo `json:"medium" yaml:"medium"`
	DeleteURL  string  `json:"delete_url" yaml:"delete_url"`
}

type outInfo struct {
	Filename  string `json:"filename" yaml:"filename"`
	Name      string `json:"name" yaml:"name"`
	Mime      string `json:"mime" yaml:"mime"`
	Extension string `json:"extension" yaml:"extension"`
	URL       string `json:"url" yaml:"url"`
}

func (od outData) titles() []string {
	return []string{
		"source",
		"id",
		"title",
		"url_viewer",
		"url",
		"display_url",
		"width",
		"height",
		"size",
		"time",
		"expiration",
		"image_filename",
		"image_name",
		"image_mime",
		"image_extension",
		"image_url",
		"thumb_filename",
		"thumb_name",
		"thumb_mime",
		"thumb_extension",
		"thumb_url",
		"medium_filename",
		"medium_name",
		"medium_mime",
		"medium_extension",
		"medium_url",
		"delete_url",
	}
}

func (od outData) toStrings() []string {
	return []string{
		od.Source,
		od.ID,
		od.Title,
		od.URLViewer,
		od.URL,
		od.DisplayURL,
		strconv.Itoa(od.Width),
		strconv.Itoa(od.Height),
		strconv.Itoa(od.Size),
		strconv.FormatInt(od.Time, 10),
		strconv.FormatInt(od.TTL, 10),
		od.Image.Filename,
		od.Image.Name,
		od.Image.Mime,
		od.Image.Extension,
		od.Image.URL,
		od.Thumb.Filename,
		od.Thumb.Name,
		od.Thumb.Mime,
		od.Thumb.Extension,
		od.Thumb.URL,
		od.Medium.Filename,
		od.Medium.Name,
		od.Medium.Mime,
		od.Medium.Extension,
		od.Medium.URL,
		od.DeleteURL,
	}
}

func convertToOutData(source string, resp imgBB.Response) outData {
	return outData{
		Source:     source,
		ID:         resp.Data.ID,
		Title:      resp.Data.Title,
		URLViewer:  resp.Data.URLViewer,
		URL:        resp.Data.URL,
		DisplayURL: resp.Data.DisplayURL,
		Width:      resp.Data.Width,
		Height:     resp.Data.Height,
		Size:       resp.Data.Size,
		Time:       resp.Data.Time,
		TTL:        resp.Data.TTL,
		Image: outInfo{
			Filename:  resp.Data.Image.Filename,
			Name:      resp.Data.Image.Name,
			Mime:      resp.Data.Image.Mime,
			Extension: resp.Data.Image.Extension,
			URL:       resp.Data.Image.URL,
		},
		Thumb: outInfo{
			Filename:  resp.Data.Thumb.Filename,
			Name:      resp.Data.Thumb.Name,
			Mime:      resp.Data.Thumb.Mime,
			Extension: resp.Data.Thumb.Extension,
			URL:       resp.Data.Thumb.URL,
		},
		Medium: outInfo{
			Filename:  resp.Data.Medium.Filename,
			Name:      resp.Data.Medium.Name,
			Mime:      resp.Data.Medium.Mime,
			Extension: resp.Data.Medium.Extension,
			URL:       resp.Data.Medium.URL,
		},
		DeleteURL: resp.Data.DeleteURL,
	}
}
