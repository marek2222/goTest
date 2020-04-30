curl -X GET localhost/samoloty-miejsca/2
curl -X POST --data '{"dane":"ala ma kota"}' localhost/samoloty-miejsca/2
curl -X GET localhost/samoloty-miejsca/2
curl -X DELETE localhost/samoloty-miejsca/2
curl -X GET localhost/samoloty-miejsca/2

# http://zergling2:56743
curl -X GET http://zergling2:56743/samoloty-miejsca/2
curl -X POST --data '{"dane":"ala ma kota"}' http://zergling2:56743/samoloty-miejsca/2
curl -X GET http://zergling2:56743/samoloty-miejsca/2
curl -X DELETE http://zergling2:56743/samoloty-miejsca/2
curl -X GET http://zergling2:56743/samoloty-miejsca/2

# http://bilety-czarterowe-test.r.pl:8413
curl -X GET http://bilety-czarterowe-test.r.pl:8413/samoloty-miejsca/1
curl -X POST --data '{"dane":"ala ma kota"}' http://bilety-czarterowe-test.r.pl:8413/samoloty-miejsca/1
curl -X GET http://bilety-czarterowe-test.r.pl:8413/samoloty-miejsca/1
curl -X DELETE http://bilety-czarterowe-test.r.pl:8413/samoloty-miejsca/1
curl -X GET http://bilety-czarterowe-test.r.pl:8413/samoloty-miejsca/1

