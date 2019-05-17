// 宣告程式屬於哪個 package
package main

// 引入套件
import (
    "fmt"
    bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Constants to be used throughout the example
const (
	URI          = "bolt://username:password@localhost:7687"
	CreateNode   = "CREATE (n:Person {name : {name}})"  
	GetNode      = "MATCH (n:Person) RETURN n.name"
	RelationNode = "MATCH path=(n:Person)-[:REL]->(m) RETURN path"
	DeleteNodes  = "MATCH (n:Person) DETACH DELETE n"
)

func main() {
	con := createConnection()
	defer con.Close()


	fmt.Printf("----- CreateNode -----\n")
	st := prepareSatement(CreateNode, con)
	executeStatement(st)

	fmt.Printf("----- GetNode -----\n")
	st := prepareSatement(GetNode, con)
	rows := queryStatement(st)
	consumeRows(rows, st)

	cleanUp(DeleteNodes, con);

}
func createConnection() bolt.Conn {
	driver := bolt.NewDriver()
	con, err := driver.OpenNeo(URI)
	handleError(err)
	return con
}

//CreateNode
func prepareSatement(query string, con bolt.Conn) bolt.Stmt {

	st, err := con.PrepareNeo(query)
    handleError(err)
	return st
}

func executeStatement(st bolt.Stmt) {
	result, err := st.ExecNeo(map[string]interface{}{"name": "Arnold"})
	handleError(err)
	numResult, err := result.RowsAffected()
	handleError(err)
	fmt.Printf("CREATED ROWS: %d\n", numResult) // CREATED ROWS: 1

	// Closing the statment will also close the rows
	st.Close()
}
//CreateNode

func queryStatement(st bolt.Stmt) bolt.Rows {
	// Even once I get the rows, if I do not consume them and close the
	// rows, Neo will discard and not send the data
	rows, err := st.QueryNeo(nil)
	handleError(err)
	return rows
}

func consumeRows(rows bolt.Rows, st bolt.Stmt) {

	fmt.Printf("COLUMNS: %#v\n", rows.Metadata()["fields"].([]interface{})) // COLUMNS: n.name
	var err error
	err = nil
	for err == nil {
		var data []interface{}
		data, _, err = rows.NextNeo()
		if err == nil {
			fmt.Printf("FIELDS: %s \n", data[0].(string)); 
		}
		
	}

	st.Close()
}

func cleanUp(query string, con bolt.Conn) {
	result, _ := con.ExecNeo(query, nil)
	fmt.Println(result)
	numResult, _ := result.RowsAffected()
	fmt.Printf("Rows Deleted: %d", numResult) // Rows Deleted: 13
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}