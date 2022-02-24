package response

import (
    "topsdk/ability411/domain"
)

type TaobaoTbkScOrderDetailsGetResponse struct {

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
    Data  domain.TaobaoTbkScOrderDetailsGetOrderPage `json:"data,omitempty" `
}