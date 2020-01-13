// grab chart holder DOM element
const chartHolder = document.getElementById('sp-graphic-zone-pie');

const pieChartData = {
    labels : ['Occupées', 'Libres'],
    datasets : [
        {
            data : [window.spotsTaken, window.spotsCount-window.spotsTaken],
            fillColors : ['#fa4d56', '#42be65']
        }
    ]
};

const pieChartOptions = {
    //title : 'Actuellement',
    theme : 'g10', // not working, please fix !
    height : '230px',
    //width : '223px'
};

new Charts.PieChart(chartHolder, {
    data : pieChartData,
    options : pieChartOptions,
});