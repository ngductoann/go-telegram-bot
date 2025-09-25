package entity

// // IPInfo represents the local and WAN IP addresses of a machine.
// type IPInfo struct {
// 	LocalIP  string `json:"local_ip"`
// 	PublicIP string `json:"public_ip"`
// }
//
// // IsValid validates the IP information
// func (ip *IPInfo) IsValid() bool {
// 	return ip.isValidIP(ip.LocalIP) && ip.isValidIP(ip.PublicIP)
// }
//
// // isValidIP checks if the provided string is a valid IP address
// func (ip *IPInfo) isValidIP(ipStr string) bool {
// 	return net.ParseIP(strings.TrimSpace(ipStr)) != nil
// }
//
// // String returns a formatted string representation of the IP information
// func (ip *IPInfo) String() string {
// 	return fmt.Sprintf("Local IP: %s, Public IP: %s", ip.LocalIP, ip.PublicIP)
// }
