# Usage

🌏 [**English**](./usage.md)
🌏 [**日本語**](./usage.ja.md)

> [!CAUTION]
>
> - **Using this container image will make your container host extremely vulnerable to attacks from malicious code.**
> - **Do not use this container image unless you understand what you are trying to do.**
> - **Do not use in production environment.**
> - **Do not use in Dify environment shared with others.**
> - **Use only in an environment that only you can use, for experiments and tests only.**

Replace `image` for `sandbox` service in your `docker-compose.yaml`.

```yaml
...
  sandbox:
    image: ghcr.io/kurokobo/dify-unsandboxed-sandbox:0.1.0
...
```

Then restart the `sandbox` service.

```bash
docker compose down sandbox
docker compose pull sandbox
docker compose up -d sandbox
```