<template>
  <section>
    <div>
      <h3>Trending Cryto Currencies</h3>
      <trending-chart @success="showCandleStickCharts" />
    </div>
    <div v-if="showTrendingCandleStick">
      <h2>Candle Stick data for Trending Currencies</h2>
      <div>
        <trending-gainers-and-losers-vue />
      </div>
    </div>
  </section>
</template>

<script>
import TrendingChart from '@/components/SubCharts/TrendingChart.vue'
import { ref } from "vue"
import TrendingGainersAndLosersVue from './SubCharts/TrendingGainersAndLosers.vue'

export default {
  name: "Chart",
  components: {
    TrendingChart,
    TrendingGainersAndLosersVue
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

