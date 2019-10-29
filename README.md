# SmartPark

## Organisation

- `backend/`: Backend principal en Go connecté à une base postgres
- `web/`: Interface web en Python
- `infra/`: Playbook ansible pour déployer l'infrastructure

## Docker

Il y a des `Dockerfile` pour chacun des services.
Pour développer en local, il y a un fichier `docker-compose.yaml` lançant tous les services, avec NATS et Postgres.

`docker-compose up` pour tout lancer. L'interface devrait être accessible sur <http://localhost:8080>.

Les ports de NATS et Postgres sont exposés localement. Il est donc possible de juste lancer NATS et Postgres dans docker et faire tourner le reste localement.

 - `docker-compose up nats postgres`
 - `export DATABASE=postgres://postgres:postgres@localhost/postgres`
 - `export NATS_URL=nats://localhost:4222`
 - `cd backend && go run main.go`
 - `cd web && python run.py`
