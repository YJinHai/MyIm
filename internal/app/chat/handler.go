package chat

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/YJinHai/MyIm/internal"
	"github.com/YJinHai/MyIm/internal/pkg/errno"
)

func Register(c *gin.Context)  {
	b := &RegisterReq{}

	if err := c.Bind(b);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.ErrBind})
		return
	}

	chat := NewAPI()
	r,err := chat.Register(b)
	internal.SendResponse(c,err, r)

}


func ListMembers(c *gin.Context)  {
	name := c.Param("name")
	secret := c.Query("secret")

	chat := NewAPI()
	r,err := chat.ListMembers(name, secret)
	internal.SendResponse(c,err, r)

}

func CreateChannel(c *gin.Context)  {
	b := &CreateReq{}

	if err := c.Bind(b);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.ErrBind})
		return
	}

	chat := NewAPI()
	r,err := chat.CreateChannel(b)
	internal.SendResponse(c,err, r)

}

func UnreadCount(c *gin.Context)  {
	uid := c.Param("uid")
	chanName := c.Param("chanName")

	id,error := strconv.Atoi(uid)
	if error != nil{
		log.Println(error)
	}

	chat := NewAPI()
	r,err := chat.unreadCount(id, chanName)
	internal.SendResponse(c,err, r)

}

func ListChannels(c *gin.Context)  {
	chat := NewAPI()
	r,err := chat.ListChannels()
	internal.SendResponse(c,err, r)

}

