package main

import (
	"github.com/op/go-logging"
	"github.com/skycoin/darknet/src/network"
	"github.com/skycoin/darknet/src/network/linux"
	"net"
	"os"
)

var (
	logger     = logging.MustGetLogger("darknet.main")
	logFormat  = "[%{module}:%{level}] %{message}"
	logModules = []string{
		"darknet.network",
	}
)

func main() {
	initLogging(logging.DEBUG, true)

	networkmanager := linux.NewNetworkManager()

	useNetworkManager := false
	if useNetworkManager {
		networkmanager.ServiceStart()
	} else {
		networkmanager.ServiceStop()
	}

	connectionA := network.WifiConnection{
		Mode:             "managed",
		SSID:             "darkaccesspointA",
		SecurityProtocol: "WEP",
		SecurityKey:      "secret passwordA",
		DHCPEnabled:      true,
	}
	connectionB := network.WifiConnection{
		Mode:             "managed",
		SSID:             "darkaccesspointB",
		SecurityProtocol: "WPA",
		SecurityKey:      "secret passwordB",
		DHCPEnabled:      true,
	}
	connectionC := network.WifiConnection{
		Mode:             "managed",
		SSID:             "darkaccesspointC",
		SecurityProtocol: "WPA",
		SecurityKey:      "secret passwordC",
		DHCPEnabled:      true,
		Addresses:        []network.Address{},
		Routes:           []network.Route{},
		Nameservers:      []net.IP{},
		DefaultGateway:   net.IP{},
	}

	ifaces, _ := network.NewWifiInterfaces()
	if len(ifaces) > 0 {
		iface := ifaces[0]
		iface.Connection = connectionA
		iface.Start()

		iface.Stop()
		iface.Stop()
		iface.Stop()

		iface.Start()

		iface.Scan()
		iface.Stats()

		iface.Connection = connectionB
		iface.Start()

		iface.Connection = connectionC
		iface.Start()
	}
}

func initLogging(level logging.Level, color bool) {
	format := logging.MustStringFormatter(logFormat)
	logging.SetFormatter(format)
	for _, s := range logModules {
		logging.SetLevel(level, s)
	}
	stdout := logging.NewLogBackend(os.Stdout, "", 0)
	stdout.Color = color
	logging.SetBackend(stdout)
}
