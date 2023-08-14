package model

import (
	"errors"
	"fmt"
	"go_blog/crypto"

	"gorm.io/gorm"
)

type Admin struct {
	AdminID string
	Name string
	HashedPassword string
}

