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

func review(c *gin.Context) {
	// 单文件上传
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	// 上传文件到指定目录
	//fileName 脱敏
	fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
	newFileName := fileId + path.Ext(file.Filename)
	dst := fmt.Sprintf("./report/%s", newFileName)
	c.SaveUploadedFile(file, dst)
	paths := "http://localhost:8081/report/" + newFileName
	manuscript := c.PostForm("manuscript")
	region := c.PostForm("region")
	users := c.PostForm("users")
	sqll := "select report from nm where manuscript = ? and users=?"
	row, errs := db.Query(sqll, manuscript, users)
	if errs != nil {
		fmt.Println("查找数据失败", errs)
	} else {
		var p string
		for row.Next() {
			err := row.Scan(&p)
			fmt.Println(err)
		}
		sql := "update manuscript set statu = '已更新'  where id = ?"
		_, errs := db.Query(sql, manuscript)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			if p == "" {
				sql := "update nm set status = ?,report=?  where manuscript = ? and users=?"
				_, errs := db.Query(sql, region, paths, manuscript, users)
				if errs != nil {
					fmt.Println("查找数据失败", errs)
				} else {
					sql := "select identity from nm where users=? and manuscript=?"
					rows, errs := db.Query(sql, users, manuscript)
					if errs != nil {
						fmt.Println("查找数据失败", errs)
					} else {
						var identity string
						for rows.Next() {
							errs = rows.Scan(&identity)
						}
						if region == "录用" && identity == "一审" {
							sql := "update manuscript set status = '二审'  where id = ?"
							_, errs := db.Query(sql, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else if region == "拒稿" && identity == "一审" {
							sql := "update manuscript set status = '拒稿'  where id = ?"
							_, errs := db.Query(sql, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else if region == "录用" && identity == "三审" {
							sql := "update manuscript set status = '录用'  where id = ?"
							_, errs := db.Query(sql, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else if region == "拒稿" && identity == "三审" {
							sql := "update manuscript set status = '拒稿'  where id = ?"
							_, errs := db.Query(sql, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else if region == "修改" {
							sql := "update nm set statu = '未修改'  where users=? and manuscript=?"
							_, errs := db.Query(sql, users, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								sp1(manuscript, c)
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else {
							sp1(manuscript, c)
							c.JSON(200, gin.H{
								"res": "提交成功",
							})
						}
					}
				}
			} else {
				sep := "http://localhost:8081"
				arr := strings.Split(p, sep)
				e := os.Remove("." + arr[1])
				fmt.Println(e)
				sql := "update nm set status = ?,report=?  where manuscript = ? and users=?"
				_, errs := db.Query(sql, region, paths, manuscript, users)
				if errs != nil {
					fmt.Println("查找数据失败", errs)
				} else {
					sql := "select identity from nm where users=? and manuscript=?"
					rows, errs := db.Query(sql, users, manuscript)
					if errs != nil {
						fmt.Println("查找数据失败", errs)
					} else {
						var identity string
						for rows.Next() {
							errs = rows.Scan(&identity)
						}
						if region == "录用" && identity == "一审" {
							sql := "update manuscript set status = '二审'  where id = ?"
							_, errs := db.Query(sql, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else if region == "拒稿" && identity == "一审" {
							sql := "update manuscript set status = '拒稿'  where id = ?"
							_, errs := db.Query(sql, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else if region == "录用" && identity == "三审" {
							sql := "update manuscript set status = '录用'  where id = ?"
							_, errs := db.Query(sql, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else if region == "拒稿" && identity == "三审" {
							sql := "update manuscript set status = '拒稿'  where id = ?"
							_, errs := db.Query(sql, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else if region == "修改" {
							sql := "update nm set statu = '未修改'  where users=? and manuscript=?"
							_, errs := db.Query(sql, users, manuscript)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								sp1(manuscript, c)
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else {
							sp1(manuscript, c)
							c.JSON(200, gin.H{
								"res": "提交成功",
							})
						}
					}
				}
			}
		}
	}
}
func sp(c *gin.Context) {
	err := c.ShouldBindJSON(&M)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select status from manuscript where id = ?"
		rows, errs := db.Query(sqll, M.Id)
		if errs != nil {
			fmt.Println("查找数据失败", errs)
		} else {
			var status string
			for rows.Next() {
				err = rows.Scan(&status)
			}
			if status == "二审" {
				sqll := "select status from nm where manuscript = ? and identity='二审'"
				rows, errs := db.Query(sqll, M.Id)
				if errs != nil {
					fmt.Println("查找数据失败", errs)
				} else {
					var nms []nm
					for rows.Next() {
						var Nm nm
						err = rows.Scan(&Nm.Status)
						nms = append(nms, Nm)
					}
					if nms[0].Status != "" && nms[1].Status != "" && nms[2].Status != "" {
						if (nms[0].Status == "拒稿" && nms[1].Status == "拒稿") || (nms[1].Status == "拒稿" && nms[2].Status == "拒稿") || (nms[0].Status == "拒稿" && nms[2].Status == "拒稿") {
							sql := "update manuscript set status = '拒稿'  where id = ?"
							_, errs := db.Query(sql, M.Id)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						}
						if (nms[0].Status == "录用" && nms[1].Status == "录用") || (nms[1].Status == "录用" && nms[2].Status == "录用") || (nms[0].Status == "录用" && nms[2].Status == "录用") {
							sql := "update manuscript set status = '三审'  where id = ?"
							_, errs := db.Query(sql, M.Id)
							if errs != nil {
								fmt.Println("查找数据失败", errs)
							} else {
								c.JSON(200, gin.H{
									"res": "提交成功",
								})
							}
						} else {
							c.JSON(200, gin.H{
								"res": "更改密码成功",
							})
						}
					} else {
						c.JSON(200, gin.H{
							"res": "更改密码成功",
						})
					}
				}
			} else {
				c.JSON(200, gin.H{
					"res": "更改密码成功",
				})
			}
		}
	}
}

func sp1(m string, c *gin.Context) {
	sqll := "select status from manuscript where id = ?"
	rows, errs := db.Query(sqll, m)
	if errs != nil {
		fmt.Println("查找数据失败", errs)
	} else {
		var status string
		for rows.Next() {
			err := rows.Scan(&status)
			fmt.Println(err)
		}
		if status == "二审" {
			sqll := "select status from nm where manuscript = ? and identity='二审'"
			rows, errs := db.Query(sqll, m)
			if errs != nil {
				fmt.Println("查找数据失败", errs)
			} else {
				var nms []nm
				for rows.Next() {
					var Nm nm
					err := rows.Scan(&Nm.Status)
					nms = append(nms, Nm)
					fmt.Println(err)
				}
				if nms[0].Status != "" && nms[1].Status != "" && nms[2].Status != "" {
					if (nms[0].Status == "拒稿" && nms[1].Status == "拒稿") || (nms[1].Status == "拒稿" && nms[2].Status == "拒稿") || (nms[0].Status == "拒稿" && nms[2].Status == "拒稿") {
						sql := "update manuscript set status = '拒稿'  where id = ?"
						_, errs := db.Query(sql, m)
						if errs != nil {
							fmt.Println("查找数据失败", errs)
						} else {
						}
					}
					if (nms[0].Status == "录用" && nms[1].Status == "录用") || (nms[1].Status == "录用" && nms[2].Status == "录用") || (nms[0].Status == "录用" && nms[2].Status == "录用") {
						sql := "update manuscript set status = '三审'  where id = ?"
						_, errs := db.Query(sql, m)
						if errs != nil {
							fmt.Println("查找数据失败", errs)
						} else {
						}
					} else {
					}
				} else {
				}
			}
		} else {
		}
	}
}
