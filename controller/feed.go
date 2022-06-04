package controller

import (
	"net/http"
	"time"

	"github.com/RaymondCode/simple-demo/common/status_code"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	status_code.Status
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Status:    status_code.Status{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
