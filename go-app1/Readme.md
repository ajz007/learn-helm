## Argocd

### Use Port Forwarding in Pods in Openlens to access argocd

### Argocd Admin credentials

```sh
kubectl -n argocd get secret argocd-initial-admin-secret \
  -o jsonpath="{.data.password}" | base64 -d
```