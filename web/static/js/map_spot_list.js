/* Create map */
mapboxgl.accessToken = 'pk.eyJ1IjoibGlvbmVsanVuZyIsImEiOiJjazI2azYxY3QwMGtyM2ZvYnJ4ZGY0Mjd0In0.fjgEahiiwwH58znbPwQShA';
var map = new mapboxgl.Map({
    container: 'map',
    style: 'mapbox://styles/mapbox/streets-v11',
    center: [7.7475, 48.5827],
    zoom: 14
});

/* The parking spot markers list */
var markerList = {
    'type': 'FeatureCollection',
    'features': [{
        'type': 'Feature',
        'geometry': {
            'type': 'Point',
            'coordinates': [7.7475, 48.5827]
        },
        'properties': {
            'title': 'Parking#001',
            'description': 'État du parking: OK'
        }
    },
    {
        'type': 'Feature',
        'geometry': {
            'type': 'Point',
            'coordinates': [7.7490, 48.5827]
        },
        'properties': {
            'title': 'Parking#002',
            'description': 'État du parking: OK'
        }
    }]
};

// Add markers to map
markerList.features.forEach(function(marker) {
 
    // create a HTML element for each feature
    var el = document.createElement('div');
    el.className = 'black-marker';
 
    // make a marker for each feature and add to the map
    new mapboxgl.Marker(el)
        .setLngLat(marker.geometry.coordinates)
        .setPopup(new mapboxgl.Popup({ offset: 25 })
            .setHTML(
                '<h3>' + marker.properties.title 
                + '</h3><p>' + marker.properties.description + '</p>'
                + "<a class=\"bx--btn bx--btn--primary\" href=\"#\">Éditer</a>"
                + "<a class=\"bx--btn bx--btn--tertiary\" href=\"#\">Statistiques</a>"
                + "<a class=\"bx--btn bx--btn--danger \" href=\"#\">Supprimer</a>"
            )
        )
        .addTo(map);
});