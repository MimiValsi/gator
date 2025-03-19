# bootdev Blog Aggregator

## Programs
Postgresql (psql) 17
go 1.24

## Installation

- go on $PATH inside .zshenv or equivalent
exemple: export PATH=$PATH:$HOME/go/bin

Simply go install may suffice
If not then go build after install

## Usage
Here's the list of all commands available:\
login - login with registered user\
register - create a new user\
reset - delete every information in DB\
users - display every registered user and show which is currently in use\
agg - fetch infos from known URL (already registered in DB)\
addfeed - create a new feed by inserting it's name and URL\
feeds - display all known feeds\
follow - follow a specific feed (let's multiple user follow a feed)\
following - display every user following a feed\
unfollow - unfollow a feed\
browse - display every feed following by current userx\

First need to register a user:
./gator register foo

After that you can add a feed:
./gator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
nota: if the name of the feed has spaces, then make sure it's in between ""

You can add has many feeds has your disk has space...

If you want to see every known feed:
./gator feeds

and if you want specifics feeds, you can follow them:
./gator follow <feedname>

and ofc, display all following feeds:
./gator following

If you want to display limited feeds:
./gator browse <limit>

By default, it will display 2.