package main

import (
	"context"
	"os"

	"github.com/JohnNON/ImgBB-cli/uploader"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Uploads given images to imgbb.com",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())

		if apiKey == "" {
			apiKey = os.Getenv(envKeyName)
		}

		opts := []uploader.Option{
			uploader.WithTTL(ttl),
			uploader.WithImageSources(images),
			uploader.WithPaths(paths),
			uploader.WithGenerate(generate),
			uploader.WithRecursive(recursive),
			uploader.WithFormatOut(outFormat, csvComma),
		}

		u := uploader.New(apiKey, opts...)

		err := u.Do(ctx)
		cobra.CheckErr(err)

		cancel()
	},
}
