package feeds

import "html"

func (f *RSSFeed) unescapeTitelDescription() {

	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)
	for i, item := range f.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		f.Channel.Item[i] = item
	}
}
