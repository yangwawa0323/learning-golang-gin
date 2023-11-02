package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// golang struct field default value
// binding : max , min , lte ( less then equal) , gte ( great then equal) , email , required
type User struct {
	gorm.Model
	Name         string  `form:"name" json:"name" uri:"name" binding:"required"`
	Email        *string `form:"email" json:"email" uri:"email" binding:"required,email"`
	Age          uint8   `form:"age" json:"age" uri:"age" binding:"required,numeric,lte=100,gte=18"`
	Birthday     *time.Time
	MemberNumber sql.NullString // golang struct , database table
}
