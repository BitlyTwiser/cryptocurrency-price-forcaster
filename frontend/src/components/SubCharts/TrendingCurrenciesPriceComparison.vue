
<template>
    <div>
        <Chart type="bar" :data="chartData" :options="chartOptions" class="chart-size" />
    </div>
</template>

<script>
import { ref } from 'vue'
import Chart from 'primevue/chart'

export default {
    name: 'TrendingCurrenciesPriceComparison',
    components: { Chart },
    props: {
        trendingData: Array
    },
    setup() {
        const chartData = ref({
            labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
            datasets: [{
                type: 'bar',
                label: 'Dataset 1',
                backgroundColor: '#66BB6A',
                data: [21,84,24,75,37,65,34],
                borderColor: 'white',
                borderWidth: 2
            }, {
                type: 'bar',
                label: 'Dataset 2',
                backgroundColor: '#FFA726',
                data: [41,52,24,74,23,21,32]
            }]
        });

        const chartOptions = ref({
                plugins: {
                    legend: {
                        labels: {
                            color: '#495057'
                        }
                    }
                },
                scales: {
                    x: {
                        ticks: {
                            color: '#495057'
                        },
                        grid: {
                            color: '#ebedef'
                        }
                    },
                    y: {
                        ticks: {
                            color: '#495057'
                        },
                        grid: {
                            color: '#ebedef'
                        }
                    }
                }
            });

		return { chartData, chartOptions }
    }
}
</script>

<style>
  /* .chart-size{
    width: 20%; 
    height: 20%
  } */
</style>