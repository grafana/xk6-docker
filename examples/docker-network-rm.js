import networks from 'k6/x/docker/networks';

export default function () {
  const networkId = 'YOUR_NETWORK_ID'
  networks.remove(networkId);
  console.log(`${networkId} network removed`);
}
