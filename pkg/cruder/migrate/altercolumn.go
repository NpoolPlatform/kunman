package migrate

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

func tableExist(ctx context.Context, conn dialect.ExecQuerier, database, table string) (bool, error) {
	rows := sql.Rows{}

	query, args := sql.
		Select("datetime_precision").
		From(sql.Table("`information_schema`.`tables`")).
		Where(
			sql.And(
				sql.EQ("table_schema", database),
				sql.EQ("table_name", table),
			),
		).
		Count().
		Query()
	if err := conn.Query(ctx, query, args, &rows); err != nil {
		return false, err
	}

	for rows.Next() {
		count := 0
		if err := rows.Scan(&count); err != nil {
			return false, err
		}
		if count == 0 {
			rows.Close()
			return false, nil
		}
	}
	rows.Close()

	return true, nil
}

func columnExist(ctx context.Context, conn dialect.ExecQuerier, database, table, column string) (bool, error) {
	rows := sql.Rows{}

	query, args := sql.
		Select("datetime_precision").
		From(sql.Table("`information_schema`.`columns`")).
		Where(
			sql.And(
				sql.EQ("table_schema", database),
				sql.EQ("table_name", table),
				sql.EQ("column_name", column),
			),
		).
		Count().
		Query()

	if err := conn.Query(ctx, query, args, &rows); err != nil {
		return false, err
	}

	for rows.Next() {
		count := 0
		if err := rows.Scan(&count); err != nil {
			return false, err
		}
		if count > 0 {
			rows.Close()
			return true, nil
		}
	}
	rows.Close()

	return false, nil
}

func alterColumn(
	ctx context.Context,
	conn dialect.ExecQuerier,
	database, table, srcColumn string,
	dstColumn *string,
	columnType string,
	unsigned, nilable, autoIncrement, unique bool,
) error {
	if exist, err := tableExist(ctx, conn, database, table); err != nil || !exist {
		return err
	}
	if exist, err := columnExist(ctx, conn, database, table, srcColumn); err != nil || !exist {
		return err
	}
	if dstColumn != nil {
		if exist, err := columnExist(ctx, conn, database, table, *dstColumn); err != nil || exist {
			return err
		}
	} else {
		dstColumn = &srcColumn
	}

	cb := sql.
		Column(*dstColumn).
		Type(columnType)

	query, args := sql.
		AlterTable(table).
		ChangeColumn(srcColumn, cb).
		Query()
	if unsigned {
		query = query + " unsigned"
	}
	if !nilable {
		query = query + " NOT NULL"
	}
	if autoIncrement {
		query = query + " AUTO_INCREMENT"
	}
	if unique {
		query = query + " UNIQUE"
	}
	if err := conn.Exec(ctx, query, args, nil); err != nil {
		return err
	}

	return nil
}

func AlterColumn(
	ctx context.Context,
	conn dialect.ExecQuerier,
	database, srcColumn string,
	dstColumn *string,
	columnType string,
	unsigned, nilable, autoIncrement, unique bool,
) error {
	rows := sql.Rows{}

	query, args := sql.
		Select("table_name").
		From(sql.Table("`information_schema`.`tables`")).
		Where(
			sql.And(
				sql.EQ("table_schema", database),
			),
		).
		Query()
	if err := conn.Query(ctx, query, args, &rows); err != nil {
		return err
	}

	tables := []string{}
	for rows.Next() {
		table := ""
		if err := rows.Scan(&table); err != nil {
			rows.Close()
			return err
		}
		tables = append(tables, table)
	}
	rows.Close()

	for _, table := range tables {
		if err := alterColumn(
			ctx,
			conn,
			database,
			table,
			srcColumn,
			dstColumn,
			columnType,
			unsigned,
			nilable,
			autoIncrement,
			unique,
		); err != nil {
			return err
		}
	}

	return nil
}

func Tables(
	ctx context.Context,
	conn dialect.ExecQuerier,
	database, srcColumn string,
	dstColumn *string,
	columnType string,
	unsigned, nilable, autoIncrement, unique bool,
) error {
	rows := sql.Rows{}

	query, args := sql.
		Select("table_name").
		From(sql.Table("`information_schema`.`tables`")).
		Where(
			sql.And(
				sql.EQ("table_schema", database),
			),
		).
		Query()
	if err := conn.Query(ctx, query, args, &rows); err != nil {
		return err
	}

	tables := []string{}
	for rows.Next() {
		table := ""
		if err := rows.Scan(&table); err != nil {
			rows.Close()
			return err
		}
		tables = append(tables, table)
	}
	rows.Close()

	for _, table := range tables {
		if err := alterColumn(
			ctx,
			conn,
			database,
			table,
			srcColumn,
			dstColumn,
			columnType,
			unsigned,
			nilable,
			autoIncrement,
			unique,
		); err != nil {
			return err
		}
	}

	return nil
}
