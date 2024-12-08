package request

type IdsRequest struct {
	Ids []int `json:"ids" form:"ids" binding:"required"` //标识
}
