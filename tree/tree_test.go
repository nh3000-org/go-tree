package tree_test

import (
	"fmt"
	"testing"

	"github.com/nh3000-org/go-tree/node"

	"github.com/nh3000-org/go-tree/tree"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := tree.New[int]()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*tree.Tree[int]", fmt.Sprintf("%T", sut))
}

func TestTree_AddRoot_WhenTreeIsEmpty_ShouldReturnTrue(t *testing.T) {
	// Arrange
	sut := tree.New[int]()

	// Act
	added := sut.AddRoot(node.New(42))

	// Assert
	assert.True(t, added)
}

func TestTree_AddRoot_WhenTreeIsNotEmpty_ShouldReturnFalse(t *testing.T) {
	// Arrange
	sut := tree.New[int]()

	// Act
	_ = sut.AddRoot(node.New(42))
	added := sut.AddRoot(node.New(43))

	// Assert
	assert.False(t, added)
}

func TestTree_GetRoot_WhenThereIsNotRoot_ShouldReturnFalse(t *testing.T) {
	// Arrange
	sut := tree.New[int]()

	// Act
	root, hasRoot := sut.GetRoot()

	// Assert
	assert.Nil(t, root)
	assert.False(t, hasRoot)
}

func TestTree_GetRoot_WhenThereIsRoot_ShouldReturnTrue(t *testing.T) {
	// Arrange
	sut := tree.New[int]()
	sut.AddRoot(node.New(42))

	// Act
	root, hasRoot := sut.GetRoot()

	// Assert
	assert.NotNil(t, root)
	assert.True(t, hasRoot)
}

func TestTree_Add_WhenThereIsNoRoot_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[int]()

	// Act
	added := tr.Add(0, node.New(42))

	// Assert
	assert.False(t, added)
}

func TestTree_Add_WhenRootIsNotRight_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[int]()

	// Act
	_ = tr.AddRoot(node.New(42))
	added := tr.Add(3, node.New(42))

	// Assert
	assert.False(t, added)
}

func TestTree_Add_WhenThereIsRootAndRootIsRight_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[int]()

	// Act
	_ = tr.AddRoot(node.New(42))
	added := tr.Add(0, node.New(42))

	// Assert
	assert.True(t, added)
}

func TestTree_Get_WhenThereIsNoRoot_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	// Act
	n, found := tr.Get(8)

	// Assert
	assert.Nil(t, n)
	assert.False(t, found)
}

func TestTree_Get_WhenThereIsNoId_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(node.New("0"))
	tr.Add(0, node.New("1.0"))

	// Act
	node, found := tr.Get(8)

	// Assert
	assert.Nil(t, node)
	assert.False(t, found)
}

func TestTree_Get_WhenThereIsIdOnRoot_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(node.New("0"))

	// Act
	node, found := tr.Get(0)

	// Assert
	assert.NotNil(t, node)
	assert.True(t, found)
}

func TestTree_Get_WhenThereIsId_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(node.New("0").WithID(0))
	tr.Add(0, node.New("1.0").WithID(1))

	// Act
	node, found := tr.Get(1)

	// Assert
	assert.NotNil(t, node)
	assert.True(t, found)
}

func TestTree_Backtrack_WhenIdNotFound_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(node.New("0.0"))

	// Act
	n, found := tr.Backtrack(1)

	// Assert
	assert.Nil(t, n)
	assert.False(t, found)
}

func TestTree_Backtrack_WhenIdFound_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(node.New("0.0").WithID(0))
	tr.Add(0, node.New("1.0").WithID(1))
	tr.Add(1, node.New("2.0").WithID(2))
	tr.Add(2, node.New("3.0").WithID(3))

	// Act
	n, found := tr.Backtrack(3)

	// Assert
	assert.NotNil(t, n)
	assert.True(t, found)
}

func TestTree_GetStructure_WhenThereIsNoRoot_ShouldReturnFalse(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	// Act
	structure, found := tr.GetStructure()

	// Assert
	assert.Nil(t, structure)
	assert.False(t, found)
}

func TestTree_GetStructure_WhenThereIsRoot_ShouldReturnTrue(t *testing.T) {
	// Arrange
	tr := tree.New[string]()

	tr.AddRoot(node.New("0.0"))
	tr.Add(0, node.New("1.0"))
	tr.Add(0, node.New("1.1"))
	tr.Add(0, node.New("1.2"))
	tr.Add(1, node.New("2.0"))
	tr.Add(1, node.New("2.1"))
	tr.Add(1, node.New("2.2"))
	tr.Add(2, node.New("2.0"))
	tr.Add(2, node.New("2.1"))
	tr.Add(2, node.New("2.2"))
	tr.Add(3, node.New("3.0"))
	tr.Add(3, node.New("3.1"))
	tr.Add(3, node.New("3.2"))
	tr.Add(4, node.New("4.0"))

	// Act
	structure, found := tr.GetStructure()

	// Assert
	assert.NotNil(t, structure)
	assert.True(t, found)
	for _, str := range structure {
		assert.NotEmpty(t, str)
	}
}

func TestTree_Filter_WhenThereIsNoRoot_ShouldNotWork(t *testing.T) {
	// Arrange
	tr := tree.New[int]()

	// Act
	newTree, ok := tr.Filter(func(obj int) bool {
		return obj%2 == 0
	})

	// Assert
	assert.False(t, ok)
	assert.Nil(t, newTree)
}

func TestTree_Filter_WhenThereIsRootButRulesDoesntApply_ShouldNotWork(t *testing.T) {
	// Arrange
	tr := tree.New[int]()
	tr.AddRoot(node.New(1).WithID(1))

	// Act
	newTree, ok := tr.Filter(func(obj int) bool {
		return obj%2 == 0
	})

	// Assert
	assert.False(t, ok)
	assert.Nil(t, newTree)
}

func TestTree_Filter_WhenEverythingIsOk_ShouldWork(t *testing.T) {
	// Arrange
	tr := tree.New[int]()

	tr.AddRoot(node.New(0).WithID(0))
	tr.Add(0, node.New(1).WithID(1))
	tr.Add(0, node.New(2).WithID(2))
	tr.Add(1, node.New(3).WithID(3))

	// Act
	newTree, ok := tr.Filter(func(obj int) bool {
		return obj%2 == 0
	})

	newN0, _ := newTree.GetRoot()

	// Assert
	assert.True(t, ok)
	assert.Equal(t, 0, newN0.GetID())

	nexts := newN0.GetNexts()
	assert.Equal(t, 2, nexts[0].GetID())
}
