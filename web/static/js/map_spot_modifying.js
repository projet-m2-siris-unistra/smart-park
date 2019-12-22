var currentMarker;
var zone_id = window.zone_id;
var spotCoordinates = window.spotCoor; 

console.debug("spotCoordinates : ", spotCoordinates);

// Displaying marker at current position
var coordinates = {
    "lng" : spotCoordinates[0], 
    "lat" : spotCoordinates[1]
};
currentMarker = new mapboxgl.Marker({
    draggable: true
    })
    .setLngLat(coordinates)
    .setPopup(new mapboxgl.Popup({ offset: 25 })
        .setHTML(
            "<a class=\"bx--btn bx--btn--primary \" onclick=\"validate_marker()\">Valider</a>"
        )
    )
    .addTo(map);


// Changing coordinates when validating
function validate_marker() {
    console.debug("validate_marker()");

    // modifying the marker
    //currentMarker.setDraggable(false);
    //currentMarker.setPopup();
    
    // adding this spot coordinates to the list
    var coordinates = currentMarker.getLngLat().wrap();
    var coorString = coordinates.lng + "," + coordinates.lat;
    console.debug(coorString);

    // push the coordinates into the input field
    document.getElementById('coordinates').value = coorString;

}