#!/bin/bash

dir=${PWD##*/}

go build -v -o ${dir} && ./${dir}
