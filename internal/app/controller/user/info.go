package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description 获取用户信息
// @Summary 获取用户信息
// @Tags User
// @Security ApiKeyAuth
// @Security OAuth2Application[write, admin]
// @Accept  json
// @Produce  json
// @Param user body user_serializer.InfoRequest true "获取用户信息"
// @Resource Info
// @Router /user/info [post]
// @Success 200 {object} user_serializer.InfoResponse
func Info(c *gin.Context)  {
	b := &user_serializer.InfoRequest{}

	if err := c.Bind(b);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.ErrBind})
		return
	}




	userService := user.NewUserService()
	r,err := userService.Info(b)
	handler.SendResponse(c,err, r)

}
