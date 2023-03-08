package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	body := `{
	"community_id":1,
	"title": "teeeeest",
	"content": "its a test"
}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, url,bytes.NewReader([]byte(body)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// 判断响应的内容是不是按预期返回需要登录的错误
	//assert.Contains(t,  w.Body.String(), "需要登录")

	//将响应的内容反序列化到responseData, 然后判断与预期是否一致
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil{
		t.Fatalf("json.Unmarshal w.Body failed, err: %v\n", err)
	}
	assert.Equal(t, res.Code,CodeNeedLogin)

}