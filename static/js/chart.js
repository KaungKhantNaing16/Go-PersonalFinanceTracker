
// fetch chart json data
let url = 'http://localhost:8080/dashboard/api';

fetch(url)
.then(res => res.json())
.then(out => LineChartCreate(out))
.catch(err => { throw err });

// create Line chart 
function LineChartCreate(dataObj) {
	console.log('Checkout this JSON! ', dataObj)
	const chartData = [
		{ day: "Mon", Amount: dataObj["Monday"] },
		{ day: "Tue", Amount: dataObj["Tuesday"] },
		{ day: "Wed", Amount: dataObj["Wednesday"] },
		{ day: "Thu", Amount: dataObj["Thursday"] },
		{ day: "Fri", Amount: dataObj["Friday"] },
		{ day: "Sat", Amount: dataObj["Saturday"] },
		{ day: "Sun", Amount: dataObj["Sunday"] },
	];
	// prepare data
	const data = {
		datasets: [{
			label: 'Income',
			data: chartData,
			backgroundColor: 'rgb(75, 192, 192)',
			borderColor: 'rgb(75, 192, 192)',
			tension: 0.4,
			parsing: {
				yAxisKey: 'Amount.Incomes',
			}
		},{
			label: 'Expense',
			data: chartData,
			backgroundColor: 'rgb(255, 99, 132)',
			borderColor: 'rgb(255, 99, 132)',
			tension: 0.4,
			parsing: {
				yAxisKey: 'Amount.Expenses',
			}
		}]
	}

	// set config
	const config = {
		type: 'line',
		data,
		options: {
			parsing: {
				xAxisKey: 'day',
			},
			scales: {
				y: {
					beginAtZero: true
				}
			},
			plugins: {
				legend: {
				  position: 'bottom',
				},
			}
		}
	};

	// render init block
	const myChart = new Chart(document.getElementById('myChart'), config);
}