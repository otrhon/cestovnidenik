class Flickr {
	user = "131912190@N07"
	album = "72157672119809916"

	constructor() {
		this.getJson()
	}

	getJson() {
		var src = "http://api.flickr.com/services/feeds/photoset.gne?nsid=" + this.user + "&set=" + this.album + "&format=json&jsoncallback=getPhotos";
		var script = document.createElement("script");
		script.src = src;
		document.body.appendChild(script);
	}
}

function getPhotos(data) {

	if (data && data.items) {
		var title = data.title;
		var items = data.items;
		var albumTitle = title.replace("Content from ", "");
		var html = "<h3>" + albumTitle + "</h3>";
		html += "<div class='images'>";

		for (var i = 0; i < items.length; ++i) {
			var item = items[i];
			html += "<a href='" + item.link + "'><img src='" + item.media.m + "' alt='' /></a>";
		}

		html += "</div>";
		document.querySelector("#flickr").innerHTML = html;
	}
}

document.addEventListener("DOMContentLoaded", function () {
	var flickrFeed = new Flickr();
});