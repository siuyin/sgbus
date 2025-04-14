package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/siuyin/sgbus/internal/ca"
	"github.com/siuyin/sgbus/internal/memg"
)

func main() {
	ctx := context.Background()
	driver := memg.NewDriver()
	defer driver.Close(ctx)

	if err := driver.VerifyConnectivity(ctx); err != nil {
		log.Fatal(err)
	}

	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: ""})
	defer session.Close(ctx)
	createIndexes(ctx, session)
	mergeStops(ctx, driver)
	addRoutes(ctx, driver)
}

func createIndexes(ctx context.Context, session neo4j.SessionWithContext) {
	indexes := []string{
		"CREATE INDEX ON :Stop(code);",
	}
	for _, index := range indexes {
		if _, err := session.Run(ctx, index, nil); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Indexes created")
}

func mergeStops(ctx context.Context, driver neo4j.DriverWithContext) {
	stop := ca.ParseStops("./data/stops.json")
	for k, v := range stop {
		node := fmt.Sprintf(
			`MERGE (:Stop {code:"%s", lng:%f, lat:%f, desc:"%s", road:"%s"})`,
			k, v.Lng, v.Lat, v.Desc, v.Road,
		)
		if _, err := neo4j.ExecuteQuery(ctx, driver, node, nil,
			neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("")); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Stops merged")
}

func addRoutes(ctx context.Context, driver neo4j.DriverWithContext) {
	s := ca.ParseServices("./data/services.json")
	for k, v := range s {
		for n := 0; n < len(v.Route); n++ {
			d := v.Route[n]
			for i := 0; i < len(d)-1; i++ {
				stop := d[i]
				nextStop := d[i+1]
				node := fmt.Sprintf(
					`MATCH (a:Stop {code:"%s"}),(b:Stop {code:"%s"})
				 CREATE (a)-[:NEXT {service_no:"%s", dir:%d, desc:"%s"}]->(b)
				`, stop, nextStop, k, n, v.Name,
				)
				if _, err := neo4j.ExecuteQuery(ctx, driver, node, nil,
					neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("")); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	log.Println("routes added")
}
