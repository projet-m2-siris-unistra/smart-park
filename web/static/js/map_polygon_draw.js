var polygon = window.polygon;
var color = window.color;
var zone_id = window.zone_id;

if (polygon != null) {
    // Loading elements on map 
    map.on('load', function() {
        
        // Add zone on map
        polygonGeoJson = {
            'id': 'zone-polygon-' + zone_id,
            'type': 'fill',
            'source': {
                'type': 'geojson',
                'data': {
                    'type': 'Feature',
                    'geometry': {
                        'type': 'Polygon',
                        'coordinates': [JSON.parse(polygon)]
                    }
                }
            },
            'paint': {
                'fill-color': color,
                'fill-opacity': 0.2
            }
        }
        map.addLayer(polygonGeoJson);
    });
}
else {
    console.debug("Warning: " + "polygon=" + polygon);
}