package core

// Label of the nonpq build, shown by "xray version".
func init() {
	if build == "Custom" {
		build = "nonpq"
	} else {
		build = "nonpq-" + build
	}
}
