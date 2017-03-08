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
	"strings"
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
	Title            string //文章标题
	Category         string //文章分类名称
	Labels           string
	Content          string    `orm:"size(5000)"` //文章内容
	Attachment       string    //
	Created          time.Time `orm:"index"` //
	Updated          time.Time `orm:"index"` //
	Views            int64     `orm:"index"` //浏览
	Author           string    //
	ReplyTime        time.Time `orm:"index"` //回复时间
	ReplyCount       int64     //回复数
	RepleyLastUserId int64     //
}
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

func RegisterDB() {
	// 检查数据库文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	// 注册模型
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}
func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}
func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replies = make([]*Comment, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
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
	//topics, err := models.GetAllTopics(this.Input().Get("cate"), true)
	//cate.TopicCount = len(topics)
	//cate.TopicCount--
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

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	var oldCate string
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	return err
}
func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	var tidNum int64
	reply := &Comment{Id: ridNum}
	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}
	replies := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = replies[0].Created
		topic.ReplyCount = int64(len(replies))
		_, err = o.Update(topic)
	}
	return err
}

func GetAllTopics(cate, label string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	// 查询table
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0 {
			qs.Filter("labels__contains", "$"+"label"+"#")
		}
		//show topics
		_, err = qs.OrderBy("-created").All(&topics)
		fmt.Printf("4444444444\n")
	} else {
		_, err = qs.All(&topics)
		fmt.Printf("5555555555\n")
	}
	return topics, err
}

func AddTopic(title, category, label, content string) error {
	//处理标签
	label = "$" + strings.Join(
		strings.Split(label, " "), "#$") + "#"
	//空格作为多个标签的分隔符
	//beego           $beego#
	//orm             $beego#$orm#
	o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Category:  category,
		Labels:    label,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	//插入
	_, err := o.Insert(topic)
	if err != nil {
		return err
	}
	//更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		//如果不存在，简单地忽略更新操作
		cate.TopicCount++
		_, err = o.Update(cate)
	}

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

	topic.Labels = strings.Replace(strings.Replace(
		topic.Labels, "#", " ", -1), "$", " ", -1)
	//fmt.Printf("888888888888888\n")
	return topic, err
}

func ModifyTopic(tid, title, category, label, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	var oldCate string
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		topic.Title = title
		topic.Category = category
		topic.Labels = label
		topic.Content = content
		topic.Updated = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}
	//更新分类统计
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	return nil
}
