package nuagewrapper

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

// NuageDomainTemplate is a wrapper to create nuage domain template in a declaritive way
func NuageDomainTemplate(domainTemplateCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.DomainTemplate {
	domainTemplate := &vspk.DomainTemplate{}

	domainTemplates, err := parent.DomainTemplates(&bambou.FetchingInfo{
		Filter: domainTemplateCfg["Name"].(string)})
	handleError(err, "domainTemplate", "READ")

	fmt.Println("################" + domainTemplateCfg["Name"].(string) + "###############")
	fmt.Println(domainTemplates)

	// init the enterprise struct that will hold either the received object
	// or will be created from the domainTemplateCfg
	if domainTemplates != nil {
		fmt.Println("domainTemplate already exists")

		domainTemplate = domainTemplates[0]
		errMergo := mergo.Map(domainTemplate, domainTemplateCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		domainTemplate.Save()

	} else {
		errMergo := mergo.Map(domainTemplate, domainTemplateCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateDomainTemplate(domainTemplate)
		handleError(err, "domainTemplate", "CREATE")

		fmt.Println("domainTemplate created")
	}
	return domainTemplate
}

// NuageDomain is a wrapper to create nuage domain in a declaritive way
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

// NuageZone is a wrapper to create nuage zone in a declaritive way
func NuageZone(zoneCfg map[string]interface{}, parent *vspk.Domain) *vspk.Zone {
	fmt.Println("########################################")
	fmt.Println("#####            Zone        ##########")
	fmt.Println("########################################")

	zones, err := parent.Zones(&bambou.FetchingInfo{
		Filter: zoneCfg["Name"].(string)})
	handleError(err, "READ", "Zone")

	zone := &vspk.Zone{}

	if zones != nil {
		fmt.Println("Zone already exists")

		zone = zones[0]
		errMergo := mergo.Map(zone, zoneCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		zone.Save()
	} else {
		errMergo := mergo.Map(zone, zoneCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateZone(zone)
		handleError(err, "CREATE", "Zone")

		fmt.Println("Zone created")
	}

	fmt.Printf("%#v \n", zone)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return zone
}

// NuageSubnet is a wrapper to create nuage subnet in a declaritive way
func NuageSubnet(subnetCfg map[string]interface{}, parent *vspk.Zone) *vspk.Subnet {
	fmt.Println("########################################")
	fmt.Println("#####            Subnet       ##########")
	fmt.Println("########################################")

	subnets, err := parent.Subnets(&bambou.FetchingInfo{
		Filter: subnetCfg["Name"].(string)})
	handleError(err, "READ", "Subnet")

	subnet := &vspk.Subnet{}

	if subnets != nil {
		fmt.Println("Subnet already exists")

		subnet = subnets[0]
		errMergo := mergo.Map(subnet, subnetCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		subnet.Save()
	} else {
		errMergo := mergo.Map(subnet, subnetCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateSubnet(subnet)
		handleError(err, "CREATE", "Subnet")

		fmt.Println("Subnet created")
	}

	fmt.Printf("%#v \n", subnet)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return subnet
}

// NuageBGPNeighbor is a wrapper to create nuage subnet bgp neighbor in a declaritive way
func NuageBGPNeighbor(bgpNeighborCfg map[string]interface{}, parent *vspk.Subnet) *vspk.BGPNeighbor {
	fmt.Println("########################################")
	fmt.Println("#####      BGP Neighbor       ##########")
	fmt.Println("########################################")

	bgpNeighbors, err := parent.BGPNeighbors(&bambou.FetchingInfo{
		Filter: bgpNeighborCfg["Name"].(string)})
	handleError(err, "READ", "Subnet")

	bgpNeighbor := &vspk.BGPNeighbor{}

	if bgpNeighbors != nil {
		fmt.Println("bgpNeighbor already exists")

		bgpNeighbor = bgpNeighbors[0]
		errMergo := mergo.Map(bgpNeighbor, bgpNeighborCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		bgpNeighbor.Save()
	} else {
		errMergo := mergo.Map(bgpNeighbor, bgpNeighborCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateBGPNeighbor(bgpNeighbor)
		handleError(err, "CREATE", "bgpNeighbor")

		fmt.Println("bgpNeighbor created")
	}

	fmt.Printf("%#v \n", bgpNeighbor)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return bgpNeighbor
}

// NuageStaticRoute is a wrapper to create nuage domain static route in a declaritive way
func NuageStaticRoute(staticRouteCfg map[string]interface{}, parent *vspk.Domain) *vspk.StaticRoute {
	fmt.Println("########################################")
	fmt.Println("#####      static Route       ##########")
	fmt.Println("########################################")

	staticRoutes, err := parent.StaticRoutes(&bambou.FetchingInfo{
		Filter: staticRouteCfg["Address"].(string)})
	handleError(err, "READ", "staticRoute")

	staticRoute := &vspk.StaticRoute{}

	if staticRoutes != nil {
		fmt.Println("staticRoute already exists")

		staticRoute = staticRoutes[0]
		fmt.Printf("Static Route: %#v", staticRoute)
		errMergo := mergo.Map(staticRoute, staticRouteCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		fmt.Printf("Static Route: %#v", staticRoute)
		staticRoute.Save()
	} else {
		errMergo := mergo.Map(staticRoute, staticRouteCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		fmt.Printf("Static Route: %#v", staticRoute)

		err := parent.CreateStaticRoute(staticRoute)
		handleError(err, "CREATE", "staticRoute")

		fmt.Println("staticRoute created")
	}

	fmt.Printf("%#v \n", staticRoute)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return staticRoute
}

// NuageIngressACLTemplate is a wrapper to create nuage domain ingress ACL Template in a declaritive way
func NuageIngressACLTemplate(ingressACLTemplateCfg map[string]interface{}, parent *vspk.Domain) *vspk.IngressACLTemplate {
	fmt.Println("########################################")
	fmt.Println("#####      ingressACLTemplate ##########")
	fmt.Println("########################################")

	ingressACLTemplates, err := parent.IngressACLTemplates(&bambou.FetchingInfo{
		Filter: ingressACLTemplateCfg["Name"].(string)})
	handleError(err, "READ", "ingressACLTemplate")

	ingressACLTemplate := &vspk.IngressACLTemplate{}

	if ingressACLTemplates != nil {
		fmt.Println("ingressACLTemplate already exists")

		ingressACLTemplate = ingressACLTemplates[0]
		errMergo := mergo.Map(ingressACLTemplate, ingressACLTemplateCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		ingressACLTemplate.Save()
	} else {
		errMergo := mergo.Map(ingressACLTemplate, ingressACLTemplateCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateIngressACLTemplate(ingressACLTemplate)
		handleError(err, "CREATE", "ingressACLTemplate")

		fmt.Println("ingressACLTemplate created")
	}

	fmt.Printf("%#v \n", ingressACLTemplate)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return ingressACLTemplate
}

// NuageEgressACLTemplate is a wrapper to create nuage domain egress ACL Template in a declaritive way
func NuageEgressACLTemplate(egressACLTemplateCfg map[string]interface{}, parent *vspk.Domain) *vspk.EgressACLTemplate {
	fmt.Println("########################################")
	fmt.Println("#####       egressACLTemplate ##########")
	fmt.Println("########################################")

	egressACLTemplates, err := parent.EgressACLTemplates(&bambou.FetchingInfo{
		Filter: egressACLTemplateCfg["Name"].(string)})
	handleError(err, "READ", "egressACLTemplate")

	egressACLTemplate := &vspk.EgressACLTemplate{}

	if egressACLTemplates != nil {
		fmt.Println("egressACLTemplate already exists")

		egressACLTemplate = egressACLTemplates[0]
		errMergo := mergo.Map(egressACLTemplate, egressACLTemplateCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		egressACLTemplate.Save()
	} else {
		errMergo := mergo.Map(egressACLTemplate, egressACLTemplateCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateEgressACLTemplate(egressACLTemplate)
		handleError(err, "CREATE", "egressACLTemplate")

		fmt.Println("egressACLTemplate created")
	}

	fmt.Printf("%#v \n", egressACLTemplate)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return egressACLTemplate
}
