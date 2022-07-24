
<template>
    <div>
        <Chart type="bar" :data="chartData" :options="chartOptions" class="chart-size" />
    </div>
</template>

<script>
import { onMounted, ref } from 'vue'
import Chart from 'primevue/chart'

export default {
    name: 'TrendingCurrenciesPriceComparison',
    components: { Chart },
    props: {
        trendingData: Array
    },
    setup(props) {
        const chartData = ref({
            labels: props.trendingData.map((i) => { return i.name}),
            datasets: [{
                type: 'bar',
                label: 'Trending Crypto Data',
                backgroundColor: '#66BB6A',
                data: props.trendingData.map((i) => {return i.price}),
                borderColor: 'white',
                borderWidth: 2
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

        const setChartLabelsAndData = () => {
            props.trendingData.forEach((c) => {

            })
        };

        onMounted(() => {
            setChartLabelsAndData()
        });

		return { chartData, chartOptions }
    },
    methods: {

    }
}
</script>

<style>
  /* .chart-size{
    width: 20%; 
    height: 20%
  } */
</style>