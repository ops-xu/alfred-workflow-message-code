package main

import (
	"database/sql"
	"github.com/deanishe/awgo"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"regexp"
)

var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

const (
	dbDriverName = "sqlite3"
)

const query = "SELECT\nt.text\nFROM\nmessage t\nWHERE\nt.\"text\" LIKE \"%验证码%\"\norder by t.ROWID desc\nlimit 10"

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	wf.Run(run)
}

const code_regexp = `：\d{4,8}|\d{4,8}，|\d{4,8}。|\d{4,8}`
const nums_regexp = `\d{4,8}`

func run() {
	dir, _ := os.UserHomeDir()
	dbName := dir + "/Library/Messages/chat.db"
	db, err := sql.Open(dbDriverName, dbName)
	checkErr(err)
	rows, err := db.Query(query)
	checkErr(err)
	for rows.Next() {
		var text string
		rows.Scan(&text)
		result := regMatch(text)
		sendItem(text, result)
	}
	sendFeedback()
	db.Close()
}

func sendItem(text string, result string) {
	wf.NewItem(text).Subtitle(result).Arg(result).Valid(true)
}

func sendFeedback() {
	wf.SendFeedback()
}

func regMatch(text string) string {
	log.Println(text)
	codeReg := regexp.MustCompile(code_regexp)
	numReg := regexp.MustCompile(nums_regexp)
	match := codeReg.FindAllStringSubmatch(text, -1)
	if match == nil {
		return "未匹配到验证码"
	}
	s := match[0][0]
	code := numReg.FindAllStringSubmatch(s, -1)
	log.Println(code[0][0])
	return code[0][0]
}
