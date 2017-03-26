package main

import (
	"os"
	"log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"io/ioutil"
)

var bucket string
var accessKey string
var secretKey string
var region string

func init() {
	initOk := true

	bucket = os.Getenv("BUCKET")
	if bucket == "" {
		log.Println("BUCKET env var missing or empty")
		initOk = false
	}

	accessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	if accessKey == "" {
		log.Println("AWS_ACCESS_KEY_ID env var missing or empty")
		initOk = false
	}

	secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secretKey == "" {
		log.Println("AWS_SECRET_ACCESS_KEY env var missing or empty")
		initOk = false
	}

	region = os.Getenv("AWS_REGION")
	if region == "" {
		log.Println("AWS_REGION env var missing or empty")
		initOk = false
	}

	if initOk == false {
		log.Println("Error initializing from environment: exiting.")
		os.Exit(1)
	}

}

func main() {

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := s3.New(sess)
	params := &s3.ListObjectsInput{
		Bucket:       aws.String(bucket), // Required
	}

	resp, err := svc.ListObjects(params)
	if err != nil {
		log.Fatalf("Failure listing bucket: %s",err.Error())
	}

	for _,c := range resp.Contents {
		log.Printf("Dump %s\n",*c.Key)
		cat(svc, bucket,*c.Key)
	}

}

func cat(s3svc *s3.S3, bucket, key string) error {
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	resp, err := s3svc.GetObject(params)
	if err != nil {
		return err
	}

	if resp.Body == nil {
		//Nothing to read
		return nil
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(content))
	return nil
}
