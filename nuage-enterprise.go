package nuagewrapper

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

// NuageEnterprise is a wrapper to create nuage enterprise in a declaritive way
func NuageEnterprise(enterpriseCfg map[string]interface{}, parent *vspk.Me) *vspk.Enterprise {
	enterprise := &vspk.Enterprise{}

	enterprises, err := parent.Enterprises(&bambou.FetchingInfo{
		Filter: enterpriseCfg["Name"].(string)})
	handleError(err, "Enterprise", "READ")

	fmt.Println("################" + enterpriseCfg["Name"].(string) + "###############")
	fmt.Println(enterprises)

	// init the enterprise struct that will hold either the received object
	// or will be created from the enterpriseCfg
	if enterprises != nil {
		fmt.Println("Enterpise already exists")

		enterprise = enterprises[0]
		errMergo := mergo.Map(enterprise, enterpriseCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		enterprise.Save()

	} else {
		errMergo := mergo.Map(enterprise, enterpriseCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateEnterprise(enterprise)
		handleError(err, "Enterprise", "CREATE")

		fmt.Println("Enterprise created")
	}
	return enterprise
}

// NuageEnterpriseprofile is a wrapper to create nuage enterprise profile in a declaritive way
func NuageEnterpriseprofile(enterpriseProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.EnterpriseProfile {
	enterpriseProfile := &vspk.EnterpriseProfile{}

	enterpriseProfiles, err := parent.EnterpriseProfiles(&bambou.FetchingInfo{
		Filter: enterpriseProfileCfg["Name"].(string)})
	handleError(err, "enterpriseProfile", "READ")

	fmt.Println("################" + enterpriseProfileCfg["Name"].(string) + "###############")
	fmt.Println(enterpriseProfiles)

	// init the enterprise struct that will hold either the received object
	// or will be created from the enterpriseProfileCfg
	if enterpriseProfiles != nil {
		fmt.Println("enterpriseProfile already exists")

		enterpriseProfile = enterpriseProfiles[0]
		errMergo := mergo.Map(enterpriseProfile, enterpriseProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		enterpriseProfile.Save()

	} else {
		errMergo := mergo.Map(enterpriseProfile, enterpriseProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateEnterpriseProfile(enterpriseProfile)
		handleError(err, "enterpriseProfile", "CREATE")

		fmt.Println("enterpriseProfile created")
	}
	return enterpriseProfile
}
