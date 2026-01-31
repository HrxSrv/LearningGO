package note

import (
	"context"
	"fmt"
	"note-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct{
	Collection *mongo.Collection
}

func NewRepository (db *mongo.Client) *Repository{
	return &Repository{
		Collection: db.Database("notesdb").Collection("notes"),
	}
}

func (r *Repository) Create (ctx context.Context,note *Note) error{ 
	result,err := r.Collection.InsertOne(ctx,note)

	if err!=nil{
		return err
	}

    note.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}


func (r *Repository) ListAllNotes (ctx context.Context,notes *[]Note) (error) {
	cursor,err:= r.Collection.Find(ctx,bson.M{})

	if err!=nil {
		return err
	}

	defer cursor.Close(ctx)

	if err:=cursor.All(ctx,notes); err!=nil{
		return err
	}

	return nil

}




func (r *Repository) Get (ctx context.Context, id primitive.ObjectID)(*Note,error){
 
	var note Note

	err:= r.Collection.FindOne(ctx,bson.M{
		"_id":id,
	}).Decode(&note)

	if err!=nil {
		return nil,err
	}

	return &note,nil
  
}


func (r *Repository) Update (ctx context.Context, id primitive.ObjectID, update bson.M) (error) {

	result,err:= r.Collection.UpdateByID(ctx,id,bson.M{
		"$set":update,
	})

	if err!=nil{
		return fmt.Errorf("update note %s: %w", id.Hex(), err)
	}

	if result.MatchedCount==0{
		return fmt.Errorf("update note %s: %w", id.Hex(), domain.ErrNotFound)
	}

    return nil
}



func (r *Repository) Delete (ctx context.Context,id primitive.ObjectID) error{

	result,err:= r.Collection.DeleteOne(ctx,bson.M{
		"_id":id,
	})

	if err!=nil{
		return fmt.Errorf("delete note %s: %w", id.Hex(), err)
	}
    
	if result.DeletedCount==0{
		return fmt.Errorf("delete note %s: %w", id.Hex(), domain.ErrNotFound)
	}
	return nil
}
