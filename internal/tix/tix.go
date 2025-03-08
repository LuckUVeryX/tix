package tix

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

const rangeError = "index out of range"

type Tix struct {
	Title       string     `json:"title"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

type Tixs []Tix

func (t *Tixs) Add(title string) {
	*t = append(*t, Tix{
		Title:     title,
		CreatedAt: time.Now(),
	})
	t.Print()
}

func (t *Tixs) Delete(idx int) error {
	if idx < 0 || idx >= len(*t) {
		return errors.New(rangeError)
	}
	*t = append((*t)[:idx], (*t)[idx+1:]...)
	t.Print()
	return nil
}

func (t *Tixs) Toggle(idx int) error {
	if idx < 0 || idx >= len(*t) {
		return errors.New(rangeError)
	}
	tix := &(*t)[idx]
	if tix.CompletedAt != nil {
		tix.CompletedAt = nil
	} else {
		now := time.Now()
		tix.CompletedAt = &now
	}
	t.Print()
	return nil
}

func (t *Tixs) Edit(idx int, title string) error {
	if idx < 0 || idx >= len(*t) {
		return errors.New(rangeError)
	}
	if title != "" {
		(*t)[idx].Title = title
	}
	t.Print()
	return nil
}

func (t *Tixs) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "", "Title", "Created At")
	for index, t := range *t {
		completed := "❌"
		if t.CompletedAt != nil {
			completed = "✅"
		}

		table.AddRow(strconv.Itoa(index), completed, t.Title, t.CreatedAt.Format(time.RFC1123))
	}

	table.Render()
}
