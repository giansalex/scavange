package types

import "strings"

// query endpoints supported by the nameservice Querier
const (
	QueryListScavenges = "list"
	QueryGetScavenge   = "get"
	QueryCommit        = "commit"
)

/*
Below you will be able how to set your own queries:


// QueryResList Queries Result Payload for a query
type QueryResList []string

// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}

*/

// QueryResScavenges Queries Result Payload for a names query
type QueryResScavenges []string

// implement fmt.Stringer
func (n QueryResScavenges) String() string {
	return strings.Join(n[:], "\n")
}
