package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Sort(t *testing.T) {
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
}

func TestQueue_Dequeue(t *testing.T) {
	assert := assert.New(t)
	actual, _ := Make(1, 2, 3, 4, 6)
	actual.Dequeue()
	expected, _ := Make(2, 3, 4, 6)
	assert.Equal(expected.size, actual.size, fmt.Sprintf("%v and %v not equal, but expected", actual.size,
		expected.size))
	q, _ := Make()
	err := q.Dequeue()
	assert.Error(err)
}

func TestQueue_Enqueue(t *testing.T) {
	assert := assert.New(t)
	q, _ := Make()
	q.Enqueue(0)
	err := q.Enqueue("")
	assert.Error(err)
	expected, _ := Make()
	expected.Enqueue(1)
	expected.Enqueue(2)
	expected.Enqueue(3)
	expected.Enqueue(4)
	expected.Enqueue(6)
	actual, _ := Make(1, 2, 3, 4, 6)
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	q, _ := Make(1, 2, 3, 4, 6)
	actual := q.IsEmpty()
	expected := false
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
	q, _ = Make()
	actual = q.IsEmpty()
	expected = true
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_Peek(t *testing.T) {
	assert := assert.New(t)
	q, _ := Make(1, 2, 3, 4, 6)
	actual, _ := q.Peek()
	expected := 1
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
	q, _ = Make()
	_, err := q.Peek()
	assert.Error(err)
}

func TestQueue_IsFull(t *testing.T) {
	assert := assert.New(t)
	q, _ := Make(1, 2, 3, 4, 6)
	actual := q.IsFull()
	expected := true
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
	q, _ = Make()
	actual = q.IsFull()
	expected = false
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_Make(t *testing.T) {
	assert := assert.New(t)
	_, err := Make(true, "")
	assert.Error(err)
	expected, _ := Make()
	expected.Enqueue(1)
	expected.Enqueue(2)
	expected.Enqueue(3)
	expected.Enqueue(4)
	expected.Enqueue(6)
	actual, _ := Make(1, 2, 3, 4, 6)
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}
