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
    type: 'FeatureCollection',
    features: [{
        type: 'Feature',
        geometry: {
            type: 'Point',
            coordinates: [7.7475, 48.5827]
        },
        properties: {
            title: 'Parking#001',
            description: 'État du parking: OK'
        }
    },
    {
        type: 'Feature',
        geometry: {
            type: 'Point',
            coordinates: [7.7490, 48.5827]
        },
        properties: {
            title: 'Parking#002',
            description: 'État du parking: OK'
        }
    }]
};

/* Add markers to map */
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

/* Loading elements on map */
map.on('load', function() {

    /* Add zone on map */
    map.addLayer({
        'id': 'maine',
        'type': 'fill',
        'source': {
            'type': 'geojson',
            'data': {
                'type': 'Feature',
                'geometry': {
                    'type': 'Polygon',
                    'coordinates': [[
                    [7.739396,48.579816],[7.742014,48.579957],
                    [7.744117,48.579134],[7.747464,48.578623],
                    [7.74888,48.57885],[7.751756,48.579929],
                    [7.755189,48.581831],[7.756906,48.583251],
                    [7.754288,48.58555],[7.753558,48.586061],
                    [7.751455,48.586743],[7.748537,48.58714],
                    [7.746906,48.586828],[7.744503,48.585834],
                    [7.740769,48.584244],[7.73901,48.582967],
                    [7.738409,48.581973],[7.738495,48.580781],
                    [7.739396,48.579816]
                    ]]
                }
            }
        },
        'layout': {},
        'paint': {
            'fill-color': '#f4e628',
            'fill-opacity': 0.2
        }
    });

    // When a click event occurs on a feature in the places layer, open a popup at the
    // location of the feature, with description HTML from its properties.
    map.on('click', 'points', function (e) {
        var coordinates = e.features[0].geometry.coordinates.slice();
        var description = e.features[0].properties.description;
        
        // Ensure that if the map is zoomed out such that multiple
        // copies of the feature are visible, the popup appears
        // over the copy being pointed to.
        while (Math.abs(e.lngLat.lng - coordinates[0]) > 180) {
            coordinates[0] += e.lngLat.lng > coordinates[0] ? 360 : -360;
        }
        
        new mapboxgl.Popup()
            .setLngLat(coordinates)
            .setHTML(description)
            .addTo(map);
    });

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