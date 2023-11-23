#!/usr/bin/env bash

if [[ ! -n $1 ]]; then
  echo "No day num provided" >&2
	exit 2
fi

dayDir="day${1}"

if [[ ! -d "${dayDir}" ]]; then
	mkdir "${dayDir}"
	cp skeleton/main.go "${dayDir}/main.go"
	sed -i '' '1,2d' "${dayDir}/main.go"
	touch "${dayDir}/input.txt"
	touch "${dayDir}/test.txt"
else
	echo "Day directory already exists" >&2
	exit 3
fi