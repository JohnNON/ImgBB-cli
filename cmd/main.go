package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const envKeyName = "IMGBB_API_KEY"

var (
	buildTime string
	buildTag  string
	buildHash string

	apiKey string

	ttl uint64

	images []string

	paths []string

	generate bool

	recursive bool

	outFormat string

	csvComma string
)

var rootCmd = &cobra.Command{
	Use:     "imgbb-cli",
	Short:   "imgbb-cli is a tool, that can upload your images to imgbb.com",
	Version: version(),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func version() string {
	tag := buildHash

	if buildTag != "" {
		tag = buildTag
	}

	return fmt.Sprintf("%s %s", tag, buildTime)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&apiKey, "key", "k", "", "imgbb api key")
	rootCmd.PersistentFlags().Uint64VarP(&ttl, "ttl", "t", 0, "time to live in seconds for images (default 0 - unlimited)")
	rootCmd.PersistentFlags().StringArrayVarP(&images, "images", "i", nil, "list of images paths, urls or base64 data")
	rootCmd.PersistentFlags().StringArrayVarP(&paths, "paths", "p", nil, "paths with images for upload")
	rootCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate image name (default false - use file name)")
	rootCmd.PersistentFlags().BoolVarP(&recursive, "recursive", "r", false, "upload recursive subdirectories (default false - skip subdirectories)")
	rootCmd.PersistentFlags().StringVarP(&outFormat, "format", "f", "", "formatted out (supported: text, json, yaml, csv. default text)")
	rootCmd.PersistentFlags().StringVarP(&csvComma, "comma", "c", ",", "csv column separator (default , - comma)")
}

func main() {
	rootCmd.AddCommand(uploadCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
