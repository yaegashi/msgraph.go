package main

import (
	"fmt"
	"testing"
)

func TestLintName(t *testing.T) {
	tests := []struct {
		input, output string
	}{
		{"UserId", "UserID"},
		{"SkuId", "SKUID"},
		{"IpAddress", "IPAddress"},
		{"VpnDnsRule", "VPNDNSRule"},
		{"Oauth2PermissionScopes", "OAuth2PermissionScopes"},
		{"WindowsWifiConfiguration", "WindowsWiFiConfiguration"},
		{"SubscribedSkus", "SubscribedSKUs"},
		{"DefaultIosEnrollmentProfile", "DefaultIOSEnrollmentProfile"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d:%s", i+1, tt.input), func(t *testing.T) {
			output := lintName(tt.input)
			if tt.output != output {
				t.Errorf("Mismatch input %q want %q got %q", tt.input, tt.output, output)
			}
		})
	}
}
