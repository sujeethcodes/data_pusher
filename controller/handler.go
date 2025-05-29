package controller

import (
	"data-pusher/constant"
	"data-pusher/entity"
	"data-pusher/repository"
	"data-pusher/usecase"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type DataHandlerController struct {
	Mysql *repository.MysqlCon
}

func (h *DataHandlerController) HandleData(c echo.Context) error {
	secret := c.Request().Header.Get("CL-X-TOKEN")
	if secret == "" {
		return c.JSON(400, entity.Response{
			Status:  constant.UNAUTHORIZED,
			Message: constant.UNAUTHORIZED_MESSAGE,
		})
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(400, entity.Response{
			Status:  constant.INTERNAL_ERROR,
			Message: constant.BODY_FAILED,
		})
	}

	dataHandlerUsecase := usecase.DataUsecase{
		Mysql: h.Mysql,
	}
	err = dataHandlerUsecase.ProcessData(secret, body)
	if err != nil {
		if err.Error() == constant.UNAUTHORIZED_MESSAGE {
			return c.JSON(http.StatusUnauthorized, entity.Response{
				Status:  constant.UNAUTHORIZED,
				Message: constant.UNAUTHORIZED_MESSAGE,
				Error:   err.Error(),
			})
		}
		if err.Error() == constant.INVAILD_JSON {
			return c.JSON(400, entity.Response{
				Status:  constant.BAD_REQUEST,
				Message: constant.INVAILD_JSON,
				Error:   err.Error(),
			})
		}
		return c.JSON(500, entity.Response{
			Status:  constant.INTERNAL_ERROR,
			Message: constant.INTERNAL_ERROR_MESSAGE,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		Status:  constant.SUCCESS,
		Message: constant.DATA_FORWARD_SUCCESS,
	})
}
