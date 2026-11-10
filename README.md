# idp-test-app-multiple

Monorepo test project for the Ruru platform's self-hosted multi-component CI.

| Path | Stack | Builds to |
|------|-------|-----------|
| `api/` | Go HTTP API (checks PostgreSQL reachability) | `registry.v2.rurueuh.fr/idp-test-app-multiple-api` |
| `front/` | Next.js frontend | `registry.v2.rurueuh.fr/idp-test-app-multiple-front` |
| `worker/` | Java background worker | `registry.v2.rurueuh.fr/idp-test-app-multiple-worker` |

Each subdirectory is built independently by the in-cluster BuildKit builder from
its own `source.path` and pushed to the in-cluster private registry.
toggle kaniko-path test
retry