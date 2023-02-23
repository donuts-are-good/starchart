#!/usr/bin/env bash

# define the username and project name of the repo you're surveying
username="YOUR_USERNAME"
project="YOUR_PROJECT_NAME"

# get a list of users who've starred the project
starred_url="https://api.github.com/repos/${username}/${project}/stargazers"
users=$(curl -s ${starred_url} | jq -r '.[].login')

# look through the users and get an avatar for each
for user in ${users}; do
  avatar_url=$(curl -s "https://api.github.com/users/${user}" | jq -r '.avatar_url')
  
  # you can edit the 48 to get other sizes
  small_avatar_url="${avatar_url}&s=48"
  profile_url="https://github.com/${user}"
  
  # this outputs to the terminal
  # you can > the output to a text file if needed
  # generally it's easier to just copy paste into
  # your README with a cheesy message thanking them.
  echo "[![](${small_avatar_url})](${profile_url} \"${user}\")"
  
done
