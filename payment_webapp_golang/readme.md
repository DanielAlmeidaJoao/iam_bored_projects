# To bring up the services, use Docker Compose:
sudo docker compose up -d

# To bring down the services, use:
sudo docker compose down

go mod init auth-service

docker exec -it <container_name_or_id> /bin/bash
docker exec -it <container_name_or_id> /bin/sh