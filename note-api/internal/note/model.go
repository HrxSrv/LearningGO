package note

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Note struct {
     
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`

}

type CreateNoteRequest struct{
	Title string `json:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required"`
}

type UpdateNoteRequest struct{
	Title string `json:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required"`
}