var offset = 0;

function setToday() {
    today = new Date().toJSON().slice(0, 10);
    document.getElementById("todayDate").value = today;
};


function addWeight(i, date, weight, id) {

    html_code = '<tr><td style="opacity: 45%;">'+i+'.</td><td>'+date+'</td><td>'+weight+'</td><td><a href="/weight/?del='+id+'"><button class="btn del-set-button" title="Delete" ><i class="bi bi-x-square"></i></button></a></td></tr>';

    document.getElementById('weightList').insertAdjacentHTML('beforeend', html_code);
};

function setWeights(weights, wcolor, off, step) {
    let start = 0, end = 0;
    let dates = [], ws = [];

    offset = offset + off;
    if (offset<0) {
        offset = 0;
    };

    let arrayLength = weights.length;
    let move = step + offset*step;

    if (arrayLength > move) {
        start = arrayLength - move;
        end = start + step;
    } else {
        offset = offset - 1;
        if (arrayLength > step) {
            end = step;
        } else {
            end = arrayLength;
        }
    };

    // console.log("OFF =", offset, ", START =", start, ", END =", end)

    document.getElementById('weightList').innerHTML = "";

    for (let i = start ; i < end; i++) {
        dates.push(weights[i].Date);
        ws.push(weights[i].Weight);
        addWeight(i+1, weights[i].Date, weights[i].Weight, weights[i].ID)
    };
    weightChart('weight-chart', dates, ws, wcolor, true);
};