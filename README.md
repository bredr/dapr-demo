# dapr-demo

Demonstrator to test various things in [dapr](dapr.io).

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

## Lessons learnt...

- The state storage in Dapr is very limited in its interface lacking query and list options on its key-value interface.
- The integration to pub-sub and http invocation are neat and straight forward.
- The distributed tracing is rich and complete.
- Developing locally with a local k8s option gives best local experience over the default local development experience as you are likely to require other state storage options for example. Skaffold.dev also gives some good options for local development.
