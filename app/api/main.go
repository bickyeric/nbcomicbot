package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/controller"
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load(".env")
}

func main() {
	ctx := context.Background()

	// region    ************************** CONNECTION **************************
	client := connection.NewMongo(ctx)
	db := client.Database(os.Getenv("DB_MONGO_DATABASE"))
	// endregion    ************************** CONNECTION **************************

	// region    ************************** REPO **************************
	sourceRepo := repository.NewSource(db)
	comicRepo := repository.NewComic(db)
	episodeRepo := repository.NewEpisode(db)
	pageRepo := repository.NewPage(db)
	// endregion    ************************** REPO **************************

	// region    ************************** SERVICE **************************
	saver := episode.NewSaveUpdate(sourceRepo, comicRepo, episodeRepo, pageRepo)
	// endregion    ************************** SERVICE **************************

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	kendang := controller.NewKendang(saver)
	e.POST("/kendang/webhook", kendang.OnHandle)

	e.Logger.Fatal(e.Start(":1907"))

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	e.Close()
	client.Disconnect(ctx)
}