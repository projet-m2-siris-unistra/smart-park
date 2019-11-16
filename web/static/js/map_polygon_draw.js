var polygon = window.polygon;

if (polygon != null) {
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

    });
}
else {
    console.debug("Warning: " + "polygon=" + polygon);
}