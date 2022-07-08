package database

import (
	"context"
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
)

type database struct {
	client *Client
}

func NewStorage(client *Client) *database {
	return &database{
		client: client,
	}
}

func (db database) AddOTPClient(ctx context.Context, logger *logrus.Logger, extID string, phoneNumber string) error {
	qb := sq.
		Update("userservice.client_o2o_otp").
		Set("externalid", extID).
		Where(sq.Eq{"mobile_phone": phoneNumber})

	query, args, err := qb.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		logger.Error(err)
		return err
	}

	if _, err = db.client.ExecContext(ctx, query, args...); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (db *database) GetExternalID(ctx context.Context, logger *logrus.Logger, otpKeys string) (string, error) {

	query, args, err := sq.Insert("userservice.otp").
		SetMap(map[string]interface{}{
			"otp_keys": otpKeys,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.Error(err)
		return "", err
	}

	if _, err = db.client.ExecContext(ctx, query, args...); err != nil {
		logger.Error(err)
		return "", err
	}

	qb := sq.
		Select(
			"externalid",
		).
		From("userservice.otp").
		Where(sq.Eq{"otp_keys": otpKeys})

	query, args, err = qb.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return "", err
	}
	var extID string
	if err = db.client.QueryRowxContext(ctx, query, args...).Scan(&extID); err != nil {
		return "", err
	}
	return extID, nil
}

func (db *database) FindIDByPhone(ctx context.Context, logger *logrus.Logger, phone string) (string, error) {

	qb := sq.Select(
		"externalid",
	).From("userservice.client_o2o_otp").
		Where(sq.Eq{"mobile_phone": phone})

	query, args, err := qb.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {

		return "", err
	}

	var extID string
	if err = db.client.QueryRowxContext(ctx, query, args...).Scan(&extID); err != nil {
		err = errors.New("incorrect phone")
		return "", err
	}

	return extID, nil
}

func (db *database) UpdateOTP(ctx context.Context, logger *logrus.Logger, otpKey, extID string) error {

	qb := sq.Select(
		"otp_keys",
	).From("userservice.otp").
		Where(sq.Eq{"externalID": extID})

	query, args, err := qb.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		logger.Error(err)
		return err
	}

	var currentOTPkeys string

	if err = db.client.QueryRowxContext(ctx, query, args...).Scan(&currentOTPkeys); err != nil {
		return err
	}

	newOTPKeys := strings.Replace(currentOTPkeys, otpKey, "", -1)
	fmt.Println(currentOTPkeys, newOTPKeys)
	if newOTPKeys == currentOTPkeys {
		err = errors.New("incorrect code")
		fmt.Println(err)
		return err
	}
	fmt.Println("you reach the middle")

	qb1 := sq.
		Update("userservice.otp").
		Set("otp_keys", newOTPKeys).
		Where(sq.Eq{"externalID": extID})

	query, args, err = qb1.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		logger.Error(err)
		return err
	}

	if _, err = db.client.ExecContext(ctx, query, args...); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
