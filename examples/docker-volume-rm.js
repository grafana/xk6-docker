import volumes from 'k6/x/docker/volumes';

export default function () {
  const name = 'VOLUME_NAME';
  const force = true
  volumes.remove(name, force);
  console.log(`Volume ${name} removed`);
}
