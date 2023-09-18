// Calculate and set net amount value
const expAmt = document.getElementById('exp-amt').innerHTML
const earnAmt = document.getElementById('income-amt').innerHTML
const net = earnAmt - expAmt
document.getElementById('net-amt').innerHTML = net

// Calculate percentate
function calculate(num, deno) {
    let numerator = parseInt(num);
    let denominator = parseInt(deno);
    let perc = "";
    if (isNaN(numerator) || isNaN(denominator)) {
        perc = " ";
    } else {
        perc = ((numerator / denominator) * 100).toFixed(1);
    }

    console.log(perc)
    return perc
}

document.getElementById('expnum').innerHTML = calculate(expAmt, earnAmt)
document.getElementById('netnum').innerHTML = calculate(net, earnAmt)


// Draw percent progress circle
const block = document.querySelectorAll('.block');
window.addEventListener('load', function(){
  block.forEach(item => {
    let numElement = item.querySelector('.num');
    let num = parseInt(numElement.innerText);
    let count = 0;
    let time = 2000 / num;
    let circle = item.querySelector('.circle');
    setInterval(() => {
      if(count == num){
        clearInterval();
      } else {
        count += 1;
        numElement.innerText = count;
      }
    }, time)
    circle.style.strokeDashoffset 
      = 195 - ( 195 * ( num / 100 ));
    let dots = item.querySelector('.dots');
    dots.style.transform = 
      `rotate(${360 * (num / 100)}deg)`;
    if(num == 100){
      dots.style.opacity = 0;
    }
  })
});