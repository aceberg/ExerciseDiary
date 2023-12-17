var id = 0;
var today = null;

function addExercise(name, weight, reps) {
    // console.log('NAME =', name);

    id = id + 1;
    html_to_insert='<tr id="'+id+'"><td><input name="name" type="text" class="form-control" value="'+name+'"></td><td><input name="weight" type="number" class="form-control" value="'+weight+'"></td><td><input name="reps" type="number" class="form-control" value="'+reps+'"></td><td><button class="btn del-set-button" title="Delete" onclick="delExercise('+id+')"><i class="bi bi-x-square"></i></button></td></tr>';

    document.getElementById('todayEx').insertAdjacentHTML('beforeend', html_to_insert);
};

function setToday() {
    if (!today) {
        today = window.sessionStorage.getItem("today");
        console.log('TODAY FROM STOR =', today);
        if (today === null) {
            today = new Date().toJSON().slice(0, 10);
            window.sessionStorage.setItem("today", today);
        }
        document.getElementById("formDate").value = today;
    }

    console.log('TODAY =', today);
};

function setFormDateSets(sets) {
    today = document.getElementById("formDate").value;
    setToday();
    window.sessionStorage.setItem("today", today);

    console.log('DAY =', today);
    // console.log('SETS =', sets);

    document.getElementById('todayEx').innerHTML = "";

    let len = sets.length;
    for (let i = 0 ; i < len; i++) {
        if (sets[i].Date == today) {
            addExercise(sets[i].Name, sets[i].Weight, sets[i].Reps);
        }
    }
};

function delExercise(exID) {
    // console.log('DEL ID =', exID);

    document.getElementById(exID).remove();
};