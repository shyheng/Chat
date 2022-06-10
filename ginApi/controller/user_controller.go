package controller

import (
	"ginApi/entity"
	"ginApi/mysql_init"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VueUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
	Token string `json:"token"`
	Fid   string `json:"fid"`
}

func (v VueUser) LoginUser(ctx *gin.Context) {
	err := ctx.Bind(&v)
	if err != nil {
		return
	}

	var user []entity.User
	db := mysql_init.DB
	db.Find(&user)
	for i := range user {
		if user[i].Name == v.Name {
			if user[i].Pass == v.Pass {
				u := user[i]
				u.State = 1
				db.Save(&u)
				ctx.JSON(200, gin.H{
					"token": user[i].Id,
					"flg":   true,
					"msg":   "登录成功",
				})
				return
			}
		}
	}
	ctx.JSON(200, gin.H{
		"flg": false,
		"msg": "账号密码错误",
	})
}

func (v VueUser) AddUser(ctx *gin.Context) {
	err := ctx.Bind(&v)
	if err != nil {
		return
	}
	var user []entity.User
	mysql_init.DB.Find(&user)
	for i := range user {
		if user[i].Name == v.Name {
			ctx.JSON(200, gin.H{
				"flg": false,
				"msg": "用户名重复",
			})
			return
		}
	}
	u := entity.User{Name: v.Name, Pass: v.Pass}
	mysql_init.DB.Create(&u)
	ctx.JSON(200, gin.H{
		"flg": true,
		"msg": "注册成功",
	})
}

func (v VueUser) SelFriend(ctx *gin.Context) {
	err := ctx.Bind(&v)
	if err != nil {
		return
	}
	var friend []entity.Friend

	mysql_init.DB.Where("uid=? and sta=?", v.Token, 1).Find(&friend)
	var fri [10]entity.User
	for i, e := range friend {
		var f entity.User
		mysql_init.DB.Where("id=?", e.Fid).First(&f)
		fri[i] = f
	}
	ctx.JSON(200, gin.H{
		"fri": fri,
	})
}

func (v VueUser) AddFriend(ctx *gin.Context) {
	err := ctx.Bind(&v)
	if err != nil {
		return
	}
	//v.Fid,v.Token
	var f []entity.Friend
	mysql_init.DB.Where("uid=? and fid=?", v.Token, v.Fid).Find(&f)
	i := len(f)
	if i == 1 {
		mysql_init.DB.Where("uid=? and fid=? and sta=?", v.Token, v.Fid, 0).Find(&f)
		i2 := len(f)
		if i2 == 1 {
			ctx.JSON(200, gin.H{
				"flg": false,
				"msg": "对方未同意",
			})
		} else {
			ctx.JSON(200, gin.H{
				"flg": false,
				"msg": "已经添加过对方",
			})
		}
	} else {
		ui, _ := strconv.Atoi(v.Token)
		fi, _ := strconv.Atoi(v.Fid)
		friend := entity.Friend{Uid: ui, Fid: fi}
		mysql_init.DB.Create(&friend)
		ctx.JSON(200, gin.H{
			"flg": true,
			"msg": "添加成功，等待对方接受",
		})
	}
}
