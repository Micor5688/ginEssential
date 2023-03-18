package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数（名称’手机号‘密码）
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		//数据验证：1.验证手机号是否存在
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		//数据验证:2.验证密码
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账户密码不能少于6位"})
			return
		}
		//数据验证 3.如果名称没有传，给一个随机的10位数字
		if len(name) == 0 {
			name = RandString(10)
		}
		log.Println(name, telephone, password)

		//验证手机号是否存在

		//创建用户

		//返回结果

		ctx.JSON(200, gin.H{
			"message": "创建用户成功",
		})
	})
	panic(r.Run())

}

// 随机字符串
func RandString(n int) string {
	var leters = []byte("asdfghkjsdfuhsdfhiusdfhsidufhsn")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for l := range result {
		result[l] = leters[rand.Intn(len(leters))]
	}
	return string(result)
}
