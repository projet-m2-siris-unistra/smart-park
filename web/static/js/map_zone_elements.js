var zones = window.zones;

if (zones != null ) {

    zones.forEach(function(zone) {
    
        zone = JSON.parse(zone);
        console.debug(zone);
    
        // 1. Create the map associated with the right div (id)
        mapboxgl.accessToken = 'pk.eyJ1IjoibGlvbmVsanVuZyIsImEiOiJjazI2azYxY3QwMGtyM2ZvYnJ4ZGY0Mjd0In0.fjgEahiiwwH58znbPwQShA';
            var map = new mapboxgl.Map({
            container: 'map-' + zone.id,
            style: 'mapbox://styles/mapbox/streets-v11',
            // center the map on firt coor of zone
            center: JSON.parse(zone.coordinates)[3],
            zoom: 13.5
        });

        // 2. Draw the zone
        map.on('load', function() {
        
            // Add zone on map
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
    
    });

}
else {
    console.debug("WARNING: zones is null, can't create maps.")
}