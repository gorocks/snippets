package mysql_test

import (
	"database/sql"
	"log"
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	conn, err := sql.Open("mysql", "root:@/mysql?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalln(err)
	}
	db = conn
}

func initTables() {
	// +------------+---------------+------+-----+---------+-------+
	// | Field      | Type          | Null | Key | Default | Extra |
	// +------------+---------------+------+-----+---------+-------+
	// | emp_no     | int(11)       | NO   | PRI | <null>  |       |
	// | birth_date | date          | NO   |     | <null>  |       |
	// | first_name  | varchar(14)   | NO   |     | <null>  |       |
	// | last_name  | varchar(16)   | NO   |     | <null>  |       |
	// | gender     | enum('M','F') | NO   |     | <null>  |       |
	// | hire_date  | date          | NO   |     | <null>  |       |
	// +------------+---------------+------+-----+---------+-------+
	s := "CREATE TABLE IF NOT EXISTS `employees` (`emp_no` int(11) NOT NULL AUTO_INCREMENT,`birth_date` date NOT NULL," +
		"`first_name` varchar(14) NOT NULL,`last_name` varchar(16) NOT NULL,`gender` enum('M','F') NOT NULL,`hire_date` date NOT NULL," +
		"PRIMARY KEY (`emp_no`)) ENGINE=InnoDB DEFAULT CHARSET=utf8"
	if _, err := db.Exec(s); err != nil {
		log.Fatalln(err)
	}
}

func cleanTables() {
	s := "DROP TABLE IF EXISTS `employees`"
	if _, err := db.Exec(s); err != nil {
		log.Fatalln(err)
	}
}

func TestTableInsert(t *testing.T) {
	initTables()
	defer cleanTables()
	s, args, _ := sq.Insert("employees").SetMap(sq.Eq{
		"birth_date": "1953-09-02",
		"first_name": "Georgi",
		"last_name":  "Facello",
		"gender":     "M",
		"hire_date":  "1986-06-26",
	}).ToSql()
	if _, err := db.Exec(s, args...); err != nil {
		return
	}
	s, args, _ = sq.Select("emp_no", "birth_date", "gender").From("employees").Where(sq.Eq{"emp_no": 1}).Limit(1).ToSql()
	var employee struct {
		empNo     int
		birthDate time.Time
		gender    string
	}
	if rows, err := db.Query(s, args...); err == nil {
		rows.Next()
		if err = rows.Scan(&employee.empNo, &employee.birthDate, &employee.gender); err != nil {
			t.Error(err)
		} else {
			if employee.empNo == 0 || employee.gender == "" {
				t.Errorf("rows %v scan failed", rows)
			}
		}
	} else {
		t.Error(err)
	}
}
