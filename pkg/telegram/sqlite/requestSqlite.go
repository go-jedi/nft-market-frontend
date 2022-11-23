package sqlite

import (
	"database/sql"
	"fmt"
)

func GetUserLang(sqliteDb *sql.DB, teleId int64) (string, error) {
	type BotParams struct {
		Lang string `json:"lang"`
	}
	rows, err := sqliteDb.Query("SELECT lang FROM bot_params WHERE tele_id = $1", teleId)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	botParams := []BotParams{}
	for rows.Next() {
		bp := BotParams{}
		err := rows.Scan(&bp.Lang)
		if err != nil {
			fmt.Println("err -->", err)
			continue
		}
		botParams = append(botParams, bp)
	}

	return botParams[0].Lang, nil
}

func TurnOnListenerWatchingWritePrice(sqliteDb *sql.DB, teleId int64, userPaymentChoose string) error {
	_, err := sqliteDb.Exec("UPDATE bot_params SET watching_write_price = $1, choose_payment = $2 WHERE tele_id = $3", 1, userPaymentChoose, teleId)
	if err != nil {
		return err
	}
	return nil
}

func CheckIsListenerWritePrice(sqliteDb *sql.DB, teleId int64) (bool, string, error) {
	type BotParams struct {
		WatchingWritePrice bool   `json:"watching_write_price"`
		ChoosePayment      string `json:"choices"`
	}
	rows, err := sqliteDb.Query("SELECT watching_write_price, choose_payment FROM bot_params WHERE tele_id = $1", teleId)
	if err != nil {
		return false, "", err
	}
	botParams := []BotParams{}
	for rows.Next() {
		bp := BotParams{}
		err := rows.Scan(&bp.WatchingWritePrice, &bp.ChoosePayment)
		if err != nil {
			continue
		}
		botParams = append(botParams, bp)
	}

	return botParams[0].WatchingWritePrice, botParams[0].ChoosePayment, nil
}

func TurnOnListenerWatchingAddMam(sqliteDb *sql.DB, teleId int64) error {
	_, err := sqliteDb.Exec("UPDATE bot_params SET watching_add_mam = $1 WHERE tele_id = $2", 1, teleId)
	if err != nil {
		return err
	}
	return nil
}

func CheckIsListenerAddM(sqliteDb *sql.DB, teleId int64) (bool, error) {
	type BotParams struct {
		WatchingAddMam bool `json:"watching_add_mam"`
	}
	rows, err := sqliteDb.Query("SELECT watching_add_mam FROM bot_params WHERE tele_id = $1", teleId)
	if err != nil {
		return false, err
	}
	botParams := []BotParams{}
	for rows.Next() {
		bp := BotParams{}
		err := rows.Scan(&bp.WatchingAddMam)
		if err != nil {
			continue
		}
		botParams = append(botParams, bp)
	}

	return botParams[0].WatchingAddMam, nil
}

func TurnOnListenerWatchingFindMam(sqliteDb *sql.DB, teleId int64) error {
	_, err := sqliteDb.Exec("UPDATE bot_params SET watching_find_mam = $1 WHERE tele_id = $2", 1, teleId)
	if err != nil {
		return err
	}
	return nil
}

func CheckIsListenerFindM(sqliteDb *sql.DB, teleId int64) (bool, error) {
	type BotParams struct {
		WatchingFindMam bool `json:"watching_find_mam"`
	}
	rows, err := sqliteDb.Query("SELECT watching_find_mam FROM bot_params WHERE tele_id = $1", teleId)
	if err != nil {
		return false, err
	}
	botParams := []BotParams{}
	for rows.Next() {
		bp := BotParams{}
		err := rows.Scan(&bp.WatchingFindMam)
		if err != nil {
			continue
		}
		botParams = append(botParams, bp)
	}

	return botParams[0].WatchingFindMam, nil
}

func TurnOnListenerWatchingChangeMinLink(sqliteDb *sql.DB, teleId int64) error {
	_, err := sqliteDb.Exec("UPDATE bot_params SET watching_change_minlink = $1 WHERE tele_id = $2", 1, teleId)
	if err != nil {
		return err
	}
	return nil
}

func CheckIsListenerWatchingChangeMinLink(sqliteDb *sql.DB, teleId int64) (bool, error) {
	type BotParams struct {
		WatchingChangeMinLink bool `json:"watching_change_minlink"`
	}
	rows, err := sqliteDb.Query("SELECT watching_change_minlink FROM bot_params WHERE tele_id = $1", teleId)
	if err != nil {
		return false, err
	}
	botParams := []BotParams{}
	for rows.Next() {
		bp := BotParams{}
		err := rows.Scan(&bp.WatchingChangeMinLink)
		if err != nil {
			continue
		}
		botParams = append(botParams, bp)
	}

	return botParams[0].WatchingChangeMinLink, nil
}

func TurnOnListenerWatchingAddBalance(sqliteDb *sql.DB, teleId int64, userTeleId int64) error {
	_, err := sqliteDb.Exec("UPDATE bot_params SET watching_add_balance = $1, choose_tele_id = $2 WHERE tele_id = $3", 1, userTeleId, teleId)
	if err != nil {
		return err
	}
	return nil
}

func CheckIsListenerWatchingAddBalance(sqliteDb *sql.DB, teleId int64) (bool, int64, error) {
	type BotParams struct {
		WatchingAddBalance bool  `json:"watching_add_balance"`
		ChooseTeleId       int64 `json:"choose_tele_id"`
	}
	rows, err := sqliteDb.Query("SELECT watching_add_balance, choose_tele_id FROM bot_params WHERE tele_id = $1", teleId)
	if err != nil {
		return false, 0, err
	}
	botParams := []BotParams{}
	for rows.Next() {
		bp := BotParams{}
		err := rows.Scan(&bp.WatchingAddBalance, &bp.ChooseTeleId)
		if err != nil {
			continue
		}
		botParams = append(botParams, bp)
	}

	return botParams[0].WatchingAddBalance, botParams[0].ChooseTeleId, nil
}

func TurnOnListenerWatchingAddMinUser(sqliteDb *sql.DB, teleId int64, userTeleId int64) error {
	_, err := sqliteDb.Exec("UPDATE bot_params SET watching_change_minuser = $1, choose_tele_id = $2 WHERE tele_id = $3", 1, userTeleId, teleId)
	if err != nil {
		return err
	}
	return nil
}

func CheckIsListenerWatchingAddMinUser(sqliteDb *sql.DB, teleId int64) (bool, int64, error) {
	type BotParams struct {
		WatchingChangeMinUser bool  `json:"watching_change_minuser"`
		ChooseTeleId          int64 `json:"choose_tele_id"`
	}
	rows, err := sqliteDb.Query("SELECT watching_change_minuser, choose_tele_id FROM bot_params WHERE tele_id = $1", teleId)
	if err != nil {
		return false, 0, err
	}
	botParams := []BotParams{}
	for rows.Next() {
		bp := BotParams{}
		err := rows.Scan(&bp.WatchingChangeMinUser, &bp.ChooseTeleId)
		if err != nil {
			continue
		}
		botParams = append(botParams, bp)
	}

	return botParams[0].WatchingChangeMinUser, botParams[0].ChooseTeleId, nil
}

func TurnOnListenerWatchingChangeBalance(sqliteDb *sql.DB, teleId int64, userTeleId int64) error {
	_, err := sqliteDb.Exec("UPDATE bot_params SET watching_change_balance = $1, choose_tele_id = $2 WHERE tele_id = $3", 1, userTeleId, teleId)
	if err != nil {
		return err
	}
	return nil
}

func CheckIsListenerWatchingChangeBalance(sqliteDb *sql.DB, teleId int64) (bool, int64, error) {
	type BotParams struct {
		WatchingChangeBalance bool  `json:"watching_change_balance"`
		ChooseTeleId          int64 `json:"choose_tele_id"`
	}
	rows, err := sqliteDb.Query("SELECT watching_change_balance, choose_tele_id FROM bot_params WHERE tele_id = $1", teleId)
	if err != nil {
		return false, 0, err
	}
	botParams := []BotParams{}
	for rows.Next() {
		bp := BotParams{}
		err := rows.Scan(&bp.WatchingChangeBalance, &bp.ChooseTeleId)
		if err != nil {
			continue
		}
		botParams = append(botParams, bp)
	}

	return botParams[0].WatchingChangeBalance, botParams[0].ChooseTeleId, nil
}

func TurnOnListenerWatchingMessageUser(sqliteDb *sql.DB, teleId int64, userTeleId int64) error {
	_, err := sqliteDb.Exec("UPDATE bot_params SET watching_message_user = $1, choose_tele_id = $2 WHERE tele_id = $3", 1, userTeleId, teleId)
	if err != nil {
		return err
	}
	return nil
}

func CheckIsListenerWatchingMessageUser(sqliteDb *sql.DB, teleId int64) (bool, int64, error) {
	type BotParams struct {
		WatchingMessageUser bool  `json:"watching_message_user"`
		ChooseTeleId        int64 `json:"choose_tele_id"`
	}
	rows, err := sqliteDb.Query("SELECT watching_message_user, choose_tele_id FROM bot_params WHERE tele_id = $1", teleId)
	if err != nil {
		return false, 0, err
	}
	botParams := []BotParams{}
	for rows.Next() {
		bp := BotParams{}
		err := rows.Scan(&bp.WatchingMessageUser, &bp.ChooseTeleId)
		if err != nil {
			continue
		}
		botParams = append(botParams, bp)
	}

	return botParams[0].WatchingMessageUser, botParams[0].ChooseTeleId, nil
}

func TurnOffListeners(sqliteDb *sql.DB, teleId int64) error {
	_, err := sqliteDb.Exec("UPDATE bot_params SET watching_write_price = $1, choose_payment = $2, watching_add_mam = $3, watching_find_mam = $4, watching_change_minlink = $5, watching_add_balance = $6, watching_change_minuser = $7, watching_change_balance = $8, watching_message_user = $9, choose_tele_id = $10 WHERE tele_id = $11", 0, "", 0, 0, 0, 0, 0, 0, 0, 0, teleId)
	if err != nil {
		return err
	}
	return nil
}
