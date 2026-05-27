# Bad application

Simulates a misbehaving application to create signals for demo purposes.

- **Logs errors** and non-errors to stderr every 0.5–2.5 seconds with realistic error messages.
- **Crashes** after+ 30–90 seconds (randomized), triggering `CrashLoopBackOff`.
- **Burns CPU**(optional) on all available cores (tight math loops).
- **Hogs Memory**(optional) 
- **Emits Traces**: TODO

Together these produce correlated alerts, logs, events, and metrics that korrel8r can navigate.

## Quick start

Deploy the existing image:

```sh
make deploy
```

Open the console page for bad-app:

```sh
make browse
```
 

## Makefile targets

| Target     | Description                                         |
|------------|-----------------------------------------------------|
| `deploy`   | Create namespace `app-to-infra` and apply deployment |
| `undeploy` | Delete the `app-to-infra` namespace                 |
| `build`    | Build the container image                           |
| `push`     | Build and push the image to the registry            |
| `logs`     | Tail logs from the running pod                      |

Override the image with `make IMAGE=my-registry/bad-app:v1 build push deploy`.

## Observable signals

Once deployed and crash-looping, look for:

- **Alerts**: `KubePodCrashLooping`, `KubePodNotReady`, `CPUThrottlingHigh`
- **Logs**: error messages in pod stderr (database timeouts, TLS errors, OOM, etc.)
- **Events**: `BackOff`, `Unhealthy`, container exit events
- **Metrics**: high `container_cpu_usage_seconds_total`, restart counts
- **Traces**: TODO

## Files

| File              | Description                  |
|-------------------|------------------------------|
| `main.go`         | Application source           |
| `Dockerfile`      | Multi-stage container build  |
| `deployment.yaml` | Kubernetes deployment        |
| `Makefile`        | Build and deploy automation  |

### View in browser

https://console-openshift-console.apps.snoflake.home/k8s/ns/app-to-infra/apps~v1~Deployment/bad-app
