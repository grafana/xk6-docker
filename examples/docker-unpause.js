import containers from 'k6/x/docker/containers';

export default function () {
  const containerId = 'YOUR_CONTAINER_ID';
  containers.unpause(containerId);
}
