package proto

// 矿
type Mining struct {
	ID string `json:"id"`
	// 用户id
	UserID string `json:"userID"`
	// 计算结果hash
	ResultHash string `json:"resultHash"`
	// 计算结果
	Result string `json:"result"`
}

// 挖矿请求
type MiningReq struct {
	// 计算结果hash
	ResultHash string `json:"resultHash" binding:"required"`
	// 计算结果
	Result string `json:"result" binding:"required"`
	BodyData
}
