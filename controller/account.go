package controller

import (
	"data-pusher/constant"
	"data-pusher/entity"
	"data-pusher/repository"
	"data-pusher/usecase"
	"data-pusher/utils"
	"fmt"

	"github.com/labstack/echo"
)

type AccountController struct {
	Mysql *repository.MysqlCon
}

func (a *AccountController) CreateAccount(c echo.Context) error {
	req := entity.Accounts{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
			Error:   err.Error(),
		})
	}

	req.Token, _ = utils.GenerateSecretToken()
	req.AccountID = utils.GenerateAccountID()
	req.Status = constant.ACTIVE_STATUS

	accountUsecase := usecase.AccountUsecase{
		Mysql: a.Mysql,
	}

	email, err := accountUsecase.IsEmailExists(req.Email)
	if err != nil {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.ACCOUNT_FETCH_FAILED,
			Error:   err.Error(),
		})
	}
	if email {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.EMAIL_EXIST,
		})
	}

	if err := accountUsecase.CreateAccount(req); err != nil {
		return c.JSON(400, entity.Response{
			Status:  400,
			Message: "User fetch failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(200, entity.Response{
		Status:  200,
		Message: "User created successfully",
	})
}

func (a *AccountController) UpdateAccount(c echo.Context) error {
	req := entity.Accounts{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
			Error:   err.Error(),
		})
	}

	if req.Status != "" {
		if req.Status != constant.ACTIVE_STATUS && req.Status != constant.IN_ACTIVE_STATUS {
			return c.JSON(400, entity.Response{
				Status:  constant.BAD_REQUEST,
				Message: constant.INVAILD_STATUS_VALUE,
			})
		}
	}

	accountUsecase := usecase.AccountUsecase{
		Mysql: a.Mysql,
	}

	_, err := accountUsecase.IsEmailExists(req.Email)
	if err != nil {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.ACCOUNT_FETCH_FAILED,
			Error:   err.Error(),
		})
	}

	if err := accountUsecase.UpdateAccount(req); err != nil {
		return c.JSON(400, entity.Response{
			Status:  400,
			Message: "User fetch failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(200, entity.Response{
		Status:  200,
		Message: "update successfully",
	})
}

func (a *AccountController) DeleteAccount(c echo.Context) error {
	req := entity.DeleteReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
			Error:   err.Error(),
		})
	}
	req.Status = constant.IN_ACTIVE_STATUS
	accountUsecase := usecase.AccountUsecase{
		Mysql: a.Mysql,
	}
	if err := accountUsecase.DeleteAccount(req); err != nil {
		return c.JSON(400, entity.Response{
			Status:  400,
			Message: "account delete failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(200, entity.Response{
		Status:  200,
		Message: "delete successfully",
	})

}

func (a *AccountController) GetAccountDetails(c echo.Context) error {
	fmt.Println("enter get accoun details function")
	AccountID := c.QueryParam("account_id")
	if AccountID == "" {
		return c.JSON(400, entity.Response{
			Status:  constant.BAD_REQUEST,
			Message: constant.BAD_REQUEST_MESSAGE,
		})
	}

	accountUsecase := usecase.AccountUsecase{
		Mysql: a.Mysql,
	}

	detils, err := accountUsecase.GetAccountDetails(AccountID)
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
