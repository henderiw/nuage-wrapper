package nuagewrapper

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

// NuageInfraGwProfile is a wrapper to create nuage infrastructure GW Profile in a declaritive way
func NuageInfraGwProfile(infraGwProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.InfrastructureGatewayProfile {
	fmt.Println("########################################")
	fmt.Println("#####     Infra GW Profile #############")
	fmt.Println("########################################")

	infraGwProfile := &vspk.InfrastructureGatewayProfile{}

	infraGwProfiles, err := parent.InfrastructureGatewayProfiles(&bambou.FetchingInfo{
		Filter: infraGwProfileCfg["Name"].(string)})
	handleError(err, "Infra GW Profile", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if infraGwProfiles != nil {
		fmt.Println("Infra Profile already exists")

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

		fmt.Println("Infra GW Profile created")
	}
	fmt.Printf("%#v \n", infraGwProfile)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return infraGwProfile
}

// NuageNsgUpgradeProfile is a wrapper to create nuage NSG Upgrade Profile in a declaritive way
func NuageNsgUpgradeProfile(upgradeProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.NSGUpgradeProfile {
	fmt.Println("########################################")
	fmt.Println("#####     Upgrade Profile  #############")
	fmt.Println("########################################")

	upgradeProfile := &vspk.NSGUpgradeProfile{}

	upgradeProfiles, err := parent.NSGUpgradeProfiles(&bambou.FetchingInfo{
		Filter: upgradeProfileCfg["Name"].(string)})
	handleError(err, "Upgrade Profile", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if upgradeProfiles != nil {
		fmt.Println("Upgrade Profile already exists")

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

		fmt.Println("Upgrade Profile created")
	}
	fmt.Printf("%#v \n", upgradeProfile)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return upgradeProfile
}

// NuageInfraAccessProfile is a wrapper to create nuage infrastructure Access Profile in a declaritive way
func NuageInfraAccessProfile(infraAccessProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.InfrastructureAccessProfile {
	fmt.Println("########################################")
	fmt.Println("#####     Infra Access Profile #########")
	fmt.Println("########################################")

	infraAccessProfile := &vspk.InfrastructureAccessProfile{}

	infraAccessProfiles, err := parent.InfrastructureAccessProfiles(&bambou.FetchingInfo{
		Filter: infraAccessProfileCfg["Name"].(string)})
	handleError(err, "Infra Access Profile", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if infraAccessProfiles != nil {
		fmt.Println("Infra Access Profile already exists")

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

		fmt.Println("Infra Access Profile created")
	}
	fmt.Printf("%#v \n", infraAccessProfile)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return infraAccessProfile
}

// NuageInfraVscProfile is a wrapper to create nuage infrastructure VSC Profile in a declaritive way
func NuageInfraVscProfile(infraVscProfileCfg map[string]interface{}, parent *vspk.Me) *vspk.InfrastructureVscProfile {
	fmt.Println("########################################")
	fmt.Println("#####     Infra VSC  Profile   #########")
	fmt.Println("########################################")

	infraVscProfile := &vspk.InfrastructureVscProfile{}

	infraVscProfiles, err := parent.InfrastructureVscProfiles(&bambou.FetchingInfo{
		Filter: infraVscProfileCfg["Name"].(string)})
	handleError(err, "Infra `VSC` Profile", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if infraVscProfiles != nil {
		fmt.Println("Infra VSC Profile already exists")

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

		fmt.Println("Infra VSC Profile created")
	}
	fmt.Printf("%#v \n", infraVscProfile)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return infraVscProfile
}

// NuageDucGroup is a wrapper to create nuage DUC Group in a declaritive way
func NuageDucGroup(ducGroupCfg map[string]interface{}, parent *vspk.Me) *vspk.DUCGroup {
	fmt.Println("########################################")
	fmt.Println("#####     DUC Group            #########")
	fmt.Println("########################################")

	ducGroup := &vspk.DUCGroup{}

	ducGroups, err := parent.DUCGroups(&bambou.FetchingInfo{
		Filter: ducGroupCfg["Name"].(string)})
	handleError(err, "DUC group", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if ducGroups != nil {
		fmt.Println("DUC group already exists")

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

		fmt.Println("DUC Group created")
	}
	fmt.Printf("%#v \n", ducGroup)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return ducGroup
}

// NuagePerfMonitor is a wrapper to create nuage Performance Monitor in a declaritive way
func NuagePerfMonitor(perfMonitorCfg map[string]interface{}, parent *vspk.Me) *vspk.PerformanceMonitor {
	fmt.Println("########################################")
	fmt.Println("#####     Performance Monitor  #########")
	fmt.Println("########################################")

	perfMonitor := &vspk.PerformanceMonitor{}

	perfMonitors, err := parent.PerformanceMonitors(&bambou.FetchingInfo{
		Filter: perfMonitorCfg["Name"].(string)})
	handleError(err, "Performance Monitor", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if perfMonitors != nil {
		fmt.Println("Performance Monitor already exists")

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

		fmt.Println("Performance Monitor created")
	}
	fmt.Printf("%#v \n", perfMonitor)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return perfMonitor
}

// NuageUnderlay is a wrapper to create nuage underlay in a declaritive way
func NuageUnderlay(underlayCfg map[string]interface{}, parent *vspk.Me) *vspk.Underlay {
	fmt.Println("########################################")
	fmt.Println("#####     Underlay             #########")
	fmt.Println("########################################")

	underlay := &vspk.Underlay{}

	underlays, err := parent.Underlays(&bambou.FetchingInfo{
		Filter: underlayCfg["Name"].(string)})
	handleError(err, "Underlay", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if underlays != nil {
		fmt.Println("Underlay already exists")

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

		fmt.Println("Underlay created")
	}
	fmt.Printf("%#v \n", underlay)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return underlay
}

// NuageInfraNsgGroup is a wrapper to create nuage NSG Group in a declaritive way
func NuageInfraNsgGroup(nsgGroupCfg map[string]interface{}, parent *vspk.Me) *vspk.NSGGroup {
	fmt.Println("########################################")
	fmt.Println("#####     NSG Group            #########")
	fmt.Println("########################################")

	nsgGroup := &vspk.NSGGroup{}

	nsgGroups, err := parent.NSGGroups(&bambou.FetchingInfo{
		Filter: nsgGroupCfg["Name"].(string)})
	handleError(err, "NSG group", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if nsgGroups != nil {
		fmt.Println("NSG group already exists")

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

		fmt.Println("NSG Group created")
	}
	fmt.Printf("%#v \n", nsgGroup)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsgGroup
}

// NuageNsgGroup is a wrapper to create nuage NSG Group in a declaritive way
func NuageNsgGroup(nsgGroupCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.NSGGroup {
	fmt.Println("########################################")
	fmt.Println("#####     NSG Group            #########")
	fmt.Println("########################################")

	nsgGroup := &vspk.NSGGroup{}

	nsgGroups, err := parent.NSGGroups(&bambou.FetchingInfo{
		Filter: nsgGroupCfg["Name"].(string)})
	handleError(err, "NSG group", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if nsgGroups != nil {
		fmt.Println("NSG group already exists")

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

		fmt.Println("NSG Group created")
	}
	fmt.Printf("%#v \n", nsgGroup)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsgGroup
}
