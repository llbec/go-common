package mpt

import (
	"testing"

	common "github.com/llbec/gocommon"
	"github.com/llbec/gocommon/db"
)

func TestDatabaseMetarootFetch(t *testing.T) {
	db := NewDatabase(db.NewDB(db.MemDBBackend, "", ""))
	if _, err := db.Node(common.Hash{}); err == nil {
		t.Fatalf("metaroot retrieval succeeded")
	}
}
