// based on https://github.com/golang/go/blob/master/src/net/url/url.go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
  "crypto/cipher"
  "strings"
  "encoding/base64"
  "crypto/aes"
  "bytes"
  "fmt"
)

func shouldEscape(c byte) bool {
  if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || c == '_' || c == '-' || c == '~' || c == '.' {
    return false
  }
  return true
}
func escape(s string) string {
  hexCount := 0
  for i := 0; i < len(s); i++ {
    c := s[i]
    if shouldEscape(c) {
      hexCount++
    }
  }

  if hexCount == 0 {
    return s
  }

  t := make([]byte, len(s)+2*hexCount)
  j := 0
  for i := 0; i < len(s); i++ {
    switch c := s[i]; {
    case shouldEscape(c):
      t[j] = '%'
      t[j+1] = "0123456789ABCDEF"[c>>4]
      t[j+2] = "0123456789ABCDEF"[c&15]
      j += 3
    default:
      t[j] = s[i]
      j++
    }
  }
  return string(t)
}

func Encrypt(privateKey, data string) (string, error) {
  var result string
  key, err := base64.StdEncoding.DecodeString(privateKey)
  if err != nil {
    return result, err
  }
  crypted := AesEncrypt(data, string(key))
  return crypted, nil
}

func Base64URLDecode(data string) ([]byte, error) {
  var missing = (4 - len(data)%4) % 4
  data += strings.Repeat("=", missing)
  res, err := base64.URLEncoding.DecodeString(data)
  fmt.Println("  decodebase64urlsafe is :", string(res), err)
  return base64.URLEncoding.DecodeString(data)
}

func Base64UrlSafeEncode(source []byte) string {
  // Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
  bytearr := base64.StdEncoding.EncodeToString(source)
  safeurl := strings.Replace(string(bytearr), "/", "_", -1)
  safeurl = strings.Replace(safeurl, "+", "-", -1)
  safeurl = strings.Replace(safeurl, "=", "", -1)
  return safeurl
}

func AesDecrypt(crypted, key []byte) []byte {
  block, err := aes.NewCipher(key)
  if err != nil {
    fmt.Println("err is:", err)
  }
  blockMode := NewECBDecrypter(block)
  origData := make([]byte, len(crypted))
  blockMode.CryptBlocks(origData, crypted)
  origData = PKCS5UnPadding(origData)
  fmt.Println("source is :", origData, string(origData))
  return origData
}

func AesEncrypt(src, key string) string {
  block, err := aes.NewCipher([]byte(key))
  if err != nil {
    fmt.Println("key error1", err)
  }
  if src == "" {
    fmt.Println("plain content empty")
  }
  ecb := NewECBEncrypter(block)
  content := []byte(src)
  content = PKCS5Padding(content, block.BlockSize())
  crypted := make([]byte, len(content))
  ecb.CryptBlocks(crypted, content)
  // 普通base64编码加密 区别于urlsafe base64
  fmt.Println("base64 result:", base64.StdEncoding.EncodeToString(crypted))
  return  base64.StdEncoding.EncodeToString(crypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
  padding := blockSize - len(ciphertext)%blockSize
  padtext := bytes.Repeat([]byte{byte(padding)}, padding)
  return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
  length := len(origData)
  // 去掉最后一个字节 unpadding 次
  unpadding := int(origData[length-1])
  return origData[:(length - unpadding)]
}

type ecb struct {
  b         cipher.Block
  blockSize int
}

func newECB(b cipher.Block) *ecb {
  return &ecb{
    b:         b,
    blockSize: b.BlockSize(),
  }
}

type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
  return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
  if len(src)%x.blockSize != 0 {
    panic("crypto/cipher: input not full blocks")
  }
  if len(dst) < len(src) {
    panic("crypto/cipher: output smaller than input")
  }
  for len(src) > 0 {
    x.b.Encrypt(dst, src[:x.blockSize])
    src = src[x.blockSize:]
    dst = dst[x.blockSize:]
  }
}

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
  return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
  if len(src)%x.blockSize != 0 {
    panic("crypto/cipher: input not full blocks")
  }
  if len(dst) < len(src) {
    panic("crypto/cipher: output smaller than input")
  }
  for len(src) > 0 {
    x.b.Decrypt(dst, src[:x.blockSize])
    src = src[x.blockSize:]
    dst = dst[x.blockSize:]
  }
}
