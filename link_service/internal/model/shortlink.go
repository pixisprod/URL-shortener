package model

type ShortLink struct {
	Link string `json:"link" binding:"required"`
	TTL  int    `json:"ttl" binding:"required"`
}
