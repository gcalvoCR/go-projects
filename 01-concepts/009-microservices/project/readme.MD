### Commands

'''sh
docker swarm init
docker stack deploy -c swarm.yml NAME
docker stack deploy -c swarm.yml myapp
docker service scale NAME_OF_SERVICE=3
docker service update --image gcalvocr/logger-service:1.0.1 NAME_OF_SERVICE
docker service update --image gcalvocr/logger-service:1.0.1 myapp_logger-service
docker stack rm NAME
docker stack rm myapp
docker swarm leave --force
'''