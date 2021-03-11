import containers from 'k6/x/docker/containers';

export default function () {
  const containerId = 'YOUR_CONTAINER_ID';
  const data = containers.inspect(containerId);
  console.log(JSON.stringify(data, null, 2));
}
