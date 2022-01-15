package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func main() {

	if len(os.Args) < 7 {
		fmt.Printf("e.g. %s $AWS_ACCESS_KEY_ID $AWS_SECRET_ACCESS_KEY $AWS_SESSION_TOKEN file_for_upload s3_target bucket\n", os.Args[0])
		os.Exit(-1)
	}
	args := os.Args[1:]
	keyId := args[0]
	secret := args[1]
	token := args[2]
	file := args[3]
	target := args[4]
	bucket := args[5]

	session := connectAws(keyId, secret, token)
	upFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize = upFileInfo.Size()

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(target),
		ACL:                  aws.String("private"),
		Body:                 upFile,
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String("text/csv"),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("upload successfully\n")
}

func connectAws(id, key, token string) *session.Session {

	if id == "" {
		panic("please specify AWS_ACCESS_KEY_ID")
	}
	if key == "" {
		panic("please specify AWS_SECRET_ACCESS_KEY")
	}
	if token == "" {
		panic("please specify AWS_SESSION_TOKEN")
	}
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				id,
				key,
				token,
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
