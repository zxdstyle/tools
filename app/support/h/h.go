package h

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net"
)

// 定义一些全局辅助方法

func Failed(r *ghttp.Request, message string, code ...interface{}) {
	status := 400
	if len(code) > 0 {
		status = code[0].(int)
	}

	r.Response.WriteHeader(status)
	r.Response.WriteJsonExit(g.Map{
		"code":    status,
		"message": message,
	})
}

func Success(r *ghttp.Request, data ...interface{}) {
	if len(data) > 0 {
		r.Response.WriteJsonExit(data[0])
	}
	r.Response.WriteJsonExit(g.Map{
		"code":    200,
		"message": "success",
	})
}

// 获取服务器Ip
func GetServerIp() (ip string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
			}
		}
	}

	return
}
