package nuagewrapper

import (
	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
	log "github.com/sirupsen/logrus"
)

// InfraGwProfile is a wrapper to create nuage infrastructure GW Profile in a declaritive way
func InfraGwProfile(infraGwProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.InfrastructureGatewayProfile {
	infraGwProfile := &vspk.InfrastructureGatewayProfile{}

	infraGwProfiles, err := parent.InfrastructureGatewayProfiles(&bambou.FetchingInfo{
		Filter: infraGwProfileCfg["Name"].(string)})
	handleError(err, "Infra GW Profile", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if infraGwProfiles != nil {
		log.Infof("Infra Profile already exists")

		infraGwProfile = infraGwProfiles[0]
		errMergo := mergo.Map(infraGwProfile, infraGwProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		infraGwProfile.Save()

	} else {

		errMergo := mergo.Map(infraGwProfile, infraGwProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateInfrastructureGatewayProfile(infraGwProfile)
		handleError(err, "Infra GW Profile", "CREATE")

		log.Infof("Infra GW Profile created")
	}
	log.Infof("%#v \n", infraGwProfile)
	return infraGwProfile
}

// NsgUpgradeProfile is a wrapper to create nuage NSG Upgrade Profile in a declaritive way
func NsgUpgradeProfile(upgradeProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.NSGUpgradeProfile {
	upgradeProfile := &vspk.NSGUpgradeProfile{}

	upgradeProfiles, err := parent.NSGUpgradeProfiles(&bambou.FetchingInfo{
		Filter: upgradeProfileCfg["Name"].(string)})
	handleError(err, "Upgrade Profile", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if upgradeProfiles != nil {
		log.Infof("Upgrade Profile already exists")

		upgradeProfile = upgradeProfiles[0]
		errMergo := mergo.Map(upgradeProfile, upgradeProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		upgradeProfile.Save()

	} else {

		errMergo := mergo.Map(upgradeProfile, upgradeProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateNSGUpgradeProfile(upgradeProfile)
		handleError(err, "Upgrade Profile", "CREATE")

		log.Infof("Upgrade Profile created")
	}
	log.Infof("%#v \n", upgradeProfile)
	return upgradeProfile
}

// InfraAccessProfile is a wrapper to create nuage infrastructure Access Profile in a declaritive way
func InfraAccessProfile(infraAccessProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.InfrastructureAccessProfile {
	infraAccessProfile := &vspk.InfrastructureAccessProfile{}

	infraAccessProfiles, err := parent.InfrastructureAccessProfiles(&bambou.FetchingInfo{
		Filter: infraAccessProfileCfg["Name"].(string)})
	handleError(err, "Infra Access Profile", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if infraAccessProfiles != nil {
		log.Infof("Infra Access Profile already exists")

		infraAccessProfile = infraAccessProfiles[0]
		errMergo := mergo.Map(infraAccessProfile, infraAccessProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		infraAccessProfile.Save()

	} else {

		errMergo := mergo.Map(infraAccessProfile, infraAccessProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateInfrastructureAccessProfile(infraAccessProfile)
		handleError(err, "Infra Access Profile", "CREATE")

		log.Infof("Infra Access Profile created")
	}
	log.Infof("%#v \n", infraAccessProfile)
	return infraAccessProfile
}

// InfraVscProfile is a wrapper to create nuage infrastructure VSC Profile in a declaritive way
func InfraVscProfile(infraVscProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.InfrastructureVscProfile {
	infraVscProfile := &vspk.InfrastructureVscProfile{}

	infraVscProfiles, err := parent.InfrastructureVscProfiles(&bambou.FetchingInfo{
		Filter: infraVscProfileCfg["Name"].(string)})
	handleError(err, "Infra `VSC` Profile", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if infraVscProfiles != nil {
		log.Infof("Infra VSC Profile already exists")

		infraVscProfile = infraVscProfiles[0]
		errMergo := mergo.Map(infraVscProfile, infraVscProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		infraVscProfile.Save()

	} else {

		errMergo := mergo.Map(infraVscProfile, infraVscProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateInfrastructureVscProfile(infraVscProfile)
		handleError(err, "Infra VSC Profile", "CREATE")

		log.Infof("Infra VSC Profile created")
	}
	log.Infof("%#v \n", infraVscProfile)
	return infraVscProfile
}

// DucGroup is a wrapper to create nuage DUC Group in a declaritive way
func DucGroup(ducGroupCfg map[string]interface{}, parent *vspk.Me) *vspk.DUCGroup {
	ducGroup := &vspk.DUCGroup{}

	ducGroups, err := parent.DUCGroups(&bambou.FetchingInfo{
		Filter: ducGroupCfg["Name"].(string)})
	handleError(err, "DUC group", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if ducGroups != nil {
		log.Infof("DUC group already exists")

		ducGroup = ducGroups[0]
		errMergo := mergo.Map(ducGroup, ducGroupCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		ducGroup.Save()

	} else {

		errMergo := mergo.Map(ducGroup, ducGroupCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateDUCGroup(ducGroup)
		handleError(err, "DUC Group", "CREATE")

		log.Infof("DUC Group created")
	}
	log.Infof("%#v \n", ducGroup)
	return ducGroup
}

// PerfMonitor is a wrapper to create nuage Performance Monitor in a declaritive way
func PerfMonitor(perfMonitorCfg map[string]interface{}, parent *vspk.Me) *vspk.PerformanceMonitor {
	perfMonitor := &vspk.PerformanceMonitor{}

	perfMonitors, err := parent.PerformanceMonitors(&bambou.FetchingInfo{
		Filter: perfMonitorCfg["Name"].(string)})
	handleError(err, "Performance Monitor", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if perfMonitors != nil {
		log.Infof("Performance Monitor already exists")

		perfMonitor = perfMonitors[0]
		errMergo := mergo.Map(perfMonitor, perfMonitorCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		perfMonitor.Save()

	} else {

		errMergo := mergo.Map(perfMonitor, perfMonitorCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreatePerformanceMonitor(perfMonitor)
		handleError(err, "Performance Monitor", "CREATE")

		log.Infof("Performance Monitor created")
	}
	log.Infof("%#v \n", perfMonitor)
	return perfMonitor
}

// Underlay is a wrapper to create nuage underlay in a declaritive way
func Underlay(underlayCfg map[string]interface{}, parent *vspk.Me) *vspk.Underlay {
	underlay := &vspk.Underlay{}

	underlays, err := parent.Underlays(&bambou.FetchingInfo{
		Filter: underlayCfg["Name"].(string)})
	handleError(err, "Underlay", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if underlays != nil {
		log.Infof("Underlay already exists")

		underlay = underlays[0]
		errMergo := mergo.Map(underlay, underlayCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		underlay.Save()

	} else {

		errMergo := mergo.Map(underlay, underlayCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateUnderlay(underlay)
		handleError(err, "Underlay", "CREATE")

		log.Infof("Underlay created")
	}
	log.Infof("%#v \n", underlay)
	return underlay
}

// InfraNsgGroup is a wrapper to create nuage NSG Group in a declaritive way
func InfraNsgGroup(nsgGroupCfg map[string]interface{}, parent *vspk.Me) *vspk.NSGGroup {
	nsgGroup := &vspk.NSGGroup{}

	nsgGroups, err := parent.NSGGroups(&bambou.FetchingInfo{
		Filter: nsgGroupCfg["Name"].(string)})
	handleError(err, "NSG group", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if nsgGroups != nil {
		log.Infof("NSG group already exists")

		nsgGroup = nsgGroups[0]
		errMergo := mergo.Map(nsgGroup, nsgGroupCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		nsgGroup.Save()

	} else {

		errMergo := mergo.Map(nsgGroup, nsgGroupCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateNSGGroup(nsgGroup)
		handleError(err, "NSG Group", "CREATE")

		log.Infof("NSG Group created")
	}
	log.Infof("%#v \n", nsgGroup)
	return nsgGroup
}

// NsgGroup is a wrapper to create nuage NSG Group in a declaritive way
func NsgGroup(nsgGroupCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.NSGGroup {
	nsgGroup := &vspk.NSGGroup{}

	nsgGroups, err := parent.NSGGroups(&bambou.FetchingInfo{
		Filter: nsgGroupCfg["Name"].(string)})
	handleError(err, "NSG group", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if nsgGroups != nil {
		log.Infof("NSG group already exists")

		nsgGroup = nsgGroups[0]
		errMergo := mergo.Map(nsgGroup, nsgGroupCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		nsgGroup.Save()

	} else {

		errMergo := mergo.Map(nsgGroup, nsgGroupCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateNSGGroup(nsgGroup)
		handleError(err, "NSG Group", "CREATE")

		log.Infof("NSG Group created")
	}
	log.Infof("%#v \n", nsgGroup)
	return nsgGroup
}
