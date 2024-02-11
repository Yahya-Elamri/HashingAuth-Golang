package utils

import (
	"gorm.io/gorm"
)

type (
	UserData struct {
		Idusers  int64  `json:"idusers"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

var Db *gorm.DB
