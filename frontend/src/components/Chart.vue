<template>
  <section>
    <div class="trending-chart">
      <trending-chart />
    </div>
    <div class="trending-gainers-losts-chart">
      <trending-gainers-and-losers />
    </div>
  </section>
</template>

<script>
import TrendingChart from '@/components/SubCharts/TrendingChart.vue'
import TrendingGainersAndLosers from '@/components/SubCharts/TrendingGainersAndLosers.vue'

export default {
  name: "Chart",
  components: {
    TrendingChart,
    TrendingGainersAndLosers
  }
}
</script>

<style>
  .trending-chart {
    display: inline-block;
    min-width: 40%;
    max-width: 40%;
  }

  .trending-gainers-losts-chart {
    display: inline-block;
    min-width: 40%;
    max-width: 40%;
  }
</style>

