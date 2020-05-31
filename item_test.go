package podcast_test

import (
	"testing"

	"github.com/eduncan911/podcast"
	"github.com/stretchr/testify/assert"
)

func TestItemAddSummaryTooLong(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	summary := ""
	for {
		if len(summary) >= 4051 {
			break
		}
		summary += "abc ss 5 "
	}

	// act
	i.AddSummary(summary)

	// assert
	assert.Len(t, i.ISummary.Text, 4000)
}

func TestItemAddImageEmptyUrl(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}

	// act
	i.AddImage("")

	// assert
	assert.Nil(t, i.IImage)
}

func TestItemAddDurationZero(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	d := int64(0)

	// act
	i.AddDuration(d)

	// assert
	assert.EqualValues(t, "", i.IDuration)
}

func TestItemAddDurationLessThanZero(t *testing.T) {
	t.Parallel()

	// arrange
	i := podcast.Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	d := int64(-13)

	// act
	i.AddDuration(d)

	// assert
	assert.EqualValues(t, "", i.IDuration)
}

func TestItemAddEpisodeType(t *testing.T) {
	t.Parallel()

	i := podcast.Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	// act
	i.AddEpisodeType("")
	// assert
	assert.EqualValues(t, "", i.IEpisodeType)
	// act
	i.AddEpisodeType("test")
	// assert
	assert.EqualValues(t, "test", i.IEpisodeType)
}

func TestItemAddEpisode(t *testing.T) {
	t.Parallel()

	i := podcast.Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	// act
	i.AddEpisode(1)
	// assert
	assert.EqualValues(t, 1, i.IEpisode)

}

func TestItemAddSeason(t *testing.T) {
	t.Parallel()

	i := podcast.Item{
		Title:       "item.title",
		Description: "item.desc",
		Link:        "http://example.com/article.html",
	}
	// act
	i.AddSeason(1)
	// assert
	assert.EqualValues(t, 1, i.ISeason)

}
