// Package tokenindex 提供用户 token 索引管理功能
// 索引用于维护用户所有活跃 token 的列表, 支持批量清除
package tokenindex

import (
	"context"
	"fmt"
	"time"

	"github.com/tsingsun/woocoo/pkg/cache"
)

// userTokenIndexKey 返回用户 token 索引的 cache key
func userTokenIndexKey(uid int) string {
	return fmt.Sprintf("user_tokens:%d", uid)
}

// Add 将 token ID 添加到用户的 token 索引中
func Add(ctx context.Context, c cache.Cache, uid int, tokenID string, ttl time.Duration) error {
	if c == nil {
		return nil
	}
	key := userTokenIndexKey(uid)

	var tokens []string
	if err := c.Get(ctx, key, &tokens); err != nil {
		tokens = []string{}
	}

	tokens = append(tokens, tokenID)
	return c.Set(ctx, key, tokens, cache.WithTTL(ttl))
}

// Remove 从用户的 token 索引中移除指定的 token ID
func Remove(ctx context.Context, c cache.Cache, uid int, tokenID string) error {
	if c == nil {
		return nil
	}
	key := userTokenIndexKey(uid)

	var tokens []string
	if err := c.Get(ctx, key, &tokens); err != nil {
		return nil
	}

	newTokens := make([]string, 0, len(tokens))
	for _, t := range tokens {
		if t != tokenID {
			newTokens = append(newTokens, t)
		}
	}

	if len(newTokens) == 0 {
		return c.Del(ctx, key)
	}
	return c.Set(ctx, key, newTokens)
}

// Get 获取用户的所有 token ID
func Get(ctx context.Context, c cache.Cache, uid int) ([]string, error) {
	if c == nil {
		return []string{}, nil
	}
	key := userTokenIndexKey(uid)

	var tokens []string
	if err := c.Get(ctx, key, &tokens); err != nil {
		return []string{}, nil
	}
	return tokens, nil
}

// ClearIndex 清除用户的所有 token 索引
// 注意: 这只是清除索引, 实际的 token 需要调用方逐个删除
func ClearIndex(ctx context.Context, c cache.Cache, uid int) error {
	if c == nil {
		return nil
	}
	key := userTokenIndexKey(uid)
	return c.Del(ctx, key)
}

// ClearAll 清除用户的所有 token 和索引
// 先获取索引中的所有 token ID, 逐个删除, 最后清除索引
func ClearAll(ctx context.Context, c cache.Cache, uid int) error {
	if c == nil {
		return nil
	}
	tokenIDs, err := Get(ctx, c, uid)
	if err != nil {
		return err
	}
	// 逐个删除 token
	for _, tokenID := range tokenIDs {
		_ = c.Del(ctx, tokenID)
	}
	// 清除索引
	return ClearIndex(ctx, c, uid)
}

// ClearExcept 清除用户除指定 token 外的所有 token 和索引
// exceptTokenID 是要保留的 token ID (通常为当前登录的 token)
func ClearExcept(ctx context.Context, c cache.Cache, uid int, exceptTokenID string) error {
	if c == nil {
		return nil
	}
	tokenIDs, err := Get(ctx, c, uid)
	if err != nil {
		return err
	}
	// 逐个删除 token, 排除指定的 token
	for _, tokenID := range tokenIDs {
		if tokenID != exceptTokenID {
			_ = c.Del(ctx, tokenID)
		}
	}
	// 清除索引
	return ClearIndex(ctx, c, uid)
}
