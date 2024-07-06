package test

import (
	"bytes"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_Registered(t *testing.T) {

	r := SetupGinRouter()

	// 创建请求
	req := httptest.NewRequest(http.MethodPost, "/registered", bytes.NewBufferString("userAddress=0x123&contractAddress=0xabc"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 执行请求并获取响应
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 验证响应状态码
	//assert.Equal(t, http.StatusOK, w.Code)

	t.Logf("Response body: %s", w.Body.String())
}

func TestUserHandler_Participation(t *testing.T) {
	r := SetupGinRouter()

	// 创建请求
	req := httptest.NewRequest(http.MethodPost, "/participation", bytes.NewBufferString("userAddress=0x123&amount=12&contractAddress=0xabc"))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 执行请求并获取响应
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 验证响应状态码
	//assert.Equal(t, http.StatusOK, w.Code)

	t.Logf("Response body: %s", w.Body.String())
}
