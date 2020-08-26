package types

// scavenge module event types
const (
	EventTypeCreateScavenge = "CreateScavenge"
	EventTypeCommitSolution = "CommitSolution"
	EventTypeSolveScavenge  = "SolveScavenge"
	EventTypeDeleteScavenge = "DeleteScavenge"

	AttributeDescription           = "description"
	AttributeSolution              = "solution"
	AttributeSolutionHash          = "solutionHash"
	AttributeReward                = "reward"
	AttributeScavenger             = "scavenger"
	AttributeSolutionScavengerHash = "solutionScavengerHash"

	AttributeValueCategory = ModuleName
)
