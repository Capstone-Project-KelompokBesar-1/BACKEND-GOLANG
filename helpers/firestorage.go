package helpers

import (
	"context"
	"io"
	"log"
	"mime/multipart"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func UploadImage(photo *multipart.FileHeader, folder string) string {
	fileImage, _ := photo.Open()

	config := &firebase.Config{
		StorageBucket: "movoo-66c7d.appspot.com",
	}

	cntx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(cntx, config, opt)
	if err != nil {
		log.Println(err)
		return ""
	}

	client, err := app.Storage(cntx)
	if err != nil {
		log.Println(err)
		return ""
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Println(err)
		return ""
	}

	wc := bucket.Object(folder + "/" + photo.Filename).NewWriter(cntx)

	if _, err = io.Copy(wc, fileImage); err != nil {
		log.Println(err)
		return ""
	}
	if err := wc.Close(); err != nil {
		log.Println(err)
		return ""
	}

	url := "https://storage.cloud.google.com/movoo-66c7d.appspot.com/" + folder + "/" + photo.Filename

	return url
}
