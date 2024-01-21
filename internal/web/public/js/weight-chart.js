var wChart = null;

function splitWeight(weight, show) {
    var dates = [];
    var ws = [];

    weight = weight.slice(show)
    let arrayLength = weight.length;
    for (let i = 0 ; i < arrayLength; i++) {
        dates.push(weight[i].Date);
        ws.push(weight[i].Weight);
    }
    // console.log('LDATA =', dates, ws);
    return { dates, ws };
};

function weightChart(id, dates, ws, wcolor, xticks) {
    
    const ctx = document.getElementById(id);

    if (wChart){
      wChart.clear();
      wChart.destroy();
    };

    wChart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: dates,
        datasets: [{
          data: ws,
          borderColor: wcolor,
          borderWidth: 1
        }]
      },
      options: {
        scales: {
          x: {
            ticks: {
                display: xticks
            },
          },
          y: {
            beginAtZero: false
          }
        },
        plugins:{
            legend: {
             display: false
            }
        }
      }
    });
};

function generateWeightChart(weight, wcolor, show) {
  if (weight) {
    let { dates, ws } = splitWeight(weight, show);

    weightChart('weight-chart', dates, ws, wcolor, false);
  };
};