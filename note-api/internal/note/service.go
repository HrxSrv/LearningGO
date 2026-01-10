package note

import (
	"context"
	"errors"
	"note-api/internal/domain"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {

	Repository *Repository

}

func NewService (repo *Repository) *Service{
	return &Service{
		Repository: repo,
	}
}

func (s *Service) CreateNote (c context.Context, req CreateNoteRequest) (*Note,error){

	note:= Note{
		Title: req.Title,
		Content: req.Content,
		CreatedAt: time.Now(),
	}

	err:= s.Repository.Create(c,&note)

	if err!=nil{
		return nil,err;
	}

	return &note,nil

}

func (s *Service) ListAllNotes (ctx context.Context) (*[]Note,error) {

	var notes []Note
    
	err:=s.Repository.ListAllNotes(ctx,&notes)

	if err!=nil{
		return nil,err
	}
    
	return &notes,nil

}


func (s *Service) GetNoteById (ctx context.Context,id primitive.ObjectID) (*Note,error){
 
     note,err:= s.Repository.Get(ctx,id)

	 if (err!=nil) {

		if (errors.Is(err,mongo.ErrNoDocuments)){
           return nil,domain.MongoDocNotFound
		}

		return nil, err
		
	 }

	 return note,nil

}


func (s *Service) UpdateNote (ctx context.Context, req *UpdateNoteRequest,id primitive.ObjectID) error {

    update:= bson.M{
		"title":req.Title,
		"content":req.Content,
	}
	err := s.Repository.Update(ctx,id,update)

	if err!=nil{
		return err;
	}
	return err
}

func (s *Service) DeleteNote (ctx context.Context, id primitive.ObjectID) error{

	err:= s.Repository.Delete(ctx,id)

	if err!=nil{
		return err
	}
	
	return nil

}