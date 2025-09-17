package entity

// IPInfo represents the local and WAN IP addresses of a machine.
type IPInfo struct {
	LocalIP  string `json:"local_id"`
	PublicIP string `json:"public_ip"`
}
