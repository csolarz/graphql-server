package controller

import (
	"net/http"
	"strconv"

	"github.com/csolarz/graphql-server/entities"
	"github.com/csolarz/graphql-server/usecase/api"
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	service api.Usecase
}

func NewApiController(service api.Usecase) *ApiController {
	return &ApiController{service: service}
}

func (lc ApiController) NewLoan(c *gin.Context) {
	var request entities.LoanRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload")
		return
	}

	loan, err := lc.service.NewLoan(c.Request.Context(), request)

	if err != nil {

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, loan)
}

func (lc ApiController) GetLoan(c *gin.Context) {
	id := c.Param("id")
	loanID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid loan ID")
		return
	}

	loan, err := lc.service.GetLoan(c.Request.Context(), loanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, loan)
}

func (lc ApiController) NewUser(c *gin.Context) {
	var request entities.UserRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := lc.service.NewUser(c.Request.Context(), request)

	if err != nil {

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}
