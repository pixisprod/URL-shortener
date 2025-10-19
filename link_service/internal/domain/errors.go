package domain

import "errors"

var (
	ErrLinkNotFound = errors.New("Link not found")
	ErrLinkExpired  = errors.New("Link expired")
)
