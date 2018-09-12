// Package core implements the essential features.
package core

import (
	"./config"
	"./inbound"
	"./outbound"
)

// Initiate the server with config file
func InitServer(configFilePath string) {

	config := config.NewConfig(configFilePath)

	// New dispatcher without ClientBundle, ClientBundle must be initiated when server is running
	d := outbound.Dispatcher{
		PrimaryDNS:               config.PrimaryDNS,
		AlternativeDNS:           config.AlternativeDNS,
		OnlyPrimaryDNS:           config.OnlyPrimaryDNS,
		IPNetworkPrimaryList:     config.IPNetworkPrimaryList,
		IPNetworkAlternativeList: config.IPNetworkAlternativeList,
		DomainPrimaryList:        config.DomainPrimaryList,
		DomainAlternativeList:    config.DomainAlternativeList,
		RedirectIPv6Record:       config.IPv6UseAlternativeDNS,
		Hosts:                    config.Hosts,
		Cache:                    config.Cache,
	}

	s := &inbound.Server{
		BindAddress: config.BindAddress,
		Dispatcher:  d,
		MinimumTTL:  config.MinimumTTL,
		RejectQtype: config.RejectQtype,
	}

	s.Run()
}
