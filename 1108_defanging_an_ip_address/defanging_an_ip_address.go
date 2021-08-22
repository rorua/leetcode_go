package defanging_an_ip_address

func defangIPaddr(address string) string {
	var result = ""
	for _, v := range address {
		if v == '.' {
			result += "[.]"
		} else {
			result += string(v)
		}
	}
	return result
}
