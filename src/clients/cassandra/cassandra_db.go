package cassandra

import (
	"github.com/gocql/gocql"
)

var(
	cluster *gocql.ClusterConfig
)

func init(){
	cluster = gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session,error){
	return cluster.CreateSession()
}