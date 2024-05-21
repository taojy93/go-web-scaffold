package handlers

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"go-web-scaffold/internal/models"
	"go-web-scaffold/internal/response"
	"go-web-scaffold/internal/service"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	newsService service.NewsService
}

func NewNewsHandler(newsService service.NewsService) *NewsHandler {
	return &NewsHandler{newsService: newsService}
}

func (h *NewsHandler) CreateNews(c *gin.Context) {
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		response.JSONResponse(c, http.StatusBadRequest, response.Error(http.StatusBadRequest, "Invalid input"))
		return
	}

	createdNews, err := h.newsService.CreateNews(&news)
	if err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(http.StatusInternalServerError, "Failed to create news"))
		return
	}

	response.JSONResponse(c, http.StatusOK, response.Success(createdNews))
}

func (h *NewsHandler) GetNewsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JSONResponse(c, http.StatusBadRequest, response.Error(http.StatusBadRequest, "Invalid ID"))
		return
	}

	news, err := h.newsService.GetNewsByID(uint(id))
	if err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(http.StatusInternalServerError, "Failed to get news"))
		return
	}

	response.JSONResponse(c, http.StatusOK, response.Success(news))
}

func (h *NewsHandler) GetAllNews(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	newsList, err := h.newsService.GetAllNews()
	if err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(http.StatusInternalServerError, "Failed to get news list"))
		return
	}
	// 获取 http.Flusher 接口以便刷新缓冲区
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(http.StatusInternalServerError, "Streaming unsupported!"))
		return
	}

	c.Stream(func(w io.Writer) bool {
		for _, news := range newsList {
			// 使用SSEvent方法发送事件
			c.SSEvent("data: ", news)
			// 刷新缓冲区，确保数据立即发送到客户端
			flusher.Flush()
			// 为了演示效果，每条数据之间加一个延时
			time.Sleep(1 * time.Second)
		}
		return false
	})
}

func (h *NewsHandler) UpdateNews(c *gin.Context) {
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		response.JSONResponse(c, http.StatusBadRequest, response.Error(http.StatusBadRequest, "Invalid input"))
		return
	}

	updatedNews, err := h.newsService.UpdateNews(&news)
	if err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(http.StatusInternalServerError, "Failed to update news"))
		return
	}

	response.JSONResponse(c, http.StatusOK, response.Success(updatedNews))
}

func (h *NewsHandler) DeleteNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JSONResponse(c, http.StatusBadRequest, response.Error(http.StatusBadRequest, "Invalid ID"))
		return
	}

	if err := h.newsService.DeleteNews(uint(id)); err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(http.StatusInternalServerError, "Failed to delete news"))
		return
	}

	response.JSONResponse(c, http.StatusOK, response.Success(nil))
}
