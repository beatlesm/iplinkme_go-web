package models

import (
	//"fmt"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	// 设置数据库路径
	_DB_NAME = "data/beeblog.db"
	// 设置数据库名称
	_SQLITE3_DRIVER = "sqlite3"
)

// 分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	View            int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:"size(5000)"` //文章内容
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64     `orm:"index"`
	Author           string
	ReplyTime        time.Time `orm:"index"`
	ReplyCount       int64
	RepleyLastUserId int64
}

func RegisterDB() {
	// 检查数据库文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	// 注册模型
	orm.RegisterModel(new(Category), new(Topic))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}
func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}
	//fmt.Printf("CC run AddCategory AAAAAAAAAAAAAAAA\n")
	// 查询数据
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	//fmt.Printf("CC run AddCategory BBBBBBBBBBBBBBBB\n")
	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	// 查询table
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	//fmt.Printf("CC run GetAllCategories \n")
	return cates, err
}
func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}
func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	// 查询table
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		//show topics
		_, err = qs.OrderBy("-created").All(&topics)
		fmt.Printf("4444444444\n")
	} else {
		_, err = qs.All(&topics)
		fmt.Printf("5555555555\n")
	}
	return topics, err
}

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	//插入
	_, err := o.Insert(topic)
	return err
}

//func GetTopic(tid string) (*Topic, error) {
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	//fmt.Printf("66666666666666\n")
	//fmt.Println("6666", tidNum)
	//fmt.Printf("66666666666666\n")
	//_, err = qs.OrderBy("-created").All(&topic)
	// Filter show topics
	err = qs.Filter("id", tidNum).One(topic)
	fmt.Println(err)
	if err != nil {
		fmt.Printf("77777777777\n")
		return nil, err
	}
	topic.Views++
	//fmt.Println("6666yyyyyyyyyyyyyyyyyyyyyyyy", topic.Views)
	_, err = o.Update(topic)
	//fmt.Printf("888888888888888\n")
	return topic, err
}

func ModifyTopic(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return nil
}
