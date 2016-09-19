(function() {
	function Flickr() {
		this.init();
	}
	
	Flickr.prototype = {
		init: function() {
			this.user = "131912190@N07";
			this.album = "72157672119809916";
			
			window.getPhotos = this.getPhotos;
			
			this.getJSON();	
		},
		getJSON: function() {
			var src = "http://api.flickr.com/services/feeds/photoset.gne?nsid=" + this.user + "&set=" + this.album + "&format=json&jsoncallback=getPhotos";	
			var script = document.createElement( "script" );
				script.src = src;
				document.body.appendChild( script );
		},
		getPhotos: function( data ) {
			var limit = 1000;
			
			if( data && data.items ) {
				var title = data.title;
				var items = data.items;
				var albumTitle = title.replace( "Content from ", "" );
				var html = "<h3>" + albumTitle + "</h3>";
				    html += "<div class='images'>";
				
				for( var i = 0; i < items.length; ++i ) {
					var item = items[i];
					var n = i + 1;
					if( n <= limit ) {
						html += "<a href='" + item.link + "'><img src='" + item.media.m + "' alt='' /></a>";
					}	
				}
				
				html += "</div>";
				
				document.querySelector( "#flickr" ).innerHTML = html;
			}	
		}
	};
	
	document.addEventListener( "DOMContentLoaded", function() {
		var flickrFeed = new Flickr();
	});
})();
