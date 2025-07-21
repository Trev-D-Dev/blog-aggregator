## Description 

Scrapes posts from websites and displays them to the user. Many users can be registered and follow different feeds.

## Prerequisites

Go and Postgres must be installed to run the program

## Installation

Run this Go install command: 
``` go install github.com/Trev-D-Dev/blog-aggregator/cmd/gator@latest ```

## Configuration

Create a .gatorconfig.json file at the home of your system. Inside put the database link.

## Usage
Must use "go run . " at the start of every input
Here are the commands:
 - login "username": logs into the given user if they exist
 - register "username": registers the given user and sets the current user to them
 - reset: wipes the user database, which also wipes the feeds database
 - users: displays a list of the users and indicates the current one
 - agg [time value (ex: 1m30s)]: aggregates posts from the current user's feeds at intervals of the given time
 - addfeed "Title" "url": adds the given feed to the database and automatically has the current user follow the feed
 - follow "url": if the given feed url is already added, the current user will then follow it
 - following: displays all the lists the current user is following
 - unfollow "url": unfollows the given feed from the current user
 - browse [optional number of posts]: displays the most recent posts of the user's feeds. An optional number can be passed to display that amount of posts. By default 2 will be displayed