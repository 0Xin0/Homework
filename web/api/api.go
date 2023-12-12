package api

import (
	"gin-demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(c *gin.Contaxt) {
	//传入用户名和密码
	username := c.PostForm("username")
	password := c.Postform("password")

	//验证用胡名是否重复
	flag := dao.SelectUser(username)
	//重复则退出
	if flag {
		//以JSON格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"massage": "user already exists",
		})
		return
	}

	dao.AddUser(username, password)
	//以JSON格式返回信息
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  200,
		"massage": "add user name successful",
	})
}

func login(c *gin.Context) {
	//传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	}

	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}

	// 正确则登录成功 设置 cookie
	c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successful",
	})
}
