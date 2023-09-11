package types

import "context"

type Driver interface {
	// Create a new instance of the scenario
	Create(Scenario) (Instance, error)
	// Run a new instance
	// Pass it the instance and the play id
	Run(context.Context, Instance, int64) error
	// Get the connection command
	ConnectionCommand(Instance) string
	// Kill and remove the instance
	// Return the play id and an error
	Kill(context.Context, Instance) (int64, error)
	// Check the work in the instance
	// Return the play id and an error
	Check(context.Context, Instance) (int64, bool)
}
