module local.packaegs/go

go 1.12

require (
	github.com/ibmdb/go_ibm_db v0.3.0
	github.com/pkg/errors v0.9.1
	gopkg.in/yaml.v2 v2.3.0
	local.packages/config v0.0.0-00010101000000-000000000000
	local.packages/db v0.0.0-00010101000000-000000000000
)

replace local.packages/db => ../common/db

replace local.packages/config => ../common/config
