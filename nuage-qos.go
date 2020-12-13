package nuagewrapper

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

// NuageRateLimiter is a wrapper to create a Rate Limiter in a declaritive way
func NuageRateLimiter(rateLimiterCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.RateLimiter {
	fmt.Println("########################################")
	fmt.Println("#####   Rate Limiter         ##########")
	fmt.Println("########################################")

	rateLimiters, err := parent.RateLimiters(&bambou.FetchingInfo{
		Filter: rateLimiterCfg["Name"].(string)})
	handleError(err, "READ", "rateLimiters")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the Cfg
	rateLimiter := &vspk.RateLimiter{}

	if rateLimiters != nil {
		fmt.Println("Rate Limiter already exists")

		rateLimiter = rateLimiters[0]
		errMergo := mergo.Map(rateLimiter, rateLimiterCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		rateLimiter.Save()
	} else {
		errMergo := mergo.Map(rateLimiter, rateLimiterCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateRateLimiter(rateLimiter)
		handleError(err, "CREATE", "rateLimiter")

		fmt.Println("rateLimiter created")
	}

	fmt.Printf("%#v \n", rateLimiter)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return rateLimiter
}

// NuageEgressQoSPolicy is a wrapper to create a QoS Policy in a declaritive way
func NuageEgressQoSPolicy(egressQoSPolicyCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.EgressQOSPolicy {
	fmt.Println("########################################")
	fmt.Println("#####   Egress QoS Policy     ##########")
	fmt.Println("########################################")

	egressQoSPolicies, err := parent.EgressQOSPolicies(&bambou.FetchingInfo{
		Filter: egressQoSPolicyCfg["Name"].(string)})
	handleError(err, "READ", "egressQoSPolicies")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	egressQoSPolicy := &vspk.EgressQOSPolicy{}

	if egressQoSPolicies != nil {
		fmt.Println("QoS Policy already exists")

		egressQoSPolicy = egressQoSPolicies[0]
		errMergo := mergo.Map(egressQoSPolicy, egressQoSPolicyCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		egressQoSPolicy.Save()
	} else {
		errMergo := mergo.Map(egressQoSPolicy, egressQoSPolicyCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateEgressQOSPolicy(egressQoSPolicy)
		handleError(err, "CREATE", "egressQoSPolicy")

		fmt.Println("egressQoSPolicy created")
	}

	fmt.Printf("%#v \n", egressQoSPolicy)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return egressQoSPolicy
}

// NuageCOSRemarkingPolicy is a wrapper to create a COS remarking Policy in a declaritive way
func NuageCOSRemarkingPolicy(cosRemarkingPolicyCfg map[string]interface{}, parent *vspk.COSRemarkingPolicyTable) *vspk.COSRemarkingPolicy {
	fmt.Println("########################################")
	fmt.Println("#####   COS Remarking Policy  ##########")
	fmt.Println("########################################")

	cosRemarkingPolicies, err := parent.COSRemarkingPolicies(&bambou.FetchingInfo{
		Filter: cosRemarkingPolicyCfg["DSCP"].(string)})
	handleError(err, "READ", "cosRemarkingPolicy")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	cosRemarkingPolicy := &vspk.COSRemarkingPolicy{}

	if cosRemarkingPolicies != nil {
		fmt.Println("COS Remark Policy Table already exists")
		cosRemarkingPolicy = cosRemarkingPolicies[0]
		return cosRemarkingPolicy
	}

	errMergo := mergo.Map(cosRemarkingPolicy, cosRemarkingPolicyCfg, mergo.WithOverride)
	if errMergo != nil {
		log.Fatal(errMergo)
	}

	err = parent.CreateCOSRemarkingPolicy(cosRemarkingPolicy)
	handleError(err, "CREATE", "cosRemarkingPolicy")

	fmt.Println("cosRemarkingPolicy created")

	fmt.Printf("%#v \n", cosRemarkingPolicy)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return cosRemarkingPolicy
}

// NuageCOSRemarkingPolicyTable is a wrapper to create a COS remarking Policy in a declaritive way
func NuageCOSRemarkingPolicyTable(cosRemarkingPolicyTableCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.COSRemarkingPolicyTable {
	fmt.Println("########################################")
	fmt.Println("#####   COS Remarking Policy Table #####")
	fmt.Println("########################################")

	cosRemarkingPolicyTables, err := parent.COSRemarkingPolicyTables(&bambou.FetchingInfo{
		Filter: cosRemarkingPolicyTableCfg["Name"].(string)})
	handleError(err, "READ", "cosRemarkingPolicyTables")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	cosRemarkingPolicyTable := &vspk.COSRemarkingPolicyTable{}

	if cosRemarkingPolicyTables != nil {
		fmt.Println("COS Remark Policy Table already exists")

		cosRemarkingPolicyTable = cosRemarkingPolicyTables[0]
		errMergo := mergo.Map(cosRemarkingPolicyTable, cosRemarkingPolicyTableCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		cosRemarkingPolicyTable.Save()
	} else {
		errMergo := mergo.Map(cosRemarkingPolicyTable, cosRemarkingPolicyTableCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateCOSRemarkingPolicyTable(cosRemarkingPolicyTable)
		handleError(err, "CREATE", "cosRemarkingPolicyTable")

		fmt.Println("cosRemarkingPolicyTable created")
	}

	fmt.Printf("%#v \n", cosRemarkingPolicyTable)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return cosRemarkingPolicyTable
}

// NuageDSCPRemarkingPolicy is a wrapper to create a DSCP remarking Policy in a declaritive way
func NuageDSCPRemarkingPolicy(dscpRemarkingPolicyCfg map[string]interface{}, parent *vspk.DSCPRemarkingPolicyTable) *vspk.DSCPRemarkingPolicy {
	fmt.Println("########################################")
	fmt.Println("#####  DSCP Remarking Policy  ##########")
	fmt.Println("########################################")

	dscpRemarkingPolicies, err := parent.DSCPRemarkingPolicies(&bambou.FetchingInfo{
		Filter: dscpRemarkingPolicyCfg["DSCP"].(string)})
	handleError(err, "READ", "dscpRemarkingPolicy")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	dscpRemarkingPolicy := &vspk.DSCPRemarkingPolicy{}

	if dscpRemarkingPolicies != nil {
		fmt.Println("DSCP Remark Policy  already exists")
		dscpRemarkingPolicy = dscpRemarkingPolicies[0]
		return dscpRemarkingPolicy
	}

	errMergo := mergo.Map(dscpRemarkingPolicy, dscpRemarkingPolicyCfg, mergo.WithOverride)
	if errMergo != nil {
		log.Fatal(errMergo)
	}

	err = parent.CreateDSCPRemarkingPolicy(dscpRemarkingPolicy)
	handleError(err, "CREATE", "dscpRemarkingPolicy")

	fmt.Println("dscpRemarkingPolicy created")

	fmt.Printf("%#v \n", dscpRemarkingPolicy)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return dscpRemarkingPolicy
}

// NuageDSCPRemarkingPolicyTable is a wrapper to create a DSCP remarking Policy in a declaritive way
func NuageDSCPRemarkingPolicyTable(dscpRemarkingPolicyTableCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.DSCPRemarkingPolicyTable {
	fmt.Println("########################################")
	fmt.Println("#####  DSCP Remarking Policy Table #####")
	fmt.Println("########################################")

	dscpRemarkingPolicyTables, err := parent.DSCPRemarkingPolicyTables(&bambou.FetchingInfo{
		Filter: dscpRemarkingPolicyTableCfg["Name"].(string)})
	handleError(err, "READ", "dscpRemarkingPolicyTables")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	dscpRemarkingPolicyTable := &vspk.DSCPRemarkingPolicyTable{}

	if dscpRemarkingPolicyTables != nil {
		fmt.Println("DSCP Remark Policy Table already exists")

		dscpRemarkingPolicyTable = dscpRemarkingPolicyTables[0]
		errMergo := mergo.Map(dscpRemarkingPolicyTable, dscpRemarkingPolicyTableCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		dscpRemarkingPolicyTable.Save()
	} else {
		errMergo := mergo.Map(dscpRemarkingPolicyTable, dscpRemarkingPolicyTableCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateDSCPRemarkingPolicyTable(dscpRemarkingPolicyTable)
		handleError(err, "CREATE", "dscpRemarkingPolicyTable")

		fmt.Println("dscpRemarkingPolicyTable created")
	}

	fmt.Printf("%#v \n", dscpRemarkingPolicyTable)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return dscpRemarkingPolicyTable
}
