package store

import (
	"context"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	dtsync "github.com/ipfs/go-datastore/sync"
	logging "github.com/ipfs/go-log/v2"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/kenlabs/PandoStore/pkg/config"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multicodec"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	lp = cidlink.LinkPrototype{
		Prefix: cid.Prefix{
			Version:  1,
			Codec:    uint64(multicodec.DagJson),
			MhType:   uint64(multicodec.Sha2_256),
			MhLength: 16,
		},
	}

	testdata1 = []byte("testdata1")
	testdata2 = []byte("testdata2")

	cid1, _ = lp.Sum(testdata1)
	cid2, _ = lp.Sum(testdata2)
	//cid3, _  = lp.Sum([]byte("testdata3"))
	peer1, _ = peer.Decode("12D3KooWBckWLKiYoUX4k3HTrbrSe4DD5SPNTKgP6vKTva1NaRkJ")
)

func TestRoundTripPandoStore(t *testing.T) {
	_ = logging.SetLogLevel("PandoStore", "debug")

	ctx := context.Background()
	ds := datastore.NewMapDatastore()
	mds := dtsync.MutexWrap(ds)

	cfg := &config.StoreConfig{SnapShotInterval: time.Second.String()}
	store, err := NewStoreFromDatastore(ctx, mds, cfg)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = store.Store(ctx, cid1, []byte("testdata1"), peer1, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	time.Sleep(time.Second * 2)
	info, err := store.MetaInclusion(ctx, cid1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("%#v", info)

	snapshot, _, err := store.SnapShotStore.GetSnapShotByHeight(ctx, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", snapshot)

	err = store.Store(ctx, cid1, []byte("testdata1"), peer1, nil)
	assert.Contains(t, err.Error(), "key has existed")

	snapshot, err = store.SnapShotStore.GetSnapShotByCid(ctx, cid2)
	assert.Contains(t, err.Error(), "not found")

	pinfo, err := store.StateStore.GetProviderInfo(ctx, peer1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", pinfo)

	//time.Sleep(time.Second * 999)
}

func TestRestartPandoStore(t *testing.T) {
	_ = logging.SetLogLevel("PandoStore", "debug")
	ctx := context.Background()
	testdir := t.TempDir()
	cfg := &config.StoreConfig{
		Type:             "levelds",
		StoreRoot:        "",
		Dir:              testdir,
		SnapShotInterval: "1s",
	}
	db, err := NewStoreFromConfig(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	err = db.Store(ctx, cid1, []byte("testdata1"), peer1, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	info, err := db.MetaInclusion(ctx, cid1)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 1500)
	err = db.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		t.Fatal(err)
	}
	db, err = NewStoreFromConfig(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	val, err := db.Get(ctx, cid1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, val, testdata1)
	info, err = db.MetaInclusion(ctx, cid1)
	if err != nil {
		t.Fatal(err)
	}
	_, c, err := db.SnapShotStore.GetSnapShotByHeight(ctx, 0)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, info.ID, cid1)
	assert.Equal(t, info.InPando, true)
	assert.Equal(t, info.InSnapShot, true)
	assert.Equal(t, info.SnapShotID, c)
	assert.Equal(t, info.Context, []byte(nil))
	assert.Equal(t, info.Provider, peer1)
	assert.Equal(t, info.SnapShotHeight, uint64(0))

}
