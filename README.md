![apolloapi-go](https://img.shields.io/badge/apolloapi--go-v0.1-brightgreen)
[![license](http://dmlc.github.io/img/apache2.svg)](https://github.com/feichenxue/apolloapi/blob/master/LICENSE)

apolloapi
=========

```
apollo 开放平台接口，其接口包含了增删改查等基本功能，另外还有创建cluster,namespace等等。
```

# Installation

```
go get -u github.com/feichenxue/apolloapi
```

# 如何使用？


## 简单示例



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


