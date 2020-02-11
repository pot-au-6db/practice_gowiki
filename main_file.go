package main

import (
	"io/ioutil"
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
	// 引数string titleと".txt"を結合
	filename := title + ".txt"

	// ioutilのReadFileメソッドで、bodyにfilenameを格納
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 返すのはstruct Page{}とエラー用nil
	return &Page{Title: title, Body: body}, nil
}
