import containers from 'k6/x/docker/containers';

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
  const containersList = containers.list({all: false});
  console.log(`${containersList.length} containers found:`)
  containersList.forEach(logContainer)
}
