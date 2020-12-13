package nuagewrapper

import (
	"fmt"
	"log"
	"strings"

	"github.com/imdario/mergo"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
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

// NuageCreateEntireNSG is a wrapper to create a complete NSG in a declaritive way
func NuageCreateEntireNSG(nsgCfg NuageNSGCfg, parent *vspk.Enterprise, Usr *vspk.Me) *vspk.NSGateway {
	fmt.Println("########################################")
	fmt.Println("#####  Create Entire NSG GW   ##########")
	fmt.Println("########################################")

	nsGatewayTemplateCfg := map[string]interface{}{
		"Name": nsgCfg.NSGTemplateName,
	}
	//fmt.Printf("NSG Template ID: %s \n", nsgCfg.NSGTemplateID)

	nsGatewayTemplate := NuageNSGatewayTemplate(nsGatewayTemplateCfg, Usr)
	nsgCfg.NSGTemplateID = nsGatewayTemplate.ID
	//fmt.Printf("NSG Template ID: %s \n", nsgCfg.NSGTemplateID)

	fmt.Printf("NSG Template Name: %s \n", nsgCfg.NSGTemplateName)
	fmt.Printf("NSG Template Personality: %s \n", nsGatewayTemplate.Personality)
	fmt.Printf("NSG Template ID: %s \n", nsgCfg.NSGTemplateID)

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
		nsGateway = NuageNSGRoot(nsGatewayCfg, Usr)
	} else {
		nsGateway = NuageNSG(nsGatewayCfg, parent)
	}

	//time.Sleep(15 * time.Second)

	for i, port := range nsgCfg.NetworkPorts {
		fmt.Printf("NSG Network Port %d Name: %s \n", i, port.Name)

		nsPortCfg := map[string]interface{}{
			"Name":            port.Name,
			"PhysicalName":    port.Name,
			"PortType":        "NETWORK",
			"VLANRange":       "0-4094",
			"EnableNATProbes": true,
			"NATTraversal":    "FULL_NAT",
		}
		nsPort := NuageNSGPort(nsPortCfg, nsGateway)

		fmt.Printf("Port: %#v \n", nsPort)

		for _, vlan := range port.Vlan {
			var nsVlanCfg map[string]interface{}
			if vlan.VscName != "" {
				vscProfCfg := map[string]interface{}{
					"Name": vlan.VscName,
				}
				vscProf := NuageInfraVscProfile(vscProfCfg, Usr)
				vlan.VscID = vscProf.ID

				//fmt.Printf("VSC NEEDED for vlan: %s with VSC ID: %s", vlan.VlanID, vlan.VscID)
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
				//fmt.Printf("NOOOOOOOOOOOOOOO VSC NEEDED for vlan: %s with VSC ID: %s", vlan.VlanID, vlan.VscID)
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

			fmt.Printf("VLANCfg: %#v \n", nsVlanCfg)
			nsVlan := NuageVlan(nsVlanCfg, nsPort)
			fmt.Printf("Port: %#v \n", nsVlan)

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
				fmt.Printf("BRANCH IPV4 AND STATIC ADDRESSING MODE \n")
			} else if vlan.AddressFamily == "IPV6" && vlan.Mode == "Static" {
				dnsV6 = vlan.DNS
				addressV6 = vlan.Address
				gatewayV6 = vlan.Gateway
				patEnabled = false
				underlayEnabled = false
				fmt.Printf("BRANCH IPV6 AND STATIC ADDRESSING MODE \n")
			} else if vlan.AddressFamily == "IPV6" {
				patEnabled = false
				underlayEnabled = false
				fmt.Printf("BRANCH IPV6 AND DYNAMIC ADDRESSING MODE \n")
			}

			var uplinkConnectionCfg map[string]interface{}
			if vlan.UnderlayName != "" {
				underlayCfg := map[string]interface{}{
					"Name": vlan.UnderlayName,
				}
				underlay := NuageUnderlay(underlayCfg, Usr)
				fmt.Println(underlay)
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

			fmt.Println(uplinkConnectionCfg)

			uplinkConn := NuageUplinkConnection(uplinkConnectionCfg, nsVlan)

			if strings.Contains(port.Name, "lte") {
				fmt.Println("LTE")

				customePropCfg := map[string]interface{}{
					"AttributeName":  "apn",
					"AttributeValue": vlan.LteConfiguration.Apn,
				}

				NuageCustomProperty(customePropCfg, uplinkConn)

				customePropCfg = map[string]interface{}{
					"AttributeName":  "pdp-type",
					"AttributeValue": vlan.LteConfiguration.PdpType,
				}

				NuageCustomProperty(customePropCfg, uplinkConn)

				customePropCfg = map[string]interface{}{
					"AttributeName":  "sim-pin",
					"AttributeValue": vlan.LteConfiguration.PinCode,
				}

				NuageCustomProperty(customePropCfg, uplinkConn)

				if vlan.LteConfiguration.Authentication != "" {
					customePropCfg = map[string]interface{}{
						"AttributeName":  "authentication",
						"AttributeValue": vlan.LteConfiguration.Authentication,
					}

					NuageCustomProperty(customePropCfg, uplinkConn)
				}

				if vlan.LteConfiguration.UserName != "" {
					customePropCfg = map[string]interface{}{
						"AttributeName":  "username",
						"AttributeValue": vlan.LteConfiguration.UserName,
					}

					NuageCustomProperty(customePropCfg, uplinkConn)
				}

				if vlan.LteConfiguration.PassWord != "" {
					customePropCfg = map[string]interface{}{
						"AttributeName":  "password",
						"AttributeValue": vlan.LteConfiguration.PassWord,
					}

					NuageCustomProperty(customePropCfg, uplinkConn)
				}

			} else {
				fmt.Println("ETHERNET")
			}
		}

	}
	for i, port := range nsgCfg.ShuntPorts {
		fmt.Printf("NSG Shunt Port %d Name: %s \n", i, port.Name)

		nsPortCfg := map[string]interface{}{
			"Name":            port.Name,
			"PhysicalName":    port.Name,
			"PortType":        "NETWORK",
			"VLANRange":       "0-4094",
			"EnableNATProbes": true,
			"NATTraversal":    "FULL_NAT",
			"Mtu":             2000,
		}
		nsPort := NuageNSGPort(nsPortCfg, nsGateway)

		for _, vlan := range port.Vlan {
			nsVlanCfg := map[string]interface{}{
				"Value":       vlan.VlanID,
				"Name":        "shunt",
				"Description": "shunt",
				"ShuntVLAN":   true,
			}

			fmt.Printf("VLANCfg: %s \n", nsVlanCfg)
			nsVlan := NuageVlan(nsVlanCfg, nsPort)
			fmt.Println(nsVlan)

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

			uplinkConn := NuageUplinkConnection(uplinkConnCfg, nsVlan)
			fmt.Println(uplinkConn)
		}

	}
	for i, port := range nsgCfg.AccessPorts {
		fmt.Printf("NSG Access Port %d Name: %s \n", i, port.Name)

		nsPortCfg := map[string]interface{}{
			"Name":         port.Name,
			"PhysicalName": port.Name,
			"PortType":     "ACCESS",
			"VLANRange":    "0-4094",
		}
		nsPort := NuageNSGPort(nsPortCfg, nsGateway)

		for _, vlan := range port.Vlan {
			nsVlanCfg := map[string]interface{}{
				"Value": vlan.VlanID,
			}
			fmt.Printf("Access VLANCfg: %#v \n", nsVlanCfg)
			nsVlan := NuageVlan(nsVlanCfg, nsPort)
			fmt.Println(nsVlan)
		}

	}
	for i, port := range nsgCfg.WifiPorts {
		fmt.Printf("NSG Wifi Port %d Name: %s \n", i, port.Name)

		nsPortCfg := map[string]interface{}{
			"Name":              port.Name,
			"WifiFrequencyBand": "FREQ_2_4_GHZ",
			"WifiMode":          "WIFI_B_G_N",
			"CountryCode":       "BE",
		}
		nsPort := NuageNSGWirelessPort(nsPortCfg, nsGateway)

		ssidConnCfg := map[string]interface{}{
			"Name":               port.Ssid,
			"Passphrase":         "4no*heydQ",
			"AuthenticationMode": "WPA2",
			"BroadcastSSID":      true,
		}
		ssidConn := NuageSSIDConnection(ssidConnCfg, nsPort)
		fmt.Println(ssidConn)
	}

	fmt.Printf("%#v \n", nsGateway)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsGateway
}

// NuageNSGatewayTemplate is a wrapper to create nuage NS Gateway template in a declaritive way
func NuageNSGatewayTemplate(nsGatewayTemplateCfg map[string]interface{}, parent *vspk.Me) *vspk.NSGatewayTemplate {
	fmt.Println("########################################")
	fmt.Println("#####     NSGateway template   #########")
	fmt.Println("########################################")

	nsGatewayTemplate := &vspk.NSGatewayTemplate{}

	nsGatewayTemplates, err := parent.NSGatewayTemplates(&bambou.FetchingInfo{
		Filter: nsGatewayTemplateCfg["Name"].(string)})
	handleError(err, "nsGatewayTemplate", "READ")

	// init the struct that will hold either the received object
	// or will be created from the Cfg object
	if nsGatewayTemplates != nil {
		fmt.Println("nsGatewayTemplate already exists")

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

		fmt.Println("nsGatewayTemplate created")
	}
	fmt.Printf("%#v \n", nsGatewayTemplate)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsGatewayTemplate
}

// NuageNSGRoot is a wrapper to create nuage NS Gateway in a declaritive way
func NuageNSGRoot(nsGatewayCfg map[string]interface{}, parent *vspk.Me) *vspk.NSGateway {
	fmt.Println("########################################")
	fmt.Println("#####        NSG Gateway      ##########")
	fmt.Println("########################################")

	nsGateways, err := parent.NSGateways(&bambou.FetchingInfo{
		Filter: nsGatewayCfg["Name"].(string)})
	handleError(err, "READ", "NS Gateway")

	// init the nsGateway struct that will hold either the received object
	// or will be created from the nsGatewayCfg
	nsGateway := &vspk.NSGateway{}

	if nsGateways != nil {
		fmt.Println("NS Gateway already exists")

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
		//fmt.Printf("nsGateway: %#v", nsGateway)
		//time.Sleep(15 * time.Second)

		err := parent.CreateNSGateway(nsGateway)
		handleError(err, "CREATE", "NS Gateway ")

		fmt.Println("NS Gateway created")
	}

	fmt.Printf("%#v \n", nsGateway)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsGateway
}

// NuageNSG is a wrapper to create nuage NS Gateway in a declaritive way
func NuageNSG(nsGatewayCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.NSGateway {
	fmt.Println("########################################")
	fmt.Println("#####        NSG Gateway      ##########")
	fmt.Println("########################################")

	nsGateways, err := parent.NSGateways(&bambou.FetchingInfo{
		Filter: nsGatewayCfg["Name"].(string)})
	handleError(err, "READ", "NS Gateway")

	// init the nsGateway struct that will hold either the received object
	// or will be created from the nsGatewayCfg
	nsGateway := &vspk.NSGateway{}

	if nsGateways != nil {
		fmt.Println("NS Gateway already exists")

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

		fmt.Println("NS Gateway created")
	}

	fmt.Printf("%#v \n", nsGateway)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsGateway
}

// NuageNSGRedundantGwGroup is a wrapper to create nuage NS Gateway redundant Group in a declaritive way
func NuageNSGRedundantGwGroup(nsRedundantGwGroupCfg map[string]interface{}, parent *vspk.Enterprise) *vspk.NSRedundantGatewayGroup {
	fmt.Println("############################################")
	fmt.Println("##### NSG redundant Gateway Group ##########")
	fmt.Println("############################################")

	nsRedundantGwGroups, err := parent.NSRedundantGatewayGroups(&bambou.FetchingInfo{
		Filter: nsRedundantGwGroupCfg["Name"].(string)})
	handleError(err, "READ", "NS Redundant Gateway Group")

	// init the nsGateway struct that will hold either the received object
	// or will be created from the nsRedundantGwGroupCfg
	nsRedundantGwGroup := &vspk.NSRedundantGatewayGroup{}

	if nsRedundantGwGroups != nil {
		fmt.Println("NS Gateway redudant group already exists")

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

		fmt.Println("NS Gateway redudant group created")
	}

	fmt.Printf("%#v \n", nsRedundantGwGroup)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsRedundantGwGroup
}

// NuageShuntLink is a wrapper to create a NSG shunt link in a declaritive way
func NuageShuntLink(shuntLinkCfg map[string]interface{}, parent *vspk.NSRedundantGatewayGroup) *vspk.ShuntLink {
	fmt.Println("########################################")
	fmt.Println("#####        NSG shunt Link   ##########")
	fmt.Println("########################################")

	shuntLinks, err := parent.ShuntLinks(&bambou.FetchingInfo{
		Filter: shuntLinkCfg["Name"].(string)})
	handleError(err, "READ", "NSG shuntLink")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	shuntLink := &vspk.ShuntLink{}

	if shuntLinks != nil {
		fmt.Println("NS shunt Link already exists")

		shuntLink = shuntLinks[0]
		errMergo := mergo.Map(shuntLink, shuntLinkCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		shuntLink.Save()
	} else {
		fmt.Printf("shuntLink: %#v \n", shuntLink)
		fmt.Printf("shuntLink: %#v \n", shuntLinkCfg)
		errMergo := mergo.Map(shuntLink, shuntLinkCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		fmt.Printf("shuntLink: %#v \n", shuntLink)
		fmt.Printf("shuntLink: %#v \n", shuntLinkCfg)
		err := parent.CreateShuntLink(shuntLink)
		handleError(err, "CREATE", "NS Shunt Link ")

		fmt.Println("NS Shunt Link created")
	}

	fmt.Printf("%#v \n", shuntLink)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return shuntLink
}

// NuageNSGRedundantPort is a wrapper to create a NSG redundant Port in a declaritive way
func NuageNSGRedundantPort(nsRedundantPortCfg map[string]interface{}, parent *vspk.NSRedundantGatewayGroup) *vspk.RedundantPort {
	fmt.Println("########################################")
	fmt.Println("#####   NSG Redundant Port    ##########")
	fmt.Println("########################################")

	nsRedundantPorts, err := parent.RedundantPorts(&bambou.FetchingInfo{
		Filter: nsRedundantPortCfg["Name"].(string)})
	handleError(err, "READ", "NSG Redundant Port")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsRedundantPortCfg
	nsRedundantPort := &vspk.RedundantPort{}

	if nsRedundantPorts != nil {
		fmt.Println("NS redundant Port already exists")

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
		fmt.Println(nsRedundantPortCfg)
		err := parent.CreateRedundantPort(nsRedundantPort)
		handleError(err, "CREATE", "NS Redundant Port ")

		fmt.Println("NS Redundant Port created")
	}

	fmt.Printf("%#v \n", nsRedundantPort)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsRedundantPort
}

// NuageNSGPort is a wrapper to create a NSG Port in a declaritive way
func NuageNSGPort(nsPortCfg map[string]interface{}, parent *vspk.NSGateway) *vspk.NSPort {
	fmt.Println("########################################")
	fmt.Println("#####        NSG Port         ##########")
	fmt.Println("########################################")

	nsPorts, err := parent.NSPorts(&bambou.FetchingInfo{
		Filter: nsPortCfg["Name"].(string)})
	handleError(err, "READ", "NSG Port")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	nsPort := &vspk.NSPort{}

	if nsPorts != nil {
		fmt.Println("NS Port already exists")

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
		fmt.Println(nsPortCfg)
		err := parent.CreateNSPort(nsPort)
		handleError(err, "CREATE", "NS Port ")

		fmt.Println("NS Port created")
	}

	fmt.Printf("%#v \n", nsPort)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsPort
}

// NuageNSGWirelessPort is a wrapper to create a NSG Wireless Port in a declaritive way
func NuageNSGWirelessPort(nsPortCfg map[string]interface{}, parent *vspk.NSGateway) *vspk.WirelessPort {
	fmt.Println("########################################")
	fmt.Println("#####  NSG Wireless Port      ##########")
	fmt.Println("########################################")

	nsPorts, err := parent.WirelessPorts(&bambou.FetchingInfo{
		Filter: nsPortCfg["Name"].(string)})
	handleError(err, "READ", "Wireless Port")

	// init the nsPort struct that will hold either the received object
	// or will be created from the nsPortCfg
	nsPort := &vspk.WirelessPort{}

	if nsPorts != nil {
		fmt.Println("Wireless Port already exists")

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

		fmt.Println("Wireless Port created")
	}

	fmt.Printf("%#v \n", nsPort)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsPort
}

// NuageredundantVlan is a wrapper to create a NSG VLAN in a declaritive way
func NuageredundantVlan(nsVlanCfg map[string]interface{}, parent *vspk.RedundantPort) *vspk.VLAN {
	fmt.Println("########################################")
	fmt.Println("#####        NSG Vlan         ##########")
	fmt.Println("########################################")

	fmt.Printf("VLAN Cfg: %#v \n", nsVlanCfg)
	nsVlans, err := parent.VLANs(&bambou.FetchingInfo{
		Filter: fmt.Sprintf("value == %d", nsVlanCfg["Value"])})

	handleError(err, "READ", "NSG VLAN")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the nsVlanCfg
	nsVlan := &vspk.VLAN{}

	fmt.Printf("VLANs %#v \n", nsVlans)

	if nsVlans != nil {
		fmt.Println("NS VLAN RG already exists")

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
		fmt.Printf("VLAN: %#v \n", nsVlan)
		fmt.Printf("Port: %#v \n", parent)
		err := parent.CreateVLAN(nsVlan)
		handleError(err, "CREATE", "NS VLAN ")

		fmt.Println("NS VLAN RG created")
	}

	fmt.Printf("%#v \n", nsVlan)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsVlan
}

// NuageVlan is a wrapper to create a NSG VLAN in a declaritive way
func NuageVlan(nsVlanCfg map[string]interface{}, parent *vspk.NSPort) *vspk.VLAN {
	fmt.Println("########################################")
	fmt.Println("#####        NSG Vlan         ##########")
	fmt.Println("########################################")

	fmt.Printf("VLAN Cfg: %#v \n", nsVlanCfg)
	nsVlans, err := parent.VLANs(&bambou.FetchingInfo{
		Filter: fmt.Sprintf("value == %d", nsVlanCfg["Value"])})

	handleError(err, "READ", "NSG VLAN")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the nsVlanCfg
	nsVlan := &vspk.VLAN{}

	fmt.Printf("VLANs %#v \n", nsVlans)

	if nsVlans != nil {
		fmt.Println("NS VLAN already exists")

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
		fmt.Printf("VLAN: %#v \n", nsVlan)
		fmt.Printf("Port: %#v \n", parent)
		err := parent.CreateVLAN(nsVlan)
		handleError(err, "CREATE", "NS VLAN ")

		fmt.Println("NS VLAN created")
	}

	fmt.Printf("%#v \n", nsVlan)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return nsVlan
}

// NuageSSIDConnection is a wrapper to create a NSG SSID in a declaritive way
func NuageSSIDConnection(ssidConnCfg map[string]interface{}, parent *vspk.WirelessPort) *vspk.SSIDConnection {
	fmt.Println("########################################")
	fmt.Println("#####   SSID Connection       ##########")
	fmt.Println("########################################")

	ssidConns, err := parent.SSIDConnections(&bambou.FetchingInfo{
		Filter: ssidConnCfg["Name"].(string)})
	handleError(err, "READ", "SSiD Connection")

	ssidConn := &vspk.SSIDConnection{}

	if ssidConns != nil {
		fmt.Println("SSiD Connection already exists")

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

		fmt.Println("SSiD Connection created")
	}

	fmt.Printf("%#v \n", ssidConn)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return ssidConn
}

// NuageUplinkConnection is a wrapper to create a NSG uplink connection in a declaritive way
func NuageUplinkConnection(uplinkConnCfg map[string]interface{}, parent *vspk.VLAN) *vspk.UplinkConnection {
	fmt.Println("########################################")
	fmt.Println("#####   NSG Uplink Connection ##########")
	fmt.Println("########################################")

	uplinkConns, err := parent.UplinkConnections(&bambou.FetchingInfo{})

	handleError(err, "READ", "Uplink Connection")

	uplinkConn := &vspk.UplinkConnection{}

	if uplinkConns != nil {
		fmt.Println("Uplink Connection already exists")

		uplinkConn = uplinkConns[0]
		errMergo := mergo.Map(uplinkConn, uplinkConnCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}

		uplinkConn.Save()
	} else {
		fmt.Printf("Uplink Connection Cfg: %#v \n", uplinkConnCfg)
		errMergo := mergo.Map(uplinkConn, uplinkConnCfg, mergo.WithOverride)
		if errMergo != nil {
			log.Fatal(errMergo)
		}
		fmt.Printf("Uplink Connection: %#v \n", uplinkConn)
		fmt.Printf("Uplink Connection Cfg: %#v \n", uplinkConnCfg)

		err = parent.CreateUplinkConnection(uplinkConn)
		handleError(err, "CREATE", "Uplink Connection")

		fmt.Println("Uplink Connection created")
	}

	fmt.Printf("%#v \n", uplinkConn)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return uplinkConn
}

// NuageCustomProperty is a wrapper to create a NSG Custom Property on a port in a declaritive way
func NuageCustomProperty(customePropCfg map[string]interface{}, parent *vspk.UplinkConnection) *vspk.CustomProperty {
	fmt.Println("########################################")
	fmt.Println("#####   Custom property       ##########")
	fmt.Println("########################################")

	customeProps, err := parent.CustomProperties(&bambou.FetchingInfo{
		Filter: customePropCfg["AttributeName"].(string)})
	handleError(err, "READ", "Custome Property")

	// init the nsVlan struct that will hold either the received object
	// or will be created from the nsVlanCfg
	customeProp := &vspk.CustomProperty{}

	if customeProps != nil {
		fmt.Println("Custom property already exists")

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

		fmt.Println("Custome Property created")
	}

	fmt.Printf("%#v \n", customeProp)
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	fmt.Println("****************************************")
	return customeProp
}
