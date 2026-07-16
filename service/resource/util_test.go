package resource

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomStr(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"length 6", 6},
		{"length 10", 10},
		{"length 32", 32},
		{"length 0", 0},
		{"length 1", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomStr(tt.n)
			assert.Len(t, got, tt.n)
			assert.NotContains(t, got, "0", "should not contain excluded char")
		})
	}
	assert.NotEqual(t, RandomStr(20), RandomStr(20), "should produce different results")
}

func TestSHA256(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"test", "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.want, SHA256(tt.input))
		})
	}
}

func TestSaltSecret(t *testing.T) {
	ori := "password"
	salt := "abcde"
	h := sha256.New()
	h.Write([]byte(ori))
	h.Write([]byte(salt))
	want := hex.EncodeToString(h.Sum(nil))

	assert.Equal(t, want, SaltSecret(ori, salt))
	assert.NotEqual(t, SaltSecret(ori, "salt1"), SaltSecret(ori, "salt2"), "different salts should produce different results")
}

func TestGeneralMFASecret(t *testing.T) {
	secret := GeneralMFASecret()
	assert.Len(t, secret, 32)

	_, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	assert.NoError(t, err, "should produce valid base32")
	assert.NotEqual(t, GeneralMFASecret(), GeneralMFASecret(), "should produce different results")
}

func TestRemoveSliceElement(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		assert.Equal(t, []string{"a", "c"}, RemoveSliceElement([]string{"a", "b", "c", "b"}, "b"))
	})
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, []int{1, 3, 4}, RemoveSliceElement([]int{1, 2, 3, 2, 4}, 2))
	})
	t.Run("element not found", func(t *testing.T) {
		assert.Len(t, RemoveSliceElement([]string{"a", "b"}, "c"), 2)
	})
	t.Run("empty slice", func(t *testing.T) {
		assert.Empty(t, RemoveSliceElement([]string{}, "a"))
	})
}

func TestUpdateSliceElement(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		got := UpdateSliceElement([]string{"a", "b", "c"}, "x", "b")
		assert.Equal(t, "x", got[1])
	})
	t.Run("multiple occurrences", func(t *testing.T) {
		got := UpdateSliceElement([]string{"a", "b", "b", "c"}, "x", "b")
		assert.NotContains(t, got, "b")
	})
	t.Run("element not found", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b"}, UpdateSliceElement([]string{"a", "b"}, "x", "z"))
	})
}

func TestDiffArrays(t *testing.T) {
	tests := []struct {
		name        string
		a, b        []int
		wantAdded   []int
		wantRemoved []int
	}{
		{
			name:        "basic diff",
			a:           []int{1, 2, 3},
			b:           []int{2, 3, 4},
			wantAdded:   []int{1},
			wantRemoved: []int{4},
		},
		{
			name:        "identical",
			a:           []int{1, 2, 3},
			b:           []int{1, 2, 3},
			wantAdded:   nil,
			wantRemoved: nil,
		},
		{
			name:        "empty a",
			a:           []int{},
			b:           []int{1, 2},
			wantAdded:   nil,
			wantRemoved: []int{1, 2},
		},
		{
			name:        "empty b",
			a:           []int{1, 2},
			b:           []int{},
			wantAdded:   []int{1, 2},
			wantRemoved: nil,
		},
		{
			name:        "both empty",
			a:           []int{},
			b:           []int{},
			wantAdded:   nil,
			wantRemoved: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			added, removed := DiffArrays(tt.a, tt.b)
			assert.ElementsMatch(t, tt.wantAdded, added)
			assert.ElementsMatch(t, tt.wantRemoved, removed)
		})
	}
}

func TestRemoveDuplicateElement(t *testing.T) {
	t.Run("with duplicates", func(t *testing.T) {
		assert.Len(t, RemoveDuplicateElement([]string{"a", "b", "a", "c", "b"}), 3)
	})
	t.Run("no duplicates", func(t *testing.T) {
		assert.Len(t, RemoveDuplicateElement([]int{1, 2, 3}), 3)
	})
	t.Run("nil input", func(t *testing.T) {
		assert.Nil(t, RemoveDuplicateElement[string](nil))
	})
	t.Run("empty input", func(t *testing.T) {
		assert.Empty(t, RemoveDuplicateElement([]string{}))
	})
}

func TestMaskEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  string
	}{
		{"long username", "john.doe@example.com", "joh****@example.com"},
		{"exactly 4 chars", "test@example.com", "tes****@example.com"},
		{"short username 3 chars", "bob@example.com", "***@example.com"},
		{"short username 1 char", "a@example.com", "***@example.com"},
		{"invalid email", "notanemail", "notanemail"},
		{"empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MaskEmail(tt.email))
		})
	}
}
