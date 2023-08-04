package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func thesis(c *gin.Context) {
	// 单文件上传
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	// 上传文件到指定目录
	//fileName 脱敏
	fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
	newFileName := fileId + path.Ext(file.Filename)
	dst := fmt.Sprintf("./file/%s", newFileName)
	c.SaveUploadedFile(file, dst)
	paths := "http://localhost:8081/file/" + newFileName
	name := c.PostForm("name")
	region := c.PostForm("region")
	first := c.PostForm("first")
	second := c.PostForm("second")
	third := c.PostForm("third")
	fourth := c.PostForm("fourth")
	editmailbox := c.PostForm("editmailbox")
	desc := c.PostForm("desc")
	mailbox := c.PostForm("mailbox")
	data := c.PostForm("data")
	sql := "select name from users where mailbox = ?"
	row, errs := db.Query(sql, editmailbox)
	var edit string
	for row.Next() {
		err := row.Scan(&edit)
		fmt.Println(err)
	}
	sqll := "insert into manuscript(name,data,mailbox,region,first,second,third,fourth,editmailbox,`desc`,file,edit) values (?,?,?,?,?,?,?,?,?,?,?,?)"
	_, errs = db.Query(sqll, name, data, mailbox, region, first, second, third, fourth, editmailbox, desc, paths, edit)
	if errs != nil {
		fmt.Println("查找数据失败", errs)
	} else {
		c.JSON(200, gin.H{
			"res": "提交成功",
		})
	}
}
func editor(c *gin.Context) {
	sqll := "select name,mailbox from users where identity='editor'"
	rows, errs := db.Query(sqll)
	if errs != nil {
		fmt.Println("查找数据失败", errs)
	} else {
		var users []Us
		for rows.Next() {
			var Users Us
			err := rows.Scan(&Users.Name, &Users.Mb)
			users = append(users, Users)
			fmt.Println(err)
		}
		c.JSON(200, gin.H{
			"res":  "更改密码成功",
			"rows": users,
		})
	}
}
func thesis0(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "select name,data,status,region,first,second,third,fourth,`desc`,file,edit,mailbox from manuscript where id = ?"
		rows, errs := db.Query(sql, N.Id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var ms []manuscript
			for rows.Next() {
				var Ms manuscript
				err = rows.Scan(&Ms.Name, &Ms.Data, &Ms.Status, &Ms.Region, &Ms.First, &Ms.Second, &Ms.Third, &Ms.Fourth, &Ms.Desc, &Ms.File, &Ms.Edit, &Ms.Mailbox)
				ms = append(ms, Ms)
			}
			c.JSON(200, gin.H{
				"res":  "提交成功",
				"data": ms,
			})
		}
	}
}
func thesis1(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "select status,report from nm where manuscript = ? and identity='一审'"
		rows, errs := db.Query(sql, N.Id)
		var nms []nm
		for rows.Next() {
			var Nm nm
			err = rows.Scan(&Nm.Status, &Nm.Report)
			nms = append(nms, Nm)
		}
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			c.JSON(200, gin.H{
				"res":  "提交成功",
				"data": nms,
			})
		}
	}
}
func thesis2(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "select users,status,report from nm where manuscript = ? and identity='二审'"
		rows, errs := db.Query(sql, N.Id)
		var nms []nm
		for rows.Next() {
			var Nm nm
			err = rows.Scan(&Nm.Users, &Nm.Status, &Nm.Report)
			nms = append(nms, Nm)
		}
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			c.JSON(200, gin.H{
				"res":  "提交成功",
				"data": nms,
			})
		}
	}
}
func thesis3(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "select status,report from nm where manuscript = ? and identity='三审'"
		rows, errs := db.Query(sql, N.Id)
		var nms []nm
		for rows.Next() {
			var Nm nm
			err = rows.Scan(&Nm.Status, &Nm.Report)
			nms = append(nms, Nm)
		}
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			c.JSON(200, gin.H{
				"res":  "提交成功",
				"data": nms,
			})
		}
	}
}
func cj(c *gin.Context) {
	// 单文件上传
	file, _ := c.FormFile("file")
	// 上传文件到指定目录
	//fileName 脱敏
	fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
	newFileName := fileId + path.Ext(file.Filename)
	dst := fmt.Sprintf("./file/%s", newFileName)
	c.SaveUploadedFile(file, dst)
	paths := "http://localhost:8081/file/" + newFileName
	id := c.PostForm("id")
	status := c.PostForm("status")
	if id == "" {
		fmt.Println("绑定失败", id)
	} else {
		sqll := "select file from manuscript where id = ?"
		row, errs := db.Query(sqll, id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var p string
			for row.Next() {
				err := row.Scan(&p)
				fmt.Println(err)
			}
			sep := "http://localhost:8081"
			arr := strings.Split(p, sep)
			e := os.Remove("." + arr[1])
			fmt.Println(e)
			sqll := "update manuscript set file = ?  where id = ?"
			_, errs := db.Query(sqll, paths, id)
			if errs != nil {
				fmt.Println("查找数据失败", errs)
			} else {
				sqll := "update nm set statu='已修改' where manuscript = ? and identity=?"
				_, errs := db.Query(sqll, id, status)
				if errs != nil {
					fmt.Println("查找数据失败", errs)
				} else {

					c.JSON(200, gin.H{
						"res": "提交成功",
					})
				}
			}
		}
	}
}
func th1(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "select id from users where mailbox in (select users from nm where manuscript = ? and identity='一审')"
		rows, errs := db.Query(sql, N.Id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var nms []nm
			for rows.Next() {
				var Nm nm
				err = rows.Scan(&Nm.Id)
				nms = append(nms, Nm)
			}
			c.JSON(200, gin.H{
				"res":  "提交成功",
				"data": nms,
			})
		}
	}
}
func th2(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "select id from users where mailbox in (select users from nm where manuscript = ? and identity='二审')"
		rows, errs := db.Query(sql, N.Id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var nms []nm
			for rows.Next() {
				var Nm nm
				err = rows.Scan(&Nm.Id)
				nms = append(nms, Nm)
			}
			c.JSON(200, gin.H{
				"res":  "提交成功",
				"data": nms,
			})
		}
	}
}
func th3(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "select id from users where mailbox in (select users from nm where manuscript = ? and identity='三审')"
		rows, errs := db.Query(sql, N.Id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var nms []nm
			for rows.Next() {
				var Nm nm
				err = rows.Scan(&Nm.Id)
				nms = append(nms, Nm)
			}
			c.JSON(200, gin.H{
				"res":  "提交成功",
				"data": nms,
			})
		}
	}
}
func tj(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "select mailbox from users where id=?"
		rows, errs := db.Query(sql, N.Id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var mailbax string
			for rows.Next() {
				err = rows.Scan(&mailbax)
			}
			sql := "insert into nm(users,manuscript,identity) values (?,?,?)"
			_, errs := db.Query(sql, mailbax, N.Manuscript, N.Identity)
			if errs != nil {
				fmt.Println("查找数据失败", errs)
			} else {
				if N.Identity == "一审" {
					sqll := "update manuscript set status = ?  where id = ?"
					_, errs := db.Query(sqll, N.Identity, N.Manuscript)
					if errs != nil {
						fmt.Println("查找数据失败", errs)
					} else {
						c.JSON(200, gin.H{
							"res": "提交成功",
						})
					}
				} else {
					c.JSON(200, gin.H{
						"res": "提交成功",
					})
				}
			}
		}
	}
}
func jj(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "select identity from nm where users=? and manuscript=?"
		rows, errs := db.Query(sql, N.Users, N.Manuscript)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var identity string
			for rows.Next() {
				err = rows.Scan(&identity)
			}
			sql := "delete from nm where users=? and manuscript=?"
			_, errs := db.Query(sql, N.Users, N.Manuscript)
			if errs != nil {
				fmt.Println("查找数据失败", errs)
			} else {
				if identity == "一审" {
					sqll := "update manuscript set status = '未审'  where id = ?"
					_, errs := db.Query(sqll, N.Manuscript)
					if errs != nil {
						fmt.Println("查找数据失败", errs)
					} else {
						c.JSON(200, gin.H{
							"res": "提交成功",
						})
					}
				} else {
					c.JSON(200, gin.H{
						"res": "提交成功",
					})
				}
			}
		}
	}
}
func gx(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sql := "update manuscript set statu='未更新' where id = ?"
		_, errs := db.Query(sql, N.Id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			c.JSON(200, gin.H{
				"res": "提交成功",
			})
		}
	}
}
