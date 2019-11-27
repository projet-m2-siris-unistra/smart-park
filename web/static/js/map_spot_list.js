var spots = window.spots;
var tenantCoor = window.tenantCoor;
var zone_id = window.zone_id

// Create map
mapboxgl.accessToken = 'pk.eyJ1IjoibGlvbmVsanVuZyIsImEiOiJjazI2azYxY3QwMGtyM2ZvYnJ4ZGY0Mjd0In0.fjgEahiiwwH58znbPwQShA';
var map = new mapboxgl.Map({
    container: 'map',
    style: 'mapbox://styles/mapbox/streets-v11',
    center: tenantCoor,
    zoom: 14
});

// Add zoom and rotation controls to the map.
map.addControl(new mapboxgl.NavigationControl());

if (spots != null) {

    // Add markers to map
    spots.forEach(function(marker) {
        marker = JSON.parse(marker);

        console.debug("marker.coordinates:" + marker.coordinates);

        // create a HTML element for each feature
        var el = document.createElement('div');
        el.className = 'black-marker';

        // make a marker for each feature and add to the map
        new mapboxgl.Marker(el)
            .setLngLat(marker.coordinates)
            .setPopup(new mapboxgl.Popup({ offset: 25 })
                .setHTML(
                    '<h3>' + marker.name 
                    + '</h3><p>' + 'Etat du parking: ' + marker.device.state + '</p>'
                    + "<a class=\"bx--btn bx--btn--primary\" href=\"/parking/zone/"
                        + zone_id
                        +"/spot/"
                        + marker.id 
                        +"\">Voir</a>"
                    + "<a class=\"bx--btn bx--btn--danger \" href=\"#\">Supprimer</a>"
                )
            )
            .addTo(map);
    });
}
else {
    console.debug("Warning: " + "spots = " + spots);
}