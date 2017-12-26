package user

import (
    "gitee.com/johng/gf/g/net/ghttp"
    "gitee.com/johng/gf/g/frame/gmvc"
)

// 定义业务相关的控制器对象
type ControllerUser struct {
    gmvc.Controller
}

func Test(s *ghttp.Server, r *ghttp.ClientRequest, w *ghttp.ServerResponse) {
    w.WriteString("Test")
    w.Output()
}

// 初始化控制器对象，并绑定操作到Web Server
func init() {
    //ghttp.GetServer("johng").Domain("localhost").BindHandler("/user", u.Info)
    ghttp.GetServer("johng").BindHandler("/test", Test)
    ghttp.GetServer("johng").BindController("/user", &ControllerUser{})
}

// 定义操作逻辑
func (c *ControllerUser) Info() {
    c.Response.WriteString("hello world!")
    //c.View.Assign("name", "john")
    //c.View.Display("user/index")
}



