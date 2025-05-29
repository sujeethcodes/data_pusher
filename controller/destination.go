package controller

import (
	"data-pusher/constant"
	"data-pusher/entity"
	"data-pusher/repository"
	"data-pusher/usecase"
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
)

type DestinationController struct {
	Mysql *repository.MysqlCon
}

func (d *DestinationController) CreateDestination(c echo.Context) error {
	req := entity.Destination{}

	// Bind request body (account_id, url)
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
			Error:   err.Error(),
		})
	}

	req.Status = constant.ACTIVE_STATUS

	// Set method from the HTTP request
	req.Method = c.Request().Method
	if req.Method != constant.POST {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
		})
	}

	// Extract headers from HTTP request
	headers := make(map[string]string)
	for key, values := range c.Request().Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	headersJSON, err := json.Marshal(headers)
	if err != nil {
		return err
	}

	req.Headers = string(headersJSON)
	fmt.Println("headersJSON----", req.Headers)

	destinationUsecase := usecase.DestinationUsecase{
		Mysql: d.Mysql,
	}
	err = destinationUsecase.CreateDestination(req)
	if err != nil {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.DESTINATION_FETCH_FAILED,
			Error:   err.Error(),
		})
	}
	return c.JSON(201, entity.Response{
		Status:  constant.SUCCESS,
		Message: constant.SUCCESS_MESSAGE,
	})
}

func (d *DestinationController) GetDestinationDetails(c echo.Context) error {
	fmt.Println("enter get accoun details function")
	account_id := c.QueryParam("account_id")
	if account_id == "" {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
		})
	}

	destinationUsecase := usecase.DestinationUsecase{
		Mysql: d.Mysql,
	}

	detils, err := destinationUsecase.GetDestinationDetails(account_id)
	if err != nil {
		return c.JSON(400, entity.Response{
			Status:  400,
			Message: "Account fetch failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(200, entity.Response{
		Status:  200,
		Message: constant.SUCCESS_MESSAGE,
		Data:    detils,
	})

}
