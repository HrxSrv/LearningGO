package note

import (
	"errors"
	"net/http"
	"note-api/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	Service *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{
       Service: svc,
	}
}

//GetNote returns all notes 
func (h *Handler) GetNote( c *gin.Context){
    
	ctx:=c.Request.Context()

	notes,err := h.Service.ListAllNotes(ctx)

	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, notes)

}


//CreateNote creates a note
func (h *Handler) CreateNote(c *gin.Context) {

	var req CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "Bad request son"})
		return
	}
    
	ctx:= c.Request.Context()

	note,err := h.Service.CreateNote(ctx,req)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, note)

}


//GetNoteByID gets a note by a given id

func (h *Handler) GetNoteByID (c *gin.Context){
     
    id:=c.Param("id")
    ctx:=c.Request.Context()
	objectId,err:=primitive.ObjectIDFromHex(id); 
	if err!=nil{
       c.JSON(http.StatusBadRequest,gin.H{
		"error":"bad request son",
	   })
	   return
	}
    
	note,err:=h.Service.GetNoteById(ctx,objectId)

	if err!=nil{
		switch err{
	case domain.MongoDocNotFound:
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Docs not there son",
		})
    default:
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})

	return	

	}
	}
    

    c.JSON(http.StatusOK,note)
	

}

//UpdateNote updates a note
func (h *Handler) UpdateNote (c *gin.Context){

	var updateRequest UpdateNoteRequest
    
	if err:=c.ShouldBindJSON(&updateRequest); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Bad req son",
		})
		return
       
	}

	id:=c.Param("id")
    
	objectId,err:=primitive.ObjectIDFromHex(id)

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Bad req son",
		})
		return
	}

	err = h.Service.UpdateNote(c.Request.Context(),&updateRequest,objectId)

	if err!=nil{

		switch {
		case errors.Is(err, domain.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"id":id,
		"mssg":"Updated successfully son",
	})
}

//DeleteNote deletes a note of given Id
func (h *Handler) DeleteNote (c *gin.Context){

	id:=c.Param("id")

	objectId,err:= primitive.ObjectIDFromHex(id)

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Bad req son",
		})
		return
	}
    
	err = h.Service.DeleteNote(c.Request.Context(),objectId)
    

	if err!=nil{
		switch {
		case errors.Is(err, domain.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"id":id,
		"mssg":"note nuked son",
	})

}
















// //DeleteNote

// func (h *Handler) DeleteNote (c *gin.Context){
// 	id:= c.Param("id")

// 	objectId,err := primitive.ObjectIDFromHex(id)

// 	if err!=nil {
//        c.JSON(http.StatusBadRequest, gin.H{
// 		"error":"Bad request son",
// 	   })
// 	   return
// 	}
//     ctx:=c.Request.Context()
// 	result,err:=h.Collection.DeleteOne(ctx,bson.M{
// 		"_id":objectId,
// 	})

//     if result.DeletedCount==0{
// 		c.JSON(http.StatusNotFound,gin.H{
// 			"error":"Delete Failed doc not found son",
// 		})
// 	}

// 	if err!=nil{
// 		c.JSON(http.StatusInternalServerError,err.Error())
// 	}
// 	c.JSON(http.StatusOK,result)
// }
