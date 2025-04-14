// Package memg provides function to manage memgraph.
package memg

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/siuyin/dflt"
)

func NewDriver() neo4j.DriverWithContext {
	dbUser := dflt.EnvString("DB_USER", "")
	dbPassword := dflt.EnvString("DB_PASSWORD", "")
	dbURI := dflt.EnvString("DB_URI", "bolt://localhost:7687")
	log.Printf(`DB_USER="%s" DB_PASSWD=**** DB_URI=%s`,dbUser, dbURI)
	driver, err := neo4j.NewDriverWithContext(dbURI, neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		log.Fatal(err)
	}
	return driver
}