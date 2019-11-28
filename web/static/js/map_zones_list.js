var zones = window.zones;

if (zones != null) {
        
    zones.forEach(function(zone) {

        zone = JSON.parse(zone);
        
        // drawing the zone
        map.on('load', function() {
            
            console.debug("ZONE JSON: ");
            console.debug(zone);

            polygonGeoJson = {
                'id': 'zone-polygon-' + zone.id,
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
            map.addLayer(polygonGeoJson);
        });

        // displaying popup with buttons on click
        map.on('click', 'zone-polygon-'+zone.id, function(e) {
            
            console.debug("zone-polygon-"+zone.name+" clicked");
            
            // popup with link to zone
            new mapboxgl.Popup()
                .setLngLat(e.lngLat)
                .setHTML(
                    "<h3>" + zone.name + "</h3>"
                    + "<a class=\"bx--btn bx--btn--primary\" href=\"/parking/"+ zone.name +"\">Voir</a>"
                )
                .addTo(map);

            // displaying the spots of the zone
            zone.spots.forEach(function(spot) {
                                
                console.debug("SPOT JSON:");
                console.debug(spot);

                var el = document.createElement('div');
                el.className = 'black-marker';

                new mapboxgl.Marker(el)
                    .setLngLat(JSON.parse(spot.coordinates))
                    .setPopup(new mapboxgl.Popup({ offset: 25 })
                        .setHTML(
                            '<h3>' + spot.name 
                            + '</h3><p>' + 'Etat du parking: ' + spot.state + '</p>'
                            + "<a class=\"bx--btn bx--btn--primary"
                            + "href=\"{{ url_for(\"spot.overview\", zone_id=zone.id, spot_id=spot.id) }}"
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