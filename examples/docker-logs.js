import containers from 'k6/x/docker/containers';

export default function () {
  const containerId = 'YOUR_CONTAINER_ID';
  const options = {
    show_stdout: true,
    // show_stderr: true,
    // since: '',
    // until: 10,
    // timestamp: true,
    // follow: false,
    tail: 20,
    // defaults: true
  }
  const result = containers.logs(containerId, options);
  console.log(JSON.stringify(result));
  // TODO: figure out why `console.log(result);`
  // returns empty string but it has length
}