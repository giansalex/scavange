package types

import "strings"

// query endpoints supported by the nameservice Querier
const (
	QueryListScavenges = "list"
	QueryGetScavenge   = "get"
	QueryCommit        = "commit"
	QueryListCommits   = "list-commits"
)

/*
Below you will be able how to set your own queries:
*/

// QueryResScavenges Queries Result Payload for a names query
type QueryResScavenges []string

// implement fmt.Stringer
func (n QueryResScavenges) String() string {
	return strings.Join(n[:], "\n")
}

// QueryResCommits Queries Result Payload for a names query
type QueryResCommits []string

// implement fmt.Stringer
func (n QueryResCommits) String() string {
	return strings.Join(n[:], "\n")
}
