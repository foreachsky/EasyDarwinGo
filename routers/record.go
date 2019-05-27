package routers

import (
	"github.com/EasyDarwin/EasyDarwin/models"
	"github.com/EasyDarwin/EasyDarwin/rtsp"
	"github.com/gin-gonic/gin"
)

// StartRecordResponse of StartRecordRequest
type StartRecordResponse struct {
	Code int    `form:"code";json:"code"`
	Msg  string `form:"msg";json:"msg"`
}

// StartRecord instace
func (h *APIHandler) StartRecord(c *gin.Context) {
	req := &models.Record{}
	if err := c.Bind(req); err != nil {
		c.IndentedJSON(200, &StartRecordResponse{
			Code: 400,
			Msg:  "Bad request",
		})
		return
	}

	pusher := rtsp.Instance.GetPusher(req.PlayPath)
	if nil == pusher {
		c.IndentedJSON(200, &StartRecordResponse{
			Code: 404,
			Msg:  "Media source not found according to playpath",
		})
		return
	}

	recorder := rtsp.NewRecorder(req.ID, pusher)
	if err := pusher.AddPlayer(recorder); nil != err {
		c.IndentedJSON(200, &StartRecordResponse{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	c.IndentedJSON(200, &StartRecordResponse{
		Code: 0,
		Msg:  "OK",
	})
}
