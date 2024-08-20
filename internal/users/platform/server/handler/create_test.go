package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rubenbupe/go-auth-server/kit/command/commandmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create_ServiceError(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("create.UserCommand"),
	).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/users", CreateHandler(commandBus))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		createUserReq := createRequest{
			Name: "Demo User",
		}

		b, err := json.Marshal(createUserReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createUserReq := createRequest{
			ID:   "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Name: "Demo User",
		}

		b, err := json.Marshal(createUserReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
