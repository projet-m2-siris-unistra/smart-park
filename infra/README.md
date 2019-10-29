# Infra

Ce dossier contient le playbook ansible permettant de déployer l'infrastructure.

L'installation de kubernetes se base sur [`kubespray`](https://github.com/kubernetes-sigs/kubespray), qui est donc inclu comme submodule.

Le fichier `hosts.ini` contient l'inventaire ansible et `site.yaml` est le playbook principal.

Pour lancer le playbook: `ansible-playbook -b -i hosts.ini play.yaml`.

Pour installer les dépendances: `pip install -r requirements.txt`.
