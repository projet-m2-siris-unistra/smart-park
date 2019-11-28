var spots = window.spots;
var tenantCoor = window.tenantCoor;
var zone_id = window.zone_id

if (spots != null) {

    // Add markers to map
    spots.forEach(function(marker) {
        marker = JSON.parse(marker);

        // create a HTML element for each feature
        var el = document.createElement('div');
        el.className = 'black-marker';

        // make a marker for each feature and add to the map
        new mapboxgl.Marker(el)
            .setLngLat(JSON.parse(marker.coordinates))
            .setPopup(new mapboxgl.Popup({ offset: 25 })
                .setHTML(
                    '<h3>' + marker.name 
                    + '</h3><p>' + 'Etat du parking: ' + marker.device.state + '</p>'
                    + "<a class=\"bx--btn bx--btn--primary\" href=\"/parking/zone/"
                    + zone_id + "/spot/" + marker.id + "\">Voir</a>"
                    + "<a class=\"bx--btn bx--btn--danger \" href=\"#\">Supprimer</a>"
                )
            )
            .addTo(map);
    });
}
else {
    console.debug("Warning: " + "spots = " + spots);
}