package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/csolarz/graphql-server/entities"
	mock_api "github.com/csolarz/graphql-server/usecase/api/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(t *testing.T) (*gin.Engine, *mock_api.Usecase) {
	gin.SetMode(gin.TestMode)

	mockSvc := mock_api.NewUsecase(t)
	apiCtrl := NewApiController(mockSvc)

	r := gin.Default()
	r.POST("/loan", apiCtrl.NewLoan)
	r.GET("/loan/:id", apiCtrl.GetLoan)
	r.POST("/user", apiCtrl.NewUser)

	return r, mockSvc
}

func TestNewLoan_InvalidRequest(t *testing.T) {
	router, _ := setupRouter(t)

	req := httptest.NewRequest("POST", "/loan", bytes.NewBuffer([]byte("{invalid")))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNewLoan_ServiceError(t *testing.T) {
	router, mockSvc := setupRouter(t)

	body := entities.LoanRequest{UserID: 10, Amount: 50000}
	jsonBody, _ := json.Marshal(body)

	mockSvc.On("NewLoan", mock.Anything, body).
		Return((*entities.Loan)(nil), errors.New("service error"))

	req := httptest.NewRequest("POST", "/loan", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestNewLoan_OK(t *testing.T) {
	router, mockSvc := setupRouter(t)

	body := entities.LoanRequest{UserID: 10, Amount: 50000}
	jsonBody, _ := json.Marshal(body)

	expectedLoan := &entities.Loan{ID: 1, UserID: 10, Amount: 50000}

	mockSvc.On("NewLoan", mock.Anything, body).
		Return(expectedLoan, nil)

	req := httptest.NewRequest("POST", "/loan", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var loanResp entities.Loan
	err := json.Unmarshal(w.Body.Bytes(), &loanResp)
	assert.NoError(t, err)

	assert.Equal(t, expectedLoan.ID, loanResp.ID)
}

func TestGetLoan_InvalidID(t *testing.T) {
	router, _ := setupRouter(t)

	req := httptest.NewRequest("GET", "/loan/abc", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetLoan_ServiceError(t *testing.T) {
	router, mockSvc := setupRouter(t)

	mockSvc.On("GetLoan", mock.Anything, int64(10)).
		Return((*entities.Loan)(nil), errors.New("service error"))

	req := httptest.NewRequest("GET", "/loan/10", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetLoan_OK(t *testing.T) {
	router, mockSvc := setupRouter(t)

	expectedLoan := &entities.Loan{ID: 10, UserID: 1, Amount: 20000}

	mockSvc.On("GetLoan", mock.Anything, int64(10)).
		Return(expectedLoan, nil)

	req := httptest.NewRequest("GET", "/loan/10", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var loanResp entities.Loan
	err := json.Unmarshal(w.Body.Bytes(), &loanResp)
	assert.NoError(t, err)

	assert.Equal(t, expectedLoan.ID, loanResp.ID)
}

func TestNewUser_InvalidRequest(t *testing.T) {
	router, _ := setupRouter(t)

	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer([]byte("{invalid")))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNewUser_ServiceError(t *testing.T) {
	router, mockSvc := setupRouter(t)

	body := entities.UserRequest{Name: "carlos"}
	jsonBody, _ := json.Marshal(body)

	mockSvc.On("NewUser", mock.Anything, body).
		Return((*entities.User)(nil), errors.New("service error"))

	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestNewUser_OK(t *testing.T) {
	router, mockSvc := setupRouter(t)

	body := entities.UserRequest{Name: "carlos"}
	jsonBody, _ := json.Marshal(body)

	expectedUser := &entities.User{ID: 1, Name: "carlos"}

	mockSvc.On("NewUser", mock.Anything, body).
		Return(expectedUser, nil)

	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var userResp entities.User
	err := json.Unmarshal(w.Body.Bytes(), &userResp)
	assert.NoError(t, err)

	assert.Equal(t, expectedUser.ID, userResp.ID)
}
