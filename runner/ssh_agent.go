package runner

import (
	"context"
	"fmt"

	"github.com/efritz/ij/command"
	"github.com/efritz/ij/logging"
	"github.com/efritz/ij/scratch"
	"github.com/efritz/ij/util"
)

const (
	SSHAgentImage = "alpine:3.8"
	SocketPath    = "/usr/local/lib/ij-ssh-agent/sockets" // TODO - consider something different here
)

func startSSHAgent(
	runID string,
	scratch *scratch.ScratchSpace,
	containerLists *ContainerLists,
	logger logging.Logger,
) error {
	containerName, err := util.MakeID()
	if err != nil {
		return fmt.Errorf("failed to generate container id: %s", err.Error())
	}

	builder, err := sshAgentCommandBuilderFactory(
		runID,
		scratch,
		containerName,
	)

	if err != nil {
		return fmt.Errorf("failed to build command args: %s")
	}

	args, _, err := builder.Build()
	if err != nil {
		return fmt.Errorf("failed to build command args: %s")
	}

	containerLists.ContainerStopper.Add(containerName)

	_, _, err = command.NewRunner(logger).RunForOutput(
		context.Background(),
		args,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to start ssh-agent container: %s", err.Error())
	}

	// TODO - exec or something to change perms?
	return nil
}

func sshAgentCommandBuilderFactory(
	runID string,
	scratch *scratch.ScratchSpace,
	containerName string,
) (*command.Builder, error) {
	builder := command.NewBuilder([]string{
		"docker",
		"run",
		"--rm",
	}, nil)

	// TODO - take out some of this jazz
	path, err := scratch.WriteScript(fmt.Sprintf(`
		set -ex
		apk add openssh-client
		ssh-agent -D -a %s/socket
	`, SocketPath))

	if err != nil {
		return nil, err
	}

	mount := fmt.Sprintf(
		"%s:%s",
		path,
		ScriptPath,
	)

	builder.AddArgs(SSHAgentImage)
	builder.AddArgs(ScriptPath)
	builder.AddFlagValue("-v", mount)
	builder.AddFlagValue("-v", SocketPath)
	builder.AddFlagValue("--entrypoint", "/bin/sh")
	builder.AddFlagValue("--name", containerName)
	builder.AddFlag("-d")
	builder.AddFlagValue("--network", runID)

	return builder, nil
}
