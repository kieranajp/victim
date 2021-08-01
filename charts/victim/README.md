# Installation with Helm

- Create a namespace:

```sh
kubectl create ns victim
kns victim
```

- Create your secrets:

```sh
kubectl create secret generic victim-secrets \
  --from-literal='SLACK_APP_TOKEN=xapp-.....' \
  --from-literal='SLACK_BOT_TOKEN=xoxb.....'
```

- Helm install:

```sh
helm install victim . -f values.yaml --set version=$(git rev-parse --short HEAD)
```
