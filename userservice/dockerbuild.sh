#/bin/bash!

docker build -t sadhal/contacts-be-go .

docker tag sadhal/contacts-be-go:latest 172.30.1.1:5000/contacts-be-dev/contacts-be-go

docker push 172.30.1.1:5000/contacts-be-dev/contacts-be-go