package service

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// DataFilepath is where the data input file is.
const DataFilepath = "./data/input.txt"

type Server struct {
	router *gin.Engine
	log    *slog.Logger
	port   int
	data   []int
}

type SuccessResponse struct {
	Target  int
	Index   int
	Value   int
	Message string
}

type ErrorResponse struct {
	Message string
}

// NewServer sets up and returns the server instance.
func NewServer(logger *slog.Logger, dataFilepath string) (*Server, error) {
	s := Server{
		router: gin.Default(),
		log:    logger,
		port:   viper.GetInt("port"),
	}
	data, err := LoadNumbers(dataFilepath)
	if err != nil {
		return &s, fmt.Errorf("cannot load data: %w", err)
	}
	s.log.Debug("data successfully loaded")
	s.data = data
	s.router.GET("/endpoint/:target", s.GetIndexHandler)
	return &s, nil
}

// Serve starts the server.
func (s *Server) Serve() error {
	log.Printf("Serving on port %d, log level set to %s\n", s.port, viper.GetString("log_level"))
	return s.router.Run(fmt.Sprintf(":%d", s.port))
}

// GetIndexHandler retrieves information on search result for target value.
func (s *Server) GetIndexHandler(c *gin.Context) {
	targetStr := c.Param("target")
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		s.log.Error("bad request, requested value must be an integer", "value", targetStr)
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("Bad input value: '%s'", targetStr)})
		return
	}
	index, value := FindIndex(s.data, target)
	if index == -1 {
		s.log.Error("requested value not found", "value", target)
		c.JSON(http.StatusNotFound, ErrorResponse{Message: fmt.Sprintf("Value not found: %d", target)})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{
		Target:  target,
		Index:   index,
		Value:   value,
		Message: "Value found",
	})
}
