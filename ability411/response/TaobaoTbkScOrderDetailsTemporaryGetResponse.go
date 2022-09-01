package response

import (
	"topsdk/ability411/domain"
)

type TaobaoTbkScOrderDetailsTemporaryGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   PublisherOrderDto
	*/
	Data domain.TaobaoTbkScOrderDetailsTemporaryGetOrderPage `json:"data,omitempty" `
}
