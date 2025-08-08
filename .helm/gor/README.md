# gor helm chart

Example usage:

```bash
helm upgrade gor .helm/gor --install \
    --namespace gor --create-namespace
```

To enable tailscale ingress:

```bash
helm upgrade gor .helm/gor --install \
    --namespace gor --create-namespace \
    --set ingress.enabled=true \
    --set ingress.ingressClassName=tailscale
```
