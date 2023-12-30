
function setToday() {
    today = new Date().toJSON().slice(0, 10);
    document.getElementById("todayDate").value = today;
};