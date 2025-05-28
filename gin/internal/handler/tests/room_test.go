package tests

import (
	"bytes"
	"database/sql"
	"gin/internal/handler"
	"gin/internal/models"
	"gin/internal/services"
	mock_services "gin/internal/services/mocks"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRoom_create(t *testing.T) {
	type mockBehavior func(s *mock_services.MockRoom, room models.RoomCreate)

	testTable := []struct {
		name                string
		inputBody           string
		inputRoom           models.RoomCreate
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"number": "101", "type": "Single", "description": "OkTest"}`,
			inputRoom: models.RoomCreate{
				Number:      "101",
				Type:        "Single",
				Description: "OkTest",
			},
			mockBehavior: func(s *mock_services.MockRoom, room models.RoomCreate) {
				s.EXPECT().CreateRoom(room).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			roomService := mock_services.NewMockRoom(c)
			testCase.mockBehavior(roomService, testCase.inputRoom)

			services := &services.Service{Room: roomService}
			handler := handler.NewHandler(services)

			//Test Server
			r := gin.New()
			r.POST("/room/", handler.CreateRoom)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/room/", bytes.NewBufferString(testCase.inputBody))

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestRoom_get(t *testing.T) {
	type mockBehavior func(s *mock_services.MockRoom, id int)

	testTable := []struct {
		name                 string
		roomID               string
		inputID              int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "OK",
			roomID:  "1",
			inputID: 1,
			mockBehavior: func(s *mock_services.MockRoom, id int) {
				s.EXPECT().GetRoomById(id).Return(&models.Room{
					Id:          1,
					Number:      "101",
					Type:        "Single",
					Description: "OkTest",
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"number":"101","type":"Single","description":"OkTest"}`,
		},
		{
			name:    "NotFound",
			roomID:  "2",
			inputID: 2,
			mockBehavior: func(s *mock_services.MockRoom, id int) {
				s.EXPECT().GetRoomById(id).Return(nil, sql.ErrNoRows)
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"room not found"}`,
		},
		{
			name:    "InvalidID",
			roomID:  "abc",
			inputID: 0,
			mockBehavior: func(s *mock_services.MockRoom, id int) {
				// No call expected
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"id param is not correct"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			roomService := mock_services.NewMockRoom(c)
			testCase.mockBehavior(roomService, testCase.inputID)

			services := &services.Service{Room: roomService}
			handler := handler.NewHandler(services)

			r := gin.New()
			r.GET("/room/:id", handler.GetRoomById)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/room/"+testCase.roomID, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
