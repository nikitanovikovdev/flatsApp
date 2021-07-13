package database

import (
	"bufio"
	"database/sql"
	"flatApp/pkg/flats"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/stretchr/testify/require"
)

var schemaPrefixRegexps = [...]*regexp.Regexp{
	regexp.MustCompile(`(?i)(^CREATE SEQUENCE\s)(["\w]+)(.*)`),
	regexp.MustCompile(`(?i)(^CREATE TABLE\s)(["\w]+)(\s.+)`),
	regexp.MustCompile(`(?i)(^ALTER TABLE\s)(["\w]+)(\s.+)`),
	regexp.MustCompile(`(?i)(^UPDATE\s)(["\w]+)(\s.+)`),
	regexp.MustCompile(`(?i)(^INSERT INTO\s)(["\w]+)(\s.+)`),
	regexp.MustCompile(`(?i)(^DELETE FROM\s)(["\w]+)(.*)`),
	regexp.MustCompile(`(?i)(.+\sFROM\s)(["\w]+)(.*)`),
	regexp.MustCompile(`(?i)(\sJOIN\s)(["\w]+)(.*)`),
}

func addSchemaPrefix(schemaName, query string) string {
	prefixedQuery := query
	for _, re := range schemaPrefixRegexps {
		prefixedQuery = re.ReplaceAllString(prefixedQuery, fmt.Sprintf("${1}%s ${2}${3}${4}", schemaName))
	}

	return prefixedQuery
}

func loadTestData(t *testing.T, db *sql.DB, schemaName string, testDataNames ...string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	for _, testDataName := range testDataNames {
		file, err := os.Open(fmt.Sprintf("../../../migrations/%s.sql", testDataName))
		require.NoError(t, err)

		reader := bufio.NewReader(file)
		var query string
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				query = addSchemaPrefix(schemaName, query)
				_, err := db.Exec(query)
				require.NoError(t, err)
				query = ""
				break
			}

			require.NoError(t, err)

			line = line[:len(line)-1]
			if line == "" {
				query = addSchemaPrefix(schemaName, query)
				_, err := db.Exec(query)
				require.NoError(t, err)
				query = ""
			}
			query += line
		}
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}
}

func createTestConnection(t *testing.T, schema string) *sql.DB {
	connStr := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable&search_path=%s",
		"postgres",
		"postgres",
		"localhost",
		"5432",
		"postgres",
		schema,
	)

	db, dbErr := sql.Open("pgx", connStr)
	require.NoError(t, dbErr)

	_, err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema))

	require.NoError(t, err)

	loadTestData(t, db, schema, "000001_create_flats_table.up", "000003_insert_into_cities.up", "000002_insert_into_flats.up")

	return db
}



func CreateTestFlatsRepository(t *testing.T, schema string) (*flats.RepositorySQL, func()) {
	db := createTestConnection(t, schema)
	repo := flats.NewRepository(db)

	return repo, func() {
		_, err := db.Exec(fmt.Sprintf("DROP SCHEMA %s CASCADE", schema))
		if err != nil {
			t.Fatal(err)
		}
	}
}