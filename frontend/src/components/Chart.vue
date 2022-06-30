<template>
  <section>
    <div>
      <h3>Trending Cryto Currencies</h3>
      <trending-chart @success="showAdditionalChartsOnSuccess" />
    </div>
    <div v-if="showAdditionalCharts && trendingData.length > 0">
      <h3>Charted Visualizations for Trending Currencies</h3>
      <div>
        <div class="trending-currencies-bar-chart" >
          <p><small>Price Bar Chart</small></p>
          <trending-currencies-price-comparison :trendingData="trendingData"/>
        </div>
        <div class="trending-chart-pie-chart">
          <p><small>Volume Chart</small></p>
          <trending-currencies-volume-pie-chart :trendingData="trendingData"/>
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
    const showAdditionalCharts = ref(false);
    const trendingData = ref();

    return { showAdditionalCharts, trendingData }
  },
  methods: {
    showAdditionalChartsOnSuccess(data){
      this.showAdditionalCharts = true
      this.trendingData = data
    },
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

