package server

import (
	snippetconverter "snippetapp/domain/snippet/converter"
	snippethandler "snippetapp/domain/snippet/handler"
	snippetrepository "snippetapp/domain/snippet/repository"
	snippetservice "snippetapp/domain/snippet/service"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type Server struct {
	snippetService   *snippetservice.Service
	snippetConverter *snippetconverter.Converter
}

type Runnable struct{}

func NewRunnable() *Runnable {
	return &Runnable{}
}

func (r *Runnable) Cmd() *cobra.Command {
	cmd.Run = func(_ *cobra.Command, _ []string) {
		server := r.Run()
		server.Start()
	}
	return cmd
}

func (r *Runnable) Run() *Server {
	return newServer()
}

func newServer() *Server {

	snippetRepository := snippetrepository.NewRepository()
	snippetService := snippetservice.NewSnippetService(snippetRepository)

	app := &Server{
		snippetService:   snippetService,
		snippetConverter: snippetconverter.NewSnippetConverter(),
	}

	return app
}

func (s *Server) Start() {
	router := gin.Default()
	s.registerRoutes(router)

	if err := router.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) registerRoutes(router *gin.Engine) {

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowedMethods:  []string{"*"},
		AllowedHeaders:  []string{"*"},
	}))

	snippetHandler := snippethandler.NewSnippetHandler(s.snippetService, s.snippetConverter)
	snippetHandler.RegisterRoutes(router)
}
