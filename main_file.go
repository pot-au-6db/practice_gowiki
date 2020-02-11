package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// ファイルをセーブする関数
func (p *Page) save() error {

	// 引数struct pのTitle要素と".txt"を結合
	filename := p.Title + ".txt"

	// ioutilのWriteFileメソッドで、引数struct pのボディを作成したファイル名で書き込む。
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// ファイルをロードする
func loadPage(title string) (*Page, error) {

	Dbcon, _ := sql.Open("sqlite3", "./test.sql")
	defer Dbcon.Close()

	cmd := `SELECT * FROM pages WHERE title = ?`
	rows := DbCon.QueryRow(cmd, title)

	var p Page
	err := rows.Scan(&p.Title, &p.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Fatalln(err)
		}
	}

	return &Page{Title: p.Title, Body: p.Body}, nil
}
