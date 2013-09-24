package hue

import (
	"testing"
)

func Test_GetConfiguration(t *testing.T) {
	user, stubServer := NewStubUser("get/username1/config.json", "username1")

	config, err := user.GetConfiguration()
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "GET", stubServer.method, "method is get")
	assertEqual(t, "/api/username1/config", stubServer.uri, "request uri")

	assertEqual(t, "Smartbridge 1", config.Name, "Name")
	assertEqual(t, "none", config.ProxyAddress, "ProxyAddress")
	assertEqual(t, uint16(0), *config.ProxyPort, "ProxyPort")
	assertEqual(t, "192.168.1.100", config.IpAddress, "IpAddress")
	assertEqual(t, "255.255.0.0", config.Netmask, "Netmask")
	assertEqual(t, "192.168.0.1", config.Gateway, "Gateway")
	assertEqual(t, false, *config.Dhcp, "Dhcp")

	softwareUpdate := config.SoftwareUpdate
	assertEqual(t, uint(1), *softwareUpdate.UpdateState, "UpdateState")
	assertEqual(t, "www.meethue.com/patchnotes/1453", softwareUpdate.Url, "Url")
	assertEqual(t, "This is a software update", softwareUpdate.Text, "Text")
	assertEqual(t, false, *softwareUpdate.Notify, "Notify")

	assertEqual(t, false, *config.LinkButton, "LinkButton")
	assertEqual(t, false, *config.PortalServices, "PortalServices")

	assertEqual(t, "2012-10-29T12:00:00", config.Utc, "Utc")
	assertNotNil(t, config.Whitelist, "config.Whitelist")
	assertEqual(t, "01003542", config.SoftwareVersion, "SoftwareVersion")
	assertEqual(t, "00:17:88:00:00:00", config.Mac, "Mac")
}
