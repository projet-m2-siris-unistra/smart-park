var usageRatio = window.usageRatio;
var freeRatio = 100 - usageRatio;

// grab chart holder DOM element
const chartHolder = document.getElementById('sp-graphic-zone-pie');

const pieChartData = {
    labels : ['Occupées', 'Libres'],
    datasets : [
        {
            data : [usageRatio, freeRatio],
            fillColors : ['#9b0404', '#4a9128']
        }
    ]
};

const pieChartOptions = {
    //title : 'Actuellement',
    theme : 'g10', // not working, please fix !
    height : '223px',
    width : '223px'
};

new Charts.PieChart(chartHolder, {
    data : pieChartData,
    options : pieChartOptions,
});