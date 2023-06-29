package db

import (
	"database/sql"
	"fmt"

	"github.com/tatsster/albion_killboard/config"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteHandler() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.SQLITE_PATH)
	if err != nil {
		fmt.Println("fail to init database: ", err)
		return nil, err
	}

	err = setupTables(db)
	if err != nil {
		fmt.Println("fail to setup tables: ", err)
		return nil, err
	}
	return db, nil
}

func setupTables(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS members (
		id TEXT PRIMARY KEY, 
		name TEXT,
		last_kill DATETIME,
		last_death DATETIME)`,
	)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMembers(db *sql.DB, members config.MemberInfo) error {
	// Insert multiple values in a single transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO members (id, name) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, member := range members {
		_, err = stmt.Exec(member.ID, member.Name)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func GetMembers(db *sql.DB) ([]config.Member, error) {
	var (
		members = make([]config.Member, 0)
	)

	rows, err := db.Query("SELECT * FROM members")
	if err != nil {
		return members, err
	}

	defer rows.Close()

	for rows.Next() {
		var member config.Member
		err := rows.Scan(&member.ID, &member.Name, &member.LastKill, &member.LastDeath)
		if err != nil {
			return members, err
		}
		members = append(members, member)
	}

	return members, rows.Err()
}
