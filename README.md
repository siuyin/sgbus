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
  * show relevant bus stops near destinaton corresponding to starting bus stop labels.

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
1. Stops: curl https://raw.githubusercontent.com/cheeaun/sgbusdata/refs/heads/main/data/v1/stops.json

Note: Services above include route information. routes.json is a kml path file.

Related to Stops is the first and last bus info:  curl https://raw.githubusercontent.com/cheeaun/sgbusdata/refs/heads/main/data/v1/firstlast.json

## Graph design
Let bus stops be nodes labelled Stop.  
Each Stop will have properties, for example:
```
{ code:"12345", ng:103.45,lat:1.03,desc:"Opp Bkt Mkt",road:"Upper Bukit Timah" }
```
And the relation between nodes be a :NEXT link type with properties, for example:
```
{ service_no:10",dir:0 }
```

Create with, Cypher:
```
CREATE (s1:Stop {code:"10001"})
CREATE (s2:Stop {code:"10002"})
CREATE (s3:Stop {code:"10003"})
CREATE (s1)-[:NEXT {service_no:10,dir:0}]->(s2)
CREATE (s2)-[:NEXT {service_no:10,dir:0}]->(s3)
CREATE (s3)-[:NEXT {service_no:10,dir:1}]->(s2)
CREATE (s2)-[:NEXT {service_no:10,dir:1}]->(s1)
CREATE (s1)-[:NEXT {service_no:20,dir:0}]->(s2)
CREATE (s2)-[:NEXT {service_no:20,dir:1}]->(s1)
```

Query with:
```
MATCH (n) -[r]- (m) RETURN n,r,m
```

Shortest paths:
```
MATCH path=(n {code:"10001"}) -[:NEXT *ALLSHORTEST (r, n | 1)]-> (m {code: "10003"}) RETURN path
```

Path list:
```
MATCH path=(n {code:"10001"})-[relationships:NEXT *ALLSHORTEST (r, n | 1)]->(m {code:"10003"})
UNWIND (relationships(path)) AS edge
RETURN DISTINCT edge;
```