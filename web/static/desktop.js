const menuClock = document.getElementById('menu-clock')
const menuDate = document.getElementById('menu-date')

function updateClock() {
    const now = new Date().toLocaleTimeString('en-GB');
    menuClock.textContent = now
}

updateClock()
setInterval(updateClock, 1000);
