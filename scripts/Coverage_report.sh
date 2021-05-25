#!/usr/bin/env bash

#go to directory

cd ../componnents/automate_ui

# execute npm.test inside automate_ui

$ exec npm test

#list all the files
$ exec ls

#copy lcov.info 
 
srcdir="/automate/components/automate_ui/lcov.info"                   
dstdir="automate/Coverage_report/"
source=/automate/components/automate_ui/lcov.info
destination=automate/Coverage_report/
d=$(date +%m%d%y)

if [ -f $source]; then
      dstfile=$(basename $source)
      cp "$source" "$dstdir/$dstfile"
fi