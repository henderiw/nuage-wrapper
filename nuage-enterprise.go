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
