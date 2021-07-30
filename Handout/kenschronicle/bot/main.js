// Notepad2 - Ken's Chronicle
const puppeteer = require('puppeteer');

const challName = "Notepad 2"

const thecookie = {
    name: 'id',
    value: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
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
        
        const browser = await puppeteer.launch();  // add `{ args: ['--no-sandbox'] }` if running as root
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

        resolve(quote);
    });
}


url = "<your url>"
url_visit(url)