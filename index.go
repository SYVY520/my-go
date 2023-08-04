package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //初始化
	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type Us struct {
	Name        string `json:"name"`
	Mb          string `json:"mailbox"`
	Phone       string `json:"phone"`
	Country     string `json:"country"`
	Institution string `json:"institution"`
	Pw          string `json:"password"`
	Cd          string `json:"code"`
	Cp          string `json:"checkPass"`
	Id          string `json:"id"`
	Identity    string `json:"identity"`
	Pic         string `json:"pic"`
	Count       string `json:"count"`
	Counts      string `json:"counts"`
	Token       string `json:"token"`
	Px          string `json:"px"`
}

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"mailbox"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("123456")
var User Us
var db *sql.DB

// 设置自带的 store（可以配置成redis）
var store = base64Captcha.DefaultMemStore

// BcryptPW 生成密码
func BcryptPW(password string) string {
	const cost = 10 //加密级别系数，越大越安全但性能开下也随之增大
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}
func Enroll(c *gin.Context) {
	err := c.ShouldBindJSON(&User)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		sqll := "select mailbox from users where mailbox = ?"
		rows, errs := db.Query(sqll, User.Mb)
		if errs != nil {
			fmt.Println("查找数据失败")
		} else {
			var mailbox string
			for rows.Next() {
				err = rows.Scan(&mailbox)
			}
			if mailbox == "" {
				pass := BcryptPW(User.Pw)
				sqls := "insert into users(mailbox,password) values (?,?)"
				_, errs := db.Query(sqls, User.Mb, pass)
				if errs != nil {
					fmt.Println("插入数据失败")
				} else {
					c.JSON(200, gin.H{
						"res": "注册成功",
					})
				}
			} else {
				c.JSON(204, gin.H{
					"res": "用户已存在",
				})
			}
		}
	}
}
func Login(c *gin.Context) {
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
				tokenString, _ := GenToken(User.Mb)
				c.JSON(200, gin.H{
					"res":  "登录成功",
					"data": gin.H{"token": tokenString},
				})
			} else {
				c.JSON(204, gin.H{
					"res": "密码错误",
				})
			}
		}
	}
}
func GetCaptcha(c *gin.Context) {
	// height 高度 png 像素高度
	// width  宽度 png 像素高度
	// length 验证码默认位数
	// maxSkew 单个数字的最大绝对倾斜因子
	// dotCount 背景圆圈的数量
	driver := base64Captcha.NewDriverDigit(40, 120, 4, 0.2, 20)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, content, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"content":   content,
	})
}
func CheckCapcha(c *gin.Context) {
	err := c.ShouldBindJSON(&User)
	if err != nil {
		fmt.Println("绑定失败", err)
	} else {
		res := store.Verify(User.Id, User.Cd, true)
		c.JSON(200, gin.H{
			"data": res,
		})
	}
}

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := c.ShouldBindJSON(&User)
		if err != nil {
			fmt.Println("绑定失败", err)
		} else {
			// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
			// 这里假设Token放在Header的Authorization中，并使用Bearer开头
			// 这里的具体实现方式要依据你的实际业务情况决定
			if User.Token == "" {
				c.JSON(http.StatusOK, gin.H{
					"code": 2003,
					"msg":  "Token为空",
				})
				c.Abort()
				return
			}
			// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
			mc, err := ParseToken(User.Token)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 2005,
					"msg":  "无效的Token",
				})
				c.Abort()
				return
			}
			// 将当前请求的username信息保存到请求的上下文c上
			c.Set("username", mc.Username)
			c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		}
	}
}
func homeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库链接失败：", err)
		return
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.Static("/pic", "pic")
	r.Static("/file", "file")
	r.Static("/report", "report")
	//配置session
	r.POST("/login", Login)
	r.POST("/enroll", Enroll)
	r.POST("/verify", CheckCapcha)
	r.GET("/get", GetCaptcha)
	r.POST("/password", Password)
	r.POST("/information", Information)
	r.POST("/inquiry", Inquiry)
	r.POST("/inquiry0", Inquiry0)
	r.POST("/inquiry1", Inquiry1)
	r.POST("/inquiry2", Inquiry2)
	r.POST("/inquiry3", Inquiry3)
	r.POST("/inquiry4", Inquiry4)
	r.POST("/inquiry5", Inquiry5)
	r.POST("/inquirys", Inquirys)
	r.POST("/thesis", thesis)
	r.POST("/review", review)
	r.POST("/thesis0", thesis0)
	r.POST("/thesis1", thesis1)
	r.POST("/thesis2", thesis2)
	r.POST("/thesis3", thesis3)
	r.POST("/th1", th1)
	r.POST("/th2", th2)
	r.POST("/th3", th3)
	r.POST("/tj", tj)
	r.POST("/cj", cj)
	r.POST("/jj", jj)
	r.POST("/sp", sp)
	r.POST("/gx", gx)
	r.POST("/upload", Upload)
	r.POST("/editor", editor)
	r.POST("/uploadUser", UploadUser)
	r.POST("/uploadeditor", uploadeditor)
	r.POST("/uploadreview", uploadreview)
	r.POST("/Remove", Remove)
	r.POST("/home", JWTAuthMiddleware(), homeHandler)
	r.Run(":8081")
}
