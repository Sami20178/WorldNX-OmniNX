# OmniNX integration status

## Implemented in this branch

- Go module for the server foundation.
- Minimal HTTP server entry point.
- `GET /health` endpoint.
- Explicit `foundation` stage response.
- Unimplemented root route returns HTTP 501 instead of pretending protocol support exists.

## Not implemented yet

- NEX/PRUDP transport.
- Authentication, account, matchmaking, gathering, or backend services.
- Nintendo-specific proprietary assets, certificates, keys, or binaries.
- Automatic copying of upstream Pretendo changes.

## Pretendo integration policy

Pretendo code must be reviewed file-by-file before integration. Preserve applicable license texts, copyright notices, and attribution. Upstream changes must be applied as auditable patches only after their exact source commit and compatibility are verified.

The `Buffer.Scan`/`QBuffer.Scan` and `PersistentGathering.ObjectID()` items remain review candidates; they are not claimed as integrated by this branch.
