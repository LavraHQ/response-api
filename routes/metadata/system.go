package metadata

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lavrahq/response-api/remote"
)

func getSpecialComments() ([]byte, error) {
	args := &remote.QueryRequestArgs{
		Type: "run_sql",
		Args: echo.Map{
			"sql": `
				SELECT
					t.table_name,
					split_part(get_rel_description (t.table_schema || '.' || t.table_name), '/', 1) AS table_type,
					split_part(get_rel_description (t.table_schema || '.' || t.table_name), '/', 2) AS table_comment
				FROM
					information_schema. "tables" t
				WHERE
					table_schema = 'public'
					OR table_schema = 'audit';
			`,
		},
	}

	return remote.NewQueryRequest(args)
}

type specialCommentsResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}

type specialComment struct {
	Table   string `json:"table"`
	Type    string `json:"type"`
	Comment string `json:"comment"`
}

type specialComments map[string]specialComment

// GetSpecial returns the special comments and internal
// table types.
func GetSpecial(c echo.Context) error {
	data, err := getSpecialComments()
	if err != nil {
		return err
	}

	resp := &specialCommentsResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return err
	}

	special := make(specialComments)
	for _, c := range resp.Result {
		if c[0] == "table_name" {
			continue
		}

		special[c[0]] = specialComment{
			Table:   c[0],
			Type:    c[1],
			Comment: c[2],
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"special": &special,
	})
}
