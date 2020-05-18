# distributed_programs_go
2 distributed programs in go: computing sum of a tab and finding prime numbers (with Sieve of Eratosthenes)

## Sum Tab

Pour exécuter le programme : copier le fichier sum_tab.go dans un dossier vide "sum_tab" => "go build" => "./sum_tab -n 12" où 12 corresponds au nombre de valeur dans le tableau à sommer.

Structure et fonctionnement du main : 

1. Récupération d'un nombre entré dans la console : n et création du tableau comportant n valeurs

2. Récupération du nombre de CPU disponibles : nbCPU

3. Création du tableau qui va gérer le "slice" en fonction du nombre de CPU. Sa structure est un tableau de maps, exemple : tabportion[0]["start"] va renvoyer l'entier de début de la première portion (en l'occurence 0)

4. Lancement des {nbCPU} Go Routines sommer() qui vont se charger de faire la somme du leur partie du tableau

5. Récupération des résultats intermédiaires via le canal résultat et stockage de ces résultats dans un tableau de taille {nbCPU}

6. Lancement d'une dernière go routine sommer() afin de faire la somme total

7. Récupération de la valeur total toujours via le canal resulat

 

Les go routines sont bien exécutés en parallèle comme le prouve la console ci-dessous, en exécutant le programme plusieurs fois, on remarque que l'ordre d'arrivé des résultats intermédiaires dans le canal resultat diffère.

![](https://i.imgur.com/RpGoEhX.png)



## Prime numbers

Pour exécuter le programme : copier le fichier prime_numbers.go dans un dossier vide "prime_numbers" => "go build" => "./prime_numbers -n 300" où 300 corresponds au nombre limite pour lequel on veut trouver l'ensemble des nombres premiers existant avant ce nombre.

Structure et fonctionnement du main : 

1. Récupération d'un nombre entré dans la console : n

2. Création des canaux "canal" et "resulat". Dans "canal" sera envoyé l'ensemble des nombres jusque n (à partir de 3 car 2 est premier et va servir de 1er multiple)

3. Lancement de la go routine filter(2, canal, resultat), la go routine va receptionner l'ensemble des valeurs arrivant de "canal" 

    3.1 si le nombre reçu n'est pas un multiple de 2 et si c'est le premier alors une nouvelle go routine filter va être lancé avec ce nouveau nombre premier et un nouveau canal.

    3.2 La première go routine va envoyer dans ce nouveau canal l'ensemble des valeurs non multiples.

    3.3 Ces actions s'enchainent sur ce principe dans chaque go routine créée.

4. Une fois tous les nombres envoyés dans le premier canal, on envoie 0 afin de récupérer les résultats. Les nombres premiers correspondent au numéro de la go routine (son premier paramètre).

5. On propage 0 dans l'ensemble du réseau (topologie en chemin (chaine orientée)) construit par les go routines. Chaque go routine transmet sa valeur dans le canal résultat.

6. Dans le main, ma première version récupérait ces résultats dans une go routine et il y avait un sleep afin de laisser le temps à la go routine de récupérer tous les résulats (car on ne sait pas à l'avance combien de nombre premier on attend), ce qui n'est évidemment pas optimal car si tous les résultats arrive en 100ms on doit attendre 1s avant de les afficher.

    Pour la deuxième version, la go routine qui correspond au dernier nombre premier trouvé va envoyer sa valeur puis 0 dans le canal résultat, ainsi le main sera au courant qu'il a reçu l'ensemble des nombres premiers.

    Ma deuxième version est possible car les canaux sont fifo et que chaque go routine envoie sa valeur dans le canal resultat avant de diffuser le 0 dans le canal du prochain filtre. (exécution atomique au sein de chaque go routine). C'est pour ces même raisons que le main reçoit tous les nombres premiers dans le bon ordre.

7. Affichage du résulat.

*Noter la différence de type de parallélisme avec l'exercice précédent.*

Les go routines s'executent bien en parallèle dans les 2 exercices mais dans le deuxième il y a en plus une relation de succesions des messages induite par les échanges unidirectionelles dans les canaux.

 
![](https://i.imgur.com/i6C9Y4B.png)

