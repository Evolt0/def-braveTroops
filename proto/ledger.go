package proto

type Ledger struct {
	// 账本id
	ID string `json:"id"`
	// 发起人
	SenderID string `json:"senderID"`
	// 收款人
	ReceiverID string `json:"receiverID"`
	// 金额
	Amount string `json:"amount"`
}
