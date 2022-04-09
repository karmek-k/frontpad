package routes

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/karmek-k/frontpad/pkg/db"
	"github.com/karmek-k/frontpad/pkg/db/models"
	"github.com/karmek-k/frontpad/pkg/utils"
	uuid "github.com/satori/go.uuid"
)

func SessionCreate(w http.ResponseWriter, r *http.Request) {
	// generate a session uuid
	id := uuid.NewV4()

	// hash the password
	hashed, err := utils.HashPassword("12345")
	if err != nil {
		panic(err)
	}

	// create a blank session
	session := models.Session{
		Id: id.String(),
		PasswordHashed: hashed,
	}
	
	// save it to redis
	db.RDB.HSet(r.Context(), session.Id,
		"password", session.PasswordHashed,
	)

	// return the session
	render.Render(w, r, session)
}

func SessionDelete(w http.ResponseWriter, r *http.Request) {

}