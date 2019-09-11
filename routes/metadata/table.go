package metadata

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lavrahq/response-api/remote"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

type tableColumnsResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

type tableColumns map[string]tableColumn

type tableColumn struct {
	Name     string `json:"name"`
	Order    string `json:"order"`
	Default  string `json:"default"`
	Type     string `json:"type"`
	Nullable string `json:"nullable"`
	Comment  string `json:"comment"`
}

// GetTableMetadata returns metadata for a specific table.
func GetTableMetadata(c echo.Context) error {
	table := c.Param("table")

	args := &remote.QueryRequestArgs{
		Type: "run_sql",
		Args: echo.Map{
			"sql": fmt.Sprintf(`
				SELECT
					cols.column_name,
					cols.ordinal_position,
					cols.column_default,
					cols.udt_name,
					cols.is_nullable,
					(
						SELECT
							pg_catalog.col_description(c.oid, cols.ordinal_position::int)
						FROM
							pg_catalog.pg_class c
						WHERE
							c.oid = (SELECT ('"' || cols.table_name || '"')::regclass::oid)
							AND c.relname = cols.table_name
					) AS column_comment
				FROM
					information_schema."columns" cols
				WHERE
					cols."table_schema" = 'public'
				AND
					cols."table_name" = '%s'
				ORDER BY cols."ordinal_position";
		`, table),
		},
	}

	resp, err := remote.NewQueryRequest(args)
	if err != nil {
		return err
	}

	rdata := &tableColumnsResponse{}
	err = json.Unmarshal(resp, &rdata)
	if err != nil {
		return err
	}

	columns := make(tableColumns)
	for _, v := range rdata.Result {
		if v[0] == "column_name" {
			continue
		}

		columns[v[0]] = tableColumn{
			Name:     v[0],
			Order:    v[1],
			Default:  v[2],
			Type:     v[3],
			Nullable: v[4],
			Comment:  v[5],
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"metadata": echo.Map{
			"columns": columns,
		},
	})
}
