var polygon = window.polygon;
var spots = window.spots;

// Create map
mapboxgl.accessToken = 'pk.eyJ1IjoibGlvbmVsanVuZyIsImEiOiJjazI2azYxY3QwMGtyM2ZvYnJ4ZGY0Mjd0In0.fjgEahiiwwH58znbPwQShA';
var map = new mapboxgl.Map({
    container: 'map',
    style: 'mapbox://styles/mapbox/streets-v11',
    center: [7.7475, 48.5827],
    zoom: 14
});


// Add markers to map
spots.forEach(function(marker) {

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

// Loading elements on map 
map.on('load', function() {

    // Add zone on map
    polygonGeoJson = {
        'id': 'zone-polygon',
        'type': 'fill',
        'source': {
            'type': 'geojson',
            'data': {
                'type': 'Feature',
                'geometry': {
                    'type': 'Polygon',
                    'coordinates': [polygon.coordinates]
                }
            }
        },
        'paint': {
            'fill-color': polygon.color,
            'fill-opacity': 0.2
        }
    }
    map.addLayer(polygonGeoJson);
    
    // Change the cursor to a pointer when the mouse is over the places layer.
    map.on('mouseenter', 'points', function () {
        map.getCanvas().style.cursor = 'pointer';
    });
    
    // Change it back to a pointer when it leaves.
    map.on('mouseleave', 'points', function () {
        map.getCanvas().style.cursor = '';
    });

    // Add zoom and rotation controls to the map.
    map.addControl(new mapboxgl.NavigationControl());
});