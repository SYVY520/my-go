package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func Password(c *gin.Context) {
	err := c.ShouldBindJSON(&User)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select password from users where mailbox = ?"
		rows, errs := db.Query(sqll, User.Mb)
		if errs != nil {
			fmt.Println("查找数据失败")
		} else {
			var password string
			for rows.Next() {
				err = rows.Scan(&password)
			}
			PasswordErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(User.Pw))
			if PasswordErr == nil {
				pass := BcryptPW(User.Cp)
				sqll := "update users set password = ?  where mailbox = ?"
				_, errs := db.Exec(sqll, pass, User.Mb)
				if errs != nil {
					fmt.Println("更新数据失败")
				} else {
					c.JSON(200, gin.H{
						"res": "更改密码成功",
					})
				}
			} else {
				c.JSON(204, gin.H{
					"res": "旧密码错误",
				})
			}
		}
	}
}
func Information(c *gin.Context) {
	err := c.ShouldBindJSON(&User)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select name,phone,country,institution,mailbox,identity,pic from users where mailbox = ?"
		rows, errs := db.Query(sqll, User.Mb)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var users []Us
			var Users Us
			for rows.Next() {
				err = rows.Scan(&Users.Name, &Users.Phone, &Users.Country, &Users.Institution, &Users.Mb, &Users.Identity, &Users.Pic)
				users = append(users, Users)
			}
			c.JSON(200, gin.H{
				"res":  "更改密码成功",
				"rows": Users,
			})
		}
	}
}
func Upload(c *gin.Context) {
	// 单文件上传
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	// 上传文件到指定目录
	//fileName 脱敏
	fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
	newFileName := fileId + path.Ext(file.Filename)
	dst := fmt.Sprintf("./pic/%s", newFileName)
	c.SaveUploadedFile(file, dst)
	paths := "http://localhost:8081/pic/" + newFileName
	mailbox := c.PostForm("mailbox")
	if mailbox == "" {
		fmt.Println("绑定失败", mailbox)
	} else {
		sqll := "select pic from users where mailbox = ?"
		rows, errs := db.Query(sqll, mailbox)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var p string
			for rows.Next() {
				err := rows.Scan(&p)
				fmt.Println(err)
			}
			fmt.Println(p)
			if p == "" {
				sqll := "update users set pic = ?  where mailbox = ?"
				_, errs := db.Query(sqll, paths, mailbox)
				if errs != nil {
					fmt.Println("查找数据失败", errs)
				} else {
					c.JSON(200, gin.H{
						"res": "更改密码成功",
					})
				}
			} else {
				sep := "http://localhost:8081"
				arr := strings.Split(p, sep)
				fmt.Println("arr:", "."+arr[1])
				e := os.Remove("." + arr[1])
				fmt.Println(e)
				sqll := "update users set pic = ?  where mailbox = ?"
				_, errs := db.Query(sqll, paths, User.Mb)
				if errs != nil {
					fmt.Println("查找数据失败", errs)
				} else {
					c.JSON(200, gin.H{
						"res": "更改密码成功",
					})
				}
			}
		}
	}
}
func UploadUser(c *gin.Context) {
	err := c.ShouldBindJSON(&User)
	if err != nil {
		fmt.Println("绑定失败")
	} else {
		fmt.Println("绑定成功", err)
		sqll := "update users set name = ?,phone=?,country=?,institution=?  where mailbox = ?"
		_, errs := db.Query(sqll, User.Name, User.Phone, User.Country, User.Institution, User.Mb)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			fmt.Println("查找数据成功")
			c.JSON(200, gin.H{
				"res": "更改密码成功",
			})
		}
	}
}
