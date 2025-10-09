package model

type ShortLink struct {
	Link string `json:"link" binding:"required"`
}
