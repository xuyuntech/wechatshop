package models

import (
	"github.com/go-xorm/xorm"
	"time"
	"github.com/satori/go.uuid"
)

type User struct {
	Id 	string `xorm:"pk"`
	Name 	string `xorm:"INDEX UNIQUE NOT NULL"`
	Password string `xorm:"NOT NULL"`
	Email string `xorm:"UNIQUE NOT NULL"`

	Created     time.Time `xorm:"-"`
	CreatedUnix int64
	Updated     time.Time `xorm:"-"`
	UpdatedUnix int64
}

func (me *User) BeforeInsert() {
	me.Id = uuid.NewV4().String()
	me.CreatedUnix = time.Now().Unix()
	me.UpdatedUnix = me.CreatedUnix
}

func (me *User) BeforeUpdate() {
	me.UpdatedUnix = time.Now().Unix()
}

func (me *User) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_unix":
		me.Created = time.Unix(me.CreatedUnix, 0).Local()
	case "updated_unix":
		me.Updated = time.Unix(me.UpdatedUnix, 0)
	}
}

func (me *User) APIJson() map[string]interface{} {
	return map[string]interface{}{
		"id": me.Id,
		"name": me.Name,
		"email": me.Email,
		"created": time.Unix(me.CreatedUnix, 0).Format("2006-01-02"),
		"updated": time.Unix(me.UpdatedUnix, 0).Format("2006-01-02"),
	}
}
