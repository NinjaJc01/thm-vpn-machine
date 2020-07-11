console.log("Hello, World!")
async function getData(url = '') {
    // Default options are marked with *
    const response = await fetch(url, {
        method: 'GET', // *GET, POST, PUT, DELETE, etc.
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *client
    });
    return await response.text();
}

async function onPageLoad() {
    const myIp = await getData("/whoami");
    console.log(myIp);
    document.getElementById("myIp").textContent = myIp;
}