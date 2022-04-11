package messages

import (
	"github.com/karmek-k/frontpad/pkg/db"
	"github.com/karmek-k/frontpad/pkg/db/models"
	"gopkg.in/olahol/melody.v1"
)

func HandleCodeMarkupMessage(s *melody.Session, msg *models.Message) error {
	// TODO: use namespaced strings instead of hashes
	return db.RDB.Set(db.Ctx, msg.SessionId+":markup", msg.Content, 0).Err()
}