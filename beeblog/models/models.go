package models
import (
	"time"
	"path"
	"os"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)
const(
	_DB_NAME          = "data/beeblog.db"
	_SQLITE3_DRIVER   = "sqlite3"
)
type Category struct{
	Id                  int64
	Title               string
	Created             time.Time   `orm:"index"`
	View                int64 	    `orm:"index"`
	TopicTime           time.Time   `orm:"index"`
	TopicCount          int64
	TopiclastUserId     int64
}
type Topic struct{
	Id           int64
	Uid          int64
	Title        string
	Content      string      `orm:"size(5000)"`
	Attachent    string
	Created      time.Time   `orm:"index"`
	Updated      time.Time   `orm:"index"`
	View         int64 	     `orm:"index"`
	Author       string
	ReplyTime    time.Time   `orm:"index"`
	ReplyCount   int64
	RepleyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME){
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}