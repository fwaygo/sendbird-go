#!/bin/bash

if [ -f config.env ]
then
  export $(cat config.env | sed 's/#.*//g' | xargs)
  echo "env set"
fi