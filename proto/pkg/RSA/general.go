package RSA

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenRsaKey() (private, public []byte, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	private = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	pubKeyData := x509.MarshalPKCS1PublicKey(publicKey)
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyData,
	}
	public = pem.EncodeToMemory(block)
	return private, public, nil
}
