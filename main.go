package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "testing-object.sg-sin1.upcloudobjects.com"
	accessKeyID := "UCOB0EW9K280U7U6CHUN"
	secretAccessKey := "kNxA5ABEWmQxk/kB8IOB7TYsDUhIU7pjydEsszk7"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Testing %#v\n", minioClient) // minioClient is now setup

	file, err := os.Open("cv_test-2_1.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	uploadInfo, err := minioClient.PutObject(context.Background(), "testing-bucket", file.Name(), file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
}
