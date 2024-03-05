#!/bin/bash
while read file || [ -n "$file" ] 
do  
    sed -i "/${file//\//\\/}/d" ./coverage.out 
done < ./coverage-exclusions