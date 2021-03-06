> ### ⚠️ This is a proof of concept
>
> As this is a proof of concept,  it won't be supported by the k6 team.
> It may also break in the future as xk6 evolves. USE AT YOUR OWN RISK!
> Any issues with the tool should be raised [here](https://github.com/k6io/xk6-docker/issues).

</br>
</br>

<div align="center">

# xk6-docker
A k6 extension for using of Docker CLI in testing. Built for [k6](https://github.com/loadimpact/k6) using [xk6](https://github.com/k6io/xk6).

</div>

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download `xk6`:
  ```bash
  $ go get -u github.com/k6io/xk6
  ```

2. Build the binary:
  ```bash
  $ xk6 build --with github.com/k6io/xk6-docker
  ```

## Example

```javascript
import containers from 'k6/x/docker/containers';

export default function () {
  containers.listContainers().forEach((container) => {
    console.log(container.id.substring(0, 10)));
  })
}
```

Result output:

```bash
$ ./k6 run script.js

        /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: ../xk6-docker/script.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0000] 32c2786ccb
INFO[0000] 738b15d70f
INFO[0000] 5813101144
INFO[0000] b1fda43ce2

running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU

     data_received........: 0 B 0 B/s
     data_sent............: 0 B 0 B/s
     iteration_duration...: avg=9.64ms min=9.64ms med=9.64ms max=9.64ms p(90)=9.64ms p(95)=9.64ms
     iterations...........: 1   25.017512/s

```

Inspect examples folder for more details.
