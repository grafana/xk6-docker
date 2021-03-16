import volumes from 'k6/x/docker/volumes';

export default function () {
  const name = 'VOLUME_NAME';
  const options = {
    name: name
  }
  const result = volumes.create(options);
  console.log(`Volume created ${JSON.stringify(result)}`);
}
