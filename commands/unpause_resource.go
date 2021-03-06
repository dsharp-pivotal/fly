package commands

import (
	"fmt"

	"github.com/concourse/fly/commands/internal/flaghelpers"
	"github.com/concourse/fly/rc"
)

type UnpauseResourceCommand struct {
	Resource flaghelpers.ResourceFlag `short:"r" long:"resource" required:"true" value-name:"PIPELINE/RESOURCE" description:"Name of a resource to unpause"`
}

func (command *UnpauseResourceCommand) Execute(args []string) error {
	target, err := rc.LoadTarget(Fly.Target)
	if err != nil {
		return err
	}
	err = target.Validate()
	if err != nil {
		return err
	}

	found, err := target.Team().UnpauseResource(command.Resource.PipelineName, command.Resource.ResourceName)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("pipeline '%s' or resource '%s' not found\n", command.Resource.PipelineName, command.Resource.ResourceName)
	}

	fmt.Printf("unpaused '%s'\n", command.Resource.ResourceName)
	return nil
}
