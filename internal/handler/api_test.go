package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
	mock_service "todo_sql_database/internal/service/mocks"
	"todo_sql_database/model"
)

func TestHandler_createTask(t *testing.T) {
	type mockBehavior func(todo *mock_service.MockTodoTask, a *mock_service.MockAuthorization,
		task model.Task, token string)

	testTable := []struct {
		name                 string
		inputTask            model.Task
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
		headerName           string
		headerValue          string
	}{
		{
			name: "OK",
			inputTask: model.Task{
				Name:        "go",
				Description: "study harder",
				Deadline:    "2023-06-22",
			},
			mockBehavior: func(todo *mock_service.MockTodoTask, a *mock_service.MockAuthorization,
				task model.Task, token string) {
				a.EXPECT().ParseToken(token).Return(1, nil).AnyTimes()
				todo.EXPECT().CreateTask(1, &task).Return(1, nil).AnyTimes()
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
			headerName:           "Authorization",
			headerValue:          "Alif token",
		},
		{
			name: "Empty fields",
			inputTask: model.Task{
				Name:        "",
				Description: "",
				Deadline:    "",
			},
			mockBehavior: func(todo *mock_service.MockTodoTask, a *mock_service.MockAuthorization,
				task model.Task, token string) {
				a.EXPECT().ParseToken(token).Return(1, nil).AnyTimes()
				todo.EXPECT().CreateTask(1, &task).Return(1, nil).AnyTimes()
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid task format provided"}`,
			headerName:           "Authorization",
			headerValue:          "Alif token",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			todo := mock_service.NewMockTodoTask(c)
			auth := mock_service.NewMockAuthorization(c)

			split := strings.Split(testCase.headerValue, " ")
			testCase.mockBehavior(todo, auth, testCase.inputTask, split[1])

			handler := NewHandler(auth, todo)

			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/task", handler.tokenAuthMiddleware, handler.createTask)

			marshal, err := json.Marshal(testCase.inputTask)
			if err != nil {
				log.Fatalf("error while marshaling. error is %v", err.Error())
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/task", bytes.NewBuffer(marshal))
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
