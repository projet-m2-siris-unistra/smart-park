// This var ensures that only one spot at the time is added & validated
var polygonClickEnabled = true;
// adding a limit of spots to be added
var count = 0;
var spotsLimit = 20;
var currentMarker;
var addedMarkerList = [];

// When clicking on the zone: adds a marker
map.on('click', 'zone-polygon', function(e) {
    if (polygonClickEnabled) {
        var coordinates = e.lngLat.wrap();
        currentMarker = new mapboxgl.Marker({
            draggable: true
        })
            .setLngLat(coordinates)
            .setPopup(new mapboxgl.Popup({ offset: 25 })
                .setHTML(
                    "<a class=\"bx--btn bx--btn--primary \" onclick=\"validate_marker()\">Valider</a>"
                    + "<a class=\"bx--btn bx--btn--danger \" onclick=\"delete_marker()\">Supprimer</a>"
                )
            )
            .addTo(map);
        polygonClickEnabled = false;
    }
});

function validate_marker() {
    console.debug("validate_marker()");

    // modifying the marker
    currentMarker.setDraggable(false);
    currentMarker.setPopup();
    
    // adding this spot coordinates to the list
    var coordinates = currentMarker.getLngLat().wrap();
    console.debug(coordinates);
    addedMarkerList.push(coordinates);

    // if not set, enable validating button

    count++;
    if (count < spotsLimit) {
        polygonClickEnabled = true;
    }
    else {
        alert("Vous avez atteinds la limite de places à ajouter. \
            Veuillez enregistrer votre travail.");
    }
}

function delete_marker() {
    console.debug("delete_marker()");
    currentMarker.remove();
    polygonClickEnabled = true;
}

// Showing popup when
function onDragEnd() {
    // TODO
}