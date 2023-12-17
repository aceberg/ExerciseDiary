

function addExercise(name, weight, reps) {
    // console.log('NAME =', name);

    html_to_insert='<tr><td><input name="name" type="text" class="form-control" value="'+name+'"></td><td><input name="weight" type="number" class="form-control" value="'+weight+'"></td><td><input name="reps" type="number" class="form-control" value="'+reps+'"></td></tr>';

    document.getElementById('todayEx').insertAdjacentHTML('beforeend', html_to_insert);
};

function fillTable() {
    var date = document.getElementById("realDate").value;
    // console.log('SETS =', sets);
    console.log('DATE =', date);
};

function setFormDate(sets) {
    var d = document.getElementById("realDate").value;
    document.getElementById("formDate").value = d;
    console.log('D =', d);
    // console.log('SETS =', sets);

    document.getElementById('todayEx').innerHTML = "";

    let len = sets.length;
    for (let i = 0 ; i < len; i++) {
        if (sets[i].Date == d) {
            addExercise(sets[i].Name, sets[i].Weight, sets[i].Reps);
        }
    }
};