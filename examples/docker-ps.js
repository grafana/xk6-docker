import docker from 'k6/x/docker';

export default function () {
  console.log(docker.ps().join("\n"));
}