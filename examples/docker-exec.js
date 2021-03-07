import containers from 'k6/x/docker/containers';

export default function () {
  const containerId = '2e0a7ff77434';
  const options = {
    // user:          '',                  // User that will run the command
    // privileged:    false,               // Is the container in privileged mode
    // yty:           false,               // Attach standard streams to a tty.
    // attach_stdin:  false,               // Attach the standard input, makes possible user interaction
    // attach_stderr: false,               // Attach the standard error
    // attach_stdout: false,               // Attach the standard output
    // detach:        false,               // Execute in detach mode
    // detach_keys:   '',                  // Escape keys for detach
    // env:           [''],                // Environment variables
    // working_dir:   '',                  // Working directory
    cmd: ['bundle', 'exec', 'rake', 'db:rollback']  // Execution commands and args
  }
  containers.execContainer(containerId, options);
}
