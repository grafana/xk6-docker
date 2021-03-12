import networks from 'k6/x/docker/networks';

export default function () {
  const networksList = networks.list();
  console.log(`${networksList.length} networks found:`);
  console.log(JSON.stringify(networksList, null, 2));
}
