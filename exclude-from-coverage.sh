#!/bin/bash
while read file || [ -n "$file" ] 
do  
echo $file
sed -i "/${file//\//\\/}/d" ./coverage.out 
done < ./go-test-exclusions