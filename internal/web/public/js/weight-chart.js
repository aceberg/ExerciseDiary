var dates = [];
var ws = [];

function splitWeight(weight, show) {
    dates = [];
    ws = [];

    weight = weight.slice(show)
    let arrayLength = weight.length;
    for (let i = 0 ; i < arrayLength; i++) {
        dates.push(weight[i].Date);
        ws.push(weight[i].Weight)
    }
    // console.log('LDATA =', dates, ws);
};

function weightChart(weight, wcolor, show) {
    
  if (weight) {
    splitWeight(weight, show);
    // console.log('FDATA =', dates, ws);

    const ctx = document.getElementById('weight-chart');

    new Chart(ctx, {
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
                display: false
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
  }
};