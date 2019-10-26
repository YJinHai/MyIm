package chat

import (
	"github.com/YJinHai/MyIm/internal/app/models"
)

// Chat represents private or channel chat
type Chat struct {
	Name    string           `json:"name"`
	Secret  string           `json:"secret"`
	Members map[int]*models.ImUser `json:"members"`
}
