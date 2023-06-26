#!/bin/bash

# Ask user for commit message
read -p "Enter commit message: " message

# Add all files to staging area
git add .

# Commit changes with message
git commit -m "$message"

# Push changes to master branch
git push origin master
