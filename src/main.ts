import { updateVisitorCount } from './counter.ts'
import kojiPeko from '/Kojima+Pekora.png'

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <div>
    <h1>Anthony Nguyen</h1>
    <img src="${kojiPeko}" class="meme"/>
    
    <p class="read-the-docs">
      Kojima has good tastes in oshi's.
    </p>

    <div class="card">
      <div id="counter"/>
    </div>
  </div>
`

const counterElement = document.querySelector<HTMLDivElement>('#counter');

if (counterElement) {
  updateVisitorCount(counterElement);
}
