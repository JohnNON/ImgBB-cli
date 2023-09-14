package uploader

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	imgBB "github.com/JohnNON/ImgBB"
)

const (
	csvOut  = "csv"
	jsonOut = "json"
	yamlOut = "yaml"
)

var supportedFormats = []string{".jpeg", ".jpg", ".png", ".bmp", ".gif", ".tiff", ".webp", ".heic", ".pdf"}

type Option func(u *Uploader)

func WithTTL(ttl uint64) Option {
	return func(u *Uploader) {
		u.ttl = ttl
	}
}

func WithImageSources(imageSources []string) Option {
	return func(u *Uploader) {
		u.imageSources = imageSources
	}
}
func WithPaths(paths []string) Option {
	return func(u *Uploader) {
		u.paths = paths
	}
}

func WithGenerate(generate bool) Option {
	return func(u *Uploader) {
		u.generate = generate
	}
}

func WithRecursive(recursive bool) Option {
	return func(u *Uploader) {
		u.recursive = recursive
	}
}
func WithFormatOut(outFormat string, comma string) Option {
	return func(u *Uploader) {
		u.out = initOutWriter(outFormat, comma)
	}
}

type writer interface {
	write(od outData)
}

type Uploader struct {
	waitGroup    *sync.WaitGroup
	imgBBClient  *imgBB.Client
	ttl          uint64
	imageSources []string
	paths        []string
	generate     bool
	recursive    bool
	out          writer
}

func New(apiKey string, opts ...Option) *Uploader {
	u := &Uploader{
		waitGroup:   &sync.WaitGroup{},
		imgBBClient: imgBB.NewClient(&http.Client{}, apiKey),
	}

	for _, o := range opts {
		o(u)
	}

	return u
}

func (u *Uploader) Do(ctx context.Context) error {
	u.waitGroup.Add(2)

	go func() {
		defer u.waitGroup.Done()

		u.uploadFromImageSources(ctx)
	}()

	go func() {
		defer u.waitGroup.Done()

		u.uploadFromPath(ctx, u.paths)
	}()

	u.waitGroup.Wait()

	return nil
}

func (u *Uploader) uploadFromImageSources(ctx context.Context) {
	u.waitGroup.Add(len(u.imageSources))

	for _, source := range u.imageSources {
		go func(source string) {
			defer u.waitGroup.Done()

			img, err := u.createImg(source)
			if err != nil {
				printErr(source, err)

				return
			}

			resp, err := u.imgBBClient.Upload(ctx, img)
			if err != nil {
				printErr(source, err)

				return
			}

			u.out.write(convertToOutData(source, resp))
		}(source)
	}
}

func (u *Uploader) uploadFromPath(ctx context.Context, paths []string) {
	u.waitGroup.Add(len(paths))

	for _, path := range paths {
		go func(path string) {
			defer u.waitGroup.Done()

			dirEntry, err := os.ReadDir(path)
			if err != nil {
				printErr(path, err)

				return
			}

			for _, entry := range dirEntry {
				if entry.IsDir() {
					if !u.recursive {
						continue
					}

					u.uploadFromPath(ctx, []string{filepath.Join(path, entry.Name())})
				}

				if !checkFileFormat(entry.Name()) {
					continue
				}

				u.waitGroup.Add(1)
				go func(source string) {
					defer u.waitGroup.Done()

					img, err := u.createImg(source)
					if err != nil {
						printErr(source, err)

						return
					}

					resp, err := u.imgBBClient.Upload(ctx, img)
					if err != nil {
						printErr(source, err)

						return
					}

					u.out.write(convertToOutData(source, resp))
				}(filepath.Join(path, entry.Name()))
			}
		}(path)
	}
}

func checkFileFormat(file string) bool {
	return slices.Contains[[]string, string](supportedFormats, getFileExtension(file))
}

func getFileExtension(file string) string {
	info := strings.Split(file, ".")

	return fmt.Sprintf(".%s", strings.ToLower(info[len(info)-1]))
}

func (u *Uploader) createImg(source string) (*imgBB.Image, error) {
	file, err := os.Open(source)
	if err == nil {
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}

		name := filepath.Base(source)

		if u.generate {
			name = fmt.Sprintf("%s%s", hashSum(data), getFileExtension(name))
		}

		return imgBB.NewImageFromFile(name, u.ttl, data)
	}

	return imgBB.NewImage(hashSum([]byte(source)), u.ttl, source)
}

func hashSum(b []byte) string {
	sum := md5.Sum(b)

	return hex.EncodeToString(sum[:])
}

func initOutWriter(fornat string, comma string) writer {
	switch fornat {
	case csvOut:
		return newCsvWriter([]rune(comma)[0])
	case jsonOut:
		return newJsonWriter()
	case yamlOut:
		return newYamlWriter()
	default:
		return newTextWriter()
	}
}

func printErr(source string, err error) {
	fmt.Fprintf(os.Stderr, "Error %s: %s\n", source, err)
}
