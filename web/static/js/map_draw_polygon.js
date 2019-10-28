
L.mapbox.accessToken = 'pk.eyJ1IjoibGlvbmVsanVuZyIsImEiOiJjazI2azYxY3QwMGtyM2ZvYnJ4ZGY0Mjd0In0.fjgEahiiwwH58znbPwQShA';
var map = L.mapbox.map('map')
    .setView([48.5827, 7.7475], 14)
    .addLayer(L.mapbox.styleLayer('mapbox://styles/mapbox/streets-v11'));

var featureGroup = L.featureGroup().addTo(map);

var drawControl = new L.Control.Draw({
    edit: {
        featureGroup: featureGroup
    },
    draw: {
        polygon: true,
        polyline: false,
        rectangle: false,
        circle: false,
        marker: false
    }
}).addTo(map);

map.on('draw:created', showPolygonArea);
map.on('draw:edited', showPolygonAreaEdited);

function showPolygonAreaEdited(e) {
    e.layers.eachLayer(function(layer) {
        showPolygonArea({ layer: layer });
    });
}

function showPolygonArea(e) {
    featureGroup.clearLayers();
    featureGroup.addLayer(e.layer);
    area = (LGeo.area(e.layer) / 1000000).toFixed(2);
    
    e.layer.bindPopup(area + ' km<sup>2</sup>');
    e.layer.openPopup();

    // Get the GeoJson of the drawn polygon
    var shape_string = JSON.stringify(e.layer.toGeoJSON());
    var coor = e.layer.toGeoJSON().geometry.coordinates[0];
    console.debug(shape_string);
    console.debug(coor);
    document.getElementById('coordinates-input').value = coor;
}