package gosql

import (
	"context"
	"database/sql"
	"log"
)

// get here is an example of a function that executes a query and returns the results of the query. Nothing new here.
// The difference is how the error is handled.
// First of all, the query executed is a context aware, so, the parent context can cancel execution after a given timeout for example
// Secondly, there is a defer function to handle closing the rows, which is important, this avoids a memory leak where we have resources still open for use
// Also, on that point, it is important that the defer function handles errors accordingly in case there is an error closing the rows. This can be a simple logging
// of the error or even further using monitoring tools, alerting etc.
// Third, is the use of the sql.NulXXX fields to define the types of the columns that could be returned from a query. This is more expressive and makes it clear to the
// reader that it is possible that a given column could be null and therefore it is important to check whether the field has been set with a value which is handled by checking
// whether the value itself, in this case, 'department' is valid with department.Valid.
// Fourth, an error could be encountered when reading the rows with rows.Next(). This could either happen when reading the next row or when the reading of the rows
// is complete and and error is encountered. To ensure that the error is captured, we check with rows.Err() to check whether there are any errors encountered.
// If none, then we can know that the reading of the rows completed successfully and evaluate the results returned.
func get(ctx context.Context, db *sql.DB, id string) (string, int, error) {
	rows, err := db.QueryContext(ctx, "SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v/n", err)
		}
	}()

	var (
		department sql.NullString //alternative type here could be *string
		age        int
	)

	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return "", 0, err
		}
	}

	if err := rows.Err(); err != nil {
		return "", 0, err
	}

	if department.Valid {
		return department.String, age, nil
	}

	return "", age, nil
}
