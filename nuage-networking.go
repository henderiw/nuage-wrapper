package nuagewrapper

import (
	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
	log "github.com/sirupsen/logrus"
)

// DomainTemplate is a wrapper to create nuage domain template in a declaritive way
func DomainTemplate(domainTemplateCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.DomainTemplate {
	domainTemplate := &vspk.DomainTemplate{}

	domainTemplates, err := parent.DomainTemplates(&bambou.FetchingInfo{
		Filter: domainTemplateCfg["Name"].(string)})
	handleError(err, "domainTemplate", "READ")

	log.Infof("################" + domainTemplateCfg["Name"].(string) + "###############")
	log.Infof(domainTemplates)

	// init the enterprise struct that will hold either the received object
	// or will be created from the domainTemplateCfg
	if domainTemplates != nil {
		log.Infof("domainTemplate already exists")

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

		log.Infof("domainTemplate created")
	}
	return domainTemplate
}

// Domain is a wrapper to create nuage domain in a declaritive way
func Domain(domainCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.Domain {
	domains, err := parent.Domains(&bambou.FetchingInfo{
		Filter: domainCfg["Name"].(string)})
	handleError(err, "READ", "Domain")

	domain := &vspk.Domain{}

	if domains != nil {
		log.Infof("DOmain already exists")

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

		log.Infof("Domain created")
	}

	log.Infof("%#v \n", domain)
	return domain
}

// Zone is a wrapper to create nuage zone in a declaritive way
func Zone(zoneCfg map[string]interface{}, parent *vspk.Domain) *vspk.Zone {
	zones, err := parent.Zones(&bambou.FetchingInfo{
		Filter: zoneCfg["Name"].(string)})
	handleError(err, "READ", "Zone")

	zone := &vspk.Zone{}

	if zones != nil {
		log.Infof("Zone already exists")

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

		log.Infof("Zone created")
	}

	log.Infof("%#v \n", zone)
	return zone
}

// Subnet is a wrapper to create nuage subnet in a declaritive way
func Subnet(subnetCfg map[string]interface{}, parent *vspk.Zone) *vspk.Subnet {
	subnets, err := parent.Subnets(&bambou.FetchingInfo{
		Filter: subnetCfg["Name"].(string)})
	handleError(err, "READ", "Subnet")

	subnet := &vspk.Subnet{}

	if subnets != nil {
		log.Infof("Subnet already exists")

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

		log.Infof("Subnet created")
	}

	log.Infof("%#v \n", subnet)
	return subnet
}

// BGPNeighbor is a wrapper to create nuage subnet bgp neighbor in a declaritive way
func BGPNeighbor(bgpNeighborCfg map[string]interface{}, parent *vspk.Subnet) *vspk.BGPNeighbor {
	bgpNeighbors, err := parent.BGPNeighbors(&bambou.FetchingInfo{
		Filter: bgpNeighborCfg["Name"].(string)})
	handleError(err, "READ", "Subnet")

	bgpNeighbor := &vspk.BGPNeighbor{}

	if bgpNeighbors != nil {
		log.Infof("bgpNeighbor already exists")

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

		log.Infof("bgpNeighbor created")
	}

	log.Infof("%#v \n", bgpNeighbor)
	return bgpNeighbor
}

// StaticRoute is a wrapper to create nuage domain static route in a declaritive way
func StaticRoute(staticRouteCfg map[string]interface{}, parent *vspk.Domain) *vspk.StaticRoute {
	staticRoutes, err := parent.StaticRoutes(&bambou.FetchingInfo{
		Filter: staticRouteCfg["Address"].(string)})
	handleError(err, "READ", "staticRoute")

	staticRoute := &vspk.StaticRoute{}

	if staticRoutes != nil {
		log.Infof("staticRoute already exists")

		staticRoute = staticRoutes[0]
		log.Infof("Static Route: %#v", staticRoute)
		errMergo := mergo.Map(staticRoute, staticRouteCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		log.Infof("Static Route: %#v", staticRoute)
		staticRoute.Save()
	} else {
		errMergo := mergo.Map(staticRoute, staticRouteCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		log.Infof("Static Route: %#v", staticRoute)

		err := parent.CreateStaticRoute(staticRoute)
		handleError(err, "CREATE", "staticRoute")

		log.Infof("staticRoute created")
	}

	log.Infof("%#v \n", staticRoute)
	return staticRoute
}

// IngressACLTemplate is a wrapper to create nuage domain ingress ACL Template in a declaritive way
func IngressACLTemplate(ingressACLTemplateCfg map[string]interface{}, parent *vspk.Domain) *vspk.IngressACLTemplate {
	ingressACLTemplates, err := parent.IngressACLTemplates(&bambou.FetchingInfo{
		Filter: ingressACLTemplateCfg["Name"].(string)})
	handleError(err, "READ", "ingressACLTemplate")

	ingressACLTemplate := &vspk.IngressACLTemplate{}

	if ingressACLTemplates != nil {
		log.Infof("ingressACLTemplate already exists")

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

		log.Infof("ingressACLTemplate created")
	}

	log.Infof("%#v \n", ingressACLTemplate)
	return ingressACLTemplate
}

// EgressACLTemplate is a wrapper to create nuage domain egress ACL Template in a declaritive way
func EgressACLTemplate(egressACLTemplateCfg map[string]interface{}, parent *vspk.Domain) *vspk.EgressACLTemplate {
	egressACLTemplates, err := parent.EgressACLTemplates(&bambou.FetchingInfo{
		Filter: egressACLTemplateCfg["Name"].(string)})
	handleError(err, "READ", "egressACLTemplate")

	egressACLTemplate := &vspk.EgressACLTemplate{}

	if egressACLTemplates != nil {
		log.Infof("egressACLTemplate already exists")

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

		log.Infof("egressACLTemplate created")
	}

	log.Infof("%#v \n", egressACLTemplate)
	return egressACLTemplate
}
