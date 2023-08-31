package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validate "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type LoginForm struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required,min=3,max=10"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type SignUpForm struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for filed, err := range fileds {
		rsp[filed[strings.Index(filed, ".")+1:]] = err
	}
	return rsp
}

func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎，来实现自定义
	if v, ok := binding.Validator.Engine().(*validate.Validate); ok {

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			test := field.Tag.Get("json")
			fmt.Println(test)
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "_" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()
		//第一个是备用， 后面是应该支持的环境
		uni := ut.New(zhT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)

		if !ok {
			return fmt.Errorf(locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
	}
	return
}

var trans ut.Translator

// 中间件
func MYLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Next()
		end := time.Since(t)
		fmt.Printf("耗时:%s\n", end)
	}
}

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}

	router := gin.New()
	router.Use(MYLogger(), gin.Recovery())
	router.POST("/login", func(context *gin.Context) {
		var loginForm LoginForm
		if err := context.ShouldBind(&loginForm); err != nil {
			errs, ok := err.(validate.ValidationErrors)
			if !ok {
				context.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}

			fmt.Print(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"message": removeTopStruct(errs.Translate(trans)),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	})

	router.POST("/signup", func(context *gin.Context) {
		var signForm SignUpForm

		if error := context.ShouldBind(&signForm); error != nil {
			fmt.Print(error.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"message": error.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	router.Run(":8083")
}
