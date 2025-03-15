package magnet

func runFunc(name string, rule Rule, param PathParam) string {
	switch name {
	case "u3c3Search2":
		return u3c3Search2(rule, param)
	default:
		return ""
	}
}

func u3c3Search2(rule Rule, param PathParam) string {
	return param.Default
}
