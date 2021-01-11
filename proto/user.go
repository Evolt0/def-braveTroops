package proto

type User struct {
	// 用户id（根据用户公钥Hash获得）
	ID string `json:"id"`
	// 用户公钥
	PubKey string `json:"pubKey"`
	// 账户余额
	Balance string `json:"balance"`
}

type UserReq struct {
	// 用户id（根据用户公钥Hash获得）
	ID string `json:"id" binding:"required"`
	// 用户公钥
	PubKey string `json:"pubKey" binding:"required"`
}
