package messenger_test

import (
	"sync"
	"testing"

	messenger "github.com/cozy-hosting/messenger/src"
	"github.com/stretchr/testify/assert"
)

func TestRoundrobinBalancerPickIsCorrect(t *testing.T) {
	// Arrange
	balancer := messenger.NewRoundrobinBalancer([]interface{}{1, 2, 3})

	// Act
	results := []int{0, 0, 0}
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		item1 := balancer.Pick()
		results[0] = item1.(int)
		balancer.Done()
		wg.Done()
	}()
	go func() {
		item2 := balancer.Pick()
		results[1] = item2.(int)
		balancer.Done()
		wg.Done()
	}()
	go func() {
		item3 := balancer.Pick()
		results[2] = item3.(int)
		balancer.Done()
		wg.Done()
	}()

	wg.Wait()

	// Assert
	assert.Equal(t, []int{1, 2, 3}, results)
}

func TestRoundRobinBalancerPickOnEmptyItemsResultsInNilPick(t *testing.T) {
	// Arrange
	balancer := messenger.NewRoundrobinBalancer([]interface{}{})

	// Act
	item := balancer.Pick()

	// Assert
	assert.Nil(t, item)
}

func TestRoundrobinBalancerPushIsUsedNextPick(t *testing.T) {
	// Arrange
	balancer := messenger.NewRoundrobinBalancer([]interface{}{1, 2})

	// Act
	results := []int{0, 0}
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		_ = balancer.Pick()
		balancer.Push()
		wg.Done()
	}()
	go func() {
		item1 := balancer.Pick()
		results[0] = item1.(int)
		balancer.Done()
		wg.Done()
	}()
	go func() {
		item2 := balancer.Pick()
		results[1] = item2.(int)
		balancer.Done()
		wg.Done()
	}()

	wg.Wait()

	// Assert
	assert.Equal(t, []int{1, 2}, results)
}
