package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/nornir-automation/gornir/pkg/gornir"
	"github.com/nornir-automation/gornir/pkg/plugins/connection"
	"github.com/nornir-automation/gornir/pkg/plugins/inventory"
	"github.com/nornir-automation/gornir/pkg/plugins/logger"
	"github.com/nornir-automation/gornir/pkg/plugins/output"
	"github.com/nornir-automation/gornir/pkg/plugins/runner"
	"github.com/pkg/errors"
	"os"
)

type RemoteCommand struct {
	Command string               // Command to execute
	Meta    *gornir.TaskMetadata // Task metadata

}


// Metadata returns the task metadata
func (t *RemoteCommand) Metadata() *gornir.TaskMetadata {
	return t.Meta
}

// RemoteCommandResults is the result of calling RemoteCommand
type RemoteCommandResults struct {
	Stdout []byte // Stdout written by the command
	Stderr []byte // Stderr written by the command
}

// String implemente Stringer interface
func (r RemoteCommandResults) String() string {
	return fmt.Sprintf("  - stdout: %s\n  - stderr: %s", r.Stdout, r.Stderr)
}

// Run runs a command on a remote device via ssh
func (t *RemoteCommand) Run(ctx context.Context, logger gornir.Logger, host *gornir.Host) (gornir.TaskInstanceResult, error) {
	conn, err := host.GetConnection("ssh")
	if err != nil {
		return RemoteCommandResults{}, errors.Wrap(err, "failed to retrieve connection")
	}
	sshConn := conn.(*connection.SSH)

	session, err := sshConn.Client.NewSession()
	if err != nil {
		return RemoteCommandResults{}, errors.Wrap(err, "failed to create session")
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr
	t.Command = fmt.Sprint("tftp 10.8.61.116 put vrpcfg.zip " + host.Hostname + ".zip")
	if err := session.Run(t.Command); err != nil {
		return RemoteCommandResults{}, errors.Wrap(err, "failed to execute command")
	}
	return RemoteCommandResults{Stdout: stdout.Bytes(), Stderr: stderr.Bytes()}, nil
}

func main() {

	log := logger.NewLogrus(false)

	file := "D:\\Python Project\\Nornir\\inventory\\hosts.yaml"
	plugin := inventory.FromYAML{HostsFile: file}
	inv, err := plugin.Create()
	if err != nil {
		log.Fatal(err)
	}

	gr := gornir.New().WithInventory(inv).WithLogger(log).WithRunner(runner.Parallel())

	//Open an SSH connection towards the devices
	results, err := gr.RunSync(
		context.Background(),
		&connection.SSHOpen{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// defer closing the SSH connection we just opened
	defer func() {
		results, err = gr.RunSync(
			context.Background(),
			&connection.SSHClose{},
		)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Following call is going to execute the task over all the hosts using the runner.Parallel runner.
	// Said runner is going to handle the parallelization for us. Gornir.RunS is also going to block
	// until the runner has completed executing the task over all the hosts
	results, err = gr.RunSync(
		context.Background(),
		&RemoteCommand{},
	)

	if err != nil {
		log.Fatal(err)
	}
	// next call is going to print the result on screen
	output.RenderResults(os.Stdout, results, "What is my ip?", true)

}
