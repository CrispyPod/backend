package main

import (
	"time"

	"crispypod.com/crispypod-backend/controllers"
	"crispypod.com/crispypod-backend/db"
	"crispypod.com/crispypod-backend/graph"
	"crispypod.com/crispypod-backend/helpers"
	"crispypod.com/crispypod-backend/rssfeed"
	"crispypod.com/crispypod-backend/schedule"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

func main() {
	helpers.CheckEnvVariables()
	db.ConnectDatabase()

	// // event dispatcher
	// defer event.CloseWait()

	r := gin.Default()
	r.Use(helpers.JWTMiddleWare())
	r.Use(cors.Default())

	if gin.Mode() == "debug" {
		pH := playground.Handler("GraphQL", "/graphql")
		r.GET("/graphql", func(ctx *gin.Context) {
			pH.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}

	gH := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	r.POST("/graphql", func(ctx *gin.Context) {
		gH.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.StaticFile("/rss", "./Feed/rss.xml")

	r.POST("/api/audioFile", controllers.AudioFileUpload)
	r.GET("/api/audioFile/:fileName", controllers.GetAudioFile)

	r.POST("/api/thumbnail", controllers.ThumbnailUpload)
	r.GET("/api/thumbnail/:fileName", controllers.GetThumbnailFile)

	rssfeed.GenerateRSSFeed()

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("0:00").Do(schedule.ClearAudioFile)
	s.StartAsync()

	r.Run()
}
