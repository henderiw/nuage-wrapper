package nuagewrapper

import (
	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
	log "github.com/sirupsen/logrus"
)

// NuageEnterprise is a wrapper to create nuage enterprise in a declaritive way
func NuageEnterprise(enterpriseCfg map[string]interface{}, parent *vspk.Me) *vspk.Enterprise {
	enterprise := &vspk.Enterprise{}

	enterprises, err := parent.Enterprises(&bambou.FetchingInfo{
		Filter: enterpriseCfg["Name"].(string)})
	handleError(err, "Enterprise", "READ")

	log.Infof("################" + enterpriseCfg["Name"].(string) + "###############")
	log.Infof(enterprises)

	// init the enterprise struct that will hold either the received object
	// or will be created from the enterpriseCfg
	if enterprises != nil {
		log.Infof("Enterpise already exists")

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

		log.Infof("Enterprise created")
	}
	return enterprise
}

// NuageEnterpriseprofile is a wrapper to create nuage enterprise profile in a declaritive way
func NuageEnterpriseprofile(enterpriseProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.EnterpriseProfile {
	enterpriseProfile := &vspk.EnterpriseProfile{}

	enterpriseProfiles, err := parent.EnterpriseProfiles(&bambou.FetchingInfo{
		Filter: enterpriseProfileCfg["Name"].(string)})
	handleError(err, "enterpriseProfile", "READ")

	log.Infof("################" + enterpriseProfileCfg["Name"].(string) + "###############")
	log.Infof(enterpriseProfiles)

	// init the enterprise struct that will hold either the received object
	// or will be created from the enterpriseProfileCfg
	if enterpriseProfiles != nil {
		log.Infof("enterpriseProfile already exists")

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

		log.Infof("enterpriseProfile created")
	}
	return enterpriseProfile
}
