package nuagewrapper

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func handleError(err *bambou.Error, t string, o string) {
	if err != nil {
		fmt.Println("Unable to " + o + " \"" + t + "\": " + err.Description)
		fmt.Printf("Error: %#v \n", err)
		os.Exit(1)
	}
}

// NuageLicense is a wrapper to create nuage license in a declaritive way
func NuageLicense(l string, parent *vspk.Me) {
	license := &vspk.License{}
	license.License = l

	err := parent.CreateLicense(license)
	handleError(err, "License", "CREATE")

	fmt.Println("License created")
}

// NuageUser is a wrapper to create nuage user in a declaritive way
func NuageUser(userCfg map[string]interface{}, parent *vspk.Me) *vspk.User {
	user := &vspk.User{}

	users, err := parent.Users(&bambou.FetchingInfo{
		Filter: userCfg["UserName"].(string)})
	handleError(err, "User", "READ")

	fmt.Println("################" + userCfg["UserName"].(string) + "###############")
	fmt.Println(users)

	// init the user struct that will hold either the received object
	// or will be created from the userCfg
	if users != nil {
		fmt.Println("User already exists")

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

		fmt.Println("user created")
	}
	return user
}

// NuageAssignUser is a wrapper to assign nuage user to a group
func NuageAssignUser(user *vspk.User, groupCfg map[string]interface{}, parent *vspk.Me) string {
	enterpriseCfg := map[string]interface{}{
		"Name": "Shared Infrastructure",
	}

	enterprises, err := parent.Enterprises(&bambou.FetchingInfo{
		Filter: enterpriseCfg["Name"].(string)})
	handleError(err, "User", "READ")

	fmt.Println("Number of enterprises retrieved: " + strconv.Itoa(len(enterprises)))

	for _, enterprise := range enterprises {
		fmt.Println(enterprise.Name)
	}

	groups, err := enterprises[0].Groups(&bambou.FetchingInfo{
		Filter: groupCfg["Name"].(string)})
	handleError(err, "Group", "READ")

	for _, group := range groups {
		fmt.Printf("Group: %#v \n", group)
	}

	return ""
}
