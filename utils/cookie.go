package utils

import (
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

var conn *redis.Client

// func init() {
// 	conn = redis.NewClient(&redis.Options{
// 		Addr:     "redis:6379",
// 		Password: "",
// 		DB:       0,
// 	})
// }

func NewSession(c echo.Context, cookieKey, redisValue string) {
	// b := make([]byte, 64)
	// if _, err := io.ReadFull(rand.Reader, b); err != nil {
	// 	panic("ランダムな文字作成時にエラーが発生しました。")
	// }
	// newRedisKey := base64.URLEncoding.EncodeToString(b)

	// if err := conn.Set(c.Request().Context(), newRedisKey, redisValue, 0).Err(); err != nil {
	// 	panic("Session登録時にエラーが発生：" + err.Error())
	// }
	cookie := new(http.Cookie)
	cookie.Name = cookieKey
	cookie.Value = redisValue
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	log.Print(c.Response().Header())
}

func GetSession(c echo.Context, cookieKey string) (string, error) {
	cookie, err := c.Cookie(cookieKey)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil

	// redisValue, err := conn.Get(c.Request().Context(), cookie.Value).Result()
	// switch {
	// case err == redis.Nil:
	// 	fmt.Println("SessionKeyが登録されていません。")
	// 	return "not registered", err
	// case err != nil:
	// 	fmt.Println("Session取得時にエラー発生：" + err.Error())
	// 	return "", err
	// }
	// return redisValue, nil
}

func DeleteSession(c echo.Context, cookieKey string) {
	redis, _ := c.Cookie(cookieKey)
	conn.Del(c.Request().Context(), redis.Name)
	cookie := &http.Cookie{
		Name:     "",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	c.SetCookie(cookie)
}
