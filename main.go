// command orgaudit assigns the SpaceAuditor role to all
// OrgAuditors in a given Cloud Foundry organization
package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/parnurzeal/gorequest"
	"github.com/pivotalservices/cf-mgmt/cloudcontroller"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	orgAuditor   = "OrgAuditor"
	spaceAuditor = "SpaceAuditor"
)

var (
	sysDomain    = kingpin.Flag("system-domain", "CF system domain").Short('d').Required().String()
	userID       = kingpin.Flag("user-id", "user ID").Envar("USER_ID").Required().String()
	clientSecret = kingpin.Flag("client-secret", "client secret").Envar("CLIENT_SECRET").Required().String()
	org          = kingpin.Flag("org", "the name of the org").Short('o').Required().String()
	debug        = kingpin.Flag("debug", "verbose logging").Envar("ORGAUDIT_DEBUG").Default("false").Bool()
)

func main() {
	kingpin.Parse()

	token, err := getToken("https://uaa."+*sysDomain, *userID, *clientSecret, *debug)
	if err != nil {
		log.Fatal(err)
	}

	cc := cloudcontroller.NewManager("https://api."+*sysDomain, token)

	orgs, err := cc.ListOrgs()
	if err != nil {
		log.Fatalf("cannot list orgs: %v", err)
	}

	var orgGUID string
	for _, o := range orgs {
		if o.Entity.Name == *org {
			orgGUID = o.MetaData.GUID
			break
		}
	}

	if orgGUID == "" {
		log.Fatal("cannot find org", *org)
	}

	spaces, err := cc.ListSpaces(orgGUID)
	if err != nil {
		log.Fatalf("cannot list spaces for org %s: %v", *org, err)
	}

	users, err := cc.GetCFUsers(orgGUID, "organizations", "auditors")
	if err != nil {
		log.Fatalf("cannot find org auditors for %s: %v", *org, err)
	}

	for user := range users {
		for _, space := range spaces {
			log.Println("assigning SpaceAuditor to", user, "for space", space.Entity.Name)
			err = cc.AddUserToSpaceRole(user, "auditors", space.MetaData.GUID)
			if err != nil {
				log.Fatalf("error assigning SpaceAuditor to user %s: %v", user, err)
			}
		}
	}
}

func getToken(uaaURL, user, secret string, debug) (string, error) {
	
	if debug != false {
		log.Println("getting token", uaaURL, user, secret)
	}
	
	request := gorequest.New()
	request.TargetType = "form"

	post := request.Post(fmt.Sprintf("%s/oauth/token", uaaURL))
	post.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	post.BasicAuth.Username = user
	post.BasicAuth.Password = secret

	params := url.Values{
		"grant_type":    {"client_credentials"},
		"response_type": {"token"},
	}
	post.Send(params.Encode())

	cmd, _ := post.AsCurlCommand()
	
	if debug != false {
		log.Println("post:", cmd)
	}

	res, body, errs := post.End()
	if len(errs) > 0 {
		return "", errs[0]
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("cannot get UAAC token (%v): %s", res.StatusCode, body)
	}
	t := struct {
		AccessToken string `json:"access_token"`
	}{}
	if err := json.Unmarshal([]byte(body), &t); err != nil {
		return "", fmt.Errorf("cannot read token: %v", err)
	}
	return t.AccessToken, nil
}
