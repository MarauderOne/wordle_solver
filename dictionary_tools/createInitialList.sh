#!/bin/bash

rm -f my.dict
rm -f initialList.dict
#sudo apt install aspell
aspell -d en dump master | aspell -l en expand > my.dict
grep -v "'" my.dict | grep "^.....$" | grep -v '[^[:lower:]]' | sort | uniq >>initialList.dict
rm -f my.dict
