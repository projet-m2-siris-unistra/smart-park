// grab chart holder DOM element
const ChartHolder = document.getElementById('sp-graphic-line-daily');

const chartData = {
    //labels : ['8:00', '12:00', '16:00', '20:00'],
    labels : [],
    datasets : [
        {
            label : "Nombre de places occupées ces derniers temps",
            data : [
                {
                    date : new Date("2020-01-12T00:00:00"),
                    value : 2
                }, 
                {
                    date : new Date("2020-01-12T04:00:00"),
                    value : 2
                }, 
                {
                    date : new Date("2020-01-12T08:00:00"),
                    value : 5
                }, 
                {
                    date : new Date("2020-01-12T12:00:00"),
                    value : 10
                },
                {
                    date : new Date("2020-01-12T16:00:00"),
                    value : 6
                },
                {
                    date : new Date("2020-01-12T20:00:00"),
                    value : 8
                },
                {
                    date : new Date("2020-01-13T00:00:00"),
                    value : 2
                }, 
                {
                    date : new Date("2020-01-13T04:00:00"),
                    value : 2
                }, 
                {
                    date : new Date("2020-01-13T08:00:00"),
                    value : 2
                }, 
                {
                    date : new Date("2020-01-13T12:00:00"),
                    value : 7
                }, 
                {
                    date : new Date("2020-01-13T16:00:00"),
                    value : 3
                }, 
                {
                    date : new Date("2020-01-13T18:00:00"),
                    value : 8
                },
                {
                    date : new Date("2020-01-13T20:00:00"),
                    value : 4
                }
            ],
            fillColors : ['#8a3ffc']
        }
    ]
};


const chartOptions = {
    axes : {
        left : {
            title : "Nombre d'usagers",
            secondary : true
        },
        bottom : {
            //title : "Heure",
            scaleType : "time",
            primary : true
        }
    },
    curve : "curveMonotoneX",
    height : '223px',
    //width : '480px'
};

new Charts.LineChart(ChartHolder, {
    data : chartData,
    options : chartOptions,
});