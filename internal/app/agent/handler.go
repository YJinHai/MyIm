package agent

import (
	"github.com/gin-gonic/gin"
)

func WSConnect(c *gin.Context)  {
	ws := NewAPI()
	ws.Connect(c.Writer, c.Request)
	//internal.SendResponse(c,err, nil)

}
