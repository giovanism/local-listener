# local-listener

Demo for LocalListener issue on Istio installation post v1.10

```
istioctl x precheck
Error [IST0143] (Pod local-listener-6fb8d8dcc6-tx9m7.istio-testing) Port 8080 is exposed in a Service but listens on localhost. It will not be exposed to other pods.
Error: Issues found when checking the cluster. Istio may not be safe to install or upgrade.
See https://istio.io/v1.10/docs/reference/config/analysis for more information about causes and resolutions.
```
