package rueidis

import (
	"context"
	"testing"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/mock"

	lib_store "github.com/eko/gocache/v4/lib/store"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewRueidis(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	// rueidis mock client
	client := mock.NewClient(ctrl)

	// When
	store := NewRueidis(client, lib_store.WithExpiration(6*time.Second))

	// Then
	assert.IsType(t, new(RueidisStore), store)
	assert.Equal(t, client, store.client)
	assert.Equal(t, &lib_store.Options{Expiration: 6 * time.Second}, store.options)
}

func TestRueidisGet(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	// rueidis mock client
	client := mock.NewClient(ctrl)
	client.EXPECT().DoCache(ctx, mock.Match("GET", "my-key"), defaultExpiration).Return(mock.Result(mock.RedisString("")))

	store := NewRueidis(client)

	// When
	value, err := store.Get(ctx, "my-key")

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestRueidisSet(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	cacheKey := "my-key"
	cacheValue := "my-cache-value"

	// rueidis mock client
	client := mock.NewClient(ctrl)
	client.EXPECT().Do(ctx, mock.Match("SET", cacheKey, cacheValue, "EX", "10")).Return(mock.Result(mock.RedisString("")))

	store := NewRueidis(client)

	// When
	err := store.Set(ctx, cacheKey, cacheValue)

	// Then
	assert.Nil(t, err)
}

func TestRueidisSetWhenNoOptionsGiven(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	cacheKey := "my-key"
	cacheValue := "my-cache-value"

	client := mock.NewClient(ctrl)
	client.EXPECT().Do(ctx, mock.Match("SET", cacheKey, cacheValue, "EX", "6")).Return(mock.Result(mock.RedisString("")))

	store := NewRueidis(client, lib_store.WithExpiration(6*time.Second))

	// When
	err := store.Set(ctx, cacheKey, cacheValue)

	// Then
	assert.Nil(t, err)
}

func TestRedisSetWithTags(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	cacheKey := "my-key"
	cacheValue := "my-cache-value"

	client := mock.NewClient(ctrl)
	client.EXPECT().Do(ctx, mock.Match("SET", cacheKey, cacheValue, "EX", "10")).Return(mock.Result(mock.RedisString("")))
	client.EXPECT().Do(ctx, mock.Match("SADD", "gocache_tag_tag1", "my-key")).Return(mock.Result(mock.RedisString("")))
	client.EXPECT().Do(ctx, mock.Match("EXPIRE", "gocache_tag_tag1", "2592000")).Return(mock.Result(mock.RedisString("")))

	store := NewRueidis(client)

	// When
	err := store.Set(ctx, cacheKey, cacheValue, lib_store.WithTags([]string{"tag1"}))

	// Then
	assert.Nil(t, err)
}

func TestRedisDelete(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	cacheKey := "my-key"

	client := mock.NewClient(ctrl)
	client.EXPECT().Do(ctx, mock.Match("DEL", cacheKey)).Return(mock.Result(mock.RedisInt64(1)))

	store := NewRueidis(client)

	// When
	err := store.Delete(ctx, cacheKey)

	// Then
	assert.Nil(t, err)
}

func TestRedisInvalidate(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	client := mock.NewClient(ctrl)
	client.EXPECT().DoCache(ctx, mock.Match("SMEMBERS", "gocache_tag_tag1"), defaultExpiration).Return(mock.Result(mock.RedisArray()))
	client.EXPECT().Do(ctx, mock.Match("DEL", "gocache_tag_tag1")).Return(mock.Result(mock.RedisInt64(1)))

	store := NewRueidis(client)

	// When
	err := store.Invalidate(ctx, lib_store.WithInvalidateTags([]string{"tag1"}))

	// Then
	assert.Nil(t, err)
}

func TestRedisClear(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	client := mock.NewClient(ctrl)
	client.EXPECT().Nodes().Return(map[string]rueidis.Client{
		"client1": client,
	})
	client.EXPECT().Do(ctx, mock.Match("ROLE")).Return(mock.Result(mock.RedisArray(mock.RedisString("master"))))
	client.EXPECT().Do(ctx, mock.Match("FLUSHALL")).Return(mock.Result(mock.RedisString("")))

	store := NewRueidis(client)

	// When
	err := store.Clear(ctx)

	// Then
	assert.Nil(t, err)
}

func TestRedisGetType(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	client := mock.NewClient(ctrl)

	store := NewRueidis(client)

	// When - Then
	assert.Equal(t, RueidisType, store.GetType())
}
