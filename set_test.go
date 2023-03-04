package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xyedo/set"
)

func Test_Creation(t *testing.T) {
	t.Run("creation with vales", func(t *testing.T) {
		arr := []int{1, 2, 3}
		myHashSet := set.New(arr...)
		require.Equal(t, arr, myHashSet.Values())
	})
	t.Run("creation with no values", func(t *testing.T) {
		myHashSet := set.New[int]()
		myHashSet.Add(1, 2, 3)
		require.Equal(t, []int{1, 2, 3}, myHashSet.Values())
	})
}

func Test_Has(t *testing.T) {
	t.Run("one values", func(t *testing.T) {
		myHashSet := set.New(1)
		v := myHashSet.Has(1)
		require.True(t, v)
	})
	t.Run("false when values not there", func(t *testing.T) {
		myHashSet := set.New(1)
		v := myHashSet.Has(2)
		require.False(t, v)
	})
	t.Run("many values", func(t *testing.T) {
		myHashSet := set.New("sexy", "brain")
		v := myHashSet.Has("sexy", "brain")
		require.True(t, v)
	})
	t.Run("many values but not all there", func(t *testing.T) {
		myHashSet := set.New("sexy", "brain")
		v := myHashSet.Has("sexy", "brain", "xyedo")
		require.False(t, v)
	})
}
func Test_Deleting(t *testing.T) {
	t.Run("valid deleting", func(t *testing.T) {
		myHashSet := set.New(1)
		v := myHashSet.Has(1)
		require.True(t, v)
		myHashSet.Delete(1)
		require.False(t, myHashSet.Has(1))
		require.Empty(t, myHashSet.Values())
	})
	t.Run("valid deleting many values", func(t *testing.T) {
		myHashSet := set.New(1, 3, 4, 5)
		v := myHashSet.Has(1, 3, 4, 5)
		require.True(t, v)
		myHashSet.Delete(1, 3, 4, 5)
		require.False(t, myHashSet.Has(1, 3, 4, 5))
		require.Empty(t, myHashSet.Values())
	})
}
