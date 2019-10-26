package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/YJinHai/MyIm/internal"
	"github.com/YJinHai/MyIm/internal/pkg/errno"

)

// @Description 获取用户信息
// @Summary 获取用户信息
// @Tags User
// @Security ApiKeyAuth
// @Security OAuth2Application[write, admin]
// @Accept  json
// @Produce  json
// @Param user body InfoRequest true "获取用户信息"
// @Resource Info
// @Router /user/info [post]
// @Success 200 {object} nfoResponse
func Info(c *gin.Context)  {
	b := &InfoRequest{}

	if err := c.Bind(b);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.ErrBind})
		return
	}

	userService := NewUserService()
	r,err := userService.Info(b)
	internal.SendResponse(c,err, r)

}

func Login(c *gin.Context)  {
	b := &LoginRequest{}

	if err := c.Bind(b);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.ErrBind})
		return
	}

	log.Println("b:",b)

	userService := NewUserService()
	r,err := userService.Login(b)
	internal.SendResponse(c,err, r)

}

func Register (c *gin.Context)  {
	b := &RegisterRequest{}

	if err := c.Bind(b);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.ErrBind})
		return
	}

	log.Println("b:",b)

	userService := NewUserService()
	r,err := userService.Register(b)
	internal.SendResponse(c,err, r)

}