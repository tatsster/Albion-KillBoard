package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tatsster/albion_killboard/config"
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

func GetAllMemberID(db *sql.DB) ([]string, error) {
	var (
		memberIDs = make([]string, 0)
	)

	rows, err := db.Query("SELECT id FROM members ORDER BY id")
	if err != nil {
		return memberIDs, err
	}

	defer rows.Close()

	for rows.Next() {
		var memberID string
		err := rows.Scan(&memberID)
		if err != nil {
			return memberIDs, err
		}
		memberIDs = append(memberIDs, memberID)
	}

	return memberIDs, rows.Err()
}

func GetMemberByID(db *sql.DB, ID string) (config.Member, error) {
	var (
		member config.Member
	)

	rows, err := db.Query("SELECT * FROM members WHERE id = ?", ID)
	if err != nil {
		return member, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&member.ID, &member.Name, &member.LastKill, &member.LastDeath)
		if err != nil {
			return member, err
		}
	} else {
		return member, fmt.Errorf("no member found with ID: %s", ID)
	}
	return member, rows.Err()
}

func UpdateKillTime(db *sql.DB, kill config.Event) error {
	sql := "UPDATE members SET last_kill = ? WHERE id = ?"

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("start transaction error: %v", err)
	}

	updateKill, err := tx.Prepare(sql)
	if err != nil {
		return fmt.Errorf("prepare statement error: %v", err)
	}

	_, err = updateKill.Exec(kill.TimeStamp, kill.Killer.ID)
	if err != nil {
		return fmt.Errorf("execute sql error: %v", err)
	}
	updateKill.Close()
	return tx.Commit()
}

func UpdatDeathTime(db *sql.DB, death config.Event) error {
	sql := "UPDATE members SET last_death = ? WHERE id = ?"

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("start transaction error: %v", err)
	}

	updateKill, err := tx.Prepare(sql)
	if err != nil {
		return fmt.Errorf("prepare statement error: %v", err)
	}

	_, err = updateKill.Exec(death.TimeStamp, death.Victim.ID)
	if err != nil {
		return fmt.Errorf("execute sql error: %v", err)
	}
	updateKill.Close()
	return tx.Commit()
}
