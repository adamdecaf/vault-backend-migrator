## vault-backend-migrator

`vault-backend-migrator` is a tool to export and import (migrate) data across vault clusters.

### Exports / Imports

Supported data exports and imports:

- `app-id` and `user-id` mappings
- `generic` backend data
- `policies`

### Configuration

This tool reads all the `VAULT_*` environment variables as the vault cli does. You likely need to specify those for the address, CA certs, etc.
