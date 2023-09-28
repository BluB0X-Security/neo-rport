package capabilities

import (
	"github.com/IOTech17/neo-rport/server/chconfig"
	chshare "github.com/IOTech17/neo-rport/share"
	"github.com/IOTech17/neo-rport/share/models"
)

func NewServerCapabilities(cfg *chconfig.MonitoringConfig) *models.Capabilities {
	caps := models.Capabilities{
		ServerVersion:      chshare.BuildVersion,
		MonitoringVersion:  chshare.MonitoringVersion,
		IPAddressesVersion: chshare.IPAddressesVersion,
	}

	if !cfg.Enabled {
		caps.MonitoringVersion = 0
	}
	return &caps
}
