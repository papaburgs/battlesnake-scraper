# battlesnake-scraper
Script to scrape game data from battlesnake.com

Expect this to be a long running application. Perhaps put a API on it for managing things remotely

## Goals
### Data collection
* Scrape areana page and get list of snakes in an arena
  * an arena would include things like, spring league, global wrapped, etc
  * https://play.battlesnake.com/arena/spring-league-2022/
  * specifically pull out the id for desired snakes
* With the arenas scraped, we now have a list of IDs
  * scrape each of those IDs with the arena details api
  * https://play.battlesnake.com/arena/details/slb_w8Ym9frbSdjfPJ7vkC4qYgHd/
  * this gives us a list of recent games (last 20 I think)
  * each game has things like point delta and elimination reason
* More details on death reasons and final state, one can use the engine endpoint
  * https://engine.battlesnake.com/games/03a8becb-aae6-4c85-8269-5251195a3d48
  * This returns a json blob similar to a move call, but with a LastFrame object that describes the games end.

### Storage
End state is to put this into a splunk kvstore to enrich game data. If I can interact with it nicely, I will
use that as the local store, if not I'll use redis or an internal badger/bolt db. Perhaps just a local file, a massive json blob.

## Other notes
### Timing considerations
This doesn't need to be real time. Plan to scrape arenas daily to get updated list of snakes. 
Plan to sleep arena schedule based on game play to get snake games. For instance, if games
are once every ten minutes, can go 190 minutes between scrapes and still get all the data, probably do those once and hour.
For arenas like spring league, can just do 30 mins for the hours that the games run.
For the game details, can maintain a list of games to be scraped and put a random sleep in there up to a minute.

## Config
Store config details in a yaml file - probably use cobra and use live reload.


### Things to store
* snakes of interest
* arenas of interest
* scrape intervals
* storage options
  * write to file, internal db, splunk kvstore
