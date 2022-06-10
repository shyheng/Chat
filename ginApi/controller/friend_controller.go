package controller

import (
	"ginApi/entity"
	"ginApi/mysql_init"
	"github.com/gin-gonic/gin"
)

type VueFriend struct {
	Token string `json:"token"`
	Fid   string `json:"fid"`
}

func (v VueFriend) AddFriend(ctx *gin.Context) {
	err := ctx.Bind(&v)
	if err != nil {
		return
	}
	var f entity.Friend
	mysql_init.DB.Where("fid=? and sta=?", v.Token, 0).Find(&f)
}
