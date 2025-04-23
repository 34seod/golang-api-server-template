package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"golang-api-server-template/internal/dto"
	"golang-api-server-template/test/factories"
	. "golang-api-server-template/test/testutils"

	. "github.com/go-playground/assert/v2"
)

// GET /v1/users success user가 있을 때
func Test_Success_GET_USERS(t *testing.T) {
	// ready
	ResetTestData()
	users := factories.UsersFactory(3)

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("GET", "/v1/users", nil)
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data":  users,
		"error": nil,
	})
	Equal(t, http.StatusOK, w.Code)
	Equal(t, expect, w.Body.String())
}

// GET /v1/users success user가 없을 때
func Test_Success_GET_USERS_WithoutUsers(t *testing.T) {
	// ready
	ResetTestData()

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("GET", "/v1/users", nil)
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data":  []any{},
		"error": nil,
	})
	Equal(t, http.StatusOK, w.Code)
	Equal(t, expect, w.Body.String())
}

// GET /v1/users/:id success
func Test_Success_GET_USERS_ID(t *testing.T) {
	// ready
	ResetTestData()
	user := factories.UserFactory()

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/v1/users/%d", user.ID), nil)
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data":  user,
		"error": nil,
	})
	Equal(t, http.StatusOK, w.Code)
	Equal(t, expect, w.Body.String())
}

// GET /v1/users/:id fail 없는 id일때
func Test_Fail_GET_USERS_ID_NotExistId(t *testing.T) {
	// ready
	ResetTestData()

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("GET", "/v1/users/99999", nil)
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data":  nil,
		"error": "record not found",
	})
	Equal(t, http.StatusBadRequest, w.Code)
	Equal(t, expect, w.Body.String())
}

// DELETE /v1/users/:id success
func Test_Success_DELETE_USERS(t *testing.T) {
	// ready
	ResetTestData()
	user := factories.UserFactory()

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/v1/users/%d", user.ID), nil)
	r.ServeHTTP(w, req)

	// check
	expect := ""
	Equal(t, http.StatusNoContent, w.Code)
	Equal(t, expect, w.Body.String())
}

// DELETE /v1/users/:id fail
func Test_Fail_DELETE_USERS_NotExistId(t *testing.T) {
	// ready
	ResetTestData()

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("DELETE", "/v1/users/99999", nil)
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data":  nil,
		"error": "record not found",
	})
	Equal(t, http.StatusBadRequest, w.Code)
	Equal(t, expect, w.Body.String())
}

// POST /v1/users success name만 있을떄
func Test_Success_POST_USERS_OnlyName(t *testing.T) {
	// ready
	ResetTestData()
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Name: GenerateRandomString(5),
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("POST", "/v1/users", strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ""
	Equal(t, http.StatusNoContent, w.Code)
	Equal(t, expect, w.Body.String())
}

// POST /v1/users success name이외에도 있을 때
func Test_Success_POST_USERS_All(t *testing.T) {
	// ready
	ResetTestData()
	email := GenerateRandomString(5) + "@example.com"
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Name:  GenerateRandomString(5),
		Tel:   GenerateRandomStringPtr(11),
		Email: &email,
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("POST", "/v1/users", strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ""
	Equal(t, http.StatusNoContent, w.Code)
	Equal(t, expect, w.Body.String())
}

// POST /v1/users fail name없을때
func Test_Fail_POST_USERS_NoName(t *testing.T) {
	// ready
	ResetTestData()
	email := GenerateRandomString(5) + "@example.com"
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Tel:   GenerateRandomStringPtr(11),
		Email: &email,
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("POST", "/v1/users", strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data": nil,
		"error": []map[string]any{
			{
				"field":   "name",
				"message": "This field is required",
			},
		},
	})
	Equal(t, http.StatusBadRequest, w.Code)
	Equal(t, expect, w.Body.String())
}

// POST /v1/users fail tel 자릿수 부족
func Test_Fail_POST_USERS_NotEnoughTelLen(t *testing.T) {
	// ready
	ResetTestData()
	email := GenerateRandomString(5) + "@example.com"
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Name:  GenerateRandomString(5),
		Tel:   GenerateRandomStringPtr(10),
		Email: &email,
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("POST", "/v1/users", strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data": nil,
		"error": []map[string]any{
			{
				"field":   "tel",
				"message": "This field needs 11 digit",
			},
		},
	})
	Equal(t, http.StatusBadRequest, w.Code)
	Equal(t, expect, w.Body.String())
}

// POST /v1/users fail email 메일 형식이 아닐때
func Test_Fail_POST_USERS_NotEmailForm(t *testing.T) {
	// ready
	ResetTestData()
	email := GenerateRandomString(5) + "example.com"
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Name:  GenerateRandomString(5),
		Tel:   GenerateRandomStringPtr(11),
		Email: &email,
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("POST", "/v1/users", strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data": nil,
		"error": []map[string]any{
			{
				"field":   "email",
				"message": "This field is in email format only",
			},
		},
	})
	Equal(t, http.StatusBadRequest, w.Code)
	Equal(t, expect, w.Body.String())
}

// PUT /v1/users/:id success name만 있을떄
func Test_Success_PUT_USERS_OnlyName(t *testing.T) {
	// ready
	ResetTestData()
	user := factories.UserFactory()
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Name: GenerateRandomString(5),
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/users/%d", user.ID), strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ""
	Equal(t, http.StatusNoContent, w.Code)
	Equal(t, expect, w.Body.String())
}

// PUT /v1/users/:id success name이외에도 있을 때
func Test_Success_PUT_USERS_All(t *testing.T) {
	// ready
	ResetTestData()
	user := factories.UserFactory()
	email := GenerateRandomString(5) + "@example.com"
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Name:  GenerateRandomString(5),
		Tel:   GenerateRandomStringPtr(11),
		Email: &email,
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/users/%d", user.ID), strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ""
	Equal(t, http.StatusNoContent, w.Code)
	Equal(t, expect, w.Body.String())
}

// PUT /v1/users/:id fail name없을때
func Test_Fail_PUT_USERS_NoName(t *testing.T) {
	// ready
	ResetTestData()
	user := factories.UserFactory()
	email := GenerateRandomString(5) + "@example.com"
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Tel:   GenerateRandomStringPtr(11),
		Email: &email,
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/users/%d", user.ID), strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data": nil,
		"error": []map[string]any{
			{
				"field":   "name",
				"message": "This field is required",
			},
		},
	})
	Equal(t, http.StatusBadRequest, w.Code)
	Equal(t, expect, w.Body.String())
}

// PUT /v1/users/:id fail tel 자릿수 부족
func Test_Fail_PUT_USERS_NotEnoughTelLen(t *testing.T) {
	// ready
	ResetTestData()
	user := factories.UserFactory()
	email := GenerateRandomString(5) + "@example.com"
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Name:  GenerateRandomString(5),
		Tel:   GenerateRandomStringPtr(10),
		Email: &email,
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/users/%d", user.ID), strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data": nil,
		"error": []map[string]any{
			{
				"field":   "tel",
				"message": "This field needs 11 digit",
			},
		},
	})
	Equal(t, http.StatusBadRequest, w.Code)
	Equal(t, expect, w.Body.String())
}

// PUT /v1/users/:id fail email 메일 형식이 아닐때
func Test_Fail_PUT_USERS_NotEmailForm(t *testing.T) {
	// ready
	ResetTestData()
	user := factories.UserFactory()
	email := GenerateRandomString(5) + "example.com"
	userJson, _ := json.Marshal(dto.UserBodyFromCreateRequest{
		Name:  GenerateRandomString(5),
		Tel:   GenerateRandomStringPtr(11),
		Email: &email,
	})

	// execute
	r, w := ApiSetup()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/users/%d", user.ID), strings.NewReader(string(userJson)))
	req.Header["Content-Type"] = []string{"application/json"}
	r.ServeHTTP(w, req)

	// check
	expect := ToJson(map[string]any{
		"data": nil,
		"error": []map[string]any{
			{
				"field":   "email",
				"message": "This field is in email format only",
			},
		},
	})
	Equal(t, http.StatusBadRequest, w.Code)
	Equal(t, expect, w.Body.String())
}
