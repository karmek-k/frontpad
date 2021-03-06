package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/karmek-k/frontpad/pkg/db"
	"github.com/karmek-k/frontpad/pkg/db/models"
	"github.com/karmek-k/frontpad/utils"
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
	err = db.RDB.Set(db.Ctx, session.Id, session.PasswordHashed, time.Hour).Err()
	if err != nil {
		panic(err)
	}

	// return the session
	renderer.JSON(w, http.StatusCreated, session)
}

func SessionDelete(w http.ResponseWriter, r *http.Request) {
	// find session in redis
	id := chi.URLParam(r, "id")
	hashed, err := db.RDB.Get(db.Ctx, id).Result()

	if err == redis.Nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	// compare passwords
	matching := utils.CheckPassword("12345", hashed)
	if !matching {
		w.WriteHeader(http.StatusForbidden)

		return
	}
	
	// delete the session
	_ = db.RDB.HDel(db.Ctx, id).Err()

	w.WriteHeader(http.StatusNoContent)
}