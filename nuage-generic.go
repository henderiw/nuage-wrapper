package nuagewrapper

import (
	"os"
	"strconv"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
	log "github.com/sirupsen/logrus"
)

func handleError(err *bambou.Error, t string, o string) {
	if err != nil {
		log.Errorf("Unable to " + o + " \"" + t + "\": " + err.Description)
		log.Errorf("Error: %#v \n", err)
		os.Exit(1)
	}
}

// License is a wrapper to create nuage license in a declaritive way
func License(l string, parent *vspk.Me) {
	license := &vspk.License{}
	license.License = l

	err := parent.CreateLicense(license)
	handleError(err, "License", "CREATE")

	log.Infof("License created")
}

// User is a wrapper to create nuage user in a declaritive way
func User(userCfg map[string]interface{}, parent *vspk.Me) *vspk.User {
	user := &vspk.User{}

	users, err := parent.Users(&bambou.FetchingInfo{
		Filter: userCfg["UserName"].(string)})
	handleError(err, "User", "READ")

	log.Debugf("################" + userCfg["UserName"].(string) + "###############")
	log.Debugf("Users %v", users)

	// init the user struct that will hold either the received object
	// or will be created from the userCfg
	if users != nil {
		log.Infof("User already exists")

		user = users[0]
		errMergo := mergo.Map(user, userCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		user.Save()

	} else {
		errMergo := mergo.Map(user, userCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateUser(user)
		handleError(err, "User", "CREATE")

		log.Infof("user created")
	}
	return user
}

// AssignUser is a wrapper to assign nuage user to a group
func AssignUser(user *vspk.User, groupCfg map[string]interface{}, parent *vspk.Me) string {
	enterpriseCfg := map[string]interface{}{
		"Name": "Shared Infrastructure",
	}

	enterprises, err := parent.Enterprises(&bambou.FetchingInfo{
		Filter: enterpriseCfg["Name"].(string)})
	handleError(err, "User", "READ")

	log.Infof("Number of enterprises retrieved: " + strconv.Itoa(len(enterprises)))

	for _, enterprise := range enterprises {
		log.Infof(enterprise.Name)
	}

	groups, err := enterprises[0].Groups(&bambou.FetchingInfo{
		Filter: groupCfg["Name"].(string)})
	handleError(err, "Group", "READ")

	for _, group := range groups {
		log.Infof("Group: %#v \n", group)
	}

	return ""
}
