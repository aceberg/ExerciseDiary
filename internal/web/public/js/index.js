

function addExercise(name, weight, reps) {
    console.log('NAME =', name);

    // html_to_insert='<p class="card-text">'+name+'</p>';
    html_to_insert='<tr><td><input name="name" type="text" class="form-control" value="'+name+'"></td><td><input name="weight" type="number" class="form-control" value="'+weight+'"></td><td><input name="reps" type="number" class="form-control" value="'+reps+'"></td></tr>';

    document.getElementById('todayEx').insertAdjacentHTML('beforeend', html_to_insert);
}