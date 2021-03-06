package goutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// MD5 对字符串进行MD5哈希
func MD5(data string) string {
	t := md5.New()
	_, err := io.WriteString(t, data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", t.Sum(nil))
}

// SHA1 对字符串进行SHA1哈希
func SHA1(data string) string {
	t := sha1.New()
	_, err := io.WriteString(t, data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", t.Sum(nil))
}

// Crmd5 MD5加密算法
func Crmd5(s string) string {
	h := md5.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		return ""
	}

	return hex.EncodeToString(h.Sum(nil))
}

// AesEncrypt AES 加密
// @origData 加密的原始数据
// @key 密钥
// @iv 向量
func AesEncrypt(origData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)

	// fmt.Println(key)
	// fmt.Println(key[:blockSize])
	// origData = ZeroPadding(origData, block.BlockSize()) key[:blockSize]
	// blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt AES 解密
// @origData 加密的原始数据
// @key 密钥
// @iv 向量
func AesDecrypt(crypted, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("密钥错误")
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData, err = PKCS5UnPadding(origData)
	if err != nil {
		return nil, err
	}
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

// ZeroPadding 0padding
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

// ZeroUnPadding 0padding
func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// PKCS5Padding PKCS5
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding PKCS5
func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("解密失败")
	}
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	if length < unpadding {
		return nil, errors.New("解密失败")
	}
	return origData[:(length - unpadding)], nil
}
