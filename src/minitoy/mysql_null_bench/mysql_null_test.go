package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type row struct {
	name   string
	number string
}

func connect() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/my_test?"+
		"charset=utf8&multiStatements=true")
	if err != nil {
		panic(err)
	}
}

func createTable() {
	s := `
		DROP TABLE IF EXISTS not_null;
		CREATE TABLE not_null
		(
			id INT PRIMARY KEY AUTO_INCREMENT,
		    name VARCHAR(50) NOT NULL,
		    number int(10)
		);
		CREATE INDEX allow_null_key_index ON not_null (name);

		DROP TABLE IF EXISTS allow_null;
		CREATE TABLE allow_null
		(
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(50),
			number INT(10)
		);
		CREATE INDEX allow_null_name_index ON allow_null (name);
	`
	_, err := db.Exec(s)
	if err != nil {
		panic(err)
	}
}

func mockData(n int) []row {
	letter := "abcdefghijklmnopqrstuvwxyz"
	ins := []row{}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		number := strconv.Itoa(1000 + rd.Intn(9000))
		name := ""
		if rd.Intn(6) != 5 {
			for j := 0; j < 4; j++ {
				p := rd.Intn(26)
				name += letter[p : p+1]
			}
		}
		ins = append(ins, row{name, number})
	}
	fmt.Println("mock done...")
	return ins
}

func insertData(data []row, chunk int) {
	var err error
	tmp1 := "INSERT INTO not_null (name,number) VALUES "
	tmp2 := "INSERT INTO allow_null (name,number) VALUES "

	f := float64(len(data) / chunk)
	count := int(math.Floor(f) + 1)
	for i := 0; i < count; i++ {
		ins1 := tmp1
		ins2 := tmp2
		remain := chunk
		if i == count-1 {
			remain = len(data) - i*chunk
			if remain <= 0 {
				return
			}
		}
		for _, row := range data[i*chunk : i*chunk+remain] {
			if row.name == "" {
				ins1 += "(''," + row.number + "),"
				ins2 += "(" + "NULL" + "," + row.number + "),"
			} else {
				ins1 += "('" + row.name + "'," + row.number + "),"
				ins2 += "('" + row.name + "'," + row.number + "),"
			}
		}
		_, err = db.Exec(ins1[0 : len(ins1)-1])
		_, err = db.Exec(ins2[0 : len(ins2)-1])
		if err != nil {
			panic(err)
		}
		fmt.Println("insert chunk...")
	}
}

func init() {
	connect()
	defer db.Close()
	createTable()
	data := mockData(1000000)
	insertData(data, 10000)
}

func BenchmarkNotNull(b *testing.B) {
	connect()
	defer db.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		row, err := db.Query("SELECT * FROM not_null WHERE name = 'abcd'")
		if err != nil {
			log.Fatalln(err)
		}
		row.Close()
	}
}

func BenchmarkAllowNull(b *testing.B) {
	connect()
	defer db.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		row, err := db.Query("SELECT * FROM allow_null WHERE name = 'abcd'")
		if err != nil {
			log.Fatalln(err)
		}
		row.Close()
	}
}
