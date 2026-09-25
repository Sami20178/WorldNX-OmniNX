# OmniNX+ Daily Development Plan

This plan is used to track daily, review-first development for WorldNX-OmniNX.

## Daily workflow

1. Review relevant Pretendo Network repositories and progress updates.
2. Review relevant NiklasCFW repositories and changes.
3. Compare findings with the current OmniNX+ repository state.
4. Select only changes that are relevant, legally redistributable, and compatible.
5. Implement changes in a dedicated branch.
6. Add tests or validation where practical.
7. Record the exact source repository, commit, license, and integration status.
8. Open or update a pull request instead of silently changing `main`.

## Safety and quality rules

- Do not copy Nintendo proprietary files, keys, certificates, firmware, or binaries.
- Do not blindly mirror upstream repositories.
- Preserve upstream licenses, copyright notices, and attribution.
- Use placeholders for private or live server URLs and shop endpoints.
- Mark incomplete protocol services clearly; do not claim the server is finished before it is tested.
- Every daily entry must state what changed and what was not implemented.

## Current baseline

- Server foundation branch: `omninx-server-foundation`
- Pull request: #2
- Current stage: HTTP foundation only
- Not yet implemented: NEX/PRUDP protocol services, authentication, matchmaking, gathering, and production backend services.

## Daily log

### 2026-09-25

- Created the server foundation with a Go module and a health endpoint.
- Added integration-status documentation and licensing boundaries.
- Kept the work in a draft pull request rather than merging incomplete server functionality into `main`.
- Next review: inspect upstream Pretendo protocol/server changes and select one small, testable integration task.
