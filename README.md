# nyaa-torrent-crawler

App to crawl http://nyaa.si to find a latest episode of the anime that you watch and download the torrent file

my goals is to create something like `CouchPotato` but for anime

## Feature

- Subscribe to anime that you want to watch every week by spicify a keyword search on nyaa.si
- Manage your subscription
- Download the torrent file if a new episode from the list of your subscription is available

## Installation

    $ go get github.com/ricoharisin/nyaa-torrent-crawler
    $ "$GOPATH/bin/nyaa-torrent-crawler" help

## Usage

### List of commands

    nyaa-torrent-crawler subscribe <keyword> <current episode>

subscribe an anime that you want to watch every week

`<keyword>` keyword search for nyaa.si, please input the keyword as detail as possible ex: `[HorribleSubs] Sakura Quest 720` otherwise the app will find the most recent matched one with random sub

`<current episode>` starting episode or current episode that you just watched, for example if you already watch episode 1 please input `1`, if you have not watched it yet, just put `0`

    nyaa-torrent-crawler crawl

start crawling to find the next episode from current episode

    nyaa-torrent-crawler list

show your subscription list

    nyaa-torrent-crawler unsubscribe <index>
    
unsubscribe the anime from your subscription list

`<index>` index from list, you can also run this command without index parameters and app will ask for index later

### How to use it?

After you subscribed several anime now put command `nyaa-torrent-crawler crawl` on `crontab` and make it running every X minute

And that's it! app will do the rest automatically 

The next step is to decide what will you do on that torrent file

In my case, i made a bash script to scp all torrent file to my NAS and delete it

	#!/bin/sh
	if [ "$(ls -A torrent)" ]; then
        	scp  torrent/* user@NAS_IP:/mnt/HD/HD_a2/Transmission/torrent/
        	if [ $? -eq 0 ];
        	then
                	echo "Success"
                	rm torrent/*
        	else
                	echo "Error cannot connect"
        	fi
	fi
 
 
## Next Feature?

I haven't decided what the next feature is, this is several idea that i came up with

- Automatically recognized the most popular fansub
- Notification when app found the latest episode
- WebUI?
- Integration with AniList?

## Contribute

I'm fairly new to golang so any code improvement, bug fixes, or even new feature are very welcome!  

## Bugs and Suggestion

Please visit the [issue tracker](https://github.com/ricoharisin/nyaa-torrent-crawler/issues) !



    
