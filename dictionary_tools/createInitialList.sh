#!/bin/bash

rm -f my.dict
rm -f initialList.dict
#sudo apt install aspell
aspell -d en dump master | aspell -l en expand > my.dict
grep -v "'" my.dict | grep "^.....$" | grep -v '[^[:lower:]]' | tr '[:lower:]' '[:upper:]' | sort | uniq | grep -v "^HDQRS$" | >>initialList.dict
rm -f my.dict
