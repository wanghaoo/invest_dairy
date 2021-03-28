package service

import (
	"fmt"
	"invest_dairy/common"
	"invest_dairy/obs"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

const (
	FilePrefix   = "/image"
	LOGODictName = "logo"
	UploadPrefix = "https://india-loan.obs.ap-southeast-3.myhuaweicloud.com"
)

var bucket *ObjectOperationsSample

func InitOssBucket() {
	var err error
	bucket = newObjectOperationsSample("XAAC0OB7RWPLN7JWSVFT",
		"xFBe92rJyfgttqyrSqnzQkoCwQbf3stEbRhf2cBz",
		"obs.ap-southeast-3.myhuaweicloud.com",
		"india-loan", "", "")
	if err != nil {
		common.Mlog.Errorf("util.GetBucket error %v ", err)
	}
}

func UploadImage(c echo.Context) *common.ResponseData {
	// 上传本地文件，用户打开文件，传入句柄。
	file, err := c.FormFile("file")
	if err != nil {
		common.Mlog.Error(err)
		return common.SetError("file upload fail")
	}
	fd, err := file.Open()
	defer fd.Close()
	if err != nil {
		common.Mlog.Error(err)
		return common.CommonError()
	}
	imagePath, err := uploadFile(fd)
	if err != nil {
		common.Mlog.Errorf("upload file error: %s", err.Error())
	}
	return common.SetData(imagePath)
}

func uploadFile(fd io.Reader) (string, error) {
	// 图片命名规则  path_mtype_时间_随机数
	var imgName = FilePrefix + time.Now().Format("030405") + strconv.Itoa(rand.Intn(10000))
	var objectName = fmt.Sprintf("%s%s.jpg", LOGODictName, imgName)

	input := &obs.PutObjectInput{}
	input.Bucket = bucket.bucketName
	input.Key = objectName
	input.Body = fd

	objectResp, err := bucket.obsClient.PutObject(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(&objectResp)
	fmt.Printf("Create object:%s successfully!\n", objectName)
	fmt.Println()
	return fmt.Sprintf("%s/%s", UploadPrefix, objectName), nil
}

type ObjectOperationsSample struct {
	bucketName string
	objectKey  string
	location   string
	obsClient  *obs.ObsClient
}

func newObjectOperationsSample(ak, sk, endpoint, bucketName, objectKey, location string) *ObjectOperationsSample {
	obsClient, err := obs.New(ak, sk, endpoint)
	if err != nil {
		panic(err)
	}
	return &ObjectOperationsSample{obsClient: obsClient, bucketName: bucketName, objectKey: objectKey, location: location}
}

func DownloadFile(path string) *os.File {
	if len(path) <= 0 {
		return nil
	}
	path = strings.Replace(path, UploadPrefix+"/", "", -1)
	input := &obs.GetObjectInput{}
	input.Bucket = bucket.bucketName
	input.Key = path
	output, err := bucket.obsClient.GetObject(input)
	if err != nil {
		common.Mlog.Errorf("get file error: %s", err.Error())
	}
	defer output.Body.Close()
	fmt.Printf("StorageClass:%s, ETag:%s, ContentType:%s, ContentLength:%d, LastModified:%s\n",
		output.StorageClass, output.ETag, output.ContentType, output.ContentLength, output.LastModified)
	// 读取对象内容
	file, err := os.Create("test.jpg")
	if err != nil {
		common.Mlog.Errorf("create file error: %s", err.Error())
	}
	defer file.Close()
	_, err = io.Copy(file, output.Body)
	if err != nil {
		common.Mlog.Errorf("copy file error: %s", err.Error())
	}
	return file
}

func DownloadFileBytes(path string) ([]byte, string, error) {
	result := make([]byte, 0)
	if len(path) <= 0 {
		return result, "", errors.New("path is nil")
	}
	path = strings.Replace(path, UploadPrefix+"/", "", -1)
	input := &obs.GetObjectInput{}
	input.Bucket = bucket.bucketName
	input.Key = path
	output, err := bucket.obsClient.GetObject(input)
	if err != nil {
		common.Mlog.Errorf("get file error: %s", err.Error())
		return result, "", err
	}
	defer output.Body.Close()
	fmt.Printf("StorageClass:%s, ETag:%s, ContentType:%s, ContentLength:%d, LastModified:%s\n",
		output.StorageClass, output.ETag, output.ContentType, output.ContentLength, output.LastModified)
	// 读取对象内容
	result, err = ioutil.ReadAll(output.Body)
	if err != nil {
		common.Mlog.Errorf("ioutil read file error: %s", err.Error())
		return result, "", err
	}
	return result, input.Key, nil
}

func downloadNetworkFileAndUploadToOBS(url string) (string, error) {
	client := httpClient()
	resp, err := client.Get(url)
	if err != nil {
		common.Mlog.Errorf("http get file error: %s", err.Error())
	}
	defer resp.Body.Close()
	return uploadFile(resp.Body)
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	return &client
}
