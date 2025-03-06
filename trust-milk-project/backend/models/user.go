package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Username string             `bson:"username"`
    Mobile   string             `bson:"mobile"`
    Email    string             `bson:"email"`
    Password string             `bson:"password"`
}
