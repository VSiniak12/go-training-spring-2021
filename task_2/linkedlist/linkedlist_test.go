package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type person struct {
	name string
	age  int
}

func TestLinkedList_Delete(t *testing.T) {
	assert := assert.New(t)
	l, _ := Make(11, 2, 3, 10, 7, 22, 3)
	l.Delete(3)
	expected, _ := Make(11, 2, 3, 7, 22, 3)
	assert.Equal(expected, l, fmt.Sprintf("%s and %s not equal, but expected", l, expected))
	assert.NotPanics(func() { l.Delete(-1) })
	assert.NotPanics(func() { l.Delete(6) })
}

func TestLinkedList_Search(t *testing.T) {
	assert := assert.New(t)
	l, _ := Make(11, 2, 3, 10, 7, 22, 3)
	assert.NotPanics(func() { l.Search(-1) })
	assert.NotPanics(func() { l.Search(8) })
	expected := 10
	actual, _ := l.Search(3)
	assert.Equal(expected, actual.value, fmt.Sprintf("%v and %v not equal, but expected", actual.value, expected))
}

func TestLinkedList_Deletion(t *testing.T) {
	assert := assert.New(t)
	l, _ := Make()
	assert.NotPanics(func() { l.Deletion() })
}

func TestLinkedList_Insert(t *testing.T) {
	assert := assert.New(t)
	l, _ := Make()
	l.Insert(1)
	l.Insert(2)
	err := l.Insert("")
	assert.Error(err)
	assert.NotPanics(func() { l.Deletion() })
}

func TestLinkedList_Sort(t *testing.T) {
	assert := assert.New(t)
	//int
	actual, _ := Make(5, 1, 7, 3, 12, -1)
	actual.Sort()
	expected, _ := Make(12, 7, 5, 3, 1, -1)
	assert.Equal(expected, actual, fmt.Sprintf("Check type int. %v and %v not equal, but expected", actual, expected))
	//string
	actual, _ = Make("q", "a", "c", "d")
	actual.Sort()
	expected, _ = Make("q", "d", "c", "a")
	assert.Equal(expected, actual, fmt.Sprintf("Check type string. %v and %v not equal, but expected", actual,
		expected))
	//rune
	actual, _ = Make('b', 'o', 'c', 'u')
	actual.Sort()
	expected, _ = Make('u', 'o', 'c', 'b')
	assert.Equal(expected, actual, fmt.Sprintf("Check type rune. %v and %v not equal, but expected", actual, expected))
	//float64
	actual, _ = Make(2.3, 7.2, -2.5, 10.5)
	actual.Sort()
	expected, _ = Make(10.5, 7.2, 2.3, -2.5)
	assert.Equal(expected, actual, fmt.Sprintf("Check type float64. %v and %v not equal, but expected", actual,
		expected))
	//some type
	personA := person{
		name: "Slava",
		age:  25,
	}
	personB := person{
		name: "Vlad",
		age:  24,
	}
	actual, _ = Make(personA, personB)
	err := actual.Sort()
	assert.Error(err)
}

func TestLinkedList_Make(t *testing.T) {
	assert := assert.New(t)
	_, err := Make(true, "")
	assert.Error(err)
	expected, _ := Make()
	expected.Insert(1)
	expected.Insert(2)
	expected.Insert(3)
	expected.Insert(4)
	expected.Insert(6)
	actual, _ := Make(1, 2, 3, 4, 6)
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}
