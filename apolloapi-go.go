package apolloapi

import (
	"fmt"
	"strings"
	"time"

	"github.com/asmcos/requests"
)

type Apolloapi struct {
	Username  string
	Passwd    string
	ApolloURL string
	Req       *requests.Request
}

func NewApollo(apollourl, username, passwd string) *Apolloapi {
	return &Apolloapi{
		Username:  username,
		Passwd:    passwd,
		ApolloURL: apollourl,
		Req:       requests.Requests(),
	}
}

func (apollo *Apolloapi) GetappIDList() {
	url := fmt.Sprintf("%s/apps/by-owner?owner=apollo", apollo.ApolloURL)
	resp, _ := apollo.Req.Get(url, requests.Auth{apollo.Username, apollo.Passwd})
	fmt.Println(resp.Text())
	fmt.Println(resp.R.StatusCode)
	fmt.Println(resp.R.Header["Content-Type"])
}

func (apollo *Apolloapi) CreateApolloProject(appId, name, orgId, orgName, ownerName string) (bool, error) {
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

func (apollo *Apolloapi) ReleaseApollo(appid, env, comment, clustersName, namespaceName string) (bool, error) {
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
