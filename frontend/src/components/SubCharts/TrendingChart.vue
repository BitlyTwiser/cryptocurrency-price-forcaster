<template>
  <div class="content-section implementation">
      <div class="card flex justify-content-center">
          <Chart type="pie" :data="chartData" style="width: 40%" />
      </div>
  </div>
</template>

<script>
import Chart from 'primevue/chart'
import { ref } from "vue"

export default {
  name: 'TrendingChart',
  components: {
    Chart
  },
  setup(){
    const chartData = ref({
        labels: ['VueJs', 'EmberJs', 'ReactJs', 'AngularJs'],
        datasets: [
          {
            backgroundColor: ['#41B883', '#E46651', '#00D8FF', '#DD1B16'],
            data: [40, 20, 80, 10]
          }
        ]
      });

    return { chartData }
  }
}
</script>
