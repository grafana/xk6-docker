> ### ⚠️ Deprecated!!!
>
> This extension was originally created as a _proof of concept_.
> At this time, there are no maintainers available to support this extension.
>
> USE AT YOUR OWN RISK!

<br />

# xk6-docker
A k6 extension for using of Docker CLI in testing. Built for [k6](https://go.k6.io/k6) using [xk6](https://github.com/grafana/xk6).

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download `xk6`:
  ```bash
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```bash
  $ xk6 build --with github.com/grafana/xk6-docker@latest
  ```

## Example

```javascript
import containers from 'k6/x/docker/containers';

export default function () {
  containers.list().forEach((container) => {
    console.log(container.id.substring(0, 10));
  });
}
```

Result output:

```plain
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

Inspect [`examples`](examples) folder for more details.
