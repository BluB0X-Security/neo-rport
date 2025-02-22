package clients

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/IOTech17/neo-rport/db/migration/clients"
	"github.com/IOTech17/neo-rport/db/sqlite"
	"github.com/IOTech17/neo-rport/server/clients/clientdata"
)

var DataSourceOptions = sqlite.DataSourceOptions{WALEnabled: false}

func NewFakeClientProvider(t *testing.T, exp *time.Duration, cs ...*clientdata.Client) *SqliteProvider {
	db, err := sqlite.New(":memory:", clients.AssetNames(), clients.Asset, DataSourceOptions)
	require.NoError(t, err)
	p := newSqliteProvider(db, exp)
	for _, cur := range cs {
		if cur != nil {
			err = p.Save(context.Background(), cur)
			require.NoError(t, err)
		}
	}
	return p
}
