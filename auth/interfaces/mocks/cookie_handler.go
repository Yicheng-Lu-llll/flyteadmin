// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"

	service "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service"
)

// CookieHandler is an autogenerated mock type for the CookieHandler type
type CookieHandler struct {
	mock.Mock
}

// DeleteCookies provides a mock function with given fields: ctx, writer
func (_m *CookieHandler) DeleteCookies(ctx context.Context, writer http.ResponseWriter) {
	_m.Called(ctx, writer)
}

type CookieHandler_RetrieveAuthCodeRequest struct {
	*mock.Call
}

func (_m CookieHandler_RetrieveAuthCodeRequest) Return(authRequestURL string, err error) *CookieHandler_RetrieveAuthCodeRequest {
	return &CookieHandler_RetrieveAuthCodeRequest{Call: _m.Call.Return(authRequestURL, err)}
}

func (_m *CookieHandler) OnRetrieveAuthCodeRequest(ctx context.Context, request *http.Request) *CookieHandler_RetrieveAuthCodeRequest {
	c_call := _m.On("RetrieveAuthCodeRequest", ctx, request)
	return &CookieHandler_RetrieveAuthCodeRequest{Call: c_call}
}

func (_m *CookieHandler) OnRetrieveAuthCodeRequestMatch(matchers ...interface{}) *CookieHandler_RetrieveAuthCodeRequest {
	c_call := _m.On("RetrieveAuthCodeRequest", matchers...)
	return &CookieHandler_RetrieveAuthCodeRequest{Call: c_call}
}

// RetrieveAuthCodeRequest provides a mock function with given fields: ctx, request
func (_m *CookieHandler) RetrieveAuthCodeRequest(ctx context.Context, request *http.Request) (string, error) {
	ret := _m.Called(ctx, request)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) string); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type CookieHandler_RetrieveTokenValues struct {
	*mock.Call
}

func (_m CookieHandler_RetrieveTokenValues) Return(idToken string, accessToken string, refreshToken string, err error) *CookieHandler_RetrieveTokenValues {
	return &CookieHandler_RetrieveTokenValues{Call: _m.Call.Return(idToken, accessToken, refreshToken, err)}
}

func (_m *CookieHandler) OnRetrieveTokenValues(ctx context.Context, request *http.Request) *CookieHandler_RetrieveTokenValues {
	c_call := _m.On("RetrieveTokenValues", ctx, request)
	return &CookieHandler_RetrieveTokenValues{Call: c_call}
}

func (_m *CookieHandler) OnRetrieveTokenValuesMatch(matchers ...interface{}) *CookieHandler_RetrieveTokenValues {
	c_call := _m.On("RetrieveTokenValues", matchers...)
	return &CookieHandler_RetrieveTokenValues{Call: c_call}
}

// RetrieveTokenValues provides a mock function with given fields: ctx, request
func (_m *CookieHandler) RetrieveTokenValues(ctx context.Context, request *http.Request) (string, string, string, error) {
	ret := _m.Called(ctx, request)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) string); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) string); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(context.Context, *http.Request) string); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Get(2).(string)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, *http.Request) error); ok {
		r3 = rf(ctx, request)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

type CookieHandler_RetrieveUserInfo struct {
	*mock.Call
}

func (_m CookieHandler_RetrieveUserInfo) Return(_a0 *service.UserInfoResponse, _a1 error) *CookieHandler_RetrieveUserInfo {
	return &CookieHandler_RetrieveUserInfo{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *CookieHandler) OnRetrieveUserInfo(ctx context.Context, request *http.Request) *CookieHandler_RetrieveUserInfo {
	c_call := _m.On("RetrieveUserInfo", ctx, request)
	return &CookieHandler_RetrieveUserInfo{Call: c_call}
}

func (_m *CookieHandler) OnRetrieveUserInfoMatch(matchers ...interface{}) *CookieHandler_RetrieveUserInfo {
	c_call := _m.On("RetrieveUserInfo", matchers...)
	return &CookieHandler_RetrieveUserInfo{Call: c_call}
}

// RetrieveUserInfo provides a mock function with given fields: ctx, request
func (_m *CookieHandler) RetrieveUserInfo(ctx context.Context, request *http.Request) (*service.UserInfoResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *service.UserInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *service.UserInfoResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.UserInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type CookieHandler_SetAuthCodeCookie struct {
	*mock.Call
}

func (_m CookieHandler_SetAuthCodeCookie) Return(_a0 error) *CookieHandler_SetAuthCodeCookie {
	return &CookieHandler_SetAuthCodeCookie{Call: _m.Call.Return(_a0)}
}

func (_m *CookieHandler) OnSetAuthCodeCookie(ctx context.Context, writer http.ResponseWriter, authRequestURL string) *CookieHandler_SetAuthCodeCookie {
	c_call := _m.On("SetAuthCodeCookie", ctx, writer, authRequestURL)
	return &CookieHandler_SetAuthCodeCookie{Call: c_call}
}

func (_m *CookieHandler) OnSetAuthCodeCookieMatch(matchers ...interface{}) *CookieHandler_SetAuthCodeCookie {
	c_call := _m.On("SetAuthCodeCookie", matchers...)
	return &CookieHandler_SetAuthCodeCookie{Call: c_call}
}

// SetAuthCodeCookie provides a mock function with given fields: ctx, writer, authRequestURL
func (_m *CookieHandler) SetAuthCodeCookie(ctx context.Context, writer http.ResponseWriter, authRequestURL string) error {
	ret := _m.Called(ctx, writer, authRequestURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, string) error); ok {
		r0 = rf(ctx, writer, authRequestURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type CookieHandler_SetTokenCookies struct {
	*mock.Call
}

func (_m CookieHandler_SetTokenCookies) Return(_a0 error) *CookieHandler_SetTokenCookies {
	return &CookieHandler_SetTokenCookies{Call: _m.Call.Return(_a0)}
}

func (_m *CookieHandler) OnSetTokenCookies(ctx context.Context, writer http.ResponseWriter, token *oauth2.Token) *CookieHandler_SetTokenCookies {
	c_call := _m.On("SetTokenCookies", ctx, writer, token)
	return &CookieHandler_SetTokenCookies{Call: c_call}
}

func (_m *CookieHandler) OnSetTokenCookiesMatch(matchers ...interface{}) *CookieHandler_SetTokenCookies {
	c_call := _m.On("SetTokenCookies", matchers...)
	return &CookieHandler_SetTokenCookies{Call: c_call}
}

// SetTokenCookies provides a mock function with given fields: ctx, writer, token
func (_m *CookieHandler) SetTokenCookies(ctx context.Context, writer http.ResponseWriter, token *oauth2.Token) error {
	ret := _m.Called(ctx, writer, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, *oauth2.Token) error); ok {
		r0 = rf(ctx, writer, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type CookieHandler_SetUserInfoCookie struct {
	*mock.Call
}

func (_m CookieHandler_SetUserInfoCookie) Return(_a0 error) *CookieHandler_SetUserInfoCookie {
	return &CookieHandler_SetUserInfoCookie{Call: _m.Call.Return(_a0)}
}

func (_m *CookieHandler) OnSetUserInfoCookie(ctx context.Context, writer http.ResponseWriter, userInfo *service.UserInfoResponse) *CookieHandler_SetUserInfoCookie {
	c_call := _m.On("SetUserInfoCookie", ctx, writer, userInfo)
	return &CookieHandler_SetUserInfoCookie{Call: c_call}
}

func (_m *CookieHandler) OnSetUserInfoCookieMatch(matchers ...interface{}) *CookieHandler_SetUserInfoCookie {
	c_call := _m.On("SetUserInfoCookie", matchers...)
	return &CookieHandler_SetUserInfoCookie{Call: c_call}
}

// SetUserInfoCookie provides a mock function with given fields: ctx, writer, userInfo
func (_m *CookieHandler) SetUserInfoCookie(ctx context.Context, writer http.ResponseWriter, userInfo *service.UserInfoResponse) error {
	ret := _m.Called(ctx, writer, userInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, *service.UserInfoResponse) error); ok {
		r0 = rf(ctx, writer, userInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
