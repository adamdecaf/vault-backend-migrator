FROM scratch
COPY bin/vault-backend-migrator-linux /bin/vault-backend-migrator
ENTRYPOINT ["/bin/vault-backend-migrator"]
