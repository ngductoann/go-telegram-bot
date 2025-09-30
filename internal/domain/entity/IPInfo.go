package entity

// IPInfo represents IP address information
type IPInfo struct {
	LocalIP  string `json:"local_ip"`
	PublicIP string `json:"public_ip"`
}

// NewIPInfo creates a new IPInfo instance
func NewIPInfo(localIP, publicIP string) *IPInfo {
	return &IPInfo{
		LocalIP:  localIP,
		PublicIP: publicIP,
	}
}

// IsValid checks if both IP addresses are present
func (ip *IPInfo) IsValid() bool {
	return ip.LocalIP != "" && ip.PublicIP != ""
}
