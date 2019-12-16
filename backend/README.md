# Backend

Le backend se connecte à une base de données Postgres et à NATS.
Il répond à des requêtes d'autres services via NATS.

## Structure

- `bus/` mets en place la connection à NATS
- `database/` gère la connection et les requêtes à la base de données
- `handlers/` contient les différents "handlers" pour les différents topics
- `main.go` est l'exécutable principal

# Database

- psql -h localhost --username="postgres" : permet de se connecter en local en ligne de commande
- \d : affichage de toutes les tables
- docker-compose up migrations : mise à jour database SQL de smart-park

# Handlers disponibles

Différents handlers sont disponibles afin d'obtenir, modifier, créer ou supprimer les données de la base de données.

- devices.get : renvoie un device via l'ID du device. Pour effectuer l'appel à ce handler : `devices.get '{"device_id":1}'`
- devices.get.notassigned : renvoie tous les devices qui sont libres, c'est-à-dire non-associé à une place. Pour effectuer l'appel à ce handler : `devices.get.notassigned  '{}'`
- tenants.get : renvoie un tenant via l'ID du tenant. Pour effectuer l'appel à ce handler :  `tenants.get  '{"tenants.get":1}'`
- zones.get : renvoie une zone via l'ID de la zone. Pour effectuer l'appel à ce handler : `zones.get '{"zone_id": 1}'`
- places.get : renvoie une place via l'ID de la place. Pour effectuer l'appel à ce handler : `places.get  '{"place_id": 1}'`
- users.get : renvoie un utilisateur via son ID. Pour effectuer l'appel à ce handler : `users.get  '{"users.get": 1}'`
- devices.list : renvoie tous les devices. Pour effectuer l'appel à ce handler : `devices.list '{"limit":20,"offset":0}'`
- tenants.list : renvoie tous les tenants. Pour effectuer l'appel à ce handler : `tenants.list '{"limit":20,"offset":0}'`
- zones.list : renvoie toutes les zones en fonction de l'ID du tenant où la zone est situé. Pour effectuer l'appel à ce handler : `zones.list '{"tenant_id": 1,"limit":20,"offset":0}'`
- places.list : renvoie toutes les places via l'ID du zone où la place est situé. Pour effectuer l'appel à ce handler : `places.list '{"zone_id": 1,"limit":20,"offset":0}'`
- users.list : renvoie tous les utilisateurs. Pour effectuer l'appel à ce handler : `users.list '{"limit":20,"offset":0}'`
- devices.update : permet de modifier les données d'un device. Pour effectuer l'appel à ce handler : `devices.update '{"device_id":1,"battery":1,"state":"free","deviceEUI":"00:0C:29:0C:47:D5","tenant_id":1}'`
- tenants.update : permet de modifier les données d'un tenant. Pour effectuer l'appel à ce handler : `tenants.update '{"tenant_id":1,"name":"Schmilbligheim","geo":"[7.7475,48.5827]"}'` 
- zones.update : permet de modifier les données d'une zone. Pour effectuer l'appel à ce handler : `zones.update '{"zone_id":1,"tenant_id":1,"name":"centre","type":"paid","color":"EB3434","geo":"[[7.739396,48.579816],[7.742014,48.579957],[7.744117,48.579134],[7.747464,48.578623],[7.74888,48.57885],[7.751756,48.579929],[7.755189,48.581831],[7.756906,48.583251],[7.754288,48.58555],[7.753558,48.586061],[7.751455,48.586743],[7.748537,48.58714],[7.746906,48.586828],[7.744503,48.585834],[7.740769,48.584244],[7.73901,48.582967],[7.738409,48.581973],[7.738495,48.580781],[7.739396,48.579816]]"}'` 
- places.update : permet de modifier les données d'une place. Pour effectuer l'appel à ce handler : `places.update '{"place_id":1,"type":"car","geo":"[7.746680,48.580402]","device_id":1}'` 
- users.update : permet de modifier les données d'un utilisateur. Pour effectuer l'appel à ce handler :  `users.update '{"user_id":1,"tenant_id":1,"username":"constantin","email":"divriotis.constantin@gmail.com","password":"sydney"}'`
- devices.new : crée une nouveau device. Pour effectuer l'appel à ce handler : `devices.new '{"battery":1,"state":"free","deviceEUI":"00:0C:29:0C:47:D5","tenant_id":1}'`
- places.new : crée une nouvelle place. Pour effectuer l'appel à ce handler : `places.new '{"type":"car","geo":"[7.746680,48.580402]","device_id":1}'`
- zones.new : crée une nouvelle zone. Pour effectuer l'appel à ce handler : `zones.new '{"tenant_id":1,"name":"centre","type":"paid","color":"EB3434","geo":"[[7.739396,48.579816],[7.742014,48.579957],[7.744117,48.579134],[7.747464,48.578623],[7.74888,48.57885],[7.751756,48.579929],[7.755189,48.581831],[7.756906,48.583251],[7.754288,48.58555],[7.753558,48.586061],[7.751455,48.586743],[7.748537,48.58714],[7.746906,48.586828],[7.744503,48.585834],[7.740769,48.584244],[7.73901,48.582967],[7.738409,48.581973],[7.738495,48.580781],[7.739396,48.579816]]"}'`
- faker.new	: permet d'inséser des données factises dans toutes les tables de la base de données. Pour effectuer l'appel à ce handler : `faker.new '{"tenants":1,"zones":1,"devices":1,"places":1,"users":1}'`

## Remarques générales

Certains paramètres sont optionnels lors des appels des handlers : 
- `limit` : par défaut, sa valeur est égale à 20
- `offset` : par défaut, sa valeur est égale à 0
- tous les paramètres des `.update` à l'exception de l'ID du handler appelé
- il est possible de choisir un seul ou plusieurs paramètres pour le `faker` (par exemple : `faker.new '{"tenants":1}'`)