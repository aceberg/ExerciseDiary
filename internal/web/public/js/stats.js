var sChart = null;

function addSet(date, reps, weight) {

    html_code = '<tr><td>'+date+'</td><td>'+reps+'</td><td>'+weight+'</td></tr>';

    document.getElementById('stats-table').insertAdjacentHTML('beforeend', html_code);
};


function setStatsPage(sets, hcolor) {
    let dates = [], ws = []; reps = [];

    let ex = document.getElementById("ex-value").value;
    console.log("EX =", ex);

    let start = 0;
    let end = sets.length;

    document.getElementById('stats-table').innerHTML = "";

    for (let i = start ; i < end; i++) {
        if (sets[i].Name === ex) {
            addSet(sets[i].Date, sets[i].Reps, sets[i].Weight);

            dates.push(sets[i].Date);
            reps.push(sets[i].Reps);
            ws.push(sets[i].Weight);
        };
    };

    console.log("REPS =", reps);

    statsChart('stats-reps', dates, reps, hcolor, true);
    weightChart('stats-weight', dates, ws, hcolor, true);
};

function statsChart(id, dates, ws, wcolor, xticks) {
    
    const ctx = document.getElementById(id);

    if (sChart){
      sChart.clear();
      sChart.destroy();
    };

    sChart = new Chart(ctx, {
      type: 'bar',
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