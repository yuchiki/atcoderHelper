package repository

import (
	"testing"
)

func TestFetchTasks(t *testing.T) {
	tasks, err := FetchTasks("arc100")

	t.Error(tasks, err)
}
