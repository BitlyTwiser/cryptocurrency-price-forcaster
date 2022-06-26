<template>
  <div>
      <div>
          <Chart type="pie" :data="chartData" class="chart-size" />
      </div>
      <Toast />
  </div>
</template>

<script>
import Chart from 'primevue/chart'
import { ref } from "vue"
import axios from "axios"
import Toast from 'primevue/toast'

export default {
  name: 'TrendingChart',
  components: {
    Chart,
    Toast
  },
  async mounted(){
    await this.getTrendingData()
    this.setChartDataLabels()
    this.setChartData()
  },
  setup(){
    const chartData = ref({
        labels: [],
        datasets: []
      });

    return { chartData }
  },
  methods: {
    randomInt(limit){
      return Math.floor(Math.random() * limit)
    },
    async getTrendingData(){
      await axios.get('http://127.0.0.1:3005/get-trending-data')
              .then((resp) => {
                debugger
                console.log('Trendong')
              })
              .catch((error) => {
                this.createToast('error', 'Failed', 'Failed to obtain trending data')
              })
    },
    createToast(severity, summary, message){
        this.$toast.add({severity: severity, summary:  summary, detail: message, life: 3000})
    },
    randomColorGenerator(){
      let colorCode = '#'
      const ints = [...Array(10).keys()]
      const alpha = Array.from(Array(26)).map((e, i) => i + 65)
      const alphanumeric = alpha.map(x => String.fromCharCode(x))
      const allValues = [ ...ints, ...alphanumeric ]
      
      while(colorCode.length < 7){
        colorCode = colorCode + allValues[this.randomInt(allValues.length)]
      }
      console.log(colorCode)
      return colorCode
    },
    setChartDataLabels(){
      this.chartData.labels = ['VueJs', 'EmberJs', 'ReactJs', 'Poop']
    },
    chartDataGenerator(){
      // Iterate through all of the trending values, for each, push the random colors and price onto the stack.
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
  },
}
</script>

<style>
  /* .chart-size{
    width: 20%; 
    height: 20%
  } */
</style>