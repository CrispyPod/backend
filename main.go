package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	// helpers.CheckEnvVariables()
	// db.ConnectDatabase()

	// // event dispatcher
	// defer event.CloseWait()

	// eventhandler.RegisterEvent()

	// r := gin.Default()
	// r.Use(helpers.JWTMiddleWare())
	// r.Use(cors.Default())

	// if gin.Mode() == "debug" {
	// 	pH := playground.Handler("GraphQL", "/graphql")
	// 	r.GET("/graphql", func(ctx *gin.Context) {
	// 		pH.ServeHTTP(ctx.Writer, ctx.Request)
	// 	})
	// }

	// gH := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	// r.POST("/graphql", func(ctx *gin.Context) {
	// 	gH.ServeHTTP(ctx.Writer, ctx.Request)
	// })

	// r.StaticFile("/rss", "./Feed/rss.xml")

	// r.POST("/api/audioFile", controllers.AudioFileUpload)
	// r.GET("/api/audioFile/:fileName", controllers.GetAudioFile)

	// r.POST("/api/imageFile", controllers.ImageFileUpload)
	// r.GET("/api/imageFile/:fileName", controllers.GetImageFile)
	// r.POST("/api/imageFile/upload", controllers.UploadFile)

	// r.POST("/api/deployLog", controllers.DeployLogUpload)

	// go rssfeed.GenerateRSSFeed()

	// s := gocron.NewScheduler(time.UTC)
	// s.Every(1).Day().At("0:00").Do(schedule.ClearFiles)
	// s.StartAsync()

	// r.Run()
}
