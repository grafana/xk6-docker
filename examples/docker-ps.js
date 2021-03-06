import docker from 'k6/x/docker';

function logContainer(container) {
  const line = [
    container.id.substring(0, 10),
    container.image,
    container.command,
    container.created,
    container.status,
    portsLabel(container.ports),
    container.names.join(', '),    
  ]
  console.log(line.join('; '))
}

function portsLabel(ports) {
  return ports.map((port) => `${port.ip}:${port.public_port}->${port.private_port}/${port.type}`).join(', ')
}

export default function () {
  const containers = docker.listContainers({all: false});
  console.log(`${containers.length} containers found:`)
  containers.forEach(logContainer)
}
