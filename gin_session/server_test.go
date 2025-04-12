package gin_session

import (
	"fmt"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//store := memstore.NewStore([]byte("moyn8y9abnd7q4zkq2m73yw8tu9j5ixm"), []byte("o6idlo2cb9f9pb6h46fimllw481ldebi"))
	store, err := redis.NewStore(
		16,                                         // 最大空闲链接数量，过大会浪费，过小将来会触发性能瓶颈
		"tcp",                                      // 指定与Redis服务器通信的网络类型，通常为"tcp"
		"localhost:6379",                           // Redis服务器的地址，格式为"host:port"
		"",                                         // Redis服务器的密码，如果没有密码可以为空字符串
		[]byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"), // authentication key
		[]byte("0Pf2r0wZBpXVXlQNdpwCXN4ncnlnZSc3"), // encryption key
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置Session中间件
	router.Use(sessions.Sessions("mysession", store))

	// 路由示例
	router.GET("/set", func(c *gin.Context) {
		// 设置Session
		session := sessions.Default(c)
		session.Set("key", "value")

		//session.Options(sessions.Options{MaxAge: 10})

		session.Save()

		c.JSON(200, gin.H{"message": "Session set"})
	})

	router.GET("/get", func(c *gin.Context) {
		// 获取Session
		session := sessions.Default(c)
		value := session.Get("key")

		c.JSON(200, gin.H{"key": value})
	})

	router.GET("/del", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Options(sessions.Options{MaxAge: -1})

		session.Save()

		c.JSON(200, gin.H{"res": "ok"})
	})

	// 启动服务
	router.Run(":8080")
}

func TestRun(t *testing.T) {
	main()
	
}
