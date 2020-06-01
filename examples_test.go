package podcast_test

import (
	"fmt"
	"os"
	"strconv"

	"github.com/eduncan911/podcast"
)

func ExampleNew() {
	ti, l, d := "title", "link", "description"

	// instantiate a new Podcast
	p := podcast.New(ti, l, d, &pubDate, &updatedDate)

	fmt.Println(p.Title, p.Link, p.Description.Text, p.Language)
	fmt.Println(p.PubDate, p.LastBuildDate)
	// Output:
	// title link description en-us
	// Sat, 04 Feb 2017 08:21:52 +0000 Mon, 06 Feb 2017 08:21:52 +0000
}

func ExamplePodcast_AddAuthor() {
	p := podcast.New("title", "link", "description", nil, nil)

	// add the Author
	p.AddAuthor("the name", "me@test.com")

	fmt.Println(p.ManagingEditor)
	fmt.Println(p.IAuthor)
	// Output:
	// me@test.com (the name)
	// me@test.com (the name)
}

func ExamplePodcast_AddCategory() {
	p := podcast.New("title", "link", "description", nil, nil)

	// add the Category
	p.AddCategory("Bombay", nil)
	p.AddCategory("American", []string{"Longhair", "Shorthair"})
	p.AddCategory("Siamese", nil)

	fmt.Println(len(p.ICategories), len(p.ICategories[1].ICategories))
	fmt.Println(p.Category)
	// Output:
	// 3 2
	// Bombay,American,Siamese
}

func ExamplePodcast_AddImage() {
	p := podcast.New("title", "link", "description", nil, nil)

	// add the Image
	p.AddImage("http://example.com/image.jpg")

	if p.IImage != nil {
		fmt.Println(p.IImage.HREF)
	}
	// Output:
	// http://example.com/image.jpg
}

func ExamplePodcast_AddItem() {
	p := podcast.New("title", "link", "description", &pubDate, &updatedDate)
	p.AddAuthor("the name", "me@test.com")
	p.AddImage("http://example.com/image.jpg")

	// create an Item
	date := pubDate.AddDate(0, 0, 77)
	item := podcast.Item{
		Title: "Episode 1",
		Description: &podcast.Description{
			Text: "Description for Episode 1",
		},
		PubDate: &date,
	}
	item.AddEnclosure(
		"http://example.com/1.mp3",
		podcast.MP3,
		183,
	)

	// add the Item
	if _, err := p.AddItem(item); err != nil {
		fmt.Println("item validation error: " + err.Error())
	}

	if len(p.Items) != 1 {
		fmt.Println("expected 1 item in the collection")
	}
	pp := p.Items[0]
	fmt.Println(
		pp.GUID, pp.Title, pp.Link, pp.Description.Text, pp.Author,
		pp.AuthorFormatted, pp.Category, pp.Comments, pp.Source,
		pp.PubDate, pp.PubDateFormatted, *pp.Enclosure,
		pp.IAuthor, pp.IDuration, pp.IExplicit, pp.IIsClosedCaptioned,
		pp.IOrder)
	// Output:
	// http://example.com/1.mp3 Episode 1 http://example.com/1.mp3 Description for Episode 1 &{{ }  me@test.com (the name)}     2017-04-22 08:21:52 +0000 UTC Sat, 22 Apr 2017 08:21:52 +0000 {{ } http://example.com/1.mp3 183 183 audio/mpeg audio/mpeg} me@test.com (the name)
}

func ExamplePodcast_AddLastBuildDate() {
	p := podcast.New("title", "link", "description", nil, nil)
	d := pubDate.AddDate(0, 0, -7)

	p.AddLastBuildDate(&d)

	fmt.Println(p.LastBuildDate)
	// Output:
	// Sat, 28 Jan 2017 08:21:52 +0000
}

func ExamplePodcast_AddPubDate() {
	p := podcast.New("title", "link", "description", nil, nil)
	d := pubDate.AddDate(0, 0, -5)

	p.AddPubDate(&d)

	fmt.Println(p.PubDate)
	// Output:
	// Mon, 30 Jan 2017 08:21:52 +0000
}

func ExamplePodcast_Bytes() {
	p := podcast.New(
		"eduncan911 Podcasts",
		"http://eduncan911.com/",
		"An example Podcast",
		&pubDate, &updatedDate,
	)
	p.AddAuthor("Jane Doe", "me@janedoe.com")
	p.AddImage("http://janedoe.com/i.jpg")

	for i := int64(5); i < 7; i++ {
		n := strconv.FormatInt(i, 10)
		d := pubDate.AddDate(0, 0, int(i+3))

		item := podcast.Item{
			Title: "Episode " + n,
			Link:  "http://example.com/" + n + ".mp3",
			Description: &podcast.Description{
				Text: "Description for Episode " + n,
			},
			PubDate: &d,
		}
		item.AddImage("http://janedoe.com/i.jpg")
		if _, err := p.AddItem(item); err != nil {
			fmt.Println(item.Title, ": error", err.Error())
			break
		}
	}

	// call Podcast.Bytes() to return a byte array
	os.Stdout.Write(p.Bytes())

	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
	//   <channel>
	//     <title>eduncan911 Podcasts</title>
	//     <link>http://eduncan911.com/</link>
	//     <description><![CDATA[An example Podcast]]></description>
	//     <generator>go podcast v1.3.1 (github.com/eduncan911/podcast)</generator>
	//     <language>en-us</language>
	//     <lastBuildDate>Mon, 06 Feb 2017 08:21:52 +0000</lastBuildDate>
	//     <managingEditor>me@janedoe.com (Jane Doe)</managingEditor>
	//     <pubDate>Sat, 04 Feb 2017 08:21:52 +0000</pubDate>
	//     <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//     <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//     <item>
	//       <guid>http://example.com/5.mp3</guid>
	//       <title>Episode 5</title>
	//       <link>http://example.com/5.mp3</link>
	//       <description><![CDATA[Description for Episode 5]]></description>
	//       <pubDate>Sun, 12 Feb 2017 08:21:52 +0000</pubDate>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//     </item>
	//     <item>
	//       <guid>http://example.com/6.mp3</guid>
	//       <title>Episode 6</title>
	//       <link>http://example.com/6.mp3</link>
	//       <description><![CDATA[Description for Episode 6]]></description>
	//       <pubDate>Mon, 13 Feb 2017 08:21:52 +0000</pubDate>
	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
	//     </item>
	//   </channel>
	// </rss>
}

func ExampleItem_AddPubDate() {
	p := podcast.New("title", "link", "description", nil, nil)
	i := podcast.Item{
		Title: "item title",
		Description: &podcast.Description{
			Text: "item description",
		},
		Link: "item link",
	}
	d := pubDate.AddDate(0, 0, -11)

	// add the pub date
	i.AddPubDate(&d)

	// before adding
	if i.PubDate != nil {
		fmt.Println(i.PubDateFormatted, *i.PubDate)
	}

	// this should not override with Podcast.PubDate
	if _, err := p.AddItem(i); err != nil {
		fmt.Println(err)
	}

	// after adding item
	fmt.Println(i.PubDateFormatted, *i.PubDate)
	// Output:
	// Tue, 24 Jan 2017 08:21:52 +0000 2017-01-24 08:21:52 +0000 UTC
	// Tue, 24 Jan 2017 08:21:52 +0000 2017-01-24 08:21:52 +0000 UTC
}

func ExampleItem_AddDuration() {
	i := podcast.Item{
		Title: "item title",
		Description: &podcast.Description{
			Text: "item desc",
		},
		Link: "item link",
	}
	d := int64(533)

	// add the Duration in Seconds
	i.AddDuration(d)

	fmt.Println(i.IDuration)
	// Output:
	// 8:53
}
