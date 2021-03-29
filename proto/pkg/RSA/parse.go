package RSA

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func parsingPublicKeyFromMem(data []byte) (*rsa.PublicKey, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("failed to decode public key")
	}
	// 解析公钥
	result, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func parsingPrivateKeyFromMem(data []byte) (*rsa.PrivateKey, error) {
	// pem解码
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("failed to decode private key")
	}
	// der加密，返回一个私匙对象
	result, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return result, nil
}
