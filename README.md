# go-network-scanner

Ce projet est un **scanner de ports TCP simple** écrit en Go. Il permet d'identifier quels ports sont ouverts sur une cible réseau.

## Fonctionnalités 

-Scan de ports TCP de 0 à 99
-Connexions non-bloquantes grâce à l'utilisation des goroutines
-Affichage en temps réel des ports ouverts.

##Concepts utilisés

-Programmation concurrente en Go (via 'goroutines')
-Communication entre routines avec les 'channels'
-Utilisation de la librairie 'ne' pour gérer les connexions TCP 
-Timeout de connexion avec 'net.DialTimeout'

##Utilisation

1.Clone le dépôt : 

'''bash 
git clone https://github.com/kitsunabi/github-go-network-scanner.git
cd github-go-network-scanner

2.Initialise les modules Go(si besoin)
go mod init github.com/kitsunabi/github-go-network-scanner

3.Lancer le programme : 
go run main.go
