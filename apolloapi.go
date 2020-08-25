package apolloapi

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asmcos/requests"
)

type apolloApi struct {
	Username  string
	Passwd    string
	ApolloURL string
	Req       *requests.Request
}

func NewApollo(apollourl, username, passwd string) *apolloApi {
	return &apolloApi{
		Username:  username,
		Passwd:    passwd,
		ApolloURL: apollourl,
		Req:       requests.Requests(),
	}
}

func (apollo *apolloApi) GetappIdList() []string {
	url := fmt.Sprintf("%s/apps/by-owner?owner=apollo", apollo.ApolloURL)
	rspone, _ := apollo.Req.Get(url, requests.Auth{apollo.Username, apollo.Passwd})
	var (
		applist []byte
		data    []string
	)

	applist = append(applist, []byte(rspone.Text())...)

	page := 1
	for {
		purl := fmt.Sprintf("%s/apps/by-owner?owner=apollo&page=%s&size=10", apollo.ApolloURL, strconv.Itoa(page))
		resp, _ := apollo.Req.Get(purl, requests.Auth{apollo.Username, apollo.Passwd})

		page++
		if resp.Text() == "[]" {
			break
		}
		applist = append(applist, []byte(resp.Text())...)
	}
	applistData := string(applist)
	applistData = strings.Replace(strings.Replace(applistData, "[", "", -1), "]", "", -1)
	applistData = strings.Replace(applistData, "},", "} ", -1)
	data = strings.Split(applistData, " ")
	return data
}

func (apollo *apolloApi) CreateApolloProject(appId, name, orgId, orgName, ownerName string) (bool, error) {
	url := fmt.Sprintf("%s/apps", apollo.ApolloURL)
	data := fmt.Sprintf(`{
		"appId":%s,
		"name":%s,
		"orgId":%s,
		"orgName":%s,
		"ownerName":%s
		}`, appId, name, orgId, orgName, ownerName)
	resp, _ := apollo.Req.PostJson(url, data, requests.Auth{apollo.Username, apollo.Passwd})

	if resp.R.StatusCode != 200 {
		return false, fmt.Errorf(resp.Text())
	}
	return true, nil
}

func (apollo *apolloApi) ReleaseApollo(appid, env, comment, clustersName, namespaceName string) (bool, error) {
	url := fmt.Sprintf("%s/apps/%s/envs/%s/clusters/%s/namespaces/%s/releases",
		apollo.ApolloURL,
		appid,
		strings.ToUpper(env),
		clustersName,
		namespaceName,
	)
	data := fmt.Sprintf(`{
		"releaseTitle":%s,
		"releaseComment":%s,
		"releasedBy":%s
		}`, time.Now().Format("20060102150405")+"-release",
		comment,
		apollo.Username)
	resp, _ := apollo.Req.PostJson(url, data, requests.Auth{apollo.Username, apollo.Passwd})
	if resp.R.StatusCode != 200 {
		return false, fmt.Errorf(resp.Text())
	}
	return true, nil
}

func (apollo *apolloApi) CreateCluster(appId, env, clusterName string) (bool, error) {
	return true, nil
}

func (apollo *apolloApi) GetConfigData(appId, env, clusters, namespaces string) (data map[string]string, err error) {
	return
}

func (apollo *apolloApi) DelConfigKey(app, env, key, namespaceName, clustersName string) (bool, error) {
	return true, nil
}

func (apollo *apolloApi) AddAppidConfig(appid, env, key, value, comment, namespaceName, clustersName string) (bool, error) {
	return true, nil
}

func (apollo *apolloApi) UpdateConfigData(appid, env, key, value, namespaceName, clustersName string) (bool, error) {
	return true, nil
}

func (apollo *apolloApi) FindProject(appId string) (data map[string]string) {
	return
}

func (apollo *apolloApi) FindProjectInfo(appId, env, clusterName string) (data map[string]string) {
	return
}

func (apollo *apolloApi) DeleteCluster(appId, clusterName, env string) (bool, error) {
	return true, nil
}
