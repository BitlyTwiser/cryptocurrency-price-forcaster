<template>
  <div>
      <div>
        <div v-if="loading">
          <ProgressSpinner style="width:50px;height:50px" strokeWidth="8" fill="var(--surface-ground)" animationDuration=".5s"/>
        </div>
          <div class="card" v-else>
            <DataTable :value="trendingData" responsiveLayout="scroll">
                <template #header>
                    Scroll
                </template>
                <Column field="position" header="Position"></Column>
                <Column field="id" header="ID"></Column>
                <Column field="name" header="Name"></Column>
                <Column field="price" header="Price"></Column>
                <Column field="quantity" header="Total Volum"></Column>
                <Column field="dailyhigh" header="24 Hour High"></Column>
                <Column field="dailylow" header="24 Hour Low"></Column>
                <Column field="chart" header="24 hour candle stick chart">
                  <Chart type="bar" :data="chartData" class="chart-size" />
                </Column>
            </DataTable>
        </div>
      </div>
      <Toast />
  </div>
</template>

<script>
import Chart from 'primevue/chart'
import { ref } from "vue"
import axios from "axios"
import Toast from 'primevue/toast'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import ProgressSpinner from 'primevue/progressspinner'

export default {
  name: 'TrendingChart',
  components: {
    Chart,
    Toast,
    DataTable,
    Column,
    ProgressSpinner
  },
  async mounted(){
    await this.getTrendingData()
    this.setChartData()
  },
  setup(){
    const chartData = ref({
        labels: [],
        datasets: []
      });
      const loading = ref(false);
      const trendingData = ref([]);
      const trendingSymbolsData = ref([]);

    return { chartData, trendingData, trendingSymbolsData, loading }
  },
  methods: {
    randomInt(limit){
      return Math.floor(Math.random() * limit)
    },
    async getTrendingData(){
      this.loading = true
      await axios.get('http://127.0.0.1:3005/get-trending-data')
              .then(async (resp) => {
                const coinData = resp.data.coins

                await this.setTableData(coinData)
                this.setChartDataLabels(coinData) //This will be removed
              })
              .catch((error) => {
                this.createToast('error', 'Failed', 'Failed to obtain trending data')
                this.loading = false
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
    async setTableData(data){
      data.forEach((val) => {
        axios.get(`http://127.0.0.1:3005/get-coin-data?coin_id=${val.item.id}`).then((resp => {
          this.normalizeChartDataAndSetTableData(resp.data)
        })).catch(() => {
          this.createToast('error', 'Failed', 'Failed to obtain coin data')
        })
      })
      this.createToast('success', 'Data Obtained', 'Retrevied trending data')
      this.loading = false
    },
    setChartDataLabels(data){
      data.forEach((c) => this.chartData.labels.push(c.item.name))
    },
    normalizeChartDataAndSetTableData(data){
      this.trendingData.push({
        position: this.trendingData.length,
        id: data.id,
        name: data.name,
        price: data.market_data.current_price.usd,
        quantity: data.market_data.total_supply,
        dailyhigh: data.market_data.high_24h.usd,
        dailylow: data.market_data.low_24h.usd,
      })
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

<style lang="scss" scoped>
img {
    vertical-align: middle;
}
::v-deep(.p-paginator) {
    .p-paginator-current {
        margin-left: auto;
    }
}

::v-deep(.p-progressbar) {
    height: .5rem;
    background-color: #D8DADC;

    .p-progressbar-value {
        background-color: #607D8B;
    }
}

::v-deep(.p-datepicker) {
    min-width: 25rem;

    td {
        font-weight: 400;
    }
}

::v-deep(.p-datatable.p-datatable-customers) {
    .p-datatable-header {
        padding: 1rem;
        text-align: left;
        font-size: 1.5rem;
    }

    .p-paginator {
        padding: 1rem;
    }

    .p-datatable-thead > tr > th {
        text-align: left;
    }

    .p-datatable-tbody > tr > td {
        cursor: auto;
    }

    .p-dropdown-label:not(.p-placeholder) {
        text-transform: uppercase;
    }
}
</style>