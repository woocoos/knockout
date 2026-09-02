// Package security 提供安全相关功能
// token 索引用于维护用户所有活跃 token 的列表, 支持批量清除
package security

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

// AddTokenIndex 将 token ID 添加到用户的 token 索引中
func AddTokenIndex(ctx context.Context, c cache.Cache, uid int, tokenID string, ttl time.Duration) error {
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

// RemoveTokenIndex 从用户的 token 索引中移除指定的 token ID
func RemoveTokenIndex(ctx context.Context, c cache.Cache, uid int, tokenID string) error {
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

// GetTokenIndex 获取用户的所有 token ID
func GetTokenIndex(ctx context.Context, c cache.Cache, uid int) ([]string, error) {
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

// ClearTokenIndex 清除用户的所有 token 索引
// 注意: 这只是清除索引, 实际的 token 需要调用方逐个删除
func ClearTokenIndex(ctx context.Context, c cache.Cache, uid int) error {
	if c == nil {
		return nil
	}
	key := userTokenIndexKey(uid)
	return c.Del(ctx, key)
}

// ClearAllTokens 清除用户的所有 token 和索引
// 先获取索引中的所有 token ID, 逐个删除, 最后清除索引
func ClearAllTokens(ctx context.Context, c cache.Cache, uid int) error {
	if c == nil {
		return nil
	}
	tokenIDs, err := GetTokenIndex(ctx, c, uid)
	if err != nil {
		return err
	}
	// 逐个删除 token
	for _, tokenID := range tokenIDs {
		_ = c.Del(ctx, tokenID)
	}
	// 清除索引
	return ClearTokenIndex(ctx, c, uid)
}

// ClearTokensExcept 清除用户除指定 token 外的所有 token 和索引
// exceptTokenID 是要保留的 token ID (通常为当前登录的 token)
func ClearTokensExcept(ctx context.Context, c cache.Cache, uid int, exceptTokenID string) error {
	if c == nil {
		return nil
	}
	tokenIDs, err := GetTokenIndex(ctx, c, uid)
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
	return ClearTokenIndex(ctx, c, uid)
}
