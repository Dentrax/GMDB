# [GMDB](https://github.com/Dentrax/GMDB) &nbsp; [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Do%20you%20love%20watching%20movies?%20Then%20what%20are%20you%20waiting%20for?&url=https://github.com/Dentrax/GMDB&via=github&hashtags=GMDB,cli,search,watch,movies,netflix,developers)


[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.png?v=103)](https://github.com/ellerbrock/open-source-badges/)
[![Build Status](https://travis-ci.org/Dentrax/GMDB.svg?branch=master)](https://travis-ci.org/Dentrax/GMDB)
[![codecov](https://codecov.io/gh/Dentrax/GMDB/branch/master/graph/badge.svg)](https://codecov.io/gh/Dentrax/GMDB)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FDentrax%2FGMDB.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FDentrax%2FGMDB?ref=badge_shield)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg)](https://sourcegraph.com/github.com/Dentrax/GMDB)

**GMDB is the ultra-simple, open-source, cross-platform Movie Library with Features (Search, Take Note, Watch Later, Like, Import, Learn)**

**Warning:** Are you [binge watcher](https://en.wikipedia.org/wiki/Binge-watching)? Let the fun begin with: `SEARCH-WATCH-REPEAT`!

[![demo](https://asciinema.org/a/tm0T7fssB8XCSal9ll3YdoVtw.svg)](https://asciinema.org/a/tm0T7fssB8XCSal9ll3YdoVtw?autoplay=1)

[What It Is](#what-it-is)

[Features](#features)

[How To Use](#how-to-use)

[Installation](#installation)

[Requirements](#requirements)

[Dependencies](#dependencies)

[About](#about)

[Collaborators](#collaborators)

[Branches](#branches) 

[Copyright & Licensing](#copyright--licensing)

[Contributing](#contributing)

[Contact](#contact)

## What It Is

**GMDB is the painless, ultra-simple, open-source, cross-platform Movie Library built in Golang**

### Screenshots

Search
--------------------------

![IMDB](https://raw.githubusercontent.com/Dentrax/GMDB/master/screenshots/search_imdb.jpg)

Search Result
--------------------------

![IMDB](https://raw.githubusercontent.com/Dentrax/GMDB/master/screenshots/search_imdb_result.jpg)

Questions
--------------------------

![IMDB](https://raw.githubusercontent.com/Dentrax/GMDB/master/screenshots/after_search.jpg)

Note
--------------------------

![IMDB](https://raw.githubusercontent.com/Dentrax/GMDB/master/screenshots/note_sample.jpg)

## Features

* Search movies from IMDB and RottenTomatoes (without API key)

* Search magnet links from 1337.x. [LEGAL_DISCLAIMER](LEGAL_DISCLAIMER.md)

* Watch the movie Trailers with MPV that fetched from IMDB

* Play the torrent magnet links directly (peerflix)

* Take the personal notes for the movies 

* Personal movie list (Watch Later, Movies Like)

* AI Learning for your Movies list (SOON)

* Terminal compatible colorful output

* Import your movie data from your Netflix data [Learn More](https://www.netflix.com/viewingactivity)

* Store your datas in SQLite Database

* Local Movie Database with SQLite

## How To Use

1. `help` command:
```
$ ./GMDB help
```

2. `search` command:
```
$ ./GMDB search movie to search
```

3. `learn` command:
```
$ ./GMDB learn ~/path/to/your/import/data.extension
```

4. `history` command:
```
$ ./GMDB history
```

5. `list` command:
```
$ ./GMDB list
```

6. `note` command:
```
$ ./GMDB note
```

7. `torrent` command:
```
$ ./GMDB torrent movie to search
```

8. `update` command:
```
$ ./GMDB update
```

## Installation

1. Clone the project to your computer and navigate to folder by executing the following command:
```
$ git clone https://github.com/Dentrax/GMDB.git && cd ./GMDB/
```

* Run this command if you want to test project before build & install
```
$ go test -v -cover -race ./...
```

* Run this command to get required packages before build & install
```
$ go get -d ./...
```

2. Build the all project using this command
```
$ go build .
```

3. Install the GMDB with following command
```
$ go install .
```

## Known Issues

* `--no-color` argument option not working for every situation

* `mpv` and `peerflix` are hard-coded necessary to operation 

* The situation without the Internet has not been tested yet

* Not tested in Windows

* I am absolutely sure there may be many untested cases. ^^

## TO-DO

- [ ] Add comments for functions

- [ ] Add benchmark tests for service parsers

- [ ] Add tests for Netflix data import

- [ ] Encrypt/Decrypt the database

- [ ] Live tests for IMDB and RottenTomatoes

- [ ] Sync database for some services (Evernote, Keep, etc.)

- [ ] Export/Import the database as a `.org` format for `orgmode`

- [ ] New search services

- [ ] Support for media players other than MPV

- [ ] `gmdb.go` file is too long and bad. Make it simple as possible

## Requirements

* `mpv`: A free, open source, and cross-platform media player [Learn More](https://mpv.io/)

* `peerflix`: Streaming torrent clients [Learn More](https://github.com/mafintosh/peerflix)

* You will need a computer on which you have the rights to compile files

## Dependencies

* Please see `Copyright & Licensing` section for more information

## About

GMDB was created to serve three purposes:

**GMDB is your painless CLI movie browser**

1. To act as a painless CLI movie browser

2. To provide a simplest and easiest way to search and watch movies

3. There is a source for you to develop your own CLI browser in Golang

## Collaborators

**Project Manager** - Furkan Türkal (GitHub: **[Dentrax](https://github.com/dentrax)**)

## Branches

We publish source for the **[GMDB]** in single rolling branch:

The **[master branch](https://github.com/dentrax/GMDB/tree/master)** is extensively tested and makes a great starting point. Also tracks [live changes](https://github.com/dentrax/GMDB/commits/master) by commits.

## Copyright & Licensing

The base project code is copyrighted by Furkan 'Dentrax' Türkal and is covered by single licence.

All program code (i.e. .go) is licensed under GNU GPLv3 License unless otherwise specified. Please see the **[LICENSE.md](https://github.com/Dentrax/GMDB/blob/master/LICENSE)** file for more information.

| Package                                                                | Author                                            | License    |
| :---------                                                             | :--------                                         | :--------  |
| [go-sqlite3](https://github.com/mattn/go-sqlite3)                      | [mattn](https://github.com/mattn)                 | MIT        |
| [go-ini](https://github.com/go-ini/ini)                                | [Unknwon](https://github.com/Unknwon)             | Apache 2.0 |
| [goconvey](https://github.com/smartystreets/goconvey)                  | [SmartyStreets](https://github.com/smartystreets) | MIT        |
| [goquery](https://github.com/PuerkitoBio/goquery)                      | [PuerkitoBio](https://github.com/PuerkitoBio)     | BSD 3      |
| [spinner](https://github.com/briandowns/spinner)                       | [briandowns](https://github.com/briandowns)       | Apache 2.0 |
| [chalk](https://github.com/ttacon/chalk)                               | [ttacon](https://github.com/ttacon)               | MIT        |
| [parsetorrentname](https://github.com/middelink/go-parse-torrent-name) | [middelink](https://github.com/middelink)         | MIT        |
| [sqlx](https://github.com/jmoiron/sqlx)                                | [jmoiron](https://github.com/jmoiron)             | MIT        |

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FDentrax%2FGMDB.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FDentrax%2FGMDB?ref=badge_large)

**References**

While this repository is being prepared, it may have been quoted from some sources that already mentioned on specific files.

If there is an unspecified source or if you think that I made a copyright infringement, please contact with me.

**Special Thanks**

- [PuerkitoBio](https://github.com/PuerkitoBio) for creating the awesome [goquery](https://github.com/PuerkitoBio/goquery) library that made this project possible.

## Contributing

Please check the [CONTRIBUTING.md](CONTRIBUTING.md) file for contribution instructions and naming guidelines.

## Contact

GMDB was created by Furkan 'Dentrax' Türkal

 * <https://www.furkanturkal.com>
 
You can contact by URL:
    **[CONTACT](https://github.com/dentrax)**

<kbd>Best Regards</kbd>
