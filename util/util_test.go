package util

import (
  "fmt"
  "os"
  "path/filepath"
  "testing"
)

func Test_TrimBrackets(t *testing.T) {
  fmt.Println(TrimBrackets("【阿里CTO：阿里所有技术和产品输出都将必须通过阿里云进行】财联社3月21日讯，阿里巴巴首席技术官张建锋今日宣布，未来一到两年，阿里集团业务要100%跑在阿里云的公有云上，未来阿里云和阿里的技术完全拉通，同时，阿里所有的技术和产品输出都将必须通过阿里云平台，包括金融、钉钉、新零售。目前阿里巴巴集团的业务有60%到70%跑在阿里云的公有云上。（澎湃）"))
  fmt.Println(TrimBracketContent("【阿里CTO：阿里所有技术和产品输出都将必须通过阿里云进行】财联社3月21日讯，阿里巴巴首席技术官张建锋今日宣布，未来一到两年，阿里集团业务要100%跑在阿里云的公有云上，未来阿里云和阿里的技术完全拉通，同时，阿里所有的技术和产品输出都将必须通过阿里云平台，包括金融、钉钉、新零售。目前阿里巴巴集团的业务有60%到70%跑在阿里云的公有云上。（澎湃）"))
}

func TestModifyFileName(t *testing.T) {

  // 遍历文件夹，获取文件路径
  paths := make([]string, 0)
  filepath.Walk("/Users/lio/Desktop/街拍", func(path string, info os.FileInfo, err error) error {
    if !info.IsDir() {
      paths = append(paths, path)
    }
    return nil
  })

  // 遍历文件路径，修改文件名
  for i, path := range paths {
    newPath := filepath.Join(filepath.Dir(path), fmt.Sprintf("%d", i+215)+filepath.Ext(path))
    os.Rename(path, newPath)
  }

}

func TestFormatInDate(t *testing.T) {
  param := make(map[string]string)
  param["cardType"] = "PAN_FRONT"
  resp, err := NewUploadRequestLocalFile("https://in-api.advance.ai/in/openapi/face-identity/v1/id-card-ocr", param, "/Users/yu/java/1.jpeg")
  if err != nil {
    t.Fatal(err)
  }
  fmt.Println(resp)
}