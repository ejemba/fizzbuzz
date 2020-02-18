#!/usr/bin/env bash

docker build -t ejemba/fizzbuzz .

docker run --name myfizzbuzz -d -p 3000:3000 ejemba/fizzbuzz

int1=3
str1=fizz
int2=5
str2=buzz
limit=100

for i in $(seq 1 100);
do
    curl -i "http://0.0.0.0:3000/fizzbuzz?int1=${int1}&int2=${int2}&limit=${limit}&str1=${str1}&str2=${str2}"
    echo ""
done

int1=4
str1=epo
int2=6
str2=jemba
limit=90

for i in $(seq 1 80);
do
    curl -i "http://0.0.0.0:3000/fizzbuzz?int1=${int1}&int2=${int2}&limit=${limit}&str1=${str1}&str2=${str2}"
    echo ""
done

int1=5
str1=cinq
int2=7
str2=sept
limit=80

for i in $(seq 1 60);
do
    curl -i "http://0.0.0.0:3000/fizzbuzz?int1=${int1}&int2=${int2}&limit=${limit}&str1=${str1}&str2=${str2}"
    echo ""
done

echo ""
echo "Statistiques des requètes"

curl -s  http://0.0.0.0:3000/debug/vars | jq '.["top.requests"]'

docker stop myfizzbuzz
docker rm myfizzbuzz
