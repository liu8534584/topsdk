package response

import (
    "github.com/liu8534584/topsdk/ability396/domain"
)

type AlibabaEletkSettlesTagGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        接口返回model
    */
    Result  domain.AlibabaEletkSettlesTagGetPageResultDTO `json:"result,omitempty" `
}
