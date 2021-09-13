package snippethandler

import (
	"net/http"

	snippetconverter "snippetapp/domain/snippet/converter"
	snippetservice "snippetapp/domain/snippet/service"
	"snippetapp/utils"

	"github.com/gin-gonic/gin"
)

type SnippetHandler struct {
	snippetService   *snippetservice.Service
	snippetConverter *snippetconverter.Converter
}

func NewSnippetHandler(
	snippetService *snippetservice.Service,
	snippetConverter *snippetconverter.Converter,
) *SnippetHandler {
	return &SnippetHandler{
		snippetService:   snippetService,
		snippetConverter: snippetConverter,
	}
}

func (h *SnippetHandler) RegisterRoutes(router *gin.Engine) {
	snippetsApi := router.Group("/snippets")
	snippetsApi.POST("", func(c *gin.Context) { h.createSnippet(c) })
	snippetsApi.GET("/:snippet_name", func(c *gin.Context) { h.getSnippet(c) })
	snippetsApi.POST("/:snippet_name/like", func(c *gin.Context) { h.likeSnippet(c) })
}

func (h *SnippetHandler) createSnippet(c *gin.Context) {
	creationParams, err := h.snippetConverter.ConvertCreateRequestToSnippet(c.Request)
	if err != nil {
		utils.ReturnError(c.Writer, err)
		return
	}
	snippet, err := h.snippetService.Create(c, creationParams.Name, creationParams.Content, creationParams.ExpiresIn)
	if err != nil {
		utils.ReturnError(c.Writer, err)
		return
	}

	utils.ReturnJSONResponseWithStatus(c.Writer, http.StatusCreated, snippet)
}

func (h *SnippetHandler) likeSnippet(c *gin.Context) {
	snippet, err := h.snippetService.Like(c, c.Param("snippet_name"))
	if err != nil {
		utils.ReturnError(c.Writer, err)
		return
	}

	utils.ReturnJSONResponse(c.Writer, snippet)
}

func (h *SnippetHandler) getSnippet(c *gin.Context) {
	snippet, err := h.snippetService.GetByName(c, c.Param("snippet_name"))
	if err != nil {
		utils.ReturnError(c.Writer, err)
		return
	}

	utils.ReturnJSONResponse(c.Writer, snippet)
}
