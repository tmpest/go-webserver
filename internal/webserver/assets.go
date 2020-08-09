package webserver

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Path in: /public/assets/path/in/s3
// Return content in bucket/path/in/s3 or 404 Not Found

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	s3Path := r.URL.Path[len("/public/assets/"):]
	log.Printf("received request for asset %s", s3Path)

	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	if err != nil {
		log.Printf("failed to establish connection with AWS: %s", err.Error())
		http.Error(w, fmt.Sprintf("failed to establish connection with AWS: %s", err.Error()), http.StatusFailedDependency)
		return
	}

	result, err := s3.New(session).GetObject(&s3.GetObjectInput{
		Bucket: aws.String("tmpest-website-assets"),
		Key:    aws.String(s3Path),
	})
	if err != nil {
		log.Printf("failed to download asset: %s", err.Error())
		http.Error(w, fmt.Sprintf("failed to download asset: %s", err.Error()), http.StatusNotFound)
		return
	}
	log.Print("found content in remote storage")

	written, err := io.Copy(w, result.Body)
	if err != nil {
		log.Printf("Error copying asset to the http response %s", err.Error())
		http.Error(w, fmt.Sprintf("Error copying asset to the http response %s", err.Error()), http.StatusInternalServerError)
		return
	}
	log.Printf("download complete - %v bytes downloaded", written)
}
