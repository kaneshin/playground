package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	config    *aws.Config
	name      = os.Getenv("AWS_S3_BUCKET_NAME")
	region    = os.Getenv("AWS_S3_BUCKET_REGION")
	byteRange = flag.String("byte-range", "", "")
)

func init() {
	config = aws.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithHTTPClient(http.DefaultClient).
		WithMaxRetries(aws.UseServiceDefaultRetries).
		WithLogger(aws.NewDefaultLogger()).
		WithLogLevel(aws.LogOff)
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	key := flag.Args()[0]

	// initialize the S3 service.
	svc := NewS3(aws.NewConfig().WithRegion(region))
	input := s3.GetObjectInput{
		Bucket: aws.String(name),
		Key:    aws.String(key),
	}

	if len(*byteRange) > 0 {
		// Downloads the specified range bytes of an object. For more information about
		// the HTTP Range header, go to http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35.
		// Range *string `location:"header" locationName:"Range" type:"string"`
		input.Range = aws.String(*byteRange)
	}

	req, output := svc.GetObjectRequest(&input)
	if err := req.Send(); err != nil {
		log.Fatalf("error %v", err)
	}
	defer output.Body.Close()

	tempdir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatalf("error %v", err)
	}

	file, err := os.Create(filepath.Join(tempdir, filepath.Base(key)))
	if err != nil {
		log.Fatalf("error %v", err)
	}
	defer file.Close()

	// write bytes to file
	b, err := ioutil.ReadAll(output.Body)
	if _, err := file.Write(b); err != nil {
		log.Fatalf("error %v", err)
	}

	log.Printf("write file: %v", file.Name())
}

// NewS3 creates a new instance of the S3 client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a S3 client with additional configuration
//     svc := NewS3(aws.NewConfig().WithRegion("us-west-2"))
func NewS3(c *aws.Config) *s3.S3 {
	sess := session.New(config, c)
	return s3.New(sess, config, c)
}
