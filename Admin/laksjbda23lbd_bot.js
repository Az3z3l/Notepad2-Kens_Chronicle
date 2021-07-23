// Notepad 2 : ken's Chronicles
const puppeteer = require('puppeteer');
var queue = require('../redis-controller/queue');
const queueName = __filename.split(".")[0].split("/").pop();

const challName = "Notepad 2"

const thecookie = {
    name: 'id',
    value: '7JtYAMsyQYAg2ReT5PyoJeH9rVRikBx6Cp',
    domain: 'chall.notepad2.gq',
    expires: -1,
    httpOnly: true,
    secure: true,
    session: true,
    sameSite: 'Lax',
  }



async function url_visit (url) {
    var quote;
    return new Promise(async function(resolve, reject) {
        // start modification
        
        const browser = await puppeteer.launch({executablePath: '/home/bot/latest/chrome'});  // add `{ args: ['--no-sandbox'] }` if running as root
        const page = await browser.newPage();         
        await page.setCookie(thecookie)
        await page.setDefaultNavigationTimeout(1e3*15);  // Timeout duration in milliseconds    // use either this or wait for navigation
        try{
            var result = await page.goto(url);
            await page.waitForNavigation(); // wait till the page finishes loading              
        }
        catch(e){
            console.log("timeout exceeded");
        }        
        await browser.close();

        // end modification
        resolve(quote);
    });
}


function popMe(){
    queue.pop(queueName,sendUrl)
}

async function sendUrl(err, url) {
    if (err) {
        throw err;
    }

    if (!url) {
        setTimeout(popMe, 1e3*2); // if null is returned, wait for 2 secs before retrying
    } else {
        console.log(url)
        q = await url_visit(url)
        popMe();
    } 
}

console.log(`Started bot for chall ${challName} with id ${queueName}`)
popMe()
setTimeout(popMe, 1e3*2);
setTimeout(popMe, 1e3*3);