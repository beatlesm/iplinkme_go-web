beeblog00.0402
实现了login登录 、在首页显示登录、退出等功能。

但是点退出时会没有反应，只能通过网页输入http://127.0.0.1:8080/login?ex=true地址退出
在chrome只第一次能用，第二次就不能用了。
必须修改login.go
func (this *LoginController) Get(){
	fmt.Printf("222222\n")

	isExit := this.Input().Get("ex") == "true"
代码

原因未找到
