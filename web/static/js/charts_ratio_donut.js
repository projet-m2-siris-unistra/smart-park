var zones = window.zones;


zones.forEach(item => {
    
    zone = JSON.parse(item);

    // grab chart holder DOM element
    const donutChartHolder = document.getElementById('sp-graphic-ratio-donut-' + zone.id);
    
    const donutChartData = {
        labels : ["Occupées", "Libres"],
        datasets : [
            {
                data : [zone.nb_taken_spots, zone.nb_total_spots-zone.nb_taken_spots],
                fillColors : ['#fa4d56', '#42be65']
            }
        ]
    };
    
    const donutChartOptions = {
        height : '200px',
        width : '233px'
    };
    
    new Charts.DonutChart(donutChartHolder, {
        data : donutChartData,
        options : donutChartOptions,
    });

});