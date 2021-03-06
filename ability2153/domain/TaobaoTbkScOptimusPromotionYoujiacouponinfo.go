package domain


type TaobaoTbkScOptimusPromotionYoujiacouponinfo struct {
    /*
        有价券商品id     */
    ItemId  *interface{} `json:"item_id,omitempty" `

    /*
        商品链接     */
    Url  *string `json:"url,omitempty" `

}

func (s *TaobaoTbkScOptimusPromotionYoujiacouponinfo) SetItemId(v interface{}) *TaobaoTbkScOptimusPromotionYoujiacouponinfo {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkScOptimusPromotionYoujiacouponinfo) SetUrl(v string) *TaobaoTbkScOptimusPromotionYoujiacouponinfo {
    s.Url = &v
    return s
}
