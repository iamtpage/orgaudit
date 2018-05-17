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
	"strings"

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
	whitelist    = kingpin.Flag("whitelist", "comma-separated whitelist of orgs to include").Short('w').Required().String()
	blacklist    = kingpin.Flag("blacklist", "comma-separated blacklist of orgs to ignore").Short('b').Required().String()
)

func main() {
	kingpin.Parse()

	//parse whitelist/blacklist
	wslice := strings.Split(*whitelist, ",")
	bslice := strings.Split(*blacklist, ",")

	//convert to map for more efficient checking with large bl/wl
	bmap := make(map[string]bool)
	wmap := make(map[string]bool)

	for _, s := range wslice {
		wmap[s] = true
	}

	for _, s := range bslice {
		bmap[s] = true
	}

	token, err := getToken("https://uaa."+*sysDomain, *userID, *clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	cc := cloudcontroller.NewManager("https://api."+*sysDomain, token, false)

	orgs, err := cc.ListOrgs()
	if err != nil {
		log.Fatalf("cannot list orgs: %v", err)
	}

	var orgGUID string
	var orgNAME string
	for _, o := range orgs {

		if bmap[o.Entity.Name] {
			//Org is in BL, skip it
			continue
		}

		if wmap[o.Entity.Name] {
			//Org is in WL, do the thing
			orgGUID = o.MetaData.GUID
			orgNAME = o.Entity.Name

			spaces, err := cc.ListSpaces(orgGUID)

			if err != nil {
				log.Fatalf("cannot list spaces for org %s: %v", orgNAME, err)
			}

			users, err := cc.GetCFUsers(orgGUID, "organizations", "auditors")

			if err != nil {
				log.Fatalf("cannot find org auditors for %s: %v", orgNAME, err)
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
	}
}

func getToken(uaaURL, user, secret string) (string, error) {

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
