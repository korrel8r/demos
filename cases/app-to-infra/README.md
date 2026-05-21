# App-to-Infra Demo Case

A misbehaving application that generates observable signals for demonstrating
korrel8r's ability to correlate from application-level problems to infrastructure causes.

## What it does

The `bad-app` application:
- **Burns CPU** on all available cores (tight math loops).
- **Logs errors** to stderr every 0.5–2.5 seconds with realistic error messages.
- **Crashes** after 30–90 seconds (randomized), triggering `CrashLoopBackOff`.

Together these produce correlated alerts, logs, events, and metrics that
korrel8r can navigate between.

## Quick start

Build and push the container image, then deploy:

```sh
make build push deploy
```

Or deploy using the pre-built image (`quay.io/korrel8r/bad-app:latest`):

```sh
make deploy
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

## Files

| File              | Description                  |
|-------------------|------------------------------|
| `main.go`         | Application source           |
| `Dockerfile`      | Multi-stage container build  |
| `deployment.yaml` | Kubernetes deployment        |
| `Makefile`        | Build and deploy automation  |

### View in browser

https://console-openshift-console.apps.snoflake.home/k8s/ns/app-to-infra/apps~v1~Deployment/bad-app
