package verify

import (
	"encoding/base64"
	"fmt"

	"github.com/tjfoc/gmsm/sm2"
)

// sign 签名 userPub 用户公钥 message 需要验证的信息
func Verify(sign string, userPub string, message string) error {
	Pubkey, err := sm2.ReadPublicKeyFromMem([]byte(userPub), nil)
	if err != nil {
		return err
	}
	decodeString, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	verify := Pubkey.Verify([]byte(message), decodeString)
	if !verify {
		return fmt.Errorf("failed to verify sign")
	}
	return nil
}
