package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetSessionPG(strConn string) (db *gorm.DB, err error) {
	dsn := strConn
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return
}

func DBInit(db *gorm.DB) error {
	query := `CREATE TABLE IF NOT EXISTS event_soul_history (id serial primary key, uid varchar(100) NOT NULL, session varchar(100) NOT NULL, img varchar(100) NOT NULL,created numeric NOT NULL, updated numeric, "userName" varchar(50)NOT NULL, "groupId" varchar(50), "extraX" varchar(100), "unVote" boolean DEFAULT FALSE)`
	//query := `UPDATE event_soul_history SET "unVote" = TRUE WHERE "groupId" = '2306'`
	if err := db.Exec(query).Error; err != nil {
		return err
	}
	return nil
}
