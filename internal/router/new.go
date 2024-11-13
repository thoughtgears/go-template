package router

import (
	"fmt"

	"github.com/thoughtgears/go-template/internal/config"
	"github.com/thoughtgears/go-template/internal/router/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	Port   string
	Debug  bool
}

func NewRouter(config *config.Config) *Router {
	var r Router

	// Has to be run before the gin.Default() call
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initiate the gin router and add generic middleware
	// such as logging and recovery
	r.engine = gin.Default()
	r.engine.Use(middleware.Logger())
	r.engine.Use(gin.Recovery())
	r.Port = config.Port
	r.Debug = config.Debug

	return &r
}

// Run starts the gin server and returns an error if it fails to start.
// It also sets up the required clients for the server to access the data
// from the GitHub API.
func (r *Router) Run() error {
	var err error

	host := "0.0.0.0"
	if r.Debug {
		host = "127.0.0.1"
	}

	if err = r.engine.Run(fmt.Sprintf("%s:%s", host, r.Port)); err != nil {
		return fmt.Errorf("starting gin server: %w", err)
	}

	return nil
}
