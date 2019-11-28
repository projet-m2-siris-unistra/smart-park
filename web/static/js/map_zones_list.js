var zones = window.zones;

if (zones != null) {
        
    zones.forEach(function(zone) {

        zone = JSON.parse(zone);
        
        // drawing the zone
        map.on('load', function() {
            
            console.debug(zone);

            polygonGeoJson = {
                'id': 'zone-polygon-' + zone.name,
                'type': 'fill',
                'source': {
                    'type': 'geojson',
                    'data': {
                        'type': 'Feature',
                        'geometry': {
                            'type': 'Polygon',
                            'coordinates': [JSON.parse(zone.coordinates)]
                        }
                    }
                },
                'paint': {
                    'fill-color': zone.color,
                    'fill-opacity': 0.2
                }
            }
            console.debug(polygonGeoJson);
            map.addLayer(polygonGeoJson);
        });

        // displaying popup with buttons on click
        map.on('click', 'zone-polygon-'+zone.name, function(e) {
            console.debug("zone-polygon-"+zone.name+" clicked");
            new mapboxgl.Popup()
                .setLngLat(e.lngLat)
                .setHTML(
                    "<h3>" + zone.name + "</h3>"
                    + "<a class=\"bx--btn bx--btn--primary\" href=\"/parking/"+ zone.name +"\">Voir</a>"
                )
                .addTo(map);

            // display the spots
            zone.spots.forEach(function(spot) {
                var spot = JSON.parse(spot);

                var el = document.createElement('div');
                el.className = 'black-marker';

                new mapboxgl.Marker(el)
                    .setLngLat(spot.point.geometry.coordinates)
                    .setPopup(new mapboxgl.Popup({ offset: 25 })
                        .setHTML(
                            '<h3>' + spot.name 
                            + '</h3><p>' + 'Etat du parking: ' + spot.state + '</p>'
                            + "<a class=\"bx--btn bx--btn--primary\" href=\"/spot/"+ spot.name +"\">Voir</a>"
                            + "<a class=\"bx--btn bx--btn--danger \" href=\"#\">Supprimer</a>"
                        )
                    )
                    .addTo(map);

            });
        });

    });
}
else {
    console.debug("Warning: " + "zones=" + zones);
}