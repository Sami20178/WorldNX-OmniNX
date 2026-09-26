# OmniNX Server

The OmniNX server is being assembled repository-by-repository.

## Current stage

1. Source manifest created.
2. Pretendo NEX sources identified for protocol/server integration.
3. NiklasCFW_Pack identified for repository review.
4. Go server foundation added in `cmd/omninx-server`.
5. `GET /health` endpoint added for basic service verification.
6. Pretendo changes remain review candidates until exact commits, compatibility, licenses, and attribution are checked.

## Run the foundation

From this directory:

```sh
go run ./cmd/omninx-server
```

Optional bind address:

```sh
OMNINX_ADDR=127.0.0.1:8080 go run ./cmd/omninx-server
```

Health check:

```sh
curl http://127.0.0.1:8080/health
```

The root route intentionally returns HTTP 501 until actual protocol services are implemented. This foundation does not claim NEX, PRUDP, authentication, matchmaking, or Nintendo service compatibility.

## Safety and licensing

Do not copy proprietary Nintendo files, keys, certificates, or unrelated binaries. Preserve upstream license texts, copyright notices, and attribution for any legally redistributable source that is later integrated.
