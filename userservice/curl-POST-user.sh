#!/usr/bin/env bash

# http://172.30.129.162:8080/personer Java
# http://172.30.94.134:6767/personer Go
URL=http://172.30.129.162:8080/personer

echo "Resetting file created_users.csv"
echo "userId" > created_users.csv
echo "Sending POST\n"
for index in {1..200}
do
    myjsonTemplate='{ "firstName": "firstName_CHANGEME", "lastName": "lastName_CHANGEME", "email": "f_CHANGEME@l_CHANGEME.se", "twitterHandle": "tweet_CHANGEME" }'
    echo $myjsonTemplate | sed "s/CHANGEME/$index/g" > myjson.tmp
    curl ${URL} -H "Content-Type: application/json" -X POST -d @myjson.tmp | jq '.' | grep "id" | awk {'print $2'} | sed 's/"//g' >> created_users.csv
done
echo " "
echo "Send finished and userId's appended to file created_users.csv\n"
