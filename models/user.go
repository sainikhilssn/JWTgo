package models

// define all data access objects in models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID  `bson:"_id"`
	First_name    *string			  `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string             `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string			  `json:"password" validate:"requried,min= 6"`
	Email         *string             `json:"phone" validate:"required"`
	Phone         *string			  `json:"" validate:""`
	User_type     *string             `json:"user_type" valididate:"required,eq=ADMIN|USER"`
	Token         *string             `json:"token"`
	Refresh_Token *string             `json:"refresh_token"`
	Created_at    time.Time           `json:"created_at"`
	Updated_at    time.Time           `json:"updated_at"`
	User_id       *string             `json:"user_id"`
}
