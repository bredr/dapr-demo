# dapr-demo

Demonstrator to test various things in [dapr](dapr.io)

## local development

Prerequisits:

- skaffold
- dapr
- minikube

Run `./setup.sh` to configure a local k8s cluster + dapr components.

Run `skaffold dev` to start the local development environment.

### Debugging

Run `skaffold debug` to switch to debugging mode.
Port forward `kubectl port-forward <pod> 56268:56268` to the relevant pod to expose debug port.
Attach debugger with task:

```json
{
  "name": "Skaffold Debug",
  "type": "go",
  "request": "attach",
  "mode": "remote",
  "host": "localhost",
  "port": 56268,
  "cwd": "${workspaceFolder}",
  "remotePath": "/go/src"
}
```
