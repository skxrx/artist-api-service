package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/skxrx/artist-api-service/proto/parser"
)

type Handler struct {
	client pb.ParserServiceClient
}

func NewHandler(client pb.ParserServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) GetEventsHandler(c *gin.Context) {
	artistName := c.Param("artistName")
	req := &pb.GetEventsRequest{ArtistName: artistName}
	resp, err := h.client.GetEvents(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Events)
}
