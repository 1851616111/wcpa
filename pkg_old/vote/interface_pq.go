package vote

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"github.com/lib/pq"
	"sort"
	"time"
)

var TABLE_VOTER string = "GO_WEICHAT_ACTIVITY_VOTER"
var SQL_TABLE = `
CREATE TABLE IF NOT EXISTS  GO_WEICHAT_ACTIVITY_VOTER (
   openid varchar(50) PRIMARY KEY,
   voterid SERIAL,
   name varchar(15),
   image varchar(300),
   company varchar(100),
   mobile varchar(20),
   declaration varchar(50),
   votedcount integer DEFAULT 0,
   voteRecords integer[]
);
`

func NewDBInterface(ip, port, user, passwd, database string) (DBInterface, error) {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, database)

	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

type DB struct {
	*sql.DB
}

func (d DB) Init() (err error) {
	_, err = d.Exec(SQL_TABLE)
	return
}

func (d DB) Register(openid string, v *Voter) error {
	_, err := d.Exec(`INSERT INTO `+TABLE_VOTER+` (openid, name, image, company, mobile, declaration)
	VALUES($1,$2,$3,$4,$5,$6)`, openid, v.Name, v.Image, v.Company, v.Mobile, v.Declaration)
	return err
}

func (d DB) Vote(openID, votedID string) (err error) {
	voteRecords := pq.Int64Array{}
	if err = d.QueryRow(`SELECT voteRecords FROM `+TABLE_VOTER+` WHERE openid = $1`, openID).Scan(&voteRecords); err != nil {
		return
	}

	records := []int64(voteRecords)
	if !hasVoteRight(records) {
		err = errors.New("vote num are only allowed 3 times per day")
		return
	}

	var tx *sql.Tx
	tx, err = d.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec(`UPDATE `+TABLE_VOTER+` SET votedcount = votedcount + 1 where voterid = $1`, votedID)
	if err != nil {
		return
	}

	voteRecords = append(voteRecords, time.Now().Unix())
	_, err = tx.Exec(`UPDATE `+TABLE_VOTER+` SET voteRecords = $1 where openid = $2`, voteRecords, openID)
	if err != nil {
		return
	}

	return tx.Commit()
}

func (d DB) ListVoters(index, size int) ([]Voter, error) {
	rows, err := d.Query(`SELECT voterid, name, image, votedcount FROM `+TABLE_VOTER+` ORDER BY votedcount DESC LIMIT $1 OFFSET $2`, size, index-1)
	if err != nil {
		return nil, err
	}

	l := []Voter{}
	for rows.Next() {
		v := Voter{}
		if err := rows.Scan(&v.ID, &v.Name, &v.Image, &v.VotedCount); err != nil {
			glog.Errorf("list voters scan row err %v\n", err)
			continue
		}
		l = append(l, v)
	}

	return l, nil
}

func hasVoteRight(records []int64) bool {
	if len(records) >= 3 {
		sort.Sort(Int64Slice(records))
		return time.Now().Unix()-records[2] >= 3600*24
	} else {
		return true
	}

}
