package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int64
	Name     string
	Phone    string
	Token    string
	Platform string
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	// 确保数据库文件所在目录存在
	dir := filepath.Dir(dataSourceName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("创建数据库目录失败: %v", err)
	}

	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	return &DB{db}, nil
}

func (db *DB) InitTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		phone TEXT NOT NULL,
		token TEXT NOT NULL,
		platform TEXT NOT NULL
	);`

	_, err := db.Exec(query)
	return err
}

func (db *DB) AddUser(user *User) error {
	query := `
	INSERT INTO users (name, phone, token, platform)
	VALUES (?, ?, ?, ?)`

	result, err := db.Exec(query, user.Name, user.Phone, user.Token, user.Platform)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (db *DB) GetAllUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, name, phone, token, platform FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Token, &user.Platform)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
} 