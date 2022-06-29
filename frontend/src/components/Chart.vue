<template>
  <section>
    <div class="trending-chart">
      <h3>Trending Cryto Currencies</h3>
      <p><small>Top 7 trending coins</small></p>
      <trending-chart />
    </div>
  </section>
</template>

<script>
import TrendingChart from '@/components/SubCharts/TrendingChart.vue'

export default {
  name: "Chart",
  components: {
    TrendingChart
  }
}
</script>

<style>
  .trending-chart {
    display: inline-block;
    min-width: 30%;
    max-width: 30%;
  }

  .trending-gainers-losts-chart {
    display: inline-block;
    min-width: 35%;
    max-width: 35%;
  }
</style>

