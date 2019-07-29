package nuagewrapper

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

// NuageCreateIKEPSK is a wrapper to create a IKE PSK in a declaritive way
func NuageCreateIKEPSK(ikePSKCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.IKEPSK {
	//create PSK
	fmt.Println("########################################")
	fmt.Println("#####     IKE PSK          #############")
	fmt.Println("########################################")

	ikePSK := &vspk.IKEPSK{}

	ikePSKs, err := parent.IKEPSKs(&bambou.FetchingInfo{
		Filter: ikePSKCfg["Name"].(string)})
	handleError(err, "IKE PSK", "READ")

	// init the ikePSK struct that will hold either the received object
	// or will be created from the ikePSKCfg
	if ikePSKs != nil {
		fmt.Println("IKE PSK already exists")

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

		fmt.Println("IKE PSK created")
	}
	fmt.Printf("%#v \n", ikePSK)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return ikePSK
}

// NuageCreateIKEGateway is a wrapper to create a IKE GW in a declaritive way
func NuageCreateIKEGateway(ikeGatewayCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.IKEGateway {
	fmt.Println("########################################")
	fmt.Println("#####     IKE Gateway      #############")
	fmt.Println("########################################")

	ikeGateways, err := parent.IKEGateways(&bambou.FetchingInfo{
		Filter: ikeGatewayCfg["Name"].(string)})
	handleError(err, "READ", "IKE Gateway")

	// init the ikeGateway struct that will hold either the received object
	// or will be created from the ikeGatewayCfg
	ikeGateway := &vspk.IKEGateway{}

	if ikeGateways != nil {
		fmt.Println("IKE Gateway already exists")

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

		fmt.Println("IKE Gateway created")

		ikeSubnet := &vspk.IKESubnet{}
		ikeSubnet.Prefix = "0.0.0.0/0"
		ikeSubnetErr := ikeGateway.CreateIKESubnet(ikeSubnet)
		handleError(ikeSubnetErr, "CREATE", "IKE Subnet")
		fmt.Printf("IKE Subnet created: %s\n", ikeSubnet)
	}
	fmt.Printf("%#v \n", ikeGateway)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return ikeGateway
}

// NuageCreateIKEEncryptionProfile is a wrapper to create a IKE Encryption Profile in a declaritive way
func NuageCreateIKEEncryptionProfile(ikeEncryptionProfileCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.IKEEncryptionprofile {
	fmt.Println("########################################")
	fmt.Println("#####IKE Encryption Profile#############")
	fmt.Println("########################################")

	ikeEncryptionProfiles, err := parent.IKEEncryptionprofiles(&bambou.FetchingInfo{
		Filter: ikeEncryptionProfileCfg["Name"].(string)})
	handleError(err, "READ", "IKE Encryption Profile")

	// init the IKEEncryptionprofile struct that will hold either the received object
	// or will be created from the IKEEncryptionprofileCfg
	ikeEncryptionProfile := &vspk.IKEEncryptionprofile{}

	if ikeEncryptionProfiles != nil {
		fmt.Println("IKE Encryption Profile already exists")

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

		fmt.Println("IKE IKE Encryption Profile created")
	}
	fmt.Printf("%#v \n", ikeEncryptionProfile)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return ikeEncryptionProfile
}

// NuageCreateIKEGatewayProfile is a wrapper to create a IKE GW Profile in a declaritive way
func NuageCreateIKEGatewayProfile(ikeGatewayProfileCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.IKEGatewayProfile {
	fmt.Println("########################################")
	fmt.Println("#####   IKE Gateway Profile   ##########")
	fmt.Println("########################################")

	ikeGatewayProfiles, err := parent.IKEGatewayProfiles(&bambou.FetchingInfo{
		Filter: ikeGatewayProfileCfg["Name"].(string)})
	handleError(err, "READ", "IKE Gateway Profile")

	// init the ikeGatewayProfile struct that will hold either the received object
	// or will be created from the ikeGatewayProfileCfg
	ikeGatewayProfile := &vspk.IKEGatewayProfile{}

	if ikeGatewayProfiles != nil {
		fmt.Println("IKE Gateway Profile already exists")

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

		fmt.Println("IKE Gateway Profile1 created")
	}

	fmt.Printf("%#v \n", ikeGatewayProfile)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return ikeGatewayProfile
}

// NuageCreateIKEGatewayConnection is a wrapper to create a IKE GW Connection in a declaritive way
func NuageCreateIKEGatewayConnection(ikeGatewayConnCfg map[string]interface{}, parent *vspk.VLAN) *vspk.IKEGatewayConnection {
	log.Println("NuageCreateIKEGatewayConnection started")

	ikeGatewayConns, err := parent.IKEGatewayConnections(&bambou.FetchingInfo{
		Filter: ikeGatewayConnCfg["Name"].(string)})
	handleError(err, "READ", "IKE GW Connection")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	ikeGatewayConn := &vspk.IKEGatewayConnection{}

	if ikeGatewayConns != nil {
		fmt.Println("IKE GW Connection already exists")

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

		fmt.Println("IKE GW Connection created")
	}

	log.Printf("%#v \n", ikeGatewayConn)
	log.Println("NuageCreateIKEGatewayConnection finished")
	return ikeGatewayConn
}

// NuageDeleteIKEGatewayConnection is a wrapper to create a IKE GW Connection in a declaritive way
func NuageDeleteIKEGatewayConnection(ikeGatewayConnCfg map[string]interface{}, parent *vspk.VLAN) error {
	log.Println("NuageDeleteIKEGatewayConnection started")

	ikeGatewayConns, err := parent.IKEGatewayConnections(&bambou.FetchingInfo{
		Filter: ikeGatewayConnCfg["Name"].(string)})
	handleError(err, "READ", "IKE GW Connection")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	ikeGatewayConn := &vspk.IKEGatewayConnection{}

	if ikeGatewayConns != nil {
		fmt.Println("IKE GW Connection already exists")
		ikeGatewayConn = ikeGatewayConns[0]
		ikeGatewayConn.Delete()
	}
	log.Printf("%#v \n", ikeGatewayConn)
	log.Println("NuageIKEGatewayConnection finished")
	return nil
}
