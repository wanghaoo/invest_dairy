package util

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func HttpPut(url string, param map[string]interface{}, auth map[string]string) (string, error) {
  return doHttp(url, http.MethodPut, param, auth, nil)
}

func HttpPostV2(url string, param map[string]interface{}, head map[string]string) (string, error) {
  return doHttp(url, http.MethodPost, param, head, nil)
}

func HttpPostV3(url string, param map[string]interface{}, request map[string]string, head map[string]string) (string, error) {
  return doHttp(url, http.MethodPost, param, head, request)
}

func HttpPostV4(url string, param map[string]interface{}, request map[string]string, head map[string]string) (string, error) {
  return doHttpV2(url, http.MethodPost, param, head, request)
}

func HttpPostV5(url string, param map[string]string) (string, error) {
  return doHttpV3(url, param)
}

func HttpGet(url string, request map[string]string, head map[string]string) (string, error) {
  return doHttp(url, http.MethodGet, nil, head, request)
}

func doHttp(url string, method string, param map[string]interface{}, head map[string]string, requestParam map[string]string) (string, error) {
  paramJson, err := json.Marshal(param)
  if err != nil {
    return "", err
  }
  client := &http.Client{}
  request, err := http.NewRequest(method, url, strings.NewReader(string(paramJson)))
  if request == nil {
    return "", errors.New("build http request error")
  }
  if requestParam != nil {
    q := request.URL.Query()
    for k, v := range requestParam {
      q.Add(k, v)
    }
    request.URL.RawQuery = q.Encode()
  }
  request.Header.Set("Content-Type", "application/json")
  if len(head) > 0 {
    for k, v := range head {
      request.Header.Set(k, v)
    }
  }
  response, err := client.Do(request)
  if err != nil {
    return "", err
  }
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return "", err
  }
  return string(body), nil
}

func doHttpV2(url string, method string, param map[string]interface{}, head map[string]string, requestParam map[string]string) (string, error) {
  paramJson, err := json.Marshal(param)
  if err != nil {
    return "", err
  }
  client := &http.Client{}
  request, err := http.NewRequest(method, url, strings.NewReader(string(paramJson)))
  if request == nil {
    return "", errors.New("build http request error")
  }
  if requestParam != nil {
    q := request.URL.Query()
    for k, v := range requestParam {
      q.Add(k, v)
    }
    request.URL.RawQuery = q.Encode()
  }
  if len(head) > 0 {
    for k, v := range head {
      request.Header.Set(k, v)
    }
  }
  response, err := client.Do(request)
  if err != nil {
    return "", err
  }
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return "", err
  }
  return string(body), nil
}

func HttpPost(url string, param string) (string, error) {
  resp, err := http.Post(url, "application/json", strings.NewReader(param))
  if err != nil {
    return "", err
  }

  if resp.StatusCode != http.StatusOK {
    fmt.Println("resp.StatusCode = ", resp.StatusCode, param)
    return "", nil
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
  return string(body), nil
}

func DumpRequest(req *http.Request) {
  // Save a copy of this request for debugging.
  requestDump, err := httputil.DumpRequest(req, true)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(requestDump))
}

//参数为二进制格式文件
func NewUploadRequest(url string, params map[string]string, filename string, data []byte) (string, error) {

  // 实例化multipart
  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  // 创建multipart 文件字段
  part, err := writer.CreateFormFile("image", filename)
  if err != nil {
    return "", err
  }
  // 写入文件数据到multipart，和读取本地文件方法的唯一区别
  _, err = part.Write(data)
  //将额外参数也写入到multipart
  for key, val := range params {
    _ = writer.WriteField(key, val)
  }
  err = writer.Close()
  if err != nil {
    return "", err
  }

  //创建请求
  req, err := http.NewRequest("POST", url, body)
  if err != nil {
    return "", err
  }
  //不要忘记加上writer.FormDataContentType()，
  //该值等于content-type :multipart/form-data; boundary=xxxxx
  req.Header.Add("Content-Type", writer.FormDataContentType())
  req.Header.Add("X-ADVAI-KEY", "ab6a0bfdd73bd9e4")
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  respBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
  return string(respBytes), nil
}

func NewUploadRequestLocalFile(url string, params map[string]string,  path string) (string, error) {

  file, err := os.Open(path)
  if err != nil {
    return "", err
  }
  defer file.Close()

  // 实例化multipart
  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  // 创建multipart 文件字段
  part, err := writer.CreateFormFile("image", filepath.Base(path))
  if err != nil {
    return "", err
  }
  // 写入文件数据到multipart
  _, err = io.Copy(part, file)
  //将额外参数也写入到multipart
  for key, val := range params {
    _ = writer.WriteField(key, val)
  }
  err = writer.Close()
  if err != nil {
    return "", err
  }

  //创建请求
  req, err := http.NewRequest("POST", url, body)
  if err != nil {
    return "", err
  }
  //不要忘记加上writer.FormDataContentType()，
  //该值等于content-type :multipart/form-data; boundary=xxxxx
  req.Header.Add("Content-Type", writer.FormDataContentType())
  req.Header.Add("X-ADVAI-KEY", "ab6a0bfdd73bd9e4")
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  respBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
  return string(respBytes), nil
}

func HttpPostUtf8(url string, param string) (string, error) {
  client := &http.Client{}
  request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(param))
  if request == nil {
    return "", errors.New("build http request error")
  }
  request.Header.Set("Content-Type", "application/json")
  request.Header.Set("Accept-Charset","utf-8")
  response, err := client.Do(request)
  if err != nil {
    return "", err
  }
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return "", err
  }
  return string(body), nil
}

//参数为二进制格式文件
func NewUploadRequest2(url string, filename1 string, data1 []byte, filename2 string, data2 []byte) (string, error) {
  // 实例化multipart
  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  // 创建multipart 文件字段
  part, err := writer.CreateFormFile("firstImage", filename1)
  if err != nil {
    return "", err
  }
  // 写入文件数据到multipart，和读取本地文件方法的唯一区别
  _, err = part.Write(data1)
  if err != nil {
    return "", err
  }

  // 创建multipart 文件字段
  part2, err := writer.CreateFormFile("secondImage", filename2)
  if err != nil {
    return "", err
  }
  // 写入文件数据到multipart，和读取本地文件方法的唯一区别
  _, err = part2.Write(data2)
  if err != nil {
    return "", err
  }
  err = writer.Close()
  if err != nil {
    return "", err
  }

  //创建请求
  req, err := http.NewRequest("POST", url, body)
  if err != nil {
    return "", err
  }
  //不要忘记加上writer.FormDataContentType()，
  //该值等于content-type :multipart/form-data; boundary=xxxxx
  req.Header.Add("Content-Type", writer.FormDataContentType())
  req.Header.Add("X-ADVAI-KEY", "ab6a0bfdd73bd9e4")
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  respBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
  return string(respBytes), nil
}

func doHttpV3(requestUrl string, param map[string]string) (string, error) {
  data := url.Values{}
  for k, v := range param {
    data.Add(k, v)
  }
  response, err := http.PostForm(requestUrl, data)
  if err != nil {
    return "", err
  }
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return "", err
  }
  return string(body), nil
}

func HttpPostV6(url string, param []byte, request map[string]string, head map[string]string) (string, error) {
	return doHttpV4(url, http.MethodPost, param, head, request)
}

func doHttpV4(url string, method string, param []byte, head map[string]string, requestParam map[string]string) (string, error) {
	client := &http.Client{}
	var zBuf bytes.Buffer
	zw := gzip.NewWriter(&zBuf)
	if _, err := zw.Write(param); err != nil {
		return "", err
	}
	zw.Close()
	request, err := http.NewRequest(method, url, &zBuf)
	if request == nil {
		return "", errors.New("build http request error")
	}
	if requestParam != nil {
		q := request.URL.Query()
		for k, v := range requestParam {
			q.Add(k, v)
		}
		request.URL.RawQuery = q.Encode()
	}
	request.Header.Set("Content-Type", "application/json")
	if len(head) > 0 {
		for k, v := range head {
			request.Header.Set(k, v)
		}
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpsGet(url string, request map[string]string, head map[string]string) (string, error) {
	return doHttpsV2(url, http.MethodGet, nil, head, request)
}

func doHttpsV2(url string, method string, param map[string]interface{}, head map[string]string, requestParam map[string]string) (string, error) {
	read := new(strings.Reader)
	if param != nil {
		paramJson, err := json.Marshal(param)
		if err != nil {
			return "", err
		}
		read = strings.NewReader(string(paramJson))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	request, err := http.NewRequest(method, url, read)
	if request == nil {
		return "", errors.New("build http request error")
	}
	if requestParam != nil {
		q := request.URL.Query()
		for k, v := range requestParam {
			q.Add(k, v)
		}
		request.URL.RawQuery = q.Encode()
	}
	request.Header.Set("Content-Type", "application/json")
	if len(head) > 0 {
		for k, v := range head {
			request.Header.Set(k, v)
		}
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}