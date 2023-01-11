#!/bin/bash

git status | grep "nothing to commit, working tree clean"

if [ $? = 1 ] ; then
    git commit -m 'auto merge README.md'
    echo "git commit"
else
    echo "nothing to commit"
fi