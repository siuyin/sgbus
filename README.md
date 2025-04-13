# Travelling in Singapore by Bus

Given I am bus user
and I want to get from my current location to a named destination by bus,
when I use the App,
then I should see suggestions on how to get to my destination.

## User Interface
* On opening App
  * determine current location
  * get destination from user
  * show relevent nearby bus stops with options labelled A,B,C etc.
  * show relevant bus stops near destinaton corresponding to starting bus stop labels

* On selecting a destination bus stop, present estimated bus travel times and walking distances for the relevant labelled options.

## Data sources

1. Bus Services: https://datamall2.mytransport.sg/ltaodataservice/BusServices
1. Bus Routes: https://datamall2.mytransport.sg/ltaodataservice/BusRoutes
1. Bus Stops: https://datamall2.mytransport.sg/ltaodataservice/BusStops

Example usage:
```
curl -H "accountKey: $LTA_ACCOUNT_KEY" 'https://datamall2.mytransport.sg/ltaodataservice/BusServices'
curl -H "accountKey: $LTA_ACCOUNT_KEY" 'https://datamall2.mytransport.sg/ltaodataservice/BusRoutes'
curl -H "accountKey: $LTA_ACCOUNT_KEY" 'https://datamall2.mytransport.sg/ltaodataservice/BusStops'
```

These data from LTA DataMall is incomplete. Eg. Bus Service 67 is missing.

I found a bus enthusiast. Chee Aun who scraped the LTA data: https://github.com/cheeaun/sgbusdata/

The processed data is in: https://github.com/cheeaun/sgbusdata/tree/main/data/v1

1. Services: curl https://raw.githubusercontent.com/cheeaun/sgbusdata/refs/heads/main/data/v1/services.json
1. Routes: curl https://raw.githubusercontent.com/cheeaun/sgbusdata/refs/heads/main/data/v1/routes.json
1. Stops: curl https://raw.githubusercontent.com/cheeaun/sgbusdata/refs/heads/main/data/v1/stops.json

Related to Stops is the first and last bus info:  curl https://raw.githubusercontent.com/cheeaun/sgbusdata/refs/heads/main/data/v1/firstlast.json

