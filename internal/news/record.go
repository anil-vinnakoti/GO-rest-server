package news

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Record struct {
	bun.BaseModel `bun:"table:news"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Author        string    `bun:"author,nullzero,notnull"`
	Title         string    `bun:"title,nullzero,notnull"`
	Content       string    `bun:"content,nullzero,notnull"`
	Summary       string    `bun:"summary,nullzero,notnull"`
	Source        string    `source:"source,nullzero,notnull"`
	Tags          []string  `bun:"tags,nullzero,notnull,array"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `bun:"deleted_at,nullzero,notnull,soft_delete"`
}
