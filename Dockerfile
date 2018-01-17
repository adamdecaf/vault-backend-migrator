FROM scratch
COPY bin/vault-backend-migrator-linux-amd64 /bin/vault-backend-migrator
ENTRYPOINT ["/bin/vault-backend-migrator"]
