
<template>
    <div>
        <Chart type="pie" :data="returnChartdata" :options="chartOptions" class="chart-size" />
    </div>
</template>

<script>
import { ref } from 'vue'
import Chart from 'primevue/chart'

export default {
    name: 'TrendingCurrenciesVolumePieChart',
    components: { Chart },
    props: {
        trendingData: Array
    },
    mounted() {
        this.setChartDataLabels()
        this.setChartData()
    },
    computed: {
        returnChartdata(){
            return this.chartData
        }
    },
    setup() {
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

		return { chartData, chartOptions }
    },
    methods: {
        randomInt(limit){
            return Math.floor(Math.random() * limit)
        },
        setChartDataLabels(){
            this.trendingData.forEach((c) => this.chartData.labels.push(c.item.name))
        },
        setChartData(){
            this.chartData.datasets = [
                {
                backgroundColor: [`rgba(${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)})`, `rgba(${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)})`, `rgba(${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)})`,`rgba(${this.randomInt(256)},${this.randomInt(256)}, ${this.randomInt(256)},${this.randomInt(256)})`],
                data: [40, 20, 80, 10]
                },
                {
                backgroundColor: [`rgba(${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)})`, `rgba(${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)})`, `rgba(${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)},${this.randomInt(256)})`,`rgba(${this.randomInt(256)},${this.randomInt(256)}, ${this.randomInt(256)},${this.randomInt(256)})`, `rgba(${this.randomInt(256)},${this.randomInt(256)}, ${this.randomInt(256)},${this.randomInt(256)})`],
                data: [100, 90, 2, 124, 5]
                }
            ]
        }
    }
}
</script>

<style>
  /* .chart-size{
    width: 20%; 
    height: 20%
  } */
</style>