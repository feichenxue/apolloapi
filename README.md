![apolloapi-go](https://img.shields.io/badge/apolloapi--go-v0.1-brightgreen)

apolloapi
=========

如何使用？
=========

# 简单示例

* 发布接口

```golang
package main

import (
	"fmt"
	"github.com/feichenxue/apolloapi"
)

func main() {

	url := "http://x.x.x.x:8080" //你的apollo port 地址
	user := "apollo"     //apollo 用户名, 默认为 apollo
	passwds := "xxxxx"  //apollo 密码

	apollo := apolloapi.NewApollo(url, user, passwds)

	ok, err := apollo.ReleaseApollo("go-test4", "prd", "test", "default", "application")
	if err != nil {
		fmt.Println(err)
	}

	if ok {
		fmt.Println("发布成功!!!")
	}
}
```


