package nuagewrapper

import (
	"github.com/henderiw/nuage-wrapper/pkg/vspk"
	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	log "github.com/sirupsen/logrus"
)

// RateLimiter is a wrapper to create a Rate Limiter in a declaritive way
func RateLimiter(rateLimiterCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.RateLimiter {

	rateLimiters, err := parent.RateLimiters(&bambou.FetchingInfo{
		Filter: rateLimiterCfg["Name"].(string)})
	handleError(err, "READ", "rateLimiters")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the Cfg
	rateLimiter := &vspk.RateLimiter{}

	if rateLimiters != nil {
		log.Infof("Rate Limiter already exists")

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

		log.Infof("rateLimiter created")
	}

	log.Infof("%#v \n", rateLimiter)
	return rateLimiter
}

// EgressQoSPolicy is a wrapper to create a QoS Policy in a declaritive way
func EgressQoSPolicy(egressQoSPolicyCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.EgressQOSPolicy {

	egressQoSPolicies, err := parent.EgressQOSPolicies(&bambou.FetchingInfo{
		Filter: egressQoSPolicyCfg["Name"].(string)})
	handleError(err, "READ", "egressQoSPolicies")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	egressQoSPolicy := &vspk.EgressQOSPolicy{}

	if egressQoSPolicies != nil {
		log.Infof("QoS Policy already exists")

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

		log.Infof("egressQoSPolicy created")
	}

	log.Infof("%#v \n", egressQoSPolicy)
	return egressQoSPolicy
}

// COSRemarkingPolicy is a wrapper to create a COS remarking Policy in a declaritive way
func COSRemarkingPolicy(cosRemarkingPolicyCfg map[string]interface{}, parent *vspk.COSRemarkingPolicyTable) *vspk.COSRemarkingPolicy {

	cosRemarkingPolicies, err := parent.COSRemarkingPolicies(&bambou.FetchingInfo{
		Filter: cosRemarkingPolicyCfg["DSCP"].(string)})
	handleError(err, "READ", "cosRemarkingPolicy")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	cosRemarkingPolicy := &vspk.COSRemarkingPolicy{}

	if cosRemarkingPolicies != nil {
		log.Infof("COS Remark Policy Table already exists")
		cosRemarkingPolicy = cosRemarkingPolicies[0]
		return cosRemarkingPolicy
	}

	errMergo := mergo.Map(cosRemarkingPolicy, cosRemarkingPolicyCfg, mergo.WithOverride)
	if errMergo != nil {
		log.Fatal(errMergo)
	}

	err = parent.CreateCOSRemarkingPolicy(cosRemarkingPolicy)
	handleError(err, "CREATE", "cosRemarkingPolicy")

	log.Infof("cosRemarkingPolicy created")

	log.Infof("%#v \n", cosRemarkingPolicy)

	return cosRemarkingPolicy
}

// COSRemarkingPolicyTable is a wrapper to create a COS remarking Policy in a declaritive way
func COSRemarkingPolicyTable(cosRemarkingPolicyTableCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.COSRemarkingPolicyTable {

	cosRemarkingPolicyTables, err := parent.COSRemarkingPolicyTables(&bambou.FetchingInfo{
		Filter: cosRemarkingPolicyTableCfg["Name"].(string)})
	handleError(err, "READ", "cosRemarkingPolicyTables")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	cosRemarkingPolicyTable := &vspk.COSRemarkingPolicyTable{}

	if cosRemarkingPolicyTables != nil {
		log.Infof("COS Remark Policy Table already exists")

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

		log.Infof("cosRemarkingPolicyTable created")
	}

	log.Infof("%#v \n", cosRemarkingPolicyTable)

	return cosRemarkingPolicyTable
}

// DSCPRemarkingPolicy is a wrapper to create a DSCP remarking Policy in a declaritive way
func DSCPRemarkingPolicy(dscpRemarkingPolicyCfg map[string]interface{}, parent *vspk.DSCPRemarkingPolicyTable) *vspk.DSCPRemarkingPolicy {

	dscpRemarkingPolicies, err := parent.DSCPRemarkingPolicies(&bambou.FetchingInfo{
		Filter: dscpRemarkingPolicyCfg["DSCP"].(string)})
	handleError(err, "READ", "dscpRemarkingPolicy")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	dscpRemarkingPolicy := &vspk.DSCPRemarkingPolicy{}

	if dscpRemarkingPolicies != nil {
		log.Infof("DSCP Remark Policy  already exists")
		dscpRemarkingPolicy = dscpRemarkingPolicies[0]
		return dscpRemarkingPolicy
	}

	errMergo := mergo.Map(dscpRemarkingPolicy, dscpRemarkingPolicyCfg, mergo.WithOverride)
	if errMergo != nil {
		log.Fatal(errMergo)
	}

	err = parent.CreateDSCPRemarkingPolicy(dscpRemarkingPolicy)
	handleError(err, "CREATE", "dscpRemarkingPolicy")

	log.Infof("dscpRemarkingPolicy created")

	log.Infof("%#v \n", dscpRemarkingPolicy)

	return dscpRemarkingPolicy
}

// DSCPRemarkingPolicyTable is a wrapper to create a DSCP remarking Policy in a declaritive way
func DSCPRemarkingPolicyTable(dscpRemarkingPolicyTableCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.DSCPRemarkingPolicyTable {

	dscpRemarkingPolicyTables, err := parent.DSCPRemarkingPolicyTables(&bambou.FetchingInfo{
		Filter: dscpRemarkingPolicyTableCfg["Name"].(string)})
	handleError(err, "READ", "dscpRemarkingPolicyTables")

	// init the struct that will hold either the received object
	// or will be created from the Cfg
	dscpRemarkingPolicyTable := &vspk.DSCPRemarkingPolicyTable{}

	if dscpRemarkingPolicyTables != nil {
		log.Infof("DSCP Remark Policy Table already exists")

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

		log.Infof("dscpRemarkingPolicyTable created")
	}

	log.Infof("%#v \n", dscpRemarkingPolicyTable)
	return dscpRemarkingPolicyTable
}
