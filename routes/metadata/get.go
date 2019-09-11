package metadata

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lavrahq/response-api/remote"
)

// GetMetadata returns the metadata available for tables and resources
// in Response.
func GetMetadata(c echo.Context) error {
	args := &remote.QueryRequestArgs{
		Type: "export_metadata",
		Args: make(map[string]interface{}),
	}

	metadata, err := remote.NewQueryRequest(args)
	if err != nil {
		return err
	}

	resp := echo.Map{}
	err = json.Unmarshal(metadata, &resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"metadata": &resp,
	})
}
