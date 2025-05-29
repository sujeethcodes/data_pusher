package usecase

import (
	"data-pusher/constant"
	"data-pusher/entity"
	"data-pusher/repository"

	"errors"
	"fmt"

	"go.uber.org/zap"
)

type AccountUsecase struct {
	Mysql *repository.MysqlCon
}

func (u *AccountUsecase) IsEmailExists(email string) (bool, error) {
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return false, errors.New("database connection not initialized")
	}

	var count int64
	err := u.Mysql.Connection.
		Table(constant.ACCOUNT_TABLE_NAME).
		Where("email = ?", email).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (u *AccountUsecase) CreateAccount(req entity.Accounts) error {
	fmt.Println("req check create account ", req)
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return errors.New("database connection not initialized")
	}

	err := u.Mysql.Connection.Table(constant.ACCOUNT_TABLE_NAME).Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *AccountUsecase) UpdateAccount(req entity.Accounts) error {
	fmt.Println("req check account update ", req)
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return errors.New("database connection not initialized")
	}

	if req.AccountID == "" {
		return errors.New(constant.ACCOUNT_ID_REQUIRED)
	}

	updateData := make(map[string]interface{})

	if req.Email != "" {
		updateData["email"] = req.Email
	}
	if req.Name != "" {
		updateData["name"] = req.Name
	}
	if req.Status != "" {
		updateData["status"] = req.Status
	}

	if len(updateData) == 0 {
		return errors.New("no fields to update")
	}

	// Perform the update
	err := u.Mysql.Connection.
		Table(constant.ACCOUNT_TABLE_NAME).
		Where("account_id = ?", req.AccountID).
		Updates(updateData).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *AccountUsecase) DeleteAccount(req entity.DeleteReq) error {
	fmt.Println("req check account delete ", req)
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return errors.New("database connection not initialized")
	}

	// Raw SQL to update status of destinations
	updateDestinationsQuery := `
	UPDATE ` + constant.DESTINATION_TABLE_NAME + ` d
	JOIN ` + constant.ACCOUNT_TABLE_NAME + ` a
	ON d.account_id = a.account_id
	SET d.status = ?
	WHERE a.account_id = ?
	`

	err := u.Mysql.Connection.Exec(updateDestinationsQuery, constant.IN_ACTIVE_STATUS, req.AccountId).Error
	if err != nil {
		return err
	}

	err = u.Mysql.Connection.
		Table(constant.ACCOUNT_TABLE_NAME).
		Where("account_id = ?", req.AccountId).
		Updates(map[string]interface{}{"status": constant.IN_ACTIVE_STATUS}).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *AccountUsecase) GetAccountDetails(account_id string) (entity.Accounts, error) {
	fmt.Println("req check account delete ", account_id)
	if u.Mysql == nil {
		zap.L().Info("Database connection failed")
		return entity.Accounts{}, errors.New("database connection not initialized")
	}
	var details = entity.Accounts{}

	err := u.Mysql.Connection.
		Table(constant.ACCOUNT_TABLE_NAME).
		Where("account_id = ?", account_id).
		Scan(&details).Error

	if err != nil {
		return details, err
	}

	return details, nil
}
