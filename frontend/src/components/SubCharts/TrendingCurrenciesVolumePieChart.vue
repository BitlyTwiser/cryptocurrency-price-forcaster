<template>
    <div>
        <Chart type="pie" :data="returnChartdata" :options="chartOptions" class="chart-size" />
    </div>
</template>

<script>
import { onMounted, ref } from 'vue'
import Chart from 'primevue/chart'

export default {
    name: 'TrendingCurrenciesVolumePieChart',
    components: { Chart },
    props: {
        trendingData: Array
    },
    computed: {
        returnChartdata(){
            return this.chartData
        }
    },
    setup(props) {
        const chartData = ref({
            labels: [],
            datasets: ''
        });

        const chartOptions = ref({
            plugins: {
                legend: {
                    labels: {
                        color: '#495057'
                    }
                }
            }
        });

        const randomInt = (limit) => {
            return Math.floor(Math.random() * limit)
        };

        const setChartDataLabels = () =>  {
            props.trendingData.forEach((c) => chartData.value.labels.push(c.name))
        }
        
        const setBackgroundColors = () => {
            let backgroundColorArray = []
            props.trendingData.forEach(() => {backgroundColorArray.push(`rgba(${randomInt(256)},${randomInt(256)},${randomInt(256)},${randomInt(256)})`)})

            return backgroundColorArray
        };

        const trendingDataPrices = () => {
            return props.trendingData.map((i) => { return i.quantity })
        };

        const setChartData = () =>  {
            chartData.value.datasets = [
                {
                backgroundColor: setBackgroundColors(),
                data: trendingDataPrices()
                }
            ]
        };

        onMounted(() => {
            setChartDataLabels()
            setChartData()
        })

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