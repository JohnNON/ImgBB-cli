# ImgBB-cli

ImgB-cli is a console line tool that can upload given images on imgbb.com.

Supported JPG, PNG, BMP, GIF, TIFF, WEBP, HEIC, and PDF.

## Manual

### Usage

```bash
imgbb-cli [command] [flags]
```

### Commands

* **help** - Help about any command
* **upload** - Uploads given images to imgbb.com

### Flags

* -h, --help - help for imgbb-cli
* -k, --key - imgbb api key
* -f, --format - formatted out (supported: text, json, yaml, csv. default text)
* -c, --comma - csv column separator (default , - comma) (default ",")
* -g, --generate - generate image name (default false - use file name)
* -i, --images - list of images paths, urls or base64 data
* -p, --paths - paths with images for upload
* -r, --recursive - upload recursive subdirectories (default false - skip subdirectories)
* -t, --ttl - time to live in seconds for images (default 0 - unlimited)
* -v, --version - version for imgbb-cli
