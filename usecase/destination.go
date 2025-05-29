package usecase

import (
	"data-pusher/constant"
	"data-pusher/entity"
	"data-pusher/repository"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type DestinationUsecase struct {
	Mysql *repository.MysqlCon
}

func (d *DestinationUsecase) CreateDestination(req entity.Destination) error {
	fmt.Println("req check create account ", req)
	if d.Mysql == nil {
		zap.L().Info("Database connection failed")
		return errors.New("database connection not initialized")
	}

	err := d.Mysql.Connection.Table(constant.DESTINATION_TABLE_NAME).Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *DestinationUsecase) GetDestinationDetails(account_id string) ([]entity.Destination, error) {
	fmt.Println("req check account delete ", account_id)
	if d.Mysql == nil {
		zap.L().Info("Database connection failed")
		return []entity.Destination{}, errors.New("database connection not initialized")
	}
	var details = []entity.Destination{}

	err := d.Mysql.Connection.
		Table(constant.DESTINATION_TABLE_NAME).
		Where("account_id = ? AND status = in active ", account_id).
		Scan(&details).Error

	if err != nil {
		return details, err
	}
	zap.L().Info("details----", zap.Any("details", details))
	return details, nil
}
