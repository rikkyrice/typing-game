module local.packages/go

go 1.12

require (
	github.com/ibmdb/go_ibm_db v0.3.0
	local.packages/go/model v0.0.0-00010101000000-000000000000
)

replace local.packages/go/model => ./model
