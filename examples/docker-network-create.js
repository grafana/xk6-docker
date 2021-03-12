import networks from 'k6/x/docker/networks';

export default function () {
  const networkName = 'YOUR_NETWORK_ID'
  const data = networks.create(networkName);
  console.log(`${JSON.stringify(data)} network created`);
}
