import docker from 'k6/x/docker';

export default function () {
  const containerId = 'YOUR_CONTAINER_ID';
  docker.startContainer(containerId, {attach: true});
}
