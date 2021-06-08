package configuration

import (
	"employee-csv-parser/utils"

	"github.com/hashicorp/go-memdb"
)

var db *memdb.MemDB

func DbConnect() {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"employee": &memdb.TableSchema{
				Name: "employee",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "FullName"},
					},
					"email": &memdb.IndexSchema{
						Name:    "email",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"salary": &memdb.IndexSchema{
						Name:    "salary",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Salary"},
					},
				},
			},
		},
	}
	connect, err := memdb.NewMemDB(schema)
	utils.PrintError(err)
	db = connect
}

func GetConnection() *memdb.MemDB {
	return db
}
