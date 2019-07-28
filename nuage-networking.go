package nuagewrapper

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

// NuageDomain is a wrapper to create nuage idomain in a declaritive way
func NuageDomain(domainCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.Domain {
	fmt.Println("########################################")
	fmt.Println("#####            Domain       ##########")
	fmt.Println("########################################")

	domains, err := parent.Domains(&bambou.FetchingInfo{
		Filter: domainCfg["Name"].(string)})
	handleError(err, "READ", "Domain")

	domain := &vspk.Domain{}

	if domains != nil {
		fmt.Println("DOmain already exists")

		domain = domains[0]
		errMergo := mergo.Map(domain, domainCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		domain.Save()
	} else {
		errMergo := mergo.Map(domain, domainCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateDomain(domain)
		handleError(err, "CREATE", "Domain")

		fmt.Println("Domain created")
	}

	fmt.Printf("%#v \n", domain)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return domain
}
