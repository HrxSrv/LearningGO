package note

import (
	"net/http"
	"note-api/internal/domain"
	// "time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
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


//GetNoteByID

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


// //UpdateNote

// func (h *Handler) UpdateNote (c *gin.Context){

// 	id:=c.Param("id")

// 	objectId,err:=primitive.ObjectIDFromHex(id)

// 	if err!=nil{
//     c.JSON(http.StatusBadRequest,gin.H{
// 		"error":"Bad request son",
// 	})
// 	return
// 	}
    
// 	var req UpdateNoteRequest

// 	if err:=c.ShouldBindJSON(&req); err!=nil{
// 		c.JSON(http.StatusBadRequest,gin.H{
// 			"error":"Bad request son",
// 		})
// 	return	
// 	}
    
// 	update := bson.M{
// 		"$set":bson.M{
// 			"title":req.Title,
// 			"content":req.Content,
// 		},
// 	}
//     ctx:=c.Request.Context()


//     result,err:=h.Collection.UpdateByID(ctx,objectId,update)

//     if result.MatchedCount==0{
// 		c.JSON(http.StatusNotFound,gin.H{
// 			"error":"didn't find the doc son",
// 		})
// 		return
// 	}

// 	 if err!=nil{
// 		c.JSON(http.StatusNotFound,gin.H{
// 			"error":"Shit didnt workout bro",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK,gin.H{
// 		"matched":result.MatchedCount,
// 		"modified":result.ModifiedCount,
// 	})
// }

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