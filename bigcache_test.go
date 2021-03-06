package bigcache

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/allegro/bigcache"
	"github.com/stretchr/testify/assert"
)

var bc, _ = bigcache.NewBigCache(bigcache.DefaultConfig(time.Second))

var testDataMap = map[string]interface{}{
	"bigcache": []string{"item.1", "item.2"},
}

func TestBigCache(t *testing.T) {
	var bcd = NewDriver(bc)
	var ttl = time.Millisecond * 500
	for key, val := range testDataMap {
		bs, _ := json.Marshal(val)
		assert.Nil(t, bcd.Set(key, bs, ttl))

		nbs, ok, err := bcd.Get(key)
		assert.True(t, ok)
		assert.Nil(t, err)
		assert.Equal(t, bs, nbs)

		assert.Nil(t, bcd.Del(key))
		_, ok, err = bcd.Get(key)
		assert.Nil(t, err)
		assert.False(t, ok)
	}

	var bcc = NewCache(bc)
	var tmp []string

	for key, val := range testDataMap {
		assert.Nil(t, bcc.Set(key, val, ttl))
		assert.Nil(t, bcc.Get(key, &tmp))
		assert.EqualValues(t, val, tmp)

		time.Sleep(ttl)

		assert.NotNil(t, bcc.Has(key))
	}
}
