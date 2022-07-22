package model

type Data struct{
	OrderId   interface{} `json:"order_id"`
	PaymentId interface{} `json:"payment_id"`
	RebillId  interface{} `json:"rebill_id"`
	CardId    interface{} `json:"card_id"`
}

/*
type Data struct{
	OrderId   interface{} `json:"order_id"binding:"required"`
	PaymentId interface{} `json:"payment_id"binding:"required"`
	RebillId  interface{} `json:"rebill_id"binding:"required"`
	CardId    interface{} `json:"card_id"binding:"required"`
}
*/
