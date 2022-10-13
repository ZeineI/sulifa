package server

import (
	"fmt"

	"github.com/ZeineI/sulifa/internal/models"
)

func (sv *Server) isAuth(username string) (*models.User, error) {

	user, err := sv.storage.GetUserAuth(username)
	if err != nil {
		return nil, fmt.Errorf("isAuth failed: %s", err)
	}

	return user, nil
}
