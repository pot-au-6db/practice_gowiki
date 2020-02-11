package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Title string
	Body  []byte
}

var DbCon *sql.DB

// templates変数にedit、view.htmlを格納。（キャッシング）
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplete(w http.ResponseWriter, tmpl string, p *Page) {

	// templatesからtmpl引数で得たページを実行する。エラーの場合は500を返す。
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// HandleFuncの第二引数に一致した引数をもつ。
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	// 無名関数（responsewriter, request）を返す。
	return func(w http.ResponseWriter, r *http.Request) {
		// mにvalidPath変数からRequest内のURLパスに一致した文字列を格納
		m := validPath.FindStringSubmatch(r.URL.Path)
		// mがnilならNotFound。
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {

	DbConnection, _ := sql.Open("sqlite3", "./test.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS pages(
			title STRING,
			body STRING)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// NotFound以外、つまり自分の指定したURLにアクセスしたい場合、自分でハンドラーを作成する。
	// HandleFuncは第二引数に関数(responsewriter,request)Handlerをとる。
	// 今回は、短くコードを書くため、makeHandler関数に突っ込んだ。
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// ListenAndServerは(Addr, Handler)を引数にとる。
	// Handlerにnilを指定した際は、デフォルトのハンドラー（NotFound）を返す。
	// サーバーがエラーをしたときは、エラーを返す。
	log.Fatal(http.ListenAndServe(":8080", nil))
}
