
function setToday() {
    today = new Date().toJSON().slice(0, 10);
    document.getElementById("todayDate").value = today;
};


function addWeight(date, weight, id) {

    html_code = '<tr><td>'+date+'</td><td>'+weight+'</td><td><a href="/weight/?del='+id+'"><button class="btn del-set-button" title="Delete" ><i class="bi bi-x-square"></i></button></a></td></tr>';

    document.getElementById('weightList').insertAdjacentHTML('beforeend', html_code);
};

function setWeights(weights, offset) {
    var start = 0;
    let arrayLength = weights.length;

    if (arrayLength > offset) {
        start = arrayLength - offset;
    };

    for (let i = start ; i < arrayLength; i++) {
        addWeight(weights[i].Date, weights[i].Weight, weights[i].ID)
    };
};