package nuagewrapper

import (
	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
	log "github.com/sirupsen/logrus"
)

// CreateIKEPSK is a wrapper to create a IKE PSK in a declaritive way
func CreateIKEPSK(ikePSKCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.IKEPSK {
	log.Infof("CreateIKEPSK started")

	ikePSK := &vspk.IKEPSK{}

	ikePSKs, err := parent.IKEPSKs(&bambou.FetchingInfo{
		Filter: ikePSKCfg["Name"].(string)})
	handleError(err, "IKE PSK", "READ")

	// init the ikePSK struct that will hold either the received object
	// or will be created from the ikePSKCfg
	if ikePSKs != nil {
		log.Infof("IKE PSK already exists")

		ikePSK = ikePSKs[0]
		errMergo := mergo.Map(ikePSK, ikePSKCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		ikePSK.Save()

	} else {

		//ikePSK = &vspk.IKEPSK{}
		errMergo := mergo.Map(ikePSK, ikePSKCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		ikePSKErr := parent.CreateIKEPSK(ikePSK)
		handleError(ikePSKErr, "IKE PSK", "CREATE")

		log.Infof("IKE PSK created")
	}
	log.Debugf("%#v \n", ikePSK)
	log.Infof("CreateIKEPSK finished")
	return ikePSK
}

// DeleteIKEPSK is a wrapper to delete a IKE PSK in a declaritive way
func DeleteIKEPSK(ikePSKCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.IKEPSK {
	log.Infof("DeleteIKEPSK started")

	ikePSK := &vspk.IKEPSK{}

	ikePSKs, err := parent.IKEPSKs(&bambou.FetchingInfo{
		Filter: ikePSKCfg["Name"].(string)})
	handleError(err, "IKE PSK", "READ")

	// init the ikePSK struct that will hold either the received object
	// or will be created from the ikePSKCfg
	if ikePSKs != nil {
		log.Infof("IKE PSK already exists")

		ikePSK = ikePSKs[0]
		ikePSK.Delete()

	} 
	log.Infof("DeleteIKEPSK finished")
	return ikePSK
}

// CreateIKEGateway is a wrapper to create a IKE GW in a declaritive way
func CreateIKEGateway(ikeGatewayCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.IKEGateway {
	log.Infof("CreateIKEGateway started")
	ikeGateways, err := parent.IKEGateways(&bambou.FetchingInfo{
		Filter: ikeGatewayCfg["Name"].(string)})
	handleError(err, "READ", "IKE Gateway")

	// init the ikeGateway struct that will hold either the received object
	// or will be created from the ikeGatewayCfg
	ikeGateway := &vspk.IKEGateway{}

	if ikeGateways != nil {
		log.Infof("IKE Gateway already exists")

		ikeGateway = ikeGateways[0]
		errMergo := mergo.Map(ikeGateway, ikeGatewayCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		ikeGateway.Save()
	} else {
		//ikeGateway1 = &vspk.IKEGateway{}
		errMergo := mergo.Map(ikeGateway, ikeGatewayCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		ikeGatewayErr := parent.CreateIKEGateway(ikeGateway)
		handleError(ikeGatewayErr, "CREATE", "IKE Gateway")

		log.Infof("IKE Gateway created")

		ikeSubnet := &vspk.IKESubnet{}
		ikeSubnet.Prefix = "0.0.0.0/0"
		ikeSubnetErr := ikeGateway.CreateIKESubnet(ikeSubnet)
		handleError(ikeSubnetErr, "CREATE", "IKE Subnet")
		log.Infof("IKE Subnet created: %s\n", ikeSubnet)
	}
	log.Debugf("%#v \n", ikeGateway)
	log.Infof("CreateIKEGateway finished")
	return ikeGateway
}

// DeleteIKEGateway is a wrapper to delete a IKE GW in a declaritive way
func DeleteIKEGateway(ikeGatewayCfg map[string]interface{}, parent *vspk.Enterprise) error {
	log.Infof("DeleteIKEGateway started")
	
	ikeGateways, err := parent.IKEGateways(&bambou.FetchingInfo{
		Filter: ikeGatewayCfg["Name"].(string)})
	handleError(err, "READ", "IKE Gateway")

	// init the ikeGateway struct that will hold either the received object
	// or will be created from the ikeGatewayCfg
	ikeGateway := &vspk.IKEGateway{}

	if ikeGateways != nil {
		log.Infof("IKE Gateway already exists")

		ikeGateway = ikeGateways[0]
		ikeGateway.Delete()
	} 
	log.Infof("DeleteIKEGateway finished")
	return ikeGateway
}

// CreateIKEEncryptionProfile is a wrapper to create a IKE Encryption Profile in a declaritive way
func CreateIKEEncryptionProfile(ikeEncryptionProfileCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.IKEEncryptionprofile {
	log.Infof("CreateIKEEncryptionProfile started")
	ikeEncryptionProfiles, err := parent.IKEEncryptionprofiles(&bambou.FetchingInfo{
		Filter: ikeEncryptionProfileCfg["Name"].(string)})
	handleError(err, "READ", "IKE Encryption Profile")

	// init the IKEEncryptionprofile struct that will hold either the received object
	// or will be created from the IKEEncryptionprofileCfg
	ikeEncryptionProfile := &vspk.IKEEncryptionprofile{}

	if ikeEncryptionProfiles != nil {
		log.Infof("IKE Encryption Profile already exists")

		ikeEncryptionProfile = ikeEncryptionProfiles[0]
		errMergo := mergo.Map(ikeEncryptionProfile, ikeEncryptionProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		ikeEncryptionProfile.Save()
	} else {
		//ikeEncryptionProfile = &vspk.IKEEncryptionprofile{}
		errMergo := mergo.Map(ikeEncryptionProfile, ikeEncryptionProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		ikeEncryptionProfileErr := parent.CreateIKEEncryptionprofile(ikeEncryptionProfile)
		handleError(ikeEncryptionProfileErr, "CREATE", "IKE Encryption Profile")

		log.Infof("IKE IKE Encryption Profile created")
	}
	log.Debugf("%#v \n", ikeEncryptionProfile)
	log.Infof("CreateIKEEncryptionProfile finished")
	return ikeEncryptionProfile
}

// DeleteIKEEncryptionProfile is a wrapper to delete a IKE Encryption Profile in a declaritive way
func DeleteIKEEncryptionProfile(ikeEncryptionProfileCfg map[string]interface{}, parent *vspk.Enterprise) error {
	log.Infof("DeleteIKEEncryptionProfile started")

	ikeEncryptionProfiles, err := parent.IKEEncryptionprofiles(&bambou.FetchingInfo{
		Filter: ikeEncryptionProfileCfg["Name"].(string)})
	handleError(err, "READ", "IKE Encryption Profile")

	// init the IKEEncryptionprofile struct that will hold either the received object
	// or will be created from the IKEEncryptionprofileCfg
	ikeEncryptionProfile := &vspk.IKEEncryptionprofile{}

	if ikeEncryptionProfiles != nil {
		log.Infof("IKE Encryption Profile already exists")

		ikeEncryptionProfile = ikeEncryptionProfiles[0]
		ikeEncryptionProfile.Delete()
	} 
	log.Infof("DeleteIKEEncryptionProfile finished")
	return ikeEncryptionProfile
}

// CreateIKEGatewayProfile is a wrapper to create a IKE GW Profile in a declaritive way
func CreateIKEGatewayProfile(ikeGatewayProfileCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.IKEGatewayProfile {
	log.Infof("CreateIKEGatewayProfile started")
	ikeGatewayProfiles, err := parent.IKEGatewayProfiles(&bambou.FetchingInfo{
		Filter: ikeGatewayProfileCfg["Name"].(string)})
	handleError(err, "READ", "IKE Gateway Profile")

	// init the ikeGatewayProfile struct that will hold either the received object
	// or will be created from the ikeGatewayProfileCfg
	ikeGatewayProfile := &vspk.IKEGatewayProfile{}

	if ikeGatewayProfiles != nil {
		log.Infof("IKE Gateway Profile already exists")

		ikeGatewayProfile = ikeGatewayProfiles[0]
		errMergo := mergo.Map(ikeGatewayProfile, ikeGatewayProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		ikeGatewayProfile.Save()
	} else {
		//ikeGatewayProfile1 = &vspk.IKEGatewayProfile{}
		errMergo := mergo.Map(ikeGatewayProfile, ikeGatewayProfileCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		ikeGatewayProfileErr := parent.CreateIKEGatewayProfile(ikeGatewayProfile)
		handleError(ikeGatewayProfileErr, "CREATE", "IKE Gateway Profile1")

		log.Infof("IKE Gateway Profile1 created")
	}

	log.Debugf("%#v \n", ikeGatewayProfile)
	log.Infof("CreateIKEGatewayProfile finished")
	return ikeGatewayProfile
}

// DeleteIKEGatewayProfile is a wrapper to delete a IKE GW Profile in a declaritive way
func DeleteIKEGatewayProfile(ikeGatewayProfileCfg map[string]interface{}, parent *vspk.VLAN) error {
	log.Infof("DeleteIKEGatewayProfile started")

	ikeGatewayProfiles, err := parent.IKEGatewayProfiles(&bambou.FetchingInfo{
		Filter: ikeGatewayProfileCfg["Name"].(string)})
	handleError(err, "READ", "IKE Gateway Profile")

	// init the ikeGatewayProfile struct that will hold either the received object
	// or will be created from the ikeGatewayProfileCfg
	ikeGatewayProfile := &vspk.IKEGatewayProfile{}

	if ikeGatewayProfiles != nil {
		log.Infof("IKE Gateway Profile already exists")

		ikeGatewayProfile = ikeGatewayProfiles[0]
		ikeGatewayProfile.Delete()
	}
	log.Infof("DeleteIKEGatewayProfile finished")
	return nil
}

// CreateIKEGatewayConnection is a wrapper to create a IKE GW Connection in a declaritive way
func CreateIKEGatewayConnection(ikeGatewayConnCfg map[string]interface{}, parent *vspk.VLAN) *vspk.IKEGatewayConnection {
	log.Infof("NuageCreateIKEGatewayConnection started")

	ikeGatewayConns, err := parent.IKEGatewayConnections(&bambou.FetchingInfo{
		Filter: ikeGatewayConnCfg["Name"].(string)})
	handleError(err, "READ", "IKE GW Connection")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	ikeGatewayConn := &vspk.IKEGatewayConnection{}

	if ikeGatewayConns != nil {
		log.Infof("IKE GW Connection already exists")

		ikeGatewayConn = ikeGatewayConns[0]
		errMergo := mergo.Map(ikeGatewayConn, ikeGatewayConnCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		ikeGatewayConn.Save()
	} else {
		errMergo := mergo.Map(ikeGatewayConn, ikeGatewayConnCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		ikeGatewayConnErr := parent.CreateIKEGatewayConnection(ikeGatewayConn)
		handleError(ikeGatewayConnErr, "CREATE", "IKE GW Connection ")

		log.Infof("IKE GW Connection created")
	}

	log.Debugf("%#v \n", ikeGatewayConn)
	log.Infof("NuageCreateIKEGatewayConnection finished")
	return ikeGatewayConn
}

// DeleteIKEGatewayConnection is a wrapper to create a IKE GW Connection in a declaritive way
func DeleteIKEGatewayConnection(ikeGatewayConnCfg map[string]interface{}, parent *vspk.VLAN) error {
	log.Infof("NuageDeleteIKEGatewayConnection started")

	ikeGatewayConns, err := parent.IKEGatewayConnections(&bambou.FetchingInfo{
		Filter: ikeGatewayConnCfg["Name"].(string)})
	handleError(err, "READ", "IKE GW Connection")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	ikeGatewayConn := &vspk.IKEGatewayConnection{}

	if ikeGatewayConns != nil {
		log.Infof("IKE GW Connection already exists")
		ikeGatewayConn = ikeGatewayConns[0]
		ikeGatewayConn.Delete()
	}
	log.Debugf("%#v \n", ikeGatewayConn)
	log.Infof("NuageIKEGatewayConnection finished")
	return nil
}
