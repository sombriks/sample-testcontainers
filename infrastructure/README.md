# The development database

A handy compose file to spin up a development database, so you don't have to
configure everything yourself.

## Requirements

- docker 25

## How to run

```bash
# on infrastructure folder
docker compose -f docker-compose.yml up -d
```

## Noteworthy

- Normally, proper schema creation and initial data feeding should be handled by
  some migration framework. This is intentionally NOT done in any of the sample
  projects by the sake of simplicity. Eyes on TestContainers and maybe on htmx,
  nothing else more.
- The `initial-state.sql` file was supposed to dwell in this folder, but due to 
  a jvm classpath technicality it was moved to sample jvm project. 
