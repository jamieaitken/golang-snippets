
function timer() {
    element = document.getElementById("timer")
    setInterval(() => {
        d = new Date();
        element.innerHTML = d.toLocaleTimeString();
    }, 1000)
}