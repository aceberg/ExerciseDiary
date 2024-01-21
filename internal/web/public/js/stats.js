var sChart = null;
var sOffset = 0;

function addSet(i, date, reps, weight) {

    html_code = '<tr><td style="opacity: 45%;">'+i+'.</td><td>'+date+'</td><td>'+reps+'</td><td>'+weight+'</td></tr>';

    document.getElementById('stats-table').insertAdjacentHTML('beforeend', html_code);
};


function setStatsPage(sets, hcolor, off, step) {
    let start = 0, end = 0;
    let dates = [], ws = [], reps = [], exs = []; 

    let ex = document.getElementById("ex-value").value;
    for (let i = 0; i < sets.length; i++) {
        if (sets[i].Name === ex) {
            exs.push(sets[i]);
        }
    };

    sOffset = sOffset + off;
    if (sOffset<0) {
        sOffset = 0;
    };

    let arrayLength = exs.length;
    let move = step + sOffset*step;

    if (arrayLength > move) {
        start = arrayLength - move;
        end = start + step;
    } else {
        sOffset = sOffset - 1;
        if (arrayLength > step) {
            end = step;
        } else {
            end = arrayLength;
        }
    };

    document.getElementById('stats-table').innerHTML = "";


    for (let i = start ; i < end; i++) {
        addSet(i+1, exs[i].Date, exs[i].Reps, exs[i].Weight);

        dates.push(exs[i].Date);
        reps.push(exs[i].Reps);
        ws.push(exs[i].Weight);
    };

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