# jellyfin-k8s-helper

A helper utility for managing Jellyfin deployments on Kubernetes.

## Features

- Kubernetes manifests for RBAC and service accounts
- Simple Go-based command-line tool

## Getting Started

1. **Build the helper:**
   ```sh
   cd cmd
   go build -o jellyfin-k8s-helper main.go
   ```

2. **Apply Kubernetes manifests:**
   ```sh
   kubectl apply -f manifests/rbac.yaml
   kubectl apply -f manifests/serviceaccount.yaml
   ```

3. **Run the helper:**
   ```sh
   ./jellyfin-k8s-helper
   ```

## Project Structure

- `cmd/` - Go source code for the helper tool
- `manifests/` - Kubernetes YAML files for RBAC and service accounts
- `Dockerfile` - Container build instructions

## License

See [LICENSE](LICENSE) for for details.