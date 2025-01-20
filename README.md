# Simple Bank 
An application for Banking

## Some commands
```bash
 # Generating 32 byte string
 openssl rand -hex 64 | head -c 32

 # Get ecr login password and login with docker to ecr
 aws ecr get-login-password | docker login --username AWS --password-stdin <base-uri> 

 # Generating a new migration
 migrate create -ext sql -dir db/migration -seq <migration-name>
```