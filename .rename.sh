#!/bin/bash

# This script renames 'AppTemplate' to new project name in all files, exept hidden
# and in .github dir
# Usage:
# ./.rename.sh NewProjectName

function rename_in_file () {
    F=$1
    cat $F | sed "s/$OLDNAME/$NEWNAME/g" > /tmp/$OLDNAME-rename
    cat /tmp/$OLDNAME-rename > $F
}

function list_dir () {
    D=$1

    FILELIST=`ls -d $D/*`

    for F in $FILELIST; do
        if [[ -d "$F" ]]; then
            echo "Scanning $F:"
            list_dir $F
        else
            echo "- $F"
            rename_in_file $F
        fi
    done
}

OLDNAME='AppTemplate'
NEWNAME=$1

if [ "$NEWNAME" != "" ]; then
    echo "Renaming project from $OLDNAME to $NEWNAME in following:"
    echo
    DIR=`pwd`
    list_dir $DIR/.github
    list_dir $DIR
else
    echo "New project name can't be empty! Usage:"
    echo "$0 NewProjectName"
fi

