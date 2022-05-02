package netsourcetos3

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	_ "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/leaf-rain/util/tool"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"time"
)

type S3Config struct {
	AccessKey string `json:"access_key,omitempty" yaml:"access_key"`
	SecretKey string `json:"secret_key,omitempty" yaml:"secret_key"`
	Token     string `json:"token,omitempty" yaml:"token"`
	EndPoint  string `json:"end_point,omitempty" yaml:"end_point"`
	Region    string `json:"region,omitempty" yaml:"region"`
}

var defaultS3 *session.Session

func InitS3(conf S3Config) error {
	var err error
	defaultS3, err = session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(conf.AccessKey, conf.SecretKey, conf.Token),
		Endpoint:         aws.String(conf.EndPoint),
		Region:           aws.String(conf.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})
	return err
}

func Netsourcetos3ByGet(url, bucket, path string, fn func(str string) string) (string, error) {
	if defaultS3 == nil {
		logx.Errorf("s3 uninitialized !")
		return "", errors.New("s3 uninitialized !")
	}
	resp, err := http.Get(url)
	defer resp.Body.Close()
	uploader := s3manager.NewUploader(defaultS3)
	//var result *s3manager.UploadOutput
	var suffix = time.Now().Format("20060102") + "/" + strconv.FormatUint(tool.GenUUID(), 10)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path + suffix),
		Body:   resp.Body,
	})
	return fn(suffix), err
}
