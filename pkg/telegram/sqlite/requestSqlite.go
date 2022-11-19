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
