# Backend

Le backend se connecte à une base de données Postgres et à NATS.
Il répond à des requêtes d'autres services via NATS.

## Structure

- `bus/` mets en place la connection à NATS
- `database/` gère la connection et les requêtes à la base de données
- `handlers/` contient les différents "handlers" pour les différents topics
- `main.go` est l'exécutable principal
