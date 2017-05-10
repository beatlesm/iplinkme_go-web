
# beeblog  log
## 02/24/2017
**beeblog00.0402**
实现了login登录 、在首页显示登录、退出等功能。

但是点退出时会没有反应，只能通过网页输入http://127.0.0.1:8080/login?ex=true地址退出

在chrome只第一次能用，第二次就不能用了。

必须修改login.go

``` go
func (this *LoginController) Get(){
	isExit := this.Input().Get("ex") == "true"
```
代码.
原因未找到

**beeblog00.0403**

添加了帐号于密码的验证

**beeblog00.0404**

添加了登录页面的返回按钮、添加了category页面.对数据库进行了添加、读取等操作

**beeblog00.0405**

对了category页面，进行了删除操作

## 02/25/2017

**beeblog00.0501**

添加了topic.go topic.html页面

**beeblog00.0502**

添加了topic页面的数据库添加、读取等操作

**beeblog00.0503**

添加了topic页面的文章排序、操作.添加首页显示

## 02/27/2017

**beeblog00.0504**

添加了topic/view页面的文章排序、读取操作

## 02/28/2017

**beeblog00.0505**

添加了topic/view页面的修改文章操作

## 03/1/2017

**beeblog00.0506**

添加了topic/view页面的删除文章操作

## 03/3/2017 星期5
### beeblog00.0601
添加回复页面的操作 

## 03/7/2017 星期2

**beeblog00.0602**

添加删除回复页面的操作 

**beeblog00.0603**

添加home页面的文章分类 的操作 

## 03/8/2017 星期3

**beeblog00.0604**

添加文章分页面 显示文章数的操作 
添加文章标签页面 分类显示某个标签的文章（有bug） 

## 03/15/2017 星期4

**beeblog00.0701**
添加文章附件上传功能