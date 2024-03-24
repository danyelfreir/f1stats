let chart;

const updateGraph = async (raceId, driverId) => {
    const res = await fetch(`/api/races/${raceId}/drivers/${driverId}`)
    const data = await res.json();
    console.log(data);
    chart.data.labels = data["Laps"].map(lap => lap["Lap"])
    // chart.data.datasets[0].data = data["Laps"].map(lap => lap["Milliseconds"]);
    chart.data.datasets[0].data = data["Laps"].map(lap => lap["Milliseconds"])
    chart.data.datasets[1].data = [data["Average"]];
    chart.update();
};

const ctx = document.getElementById('myChart');

chart = new Chart(ctx, {
    type: 'line',
    data: {
        // labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
        datasets: [{
            label: 'Lap times',
            data: [],
            borderWidth: 1
        }, {
            label: "Average lap time",
            data: [],
            borderWidth: 1
        }]
    },
    options: {
        // parsing: {
        //     xAxisKey: "Lap",
        //     yAxisKey: "Milliseconds",
        // },
        scales: {
            y: {
                ticks: {
                    callback: (value, index, tick) => {
                        let remainder = value;
                        const minutes = Math.floor(remainder / (1000 * 60))
                        remainder = remainder - (minutes * 1000 * 60);
                        const seconds = Math.floor(remainder / 1000);
                        remainder = remainder - seconds * 1000;
                        const timestamp = `${minutes}:${seconds}.${remainder}`;
                        return timestamp;
                    },
                },
                beginAtZero: false
            }
        }
    }
});

