Pour les tests :


- lancer le server nats :
	- Dans un terminal : 
	  go run nats-server-2.1.0/main.go

- lancer le subscriber :
	- Dans un terminal :
	  go run subscriber.go

- lancer le fichier client.go avec une liste d'ID représentant les capteurs
	- Dans un terminal : 
	  DEVICE_IDS=1,2,3 go run publisher.go