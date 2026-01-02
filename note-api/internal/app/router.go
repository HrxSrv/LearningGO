package app 

import (

"github.com/gin-gonic/gin"
"note-api/internal/config"
"note-api/internal/database"
"note-api/internal/note"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func StartServer(){
     cfg:= config.LoadConfig()
	 mongouri:= cfg.MongodRI
     dbClient := database.LoadMongo(mongouri)
     
     app := &App{
		DB: dbClient,
		Config: cfg,
	 }

     r:=gin.Default();

	 r.GET("/health", func(h *gin.Context){
		h.JSON(200,gin.H{"status":"ok"})
	 })
     

	 //note routes
     repo:= note.NewRepository(app.DB)
	 svc:=note.NewService(repo)
	 noteHandler := note.NewHandler(svc)

	 r.GET("/notes", noteHandler.GetNote)
	 r.POST("/notes", noteHandler.CreateNote)
	 r.GET("/notes/:id",noteHandler.GetNoteByID)
	//  r.PUT("/notes/:id",noteHandler.UpdateNote)
	//  r.DELETE("/notes/:id",noteHandler.DeleteNote)
	 r.Run(":" + app.Config.Port)
}