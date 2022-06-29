<template>
  <section>
    <div>
      <h3>Trending Cryto Currencies</h3>
      <trending-chart @success="showCandleStickCharts" />
    </div>
    <div v-if="showTrendingCandleStick">
      <h3>Charted Visualizations for Trending Currencies</h3>
      <div>
        <div class="trending-currencies-bar-chart" >
          <trending-currencies-price-comparison />
        </div>
        <div class="trending-chart-pie-chart">
          <trending-currencies-volume-pie-chart />
        </div>
        
      </div>
    </div>
  </section>
</template>

<script>
import TrendingChart from '@/components/SubCharts/TrendingChart.vue'
import { ref } from "vue"
import TrendingCurrenciesPriceComparison from '@/components/SubCharts/TrendingCurrenciesPriceComparison.vue'
import TrendingCurrenciesVolumePieChart from '@/components/SubCharts/TrendingCurrenciesVolumePieChart.vue'

export default {
  name: "Chart",
  components: {
    TrendingChart,
    TrendingCurrenciesPriceComparison,
    TrendingCurrenciesVolumePieChart
  },
  setup(){
    const showTrendingCandleStick = ref(false);

    return { showTrendingCandleStick }
  },
  methods: {
    showCandleStickCharts(){
      this.showTrendingCandleStick = true
    }
  }
}
</script>

<style>
  .trending-chart-pie-chart {
    display: inline-block;
    min-width: 25%;
    max-width: 25%;
  }

  .trending-currencies-bar-chart {
    display: inline-block;
    min-width: 35%;
    max-width: 35%;
  }
</style>

