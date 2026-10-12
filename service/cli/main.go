/*
pjq is the command-line client for PigeonJQ. It talks to a running
pigeonjq-server over HTTP, at http://localhost:11300 by default; set
PIGEONJQ_ADDR to point it elsewhere.

Without arguments, pjq prints command help. With a command, it executes the
requested operation. Commands may define their own flags.

Usage:

	pjq <command> [command flags] [args...]

The available commands:

	status
	    Display the server status, optionally scoped to a tube with --tube.

	job-status [job_id]
	    Display the status of a single job.

	queue [body] --tube <tube>
	    Queue a job on a tube, with an optional --priority.

	cancel [job_id]
	    Cancel a queued job. Registered but not yet implemented.
*/
package main

import "pigeonjq/cli/cmd"

func main() {
	cmd.Execute()
}
