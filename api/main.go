package main

import (
	"app/auth"
	_ "app/docs"
	"app/handlers"

	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	pluginsDir string
	dbPath     string
	port       string
	envFP      string
)

func init() {

	flag.StringVar(&pluginsDir, "d", "plugins", "specify the relative path of the processes dir")
	flag.StringVar(&port, "p", "5050", "specify the port to run the api on")
	flag.StringVar(&envFP, "e", "/.env", "specify the path of the dot env file to load")
	flag.StringVar(&dbPath, "db", "/.data/db.sqlite", "specify the path of the sqlite database")

	flag.Parse()

	err := godotenv.Load(envFP)
	if err != nil {
		log.Warnf("no .env file is being used: %s", err.Error())
	}
}

// @title Process-API Server
// @version dev-8.16.23
// @description An OGC compliant process server.

// @contact.name Seth Lawler
// @contact.email slawler@dewberry.com

// @host localhost:5050
// @BasePath /
// @schemes http

// @externalDocs.description  OGC Process API
// @externalDocs.url          https://docs.ogc.org/is/18-062r2/18-062r2.html#toc0

// @externalDocs.description   Schemas
// @externalDocs.url    http://schemas.opengis.net/ogcapi/processes/part1/1.0/openapi/schemas/
func main() {

	// admin := []string{"admin"}
	allUsers := []string{"admin", "reader", "writer"}
	writer := []string{"admin", "writer"}

	// Initialize resources
	rh := handlers.NewRESTHander(pluginsDir, dbPath)
	// todo: handle this error: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running
	// todo: all non terminated job statuses should be updated to unknown
	// todo: all logs in the logs directory should be moved to storage

	// Goroutines
	go rh.StatusUpdateRoutine()
	go rh.JobCompletionRoutine()

	// Set server configuration
	e := echo.New()
	e.Static("/public", "public")

	// e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: e.Logger.Output(),
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
	}))
	e.Logger.SetLevel(log.DEBUG)
	log.SetLevel(log.DEBUG)
	e.Renderer = &rh.T

	// Server
	e.GET("/ping", auth.Authorize(rh.Ping, allUsers...))
	e.GET("/", auth.Authorize(rh.LandingPage, allUsers...))
	e.GET("/swagger/*", auth.Authorize(echoSwagger.WrapHandler, allUsers...))
	e.GET("/conformance", auth.Authorize(rh.Conformance, allUsers...))

	// Processes
	e.GET("/processes", auth.Authorize(rh.ProcessListHandler, allUsers...))
	e.GET("/processes/:processID", auth.Authorize(rh.ProcessDescribeHandler, allUsers...))
	e.POST("/processes/:processID/execution", auth.Authorize(rh.Execution, writer...))

	// TODO
	// e.Post("processes/:processID/new, rh.RegisterNewProcess, admin...))
	// e.Delete("processes/:processID", rh.RegisterNewProcess, admin...))

	// Jobs
	e.GET("/jobs", auth.Authorize(rh.ListJobsHandler, allUsers...))
	e.GET("/jobs/:jobID", auth.Authorize(rh.JobStatusHandler, allUsers...))
	e.GET("/jobs/:jobID/results", auth.Authorize(rh.JobResultsHandler, allUsers...))
	e.GET("/jobs/:jobID/logs", auth.Authorize(rh.JobLogsHandler, allUsers...))
	e.GET("/jobs/:jobID/metadata", auth.Authorize(rh.JobMetaDataHandler, allUsers...))
	e.DELETE("/jobs/:jobID", auth.Authorize(rh.JobDismissHandler, writer...))

	// Callbacks
	e.PUT("/jobs/:jobID/status", auth.Authorize(rh.JobStatusUpdateHandler, allUsers...))
	// e.POST("/jobs/:jobID/results", rh.JobResultsUpdateHandler)

	// Start server
	go func() {
		e.Logger.Info("server starting on port: ", port)
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Error("server error : ", err.Error())
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	<-quit
	e.Logger.Info("gracefully shutting down the server")

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Error(err)
		}
	}()

	// Shutdown the server
	// By default, Docker provides a grace period of 10 seconds with the docker stop command.

	// Kill any running docker containers (clean up resources)
	rh.ActiveJobs.KillAll()
	e.Logger.Info("kill command sent to all active jobs")

	// sleep so that Close() routines spawned by KillAll() can finish writing logs, and updating statuses
	// aws batch jobs close() methods take minimum of 5 seconds
	time.Sleep(5 * time.Second)

	if err := rh.DB.Handle.Close(); err != nil {
		e.Logger.Error(err)
	} else {
		e.Logger.Info("closed connection to database")
	}

	time.Sleep(4 * time.Second)

	e.Logger.Info("server gracefully shutdown")

}
