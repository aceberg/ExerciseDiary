package db

import (
	"fmt"

	"github.com/aceberg/ExerciseDiary/internal/models"
)

// Create - create table if not exists
func Create(path string) {

	sqlStatement := `CREATE TABLE IF NOT EXISTS exercises (
		"ID"		INTEGER PRIMARY KEY,
		"GR"		TEXT,
		"PLACE"		TEXT,
		"NAME"		TEXT,
		"DESCR"		TEXT,
		"IMAGE"		TEXT,
		"COLOR"		TEXT,
		"WEIGHT"	INTEGER,
		"REPS"		INTEGER
	);`
	exec(path, sqlStatement)

	sqlStatement = `CREATE TABLE IF NOT EXISTS sets (
		"ID"		INTEGER PRIMARY KEY,
		"DATE"		TEXT,
		"NAME"		TEXT,
		"COLOR"		TEXT,
		"WEIGHT"	INTEGER,
		"REPS"		INTEGER
	);`
	exec(path, sqlStatement)

	sqlStatement = `CREATE TABLE IF NOT EXISTS weight (
		"ID"		INTEGER PRIMARY KEY,
		"DATE"		TEXT,
		"WEIGHT"    INTEGER
	);`
	exec(path, sqlStatement)
}

// InsertEx - insert one exercise into DB
func InsertEx(path string, ex models.Exercise) {

	sqlStatement := `INSERT INTO exercises (GR, PLACE, NAME, DESCR, IMAGE, COLOR, WEIGHT, REPS) 
	VALUES ('%s','%s','%s','%s','%s','%s','%v','%d');`

	ex.Group = quoteStr(ex.Group)
	ex.Name = quoteStr(ex.Name)
	ex.Descr = quoteStr(ex.Descr)

	sqlStatement = fmt.Sprintf(sqlStatement, ex.Group, ex.Place, ex.Name, ex.Descr, ex.Image, ex.Color, ex.Weight, ex.Reps)

	exec(path, sqlStatement)
}

// InsertSet - insert one set into DB
func InsertSet(path string, ex models.Set) {

	sqlStatement := `INSERT INTO sets (DATE, NAME, COLOR, WEIGHT, REPS) 
	VALUES ('%s','%s','%s','%v','%d');`

	ex.Name = quoteStr(ex.Name)

	sqlStatement = fmt.Sprintf(sqlStatement, ex.Date, ex.Name, ex.Color, ex.Weight, ex.Reps)

	exec(path, sqlStatement)
}

// InsertW - insert weight
func InsertW(path string, ex models.BodyWeight) {

	sqlStatement := `INSERT INTO weight (DATE, WEIGHT) 
	VALUES ('%s','%v');`

	sqlStatement = fmt.Sprintf(sqlStatement, ex.Date, ex.Weight)

	exec(path, sqlStatement)
}

// DeleteEx - delete one exercise
func DeleteEx(path string, id int) {

	sqlStatement := `DELETE FROM exercises WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}

// DeleteSet - delete one set
func DeleteSet(path string, id int) {

	sqlStatement := `DELETE FROM sets WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}

// DeleteW - delete weight
func DeleteW(path string, id int) {

	sqlStatement := `DELETE FROM weight WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}

// ClearEx - delete all exercises from table
func ClearEx(path string) {
	sqlStatement := `DELETE FROM exercises;`
	exec(path, sqlStatement)
}

// ClearSet - delete all sets from table
func ClearSet(path string) {
	sqlStatement := `DELETE FROM sets;`
	exec(path, sqlStatement)
}
