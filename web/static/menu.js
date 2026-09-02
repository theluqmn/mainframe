const menuClock = document.getElementById('menu-clock')
const menuDate = document.getElementById('menu-date')

function updateClock() {
    const now = new Date().toLocaleTimeString('en-GB');
    menuClock.textContent = now
}

function updateDate() {
    const now = new Date().toLocaleDateString('en-GB')
    menuDate.textContent = now
}

updateClock(); setInterval(updateClock, 1000);
updateDate(); setInterval(updateDate, 1000);