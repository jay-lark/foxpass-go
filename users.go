package foxpass

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Users struct {
	Username  string `json:"username"`
	IsEngUser bool   `json:"is_eng_user"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

func (c *Client) GetAllUsers(ctx context.Context) ([]string, error) {
	res, err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/users"), nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	raw := map[string]json.RawMessage{}
	err = json.NewDecoder(res.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	var user []Users
	err = json.Unmarshal(raw["data"], &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
