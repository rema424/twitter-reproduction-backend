package dbutil

import (
	"bytes"
	"database/sql"
	"fmt"
	"time"

	"github.com/rema424/twitter-reproduction-backend/lib/logutil"

	"github.com/go-sql-driver/mysql"
)

// log ...
func (s *Session) log(query string, args []interface{}, rows int, start time.Time, err error) {
	if s == nil || s.context == nil {
		return
	}

	// Query 時間を取得
	passed := time.Since(start)
	ms := float64(passed) / float64(time.Millisecond)

	// log を書き込んでいく
	var buf bytes.Buffer
	// error
	if err != nil && err != sql.ErrNoRows {
		buf.WriteString(err.Error())
		buf.WriteString(" ")
	}
	// time
	fmt.Fprintf(&buf, "[%.2f ms] ", ms)
	// rows
	fmt.Fprintf(&buf, "[%d rows] ", rows)
	// query
	buf.WriteString(query)
	// args
	if s.secretArgs {
		// 特秘事項が含まれる場合はログ出力しない
		s.secretArgs = false
	} else if len(args) != 0 {
		buf.WriteString(" [")
		for i, arg := range args {
			if i != 0 {
				buf.WriteString(", ")
			}
			fmt.Fprint(&buf, arg)
		}
		buf.WriteString("]")
	}

	// log の種類
	logger := logutil.Debugf
	if err != nil && err != sql.ErrNoRows {
		if isDuplicateEntry(err) || isForeignKeyConstraint(err) {
			logger = logutil.Infof
		} else {
			logger = logutil.Warningf
		}
	} else if passed > s.warningTime || rows > s.warningRows {
		logger = logutil.Warningf
	}

	logger(s.context, buf.String())
}

// isDuplicateEntry ...
func isDuplicateEntry(err error) bool {
	if err == nil {
		return false
	}
	merr, ok := err.(*mysql.MySQLError)
	return ok && merr.Number == 1062
}

// isForeignKeyConstraint ...
func isForeignKeyConstraint(err error) bool {
	if err == nil {
		return false
	}
	merr, ok := err.(*mysql.MySQLError)
	return ok && merr.Number == 1452
}
