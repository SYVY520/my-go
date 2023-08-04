package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type manuscript struct {
	Name        string `json:"name"`
	Data        string `json:"data"`
	Status      string `json:"status"`
	Mailbox     string `json:"mailbox"`
	Id          string `json:"id"`
	Region      string `json:"region"`
	Edit        string `json:"edit"`
	Editmailbox string `json:"editmailbox"`
	First       string `json:"first"`
	Second      string `json:"second"`
	Third       string `json:"third"`
	Fourth      string `json:"fourth"`
	Desc        string `json:"desc"`
	File        string `json:"file"`
	Identity    string `json:"identity"`
	Nmregion    string `json:"nmregion"`
	Nmidentity  string `json:"nmidentity"`
	Count       string `json:"count"`
	Statu       string `json:"statu"`
}
type nm struct {
	Status      string `json:"status"`
	Editmailbox string `json:"editmailbox"`
	Identity    string `json:"identity"`
	Report      string `json:"report"`
	Manuscript  string `json:"manuscript"`
	Users       string `json:"users"`
	Id          string `json:"id"`
}

var M manuscript
var N nm

func Inquiry(c *gin.Context) {
	err := c.ShouldBindJSON(&User)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select id,name,data,status,region,edit,statu from manuscript where mailbox = ?"
		rows, errs := db.Query(sqll, User.Mb)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var ms []manuscript
			for rows.Next() {
				var Ms manuscript
				err = rows.Scan(&Ms.Id, &Ms.Name, &Ms.Data, &Ms.Status, &Ms.Region, &Ms.Edit, &Ms.Statu)
				ms = append(ms, Ms)
			}
			c.JSON(200, gin.H{
				"res":  "更改密码成功",
				"rows": ms,
			})
		}
	}
}

//主编未审稿件

func Inquiry0(c *gin.Context) {
	err := c.ShouldBindJSON(&M)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select id,name,data,status,region,edit from manuscript where editmailbox = ? and status!='录用' and status!='拒稿'"
		rows, errs := db.Query(sqll, M.Editmailbox)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var ms []manuscript
			for rows.Next() {
				var Ms manuscript
				err = rows.Scan(&Ms.Id, &Ms.Name, &Ms.Data, &Ms.Status, &Ms.Region, &Ms.Edit)
				sqll = "select count(*) from nm where manuscript = ? and identity=?"
				row, errss := db.Query(sqll, Ms.Id, Ms.Status)
				for row.Next() {
					err = row.Scan(&Ms.Count)
				}
				ms = append(ms, Ms)
				fmt.Println(errss)
			}
			c.JSON(200, gin.H{
				"res":  "更改密码成功",
				"rows": ms,
			})
		}
	}
}
func Inquiry1(c *gin.Context) {
	err := c.ShouldBindJSON(&M)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select id,name,data,status,region,edit from manuscript where editmailbox = ? and status='录用' or status='拒稿'"
		rows, errs := db.Query(sqll, M.Editmailbox)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var ms []manuscript
			for rows.Next() {
				var Ms manuscript
				err = rows.Scan(&Ms.Id, &Ms.Name, &Ms.Data, &Ms.Status, &Ms.Region, &Ms.Edit)
				ms = append(ms, Ms)
			}
			c.JSON(200, gin.H{
				"res":  "更改密码成功",
				"rows": ms,
			})
		}
	}
}
func Inquiry2(c *gin.Context) {
	sqll := "select id,name,country,institution,identity from users where identity!='contributors'"
	rows, errs := db.Query(sqll)

	if errs != nil {
		fmt.Println("查找数据失败", errs)
	} else {
		var users []Us
		for rows.Next() {
			var Users Us
			err := rows.Scan(&Users.Id, &Users.Name, &Users.Country, &Users.Institution, &Users.Identity)
			users = append(users, Users)
			log.Println(err)
		}
		c.JSON(200, gin.H{
			"res":  "更改密码成功",
			"rows": users,
		})
	}
}
func Inquiry3(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select manuscript.id,manuscript.name,data,edit,region,nm.status from manuscript,nm  where manuscript.id in (select manuscript from nm where users = ? and (nm.status='录用' or nm.status='拒稿'))  and nm.manuscript=manuscript.id and users=?"
		rows, errs := db.Query(sqll, N.Editmailbox, N.Editmailbox)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var ms []manuscript
			for rows.Next() {
				var Ms manuscript
				err := rows.Scan(&Ms.Id, &Ms.Name, &Ms.Data, &Ms.Edit, &Ms.Region, &Ms.Nmregion)
				ms = append(ms, Ms)
				fmt.Println(err)
			}
			c.JSON(200, gin.H{
				"res":  "更改密码成功",
				"rows": ms,
			})
		}
	}
}
func Inquiry4(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select manuscript.id,manuscript.name,data,edit,region,nm.identity from manuscript,nm where manuscript.id in (select manuscript from nm where users = ? and nm.status is null) and nm.manuscript=manuscript.id and nm.users=?"
		rows, errs := db.Query(sqll, N.Editmailbox, N.Editmailbox)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var ms []manuscript
			for rows.Next() {
				var Ms manuscript
				err := rows.Scan(&Ms.Id, &Ms.Name, &Ms.Data, &Ms.Edit, &Ms.Region, &Ms.Nmidentity)
				ms = append(ms, Ms)
				fmt.Println(err)
			}
			c.JSON(200, gin.H{
				"res":  "更改密码成功",
				"rows": ms,
			})
		}
	}
}
func Inquiry5(c *gin.Context) {
	err := c.ShouldBindJSON(&N)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select manuscript.id,manuscript.name,data,edit,region,nm.identity,nm.statu from manuscript,nm where manuscript.id in (select manuscript from nm where users = ? and nm.status='修改') and nm.manuscript=manuscript.id and nm.users=?"
		rows, errs := db.Query(sqll, N.Editmailbox, N.Editmailbox)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var ms []manuscript
			for rows.Next() {
				var Ms manuscript
				err := rows.Scan(&Ms.Id, &Ms.Name, &Ms.Data, &Ms.Edit, &Ms.Region, &Ms.Nmidentity, &Ms.Statu)
				ms = append(ms, Ms)
				fmt.Println(err)
			}
			c.JSON(200, gin.H{
				"res":  "更改密码成功",
				"rows": ms,
			})
		}
	}
}
func Inquirys(c *gin.Context) {
	err := c.ShouldBindJSON(&M)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select region from manuscript where id = ?"
		rows, errs := db.Query(sqll, M.Id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var region string
			for rows.Next() {
				err = rows.Scan(&region)
			}
			sqll := "select id,name,country,institution,mailbox from users where identity = 'review'"
			rows, errs := db.Query(sqll)
			if errs != nil {
				fmt.Println("查找数据失败", errs)
			} else {
				var users []Us
				for rows.Next() {
					var Users Us
					err = rows.Scan(&Users.Id, &Users.Name, &Users.Country, &Users.Institution, &Users.Mb)
					sqll = "select count(*) from nm where users = ? and status is null"
					row, errss := db.Query(sqll, Users.Mb)
					for row.Next() {
						err = row.Scan(&Users.Count)
					}
					sqll = "select count(*) from nm where users = ? and status is not null and manuscript in (select id from manuscript where region=?)"
					row, errss = db.Query(sqll, Users.Mb, region)
					for row.Next() {
						err = row.Scan(&Users.Counts)
						fmt.Println(errss)
					}
					count, _ := strconv.ParseFloat(Users.Count, 64)
					counts, _ := strconv.ParseFloat(Users.Counts, 64)
					Users.Px = strconv.FormatFloat(counts*1.5-count, 'f', 2, 64)
					users = append(users, Users)
				}
				c.JSON(200, gin.H{
					"res":  "更改密码成功",
					"rows": users,
				})
			}
		}
	}
}
