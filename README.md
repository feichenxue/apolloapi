![apolloapi-go](https://img.shields.io/badge/apolloapi--go-v0.1-brightgreen)
[![license](http://dmlc.github.io/img/apache2.svg)](https://github.com/feichenxue/apolloapi/blob/master/LICENSE)

apolloapi
=========

```
apollo 开放平台接口，其接口包含了增删改查等基本功能，另外还有创建cluster,namespace等等。
```

# Installation

```bash
go version >= 1.14 GO111MODULE=on
```

```bash
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

	ok, err := apollo.ReleaseApollo("go-test-project", "prd", "comment-test", "default", "application")
	if err != nil {
		fmt.Println(err)
	}

	if ok {
		fmt.Println("发布成功!!!")
	}
}
```

# 接口列表

* NewApollo

```go
<<<<<<< HEAD
func NewApollo(apollourl, username, passwd string) *apolloApi
=======
func NewApollo(apollourl, username, passwd string) *Apolloapi
>>>>>>> 60358bf1326ed4f3aff0400049abfff6ae3d14dd
```
新建apollo一个对象，用以调用此包内的其它方法

* CreateApolloProject

```go
<<<<<<< HEAD
func (apollo *apolloApi) CreateApolloProject(appId, name, orgId, orgName, ownerName string) (bool, error)
=======
func (apollo *Apolloapi) CreateApolloProject(appId, name, orgId, orgName, ownerName string) (bool, error)
>>>>>>> 60358bf1326ed4f3aff0400049abfff6ae3d14dd
```
新建apollo一个项目

* ReleaseApollo

```go
<<<<<<< HEAD
func (apollo *apolloApi) ReleaseApollo(appid, env, comment, clustersName, namespaceName string) (bool, error)
```
发布配置接口

```go
func (apollo *apolloApi) GetappIDList() []string
```
获取apollos所有appid详细信息，其返回值类型为字符串切片

```go
func (apollo *apolloApi) CreateCluster(appId, env, clusterName string) (bool, error)
```
创建Cluster

```go
func (apollo *apolloApi) GetConfigData(appId, env, clusters, namespaces string) (data map[string]string, err error)
```
获取相应appid的相应集群下相应namespaces下的配置
```

```go
未完待续······
```


=======
func (apollo *Apolloapi) ReleaseApollo(appid, env, comment, clustersName, namespaceName string) (bool, error)
```
发布配置接口

>>>>>>> 60358bf1326ed4f3aff0400049abfff6ae3d14dd
