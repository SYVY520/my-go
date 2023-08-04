package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func uploadeditor(c *gin.Context) {
	err := c.ShouldBindJSON(&User)
	if err != nil {
		fmt.Println("绑定失败")
	} else {
		sqll := "update users set identity ='editor' where mailbox = ?"
		_, errs := db.Query(sqll, User.Mb)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			c.JSON(200, gin.H{
				"res": "更改密码成功",
			})
		}
	}
}
func uploadreview(c *gin.Context) {
	err := c.ShouldBindJSON(&User)
	if err != nil {
		fmt.Println("绑定失败")
	} else {
		sqll := "update users set identity ='review' where mailbox = ?"
		_, errs := db.Query(sqll, User.Mb)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			c.JSON(200, gin.H{
				"res": "更改密码成功",
			})
		}
	}
}
func Remove(c *gin.Context) {
	err := c.ShouldBindJSON(&User)
	if err != nil {
		fmt.Println("绑定失败")
	} else {
		sqll := "update users set identity ='contributors' where id = ?"
		_, errs := db.Query(sqll, User.Id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			c.JSON(200, gin.H{
				"res": "更改密码成功",
			})
		}
	}
}
