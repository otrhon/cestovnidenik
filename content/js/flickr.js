var Flickr = (function () {
    function Flickr() {
        this.user = "131912190@N07";
    }
    Flickr.prototype.getAlbumsJson = function () {
        var src = "https://api.flickr.com/services/rest/?method=flickr.photosets.getList&api_key=5a96fe21b2b6fce8d8f8278a2eb0286f&user_id=" + this.user + "&format=json&jsoncallback=getAlbums";
        var script = document.createElement("script");
        script.src = src;
        document.body.appendChild(script);
    };
    return Flickr;
}());
var gscript = null;
function getJson(album) {
    if (gscript != null) {
        document.body.removeChild(gscript);
        document.querySelector("#flickrPhotos").innerHTML = "";
        gscript = null;
    }
    if (album == "0") {
        return;
    }
    var src = "https://api.flickr.com/services/rest/?method=flickr.photosets.getPhotos&api_key=5a96fe21b2b6fce8d8f8278a2eb0286f&photoset_id=" + album + "&user_id=131912190%40N07&extras=url_s,url_l,url_m&format=json&jsoncallback=getPhotos";
    gscript = document.createElement("script");
    gscript.src = src;
    document.body.appendChild(gscript);
}
function getAlbums(data) {
    if (data && data.photosets) {
        var items = data.photosets.photoset;
        var html = "<select onchange='getJson(this.options[this.selectedIndex].value)'>";
        html += "<option value='0'>-- Výběr alba --</option>";
        var tmpAry = new Array();
        for (var i = 0; i < items.length; ++i) {
            var item = items[i];
            tmpAry[i] = new Array();
            tmpAry[i][0] = item.title._content;
            tmpAry[i][1] = item.id;
        }
        tmpAry.sort();
        for (var i = 0; i < tmpAry.length; i++) {
            html += "<option value='" + tmpAry[i][1] + "'>" + tmpAry[i][0] + "</option>";
        }
        html += "</select>";
        document.querySelector("#flickrAlbums").innerHTML = html;
    }
}
function add(url) {
    this.tinymce.activeEditor.execCommand("mceInsertContent", false, "<img style='display: block; padding: 5px; margin: 0 auto' src='" + url + "'>");
}
function getPhotos(data) {
    if (data && data.photoset) {
        var items = data.photoset.photo;
        var html = "<div class='images'>";
        for (var i = 0; i < items.length; ++i) {
            var item = items[i];
            html += "<a href='#' onclick='add(\"" + item.url_m + "\")'><img width='" + item.width_s + "' height='" + item.height_s + "' src='" + item.url_s + "' alt='' /></a>";
        }
        html += "</div>";
        document.querySelector("#flickrPhotos").innerHTML = html;
    }
}
document.addEventListener("DOMContentLoaded", function () {
    var flickrFeed = new Flickr();
    flickrFeed.getAlbumsJson();
});
