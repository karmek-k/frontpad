package messages

import (
	"errors"
	"time"

	"github.com/karmek-k/frontpad/pkg/db"
	"github.com/karmek-k/frontpad/pkg/db/models"
	"github.com/karmek-k/frontpad/utils"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/olahol/melody.v1"
)

func SessionCreate(s *melody.Session, msg *models.Message) error {
	// generate a session uuid
	id := uuid.NewV4()

	// hash the password
	hashed, err := utils.HashPassword("12345")
	if err != nil {
		return err
	}

	// create a blank session
	session := models.Session{
		Id:             id.String(),
		PasswordHashed: hashed,
	}

	// save it to redis
	err = db.RDB.Set(db.Ctx, session.Id, session.PasswordHashed, time.Hour).Err()
	if err != nil {
		return err
	}

	// return the session
	return errors.New("session created - handle this!")
}