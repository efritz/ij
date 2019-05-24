package runner

import (
	"context"
	"fmt"
	"os/user"

	"github.com/efritz/ij/command"
	"github.com/efritz/ij/logging"
	"github.com/efritz/ij/scratch"
)

const (
	SSHAgentImage    = "efritz/ij-ssh-agent:latest"
	SocketVolumePath = "/tmp/ij/ssh-agent"
	SocketPath       = "/tmp/ij/ssh-agent/ssh-agent.sock"
)

func startSSHAgent(
	runID string,
	scratch *scratch.ScratchSpace,
	containerLists *ContainerLists,
	logger logging.Logger,
) error {
	containerName := fmt.Sprintf("%s-ssh-agent", runID)

	builder, err := sshAgentCommandBuilderFactory(
		runID,
		scratch,
		containerName,
	)

	if err != nil {
		return fmt.Errorf("failed to build command args: %s", err.Error())
	}

	args, _, err := builder.Build()
	if err != nil {
		return fmt.Errorf("failed to build command args: %s", err.Error())
	}

	containerLists.ContainerStopper.Add(containerName)

	_, errOutput, err := command.NewRunner(logger).RunForOutput(
		context.Background(),
		args,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to start ssh-agent container: %s, %s", err.Error(), errOutput)
	}

	_, errOutput, err = command.NewRunner(logger).RunForOutput(
		context.Background(),
		[]string{
			"docker",
			"exec",
			containerName,
			// TODO - perms may be an issue on windows
			"ssh-add",
		},
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to add ssh-keys: %s, %s", err.Error(), errOutput)
	}

	// TODO - need to run the ensure keys again but inside the container
	return nil
}

func sshAgentCommandBuilderFactory(
	runID string,
	scratch *scratch.ScratchSpace,
	containerName string,
) (*command.Builder, error) {
	current, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user (%s)", err.Error())
	}

	builder := command.NewBuilder([]string{
		"docker",
		"run",
		"--rm",
	}, nil)

	builder.AddArgs(SSHAgentImage)
	builder.AddFlagValue("-v", fmt.Sprintf("%s/.ssh:/root/.ssh", current.HomeDir))
	builder.AddFlagValue("--name", containerName)
	builder.AddFlag("-d")
	builder.AddFlagValue("--network", runID)

	return builder, nil
}
