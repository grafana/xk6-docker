import volumes from 'k6/x/docker/volumes';

export default function () {
  const volumesList = volumes.list();
  console.log(`${volumesList.volumes.length} volumes found:`)
  volumesList.volumes.forEach((vl) => console.log(JSON.stringify(vl, null, 2)))
}
