package nuagewrapper

import (
	"strings"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
	log "github.com/sirupsen/logrus"
)

// NuageNetworkVlanCfg defines the structure of a network VLAN for an NSG
type NuageNetworkVlanCfg struct {
	VlanID           int    `json:"vlanId"`
	Role             string `json:"Role"`
	AddressFamily    string `json:"addressFamily"`
	Mode             string `json:"mode"`
	UnderlayName     string `json:"underlayName,omitempty"`
	UnderlayID       string `json:"underlayID,omitempty"`
	VscName          string `json:"vscName,omitempty"`
	VscID            string `json:"vscId,omitempty"`
	Address          string `json:"address,omitempty"`
	Netmask          string `json:"netmask,omitempty"`
	DNS              string `json:"dns,omitempty"`
	Gateway          string `json:"gateway,omitempty"`
	DucVlan          bool   `json:"ducVLAN,omitempty"`
	LteConfiguration struct {
		Apn            string `json:"apn"`
		PdpType        string `json:"pdp-type"`
		PinCode        string `json:"pin-code"`
		Authentication string `json:"authentication"`
		UserName       string `json:"username"`
		PassWord       string `json:"password"`
	} `json:"lteConfiguration,omitempty"`
}

// NuageAccessVlanCfg defines the structure of a access VLAN for an NSG
type NuageAccessVlanCfg struct {
	VlanID int `json:"vlanId"`
}

// NuageNSGCfg defines the structure of a template for an NSG
type NuageNSGCfg struct {
	Name                string `json:"Name"`
	NSGTemplateName     string `json:"NSGTemplateName"`
	NSGTemplateID       string `json:"NSGTemplateID"`
	NetworkAcceleration string `json:"networkAcceleration,omitempty"`
	NetworkPorts        []struct {
		Name string                `json:"Name"`
		Vlan []NuageNetworkVlanCfg `json:"Vlan`
	} `json:"NetworkPorts"`
	ShuntPorts []struct {
		Name string                `json:"Name"`
		Vlan []NuageNetworkVlanCfg `json:"Vlan`
	} `json:"ShuntPorts"`
	AccessPorts []struct {
		Name string               `json:"Name"`
		Vlan []NuageAccessVlanCfg `json:"Vlan`
	} `json:"AccessPorts"`
	WifiPorts []struct {
		Name string `json:"Name"`
		Ssid string `json:"ssid"`
	} `json:"WifiPorts"`
}

// CreateEntireNSG is a wrapper to create a complete NSG in a declaritive way
func CreateEntireNSG(nsgCfg NuageNSGCfg, parent *vspk.Enterprise, Usr *vspk.Me) *vspk.NSGateway {

	nsGatewayTemplateCfg := map[string]interface{}{
		"Name": nsgCfg.NSGTemplateName,
	}
	//log.Infof("NSG Template ID: %s \n", nsgCfg.NSGTemplateID)

	nsGatewayTemplate := NSGatewayTemplate(nsGatewayTemplateCfg, Usr)
	nsgCfg.NSGTemplateID = nsGatewayTemplate.ID
	//flog.Infof("NSG Template ID: %s \n", nsgCfg.NSGTemplateID)

	log.Infof("NSG Template Name: %s ", nsgCfg.NSGTemplateName)
	log.Infof("NSG Template Personality: %s ", nsGatewayTemplate.Personality)
	log.Infof("NSG Template ID: %s ", nsgCfg.NSGTemplateID)

	networkAcceleration := nsgCfg.NetworkAcceleration
	if nsgCfg.ShuntPorts != nil {
		networkAcceleration = "PERFORMANCE"
	} else if nsGatewayTemplate.Personality == "NSGDUC" {
		networkAcceleration = "PERFORMANCE"
	}

	var functions []interface{}
	var tunnelShaping string
	tunnelShaping = "DISABLED"
	if nsGatewayTemplate.Personality == "NSGDUC" {
		functions = []interface{}{"UBR", "HUB", "GATEWAY"}
		tunnelShaping = "ENABLED"
	}

	nsGatewayCfg := map[string]interface{}{
		"Name":                  nsgCfg.Name,
		"TCPMSSEnabled":         true,
		"TCPMaximumSegmentSize": 1330,
		"NetworkAcceleration":   networkAcceleration,
		"TemplateID":            nsgCfg.NSGTemplateID,
		"Functions":             functions,
		"TunnelShaping":         tunnelShaping,
	}

	var nsGateway *vspk.NSGateway
	if nsGatewayTemplate.Personality == "NSGDUC" || nsGatewayTemplate.Personality == "NSGBR" {
		nsGateway = NSGRoot(nsGatewayCfg, Usr)
	} else {
		nsGateway = NSG(nsGatewayCfg, parent)
	}

	//time.Sleep(15 * time.Second)

	for i, port := range nsgCfg.NetworkPorts {
		log.Infof("NSG Network Port %d Name: %s \n", i, port.Name)

		nsPortCfg := map[string]interface{}{
			"Name":            port.Name,
			"PhysicalName":    port.Name,
			"PortType":        "NETWORK",
			"VLANRange":       "0-4094",
			"EnableNATProbes": true,
			"NATTraversal":    "FULL_NAT",
		}
		nsPort := NSGPort(nsPortCfg, nsGateway)

		log.Infof("Port: %#v \n", nsPort)

		for _, vlan := range port.Vlan {
			var nsVlanCfg map[string]interface{}
			if vlan.VscName != "" {
				vscProfCfg := map[string]interface{}{
					"Name": vlan.VscName,
				}
				vscProf := InfraVscProfile(vscProfCfg, Usr)
				vlan.VscID = vscProf.ID

				//log.Infof("VSC NEEDED for vlan: %s with VSC ID: %s", vlan.VlanID, vlan.VscID)
				//time.Sleep(5 * time.Second)

				nsVlanCfg = map[string]interface{}{
					"Value":                  vlan.VlanID,
					"AssociatedVSCProfileID": vlan.VscID,
				}
				if nsGatewayTemplate.Personality == "NSGDUC" {
					nsVlanCfg = map[string]interface{}{
						"Value":                  vlan.VlanID,
						"AssociatedVSCProfileID": vlan.VscID,
						"DucVlan":                false,
					}
				}
			} else {
				//log.Infof("NOOOOOOOOOOOOOOO VSC NEEDED for vlan: %s with VSC ID: %s", vlan.VlanID, vlan.VscID)
				//time.Sleep(5 * time.Second)
				nsVlanCfg = map[string]interface{}{
					"Value": vlan.VlanID,
				}
				if nsGatewayTemplate.Personality == "NSGDUC" {
					nsVlanCfg = map[string]interface{}{
						"Value":   vlan.VlanID,
						"DucVlan": true,
					}
				}
			}

			log.Infof("VLANCfg: %#v \n", nsVlanCfg)
			nsVlan := Vlan(nsVlanCfg, nsPort)
			log.Infof("Port: %#v \n", nsVlan)

			var patEnabled = true
			var underlayEnabled = true
			var dnsV4 = ""
			var addressV4 = ""
			var netmaskV4 = ""
			var gatewayV4 = ""
			var dnsV6 = ""
			var addressV6 = ""
			var gatewayV6 = ""

			if vlan.AddressFamily == "IPV4" && vlan.Mode == "Static" {
				dnsV4 = vlan.DNS
				addressV4 = vlan.Address
				netmaskV4 = vlan.Netmask
				gatewayV4 = vlan.Gateway
				log.Infof("BRANCH IPV4 AND STATIC ADDRESSING MODE \n")
			} else if vlan.AddressFamily == "IPV6" && vlan.Mode == "Static" {
				dnsV6 = vlan.DNS
				addressV6 = vlan.Address
				gatewayV6 = vlan.Gateway
				patEnabled = false
				underlayEnabled = false
				log.Infof("BRANCH IPV6 AND STATIC ADDRESSING MODE \n")
			} else if vlan.AddressFamily == "IPV6" {
				patEnabled = false
				underlayEnabled = false
				log.Infof("BRANCH IPV6 AND DYNAMIC ADDRESSING MODE \n")
			}

			var uplinkConnectionCfg map[string]interface{}
			if vlan.UnderlayName != "" {
				underlayCfg := map[string]interface{}{
					"Name": vlan.UnderlayName,
				}
				underlay := Underlay(underlayCfg, Usr)
				log.Infof("Underlay: %v", underlay)
				vlan.UnderlayID = underlay.ID

				if vlan.DucVlan == true {
					//UBR Data VLAN cannot have a Role assigned (PRIMARY/SECONDARY)
					uplinkConnectionCfg = map[string]interface{}{
						"PATEnabled":      patEnabled,
						"UnderlayEnabled": underlayEnabled,
						"Mode":            vlan.Mode,
						"AddressFamily":   vlan.AddressFamily,
						"DNSAddress":      dnsV4,
						"Gateway":         gatewayV4,
						"Address":         addressV4,
						"Netmask":         netmaskV4,
						"DNSAddressV6":    dnsV6,
						"GatewayV6":       gatewayV6,
						"AddressV6":       addressV6,
						"AssocUnderlayID": vlan.UnderlayID,
					}
				} else {
					uplinkConnectionCfg = map[string]interface{}{
						"PATEnabled":      patEnabled,
						"UnderlayEnabled": underlayEnabled,
						"Role":            vlan.Role,
						"Mode":            vlan.Mode,
						"AddressFamily":   vlan.AddressFamily,
						"DNSAddress":      dnsV4,
						"Gateway":         gatewayV4,
						"Address":         addressV4,
						"Netmask":         netmaskV4,
						"DNSAddressV6":    dnsV6,
						"GatewayV6":       gatewayV6,
						"AddressV6":       addressV6,
						"AssocUnderlayID": vlan.UnderlayID,
					}
				}

			} else {
				uplinkConnectionCfg = map[string]interface{}{
					"PATEnabled":      patEnabled,
					"UnderlayEnabled": underlayEnabled,
					"Role":            vlan.Role,
					"Mode":            vlan.Mode,
					"AddressFamily":   vlan.AddressFamily,
					"DNSAddress":      dnsV4,
					"Gateway":         gatewayV4,
					"Address":         addressV4,
					"Netmask":         netmaskV4,
					"DNSAddressV6":    dnsV6,
					"GatewayV6":       gatewayV6,
					"AddressV6":       addressV6,
				}
			}

			//log.Infof(uplinkConnectionCfg)

			uplinkConn := UplinkConnection(uplinkConnectionCfg, nsVlan)

			if strings.Contains(port.Name, "lte") {
				log.Infof("LTE")

				customePropCfg := map[string]interface{}{
					"AttributeName":  "apn",
					"AttributeValue": vlan.LteConfiguration.Apn,
				}

				CustomProperty(customePropCfg, uplinkConn)

				customePropCfg = map[string]interface{}{
					"AttributeName":  "pdp-type",
					"AttributeValue": vlan.LteConfiguration.PdpType,
				}

				CustomProperty(customePropCfg, uplinkConn)

				customePropCfg = map[string]interface{}{
					"AttributeName":  "sim-pin",
					"AttributeValue": vlan.LteConfiguration.PinCode,
				}

				CustomProperty(customePropCfg, uplinkConn)

				if vlan.LteConfiguration.Authentication != "" {
					customePropCfg = map[string]interface{}{
						"AttributeName":  "authentication",
						"AttributeValue": vlan.LteConfiguration.Authentication,
					}

					CustomProperty(customePropCfg, uplinkConn)
				}

				if vlan.LteConfiguration.UserName != "" {
					customePropCfg = map[string]interface{}{
						"AttributeName":  "username",
						"AttributeValue": vlan.LteConfiguration.UserName,
					}

					CustomProperty(customePropCfg, uplinkConn)
				}

				if vlan.LteConfiguration.PassWord != "" {
					customePropCfg = map[string]interface{}{
						"AttributeName":  "password",
						"AttributeValue": vlan.LteConfiguration.PassWord,
					}

					CustomProperty(customePropCfg, uplinkConn)
				}

			} else {
				log.Infof("ETHERNET")
			}
		}

	}
	for i, port := range nsgCfg.ShuntPorts {
		log.Infof("NSG Shunt Port %d Name: %s \n", i, port.Name)

		nsPortCfg := map[string]interface{}{
			"Name":            port.Name,
			"PhysicalName":    port.Name,
			"PortType":        "NETWORK",
			"VLANRange":       "0-4094",
			"EnableNATProbes": true,
			"NATTraversal":    "FULL_NAT",
			"Mtu":             2000,
		}
		nsPort := NSGPort(nsPortCfg, nsGateway)

		for _, vlan := range port.Vlan {
			nsVlanCfg := map[string]interface{}{
				"Value":       vlan.VlanID,
				"Name":        "shunt",
				"Description": "shunt",
				"ShuntVLAN":   true,
			}

			log.Infof("VLANCfg: %s \n", nsVlanCfg)
			nsVlan := Vlan(nsVlanCfg, nsPort)
			//log.Infof(nsVlan)

			uplinkConnCfg := map[string]interface{}{
				"PATEnabled":    true,
				"Role":          vlan.Role,
				"Mode":          "Static",
				"AddressFamily": vlan.AddressFamily,
				"DNSAddress":    vlan.DNS,
				"Gateway":       vlan.Gateway,
				"Address":       vlan.Address,
				"Netmask":       vlan.Netmask,
				//	"DNSAddressV6":    dnsV6,
				//	"GatewayV6":       gatewayV6,
				//	"AddressV6":       addressV6,
				"UnderlayEnabled": true,
				//"UnderlayID":      port.UnderlayID,
			}

			uplinkConn := UplinkConnection(uplinkConnCfg, nsVlan)
			//log.Infof(uplinkConn)
		}

	}
	for i, port := range nsgCfg.AccessPorts {
		log.Infof("NSG Access Port %d Name: %s \n", i, port.Name)

		nsPortCfg := map[string]interface{}{
			"Name":         port.Name,
			"PhysicalName": port.Name,
			"PortType":     "ACCESS",
			"VLANRange":    "0-4094",
		}
		nsPort := NSGPort(nsPortCfg, nsGateway)

		for _, vlan := range port.Vlan {
			nsVlanCfg := map[string]interface{}{
				"Value": vlan.VlanID,
			}
			log.Infof("Access VLANCfg: %#v \n", nsVlanCfg)
			nsVlan := Vlan(nsVlanCfg, nsPort)
			//log.Infof(nsVlan)
		}

	}
	for i, port := range nsgCfg.WifiPorts {
		log.Infof("NSG Wifi Port %d Name: %s \n", i, port.Name)

		nsPortCfg := map[string]interface{}{
			"Name":              port.Name,
			"WifiFrequencyBand": "FREQ_2_4_GHZ",
			"WifiMode":          "WIFI_B_G_N",
			"CountryCode":       "BE",
		}
		nsPort := NSGWirelessPort(nsPortCfg, nsGateway)

		ssidConnCfg := map[string]interface{}{
			"Name":               port.Ssid,
			"Passphrase":         "4no*heydQ",
			"AuthenticationMode": "WPA2",
			"BroadcastSSID":      true,
		}
		ssidConn := SSIDConnection(ssidConnCfg, nsPort)
		//log.Infof(ssidConn)
	}

	log.Infof("%#v \n", nsGateway)
	return nsGateway
}

// NSGatewayTemplate is a wrapper to create nuage NS Gateway template in a declaritive way
func NSGatewayTemplate(nsGatewayTemplateCfg map[string]interface{}, parent *vspk.Me) *vspk.NSGatewayTemplate {

	nsGatewayTemplate := &vspk.NSGatewayTemplate{}

	nsGatewayTemplates, err := parent.NSGatewayTemplates(&bambou.FetchingInfo{
		Filter: nsGatewayTemplateCfg["Name"].(string)})
	handleError(err, "nsGatewayTemplate", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if nsGatewayTemplates != nil {
		log.Infof("nsGatewayTemplate already exists")

		nsGatewayTemplate = nsGatewayTemplates[0]
		errMergo := mergo.Map(nsGatewayTemplate, nsGatewayTemplateCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		nsGatewayTemplate.Save()

	} else {

		errMergo := mergo.Map(nsGatewayTemplate, nsGatewayTemplateCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateNSGatewayTemplate(nsGatewayTemplate)
		handleError(err, "nsGatewayTemplate", "CREATE")

		log.Infof("nsGatewayTemplate created")
	}
	log.Infof("%#v \n", nsGatewayTemplate)
	return nsGatewayTemplate
}

// NSGRoot is a wrapper to create nuage NS Gateway in a declaritive way
func NSGRoot(nsGatewayCfg map[string]interface{}, parent *vspk.Me) *vspk.NSGateway {

	nsGateways, err := parent.NSGateways(&bambou.FetchingInfo{
		Filter: nsGatewayCfg["Name"].(string)})
	handleError(err, "READ", "NS Gateway")

	// init the nsGateway struct that will hold either the received object
	// or will be created from the nsGatewayCfg
	nsGateway := &vspk.NSGateway{}

	if nsGateways != nil {
		log.Infof("NS Gateway already exists")

		nsGateway = nsGateways[0]
		errMergo := mergo.Map(nsGateway, nsGatewayCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		nsGateway.Save()
	} else {
		errMergo := mergo.Map(nsGateway, nsGatewayCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		//log.Infof("nsGateway: %#v", nsGateway)
		//time.Sleep(15 * time.Second)

		err := parent.CreateNSGateway(nsGateway)
		handleError(err, "CREATE", "NS Gateway ")

		log.Infof("NS Gateway created")
	}

	log.Infof("%#v \n", nsGateway)
	return nsGateway
}

// NSG is a wrapper to create nuage NS Gateway in a declaritive way
func NSG(nsGatewayCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.NSGateway {
	nsGateways, err := parent.NSGateways(&bambou.FetchingInfo{
		Filter: nsGatewayCfg["Name"].(string)})
	handleError(err, "READ", "NS Gateway")

	// init the nsGateway struct that will hold either the received object
	// or will be created from the nsGatewayCfg
	nsGateway := &vspk.NSGateway{}

	if nsGateways != nil {
		log.Infof("NS Gateway already exists")

		nsGateway = nsGateways[0]
		errMergo := mergo.Map(nsGateway, nsGatewayCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		nsGateway.Save()
	} else {
		errMergo := mergo.Map(nsGateway, nsGatewayCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateNSGateway(nsGateway)
		handleError(err, "CREATE", "NS Gateway ")

		log.Infof("NS Gateway created")
	}

	log.Infof("%#v \n", nsGateway)
	return nsGateway
}

// GetNSG is a wrapper to get nuage NS Gateway in a declaritive way
func GetNSG(nsGatewayCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.NSGateway {
	nsGateways, err := parent.NSGateways(&bambou.FetchingInfo{
		Filter: nsGatewayCfg["Name"].(string)})
	handleError(err, "READ", "NS Gateway")

	// init the nsGateway struct that will hold either the received object
	// or will be created from the nsGatewayCfg
	nsGateway := &vspk.NSGateway{}

	if nsGateways == nil {
		log.Infof("NS Gateway does not exists")

		
		return nil
	} 
	nsGateway = nsGateways[0]
	return nsGateway
}

// DeleteNSG is a wrapper to delete nuage NS Gateway in a declaritive way
func DeleteNSG(nsGatewayCfg map[string]interface{}, parent *vspk.Enterprise) error {
	log.Infof("DeleteNSG started")

	nsGateways, err := parent.NSGateways(&bambou.FetchingInfo{
		Filter: nsGatewayCfg["Name"].(string)})
	handleError(err, "READ", "NS Gateway")

	// init the nsGateway struct that will hold either the received object
	// or will be created from the nsGatewayCfg
	nsGateway := &vspk.NSGateway{}

	if nsGateways != nil {
		log.Infof("NS Gateway already exists")

		nsGateway = nsGateways[0]
		nsGateway.Delete()
	} 
	log.Infof("DeleteNSG finished")
	return nil
}

// NSGRedundantGwGroup is a wrapper to create nuage NS Gateway redundant Group in a declaritive way
func NSGRedundantGwGroup(nsRedundantGwGroupCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.NSRedundantGatewayGroup {

	nsRedundantGwGroups, err := parent.NSRedundantGatewayGroups(&bambou.FetchingInfo{
		Filter: nsRedundantGwGroupCfg["Name"].(string)})
	handleError(err, "READ", "NS Redundant Gateway Group")

	// init the nsGateway struct that will hold either the received object
	// or will be created from the nsRedundantGwGroupCfg
	nsRedundantGwGroup := &vspk.NSRedundantGatewayGroup{}

	if nsRedundantGwGroups != nil {
		log.Infof("NS Gateway redudant group already exists")

		nsRedundantGwGroup = nsRedundantGwGroups[0]
		errMergo := mergo.Map(nsRedundantGwGroup, nsRedundantGwGroupCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		nsRedundantGwGroup.Save()
	} else {
		errMergo := mergo.Map(nsRedundantGwGroup, nsRedundantGwGroupCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateNSRedundantGatewayGroup(nsRedundantGwGroup)
		handleError(err, "CREATE", "NS Gateway redudant group")

		log.Infof("NS Gateway redudant group created")
	}

	log.Infof("%#v \n", nsRedundantGwGroup)
	return nsRedundantGwGroup
}

// ShuntLink is a wrapper to create a NSG shunt link in a declaritive way
func ShuntLink(shuntLinkCfg map[string]interface{}, parent *vspk.NSRedundantGatewayGroup) *vspk.ShuntLink {

	shuntLinks, err := parent.ShuntLinks(&bambou.FetchingInfo{
		Filter: shuntLinkCfg["Name"].(string)})
	handleError(err, "READ", "NSG shuntLink")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	shuntLink := &vspk.ShuntLink{}

	if shuntLinks != nil {
		log.Infof("NS shunt Link already exists")

		shuntLink = shuntLinks[0]
		errMergo := mergo.Map(shuntLink, shuntLinkCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		shuntLink.Save()
	} else {
		log.Infof("shuntLink: %#v \n", shuntLink)
		log.Infof("shuntLink: %#v \n", shuntLinkCfg)
		errMergo := mergo.Map(shuntLink, shuntLinkCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		log.Infof("shuntLink: %#v \n", shuntLink)
		log.Infof("shuntLink: %#v \n", shuntLinkCfg)
		err := parent.CreateShuntLink(shuntLink)
		handleError(err, "CREATE", "NS Shunt Link ")

		log.Infof("NS Shunt Link created")
	}

	log.Infof("%#v \n", shuntLink)
	return shuntLink
}

// NSGRedundantPort is a wrapper to create a NSG redundant Port in a declaritive way
func NSGRedundantPort(nsRedundantPortCfg map[string]interface{}, parent *vspk.NSRedundantGatewayGroup) *vspk.RedundantPort {

	nsRedundantPorts, err := parent.RedundantPorts(&bambou.FetchingInfo{
		Filter: nsRedundantPortCfg["Name"].(string)})
	handleError(err, "READ", "NSG Redundant Port")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsRedundantPortCfg
	nsRedundantPort := &vspk.RedundantPort{}

	if nsRedundantPorts != nil {
		log.Infof("NS redundant Port already exists")

		nsRedundantPort = nsRedundantPorts[0]
		errMergo := mergo.Map(nsRedundantPort, nsRedundantPortCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		nsRedundantPort.Save()
	} else {
		errMergo := mergo.Map(nsRedundantPort, nsRedundantPortCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		log.Infof(nsRedundantPortCfg)
		err := parent.CreateRedundantPort(nsRedundantPort)
		handleError(err, "CREATE", "NS Redundant Port ")

		log.Infof("NS Redundant Port created")
	}

	log.Infof("%#v \n", nsRedundantPort)

	return nsRedundantPort
}

// NSGPort is a wrapper to create a NSG Port in a declaritive way
func NSGPort(nsPortCfg map[string]interface{}, parent *vspk.NSGateway) *vspk.NSPort {

	nsPorts, err := parent.NSPorts(&bambou.FetchingInfo{
		Filter: nsPortCfg["Name"].(string)})
	handleError(err, "READ", "NSG Port")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	nsPort := &vspk.NSPort{}

	if nsPorts != nil {
		log.Infof("NS Port already exists")

		nsPort = nsPorts[0]
		errMergo := mergo.Map(nsPort, nsPortCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		nsPort.Save()
	} else {
		errMergo := mergo.Map(nsPort, nsPortCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		//log.Infof(nsPortCfg)
		err := parent.CreateNSPort(nsPort)
		handleError(err, "CREATE", "NS Port ")

		log.Infof("NS Port created")
	}

	log.Infof("%#v \n", nsPort)
	return nsPort
}

// DeleteNSGPort is a wrapper to create a NSG Port in a declaritive way
func DeleteNSGPort(nsPortCfg map[string]interface{}, parent *vspk.NSGateway) error {
	log.Infof("DeleteNSGPort started")
	nsPorts, err := parent.NSPorts(&bambou.FetchingInfo{
		Filter: nsPortCfg["Name"].(string)})
	handleError(err, "READ", "NSG Port")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	nsPort := &vspk.NSPort{}

	if nsPorts != nil {
		log.Infof("NS Port already exists")

		nsPort = nsPorts[0]
		nsPort.Delete()
	} 

	log.Infof("DeleteNSGPort finished")
	return nil
}

// GetNSGPort is a wrapper to create a NSG Port in a declaritive way
func GetNSGPort(nsPortCfg map[string]interface{}, parent *vspk.NSGateway) *vspk.NSPort {

	nsPorts, err := parent.NSPorts(&bambou.FetchingInfo{
		Filter: nsPortCfg["Name"].(string)})
	handleError(err, "READ", "NSG Port")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	nsPort := &vspk.NSPort{}

	if nsPorts == nil {
		log.Infof("NS Port does not exists")
		return nil
	} 

	nsPort = nsPorts[0]
	return nsPort
}

// NSGWirelessPort is a wrapper to create a NSG Wireless Port in a declaritive way
func NSGWirelessPort(nsPortCfg map[string]interface{}, parent *vspk.NSGateway) *vspk.WirelessPort {

	nsPorts, err := parent.WirelessPorts(&bambou.FetchingInfo{
		Filter: nsPortCfg["Name"].(string)})
	handleError(err, "READ", "Wireless Port")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	nsPort := &vspk.WirelessPort{}

	if nsPorts != nil {
		log.Infof("Wireless Port already exists")

		nsPort = nsPorts[0]
		errMergo := mergo.Map(nsPort, nsPortCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		nsPort.Save()
	} else {
		errMergo := mergo.Map(nsPort, nsPortCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		err := parent.CreateWirelessPort(nsPort)
		handleError(err, "CREATE", "Wireless Port ")

		log.Infof("Wireless Port created")
	}

	log.Infof("%#v \n", nsPort)
	return nsPort
}

// RedundantVlan is a wrapper to create a NSG VLAN in a declaritive way
func RedundantVlan(nsVlanCfg map[string]interface{}, parent *vspk.RedundantPort) *vspk.VLAN {

	log.Infof("VLAN Cfg: %#v \n", nsVlanCfg)
	nsVlans, err := parent.VLANs(&bambou.FetchingInfo{
		Filter: fmt.Sprintf("value == %d", nsVlanCfg["Value"])})

	handleError(err, "READ", "NSG VLAN")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the nsVlanCfg
	nsVlan := &vspk.VLAN{}

	log.Infof("VLANs %#v \n", nsVlans)

	if nsVlans != nil {
		log.Infof("NS VLAN RG already exists")

		nsVlan = nsVlans[0]
		errMergo := mergo.Map(nsVlan, nsVlanCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		nsVlan.Save()
	} else {
		errMergo := mergo.Map(nsVlan, nsVlanCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		//nsVlan.Value, _ = strconv.Atoi("0")
		//nsVlan.Value = 0
		log.Infof("VLAN: %#v ", nsVlan)
		log.Infof("Port: %#v ", parent)
		err := parent.CreateVLAN(nsVlan)
		handleError(err, "CREATE", "NS VLAN ")

		//log.Infof("NS VLAN RG created")
	}

	log.Infof("%#v \n", nsVlan)
	return nsVlan
}

// Vlan is a wrapper to create a NSG VLAN in a declaritive way
func Vlan(nsVlanCfg map[string]interface{}, parent *vspk.NSPort) *vspk.VLAN {

	log.Infof("VLAN Cfg: %#v \n", nsVlanCfg)
	nsVlans, err := parent.VLANs(&bambou.FetchingInfo{
		Filter: fmt.Sprintf("value == %d", nsVlanCfg["Value"])})

	handleError(err, "READ", "NSG VLAN")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the nsVlanCfg
	nsVlan := &vspk.VLAN{}

	log.Infof("VLANs %#v \n", nsVlans)

	if nsVlans != nil {
		flog.Infof("NS VLAN already exists")

		nsVlan = nsVlans[0]
		errMergo := mergo.Map(nsVlan, nsVlanCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		nsVlan.Save()
	} else {
		errMergo := mergo.Map(nsVlan, nsVlanCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		//nsVlan.Value, _ = strconv.Atoi("0")
		//nsVlan.Value = 0
		log.Infof("VLAN: %#v \n", nsVlan)
		log.Infof("Port: %#v \n", parent)
		err := parent.CreateVLAN(nsVlan)
		handleError(err, "CREATE", "NS VLAN ")

		log.Infof("NS VLAN created")
	}

	log.Infof("%#v \n", nsVlan)
	return nsVlan
}

// DeleteVlan is a wrapper to delete a NSG VLAN in a declaritive way
func DeleteVlan(nsVlanCfg map[string]interface{}, parent *vspk.NSPort) error {
	log.Infof("DeleteVlan started")
	log.Debugf("VLAN Cfg: %#v \n", nsVlanCfg)
	nsVlans, err := parent.VLANs(&bambou.FetchingInfo{
		Filter: fmt.Sprintf("value == %d", nsVlanCfg["Value"])})

	handleError(err, "READ", "NSG VLAN")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the nsVlanCfg
	nsVlan := &vspk.VLAN{}

	log.Debugf("VLANs %#v \n", nsVlans)

	if nsVlans != nil {
		log.Infof("NS VLAN already exists")

		nsVlan = nsVlans[0]
		nsVlan.Delete()
	} 
	log.Infof("DeleteVlan finished")
	return nil
}

// GetVlan is a wrapper to get a NSG VLAN in a declaritive way
func GetVlan(nsVlanCfg map[string]interface{}, parent *vspk.NSPort) *vspk.VLAN {
	log.Infof("GetVlan started")
	nsVlans, err := parent.VLANs(&bambou.FetchingInfo{
		Filter: fmt.Sprintf("value == %d", nsVlanCfg["Value"])})
	handleError(err, "READ", "NSG VLAN")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the nsVlanCfg
	nsVlan := &vspk.VLAN{}

	log.Debugf("VLANs %#v \n", nsVlans)

	if nsVlans == nil {
		log.Infof("NS VLAN already exists")
		return nil
	} 

	vlan = nsVlans[0]
	return nsVlan
}

// SSIDConnection is a wrapper to create a NSG SSID in a declaritive way
func SSIDConnection(ssidConnCfg map[string]interface{}, parent *vspk.WirelessPort) *vspk.SSIDConnection {

	ssidConns, err := parent.SSIDConnections(&bambou.FetchingInfo{
		Filter: ssidConnCfg["Name"].(string)})
	handleError(err, "READ", "SSiD Connection")

	ssidConn := &vspk.SSIDConnection{}

	if ssidConns != nil {
		log.Infof("SSiD Connection already exists")

		ssidConn = ssidConns[0]
		errMergo := mergo.Map(ssidConn, ssidConnCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		ssidConn.Save()
	} else {
		errMergo := mergo.Map(ssidConn, ssidConnCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateSSIDConnection(ssidConn)
		handleError(err, "CREATE", "SSiD Connection")

		log.Infof("SSiD Connection created")
	}

	log.Infof("%#v \n", ssidConn)
	return ssidConn
}

// UplinkConnection is a wrapper to create a NSG uplink connection in a declaritive way
func UplinkConnection(uplinkConnCfg map[string]interface{}, parent *vspk.VLAN) *vspk.UplinkConnection {

	uplinkConns, err := parent.UplinkConnections(&bambou.FetchingInfo{})

	handleError(err, "READ", "Uplink Connection")

	uplinkConn := &vspk.UplinkConnection{}

	if uplinkConns != nil {
		log.Infof("Uplink Connection already exists")

		uplinkConn = uplinkConns[0]
		errMergo := mergo.Map(uplinkConn, uplinkConnCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		uplinkConn.Save()
	} else {
		log.Infof("Uplink Connection Cfg: %#v \n", uplinkConnCfg)
		errMergo := mergo.Map(uplinkConn, uplinkConnCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		log.Infof("Uplink Connection: %#v \n", uplinkConn)
		log.Infof("Uplink Connection Cfg: %#v \n", uplinkConnCfg)

		err = parent.CreateUplinkConnection(uplinkConn)
		handleError(err, "CREATE", "Uplink Connection")

		log.Infof("Uplink Connection created")
	}

	log.Infof("%#v \n", uplinkConn)
	return uplinkConn
}

// CustomProperty is a wrapper to create a NSG Custom Property on a port in a declaritive way
func CustomProperty(customePropCfg map[string]interface{}, parent *vspk.UplinkConnection) *vspk.CustomProperty {

	customeProps, err := parent.CustomProperties(&bambou.FetchingInfo{
		Filter: customePropCfg["AttributeName"].(string)})
	handleError(err, "READ", "Custome Property")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the nsVlanCfg
	customeProp := &vspk.CustomProperty{}

	if customeProps != nil {
		log.Infof("Custom property already exists")

		customeProp = customeProps[0]
		errMergo := mergo.Map(customeProp, customePropCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		customeProp.Save()
	} else {
		errMergo := mergo.Map(customeProp, customePropCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		err := parent.CreateCustomProperty(customeProp)
		handleError(err, "CREATE", "Custome Property")

		log.Infof("Custome Property created")
	}

	log.Infof("%#v \n", customeProp)

	return customeProp
}
