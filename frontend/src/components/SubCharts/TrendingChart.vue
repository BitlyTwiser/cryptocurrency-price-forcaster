<template>
  <div>
      <div>
        <div v-if="loading">
          <ProgressSpinner style="width:50px;height:50px" strokeWidth="8" fill="var(--surface-ground)" animationDuration=".5s"/>
        </div>
          <div class="card" v-else>
            <DataTable :value="trendingData" responsiveLayout="scroll">
                <template #header>
                    Top 7 trending coins
                </template>
                <Column field="position" header="Position"></Column>
                <Column field="id" header="ID"></Column>
                <Column field="name" header="Name"></Column>
                <Column field="price" header="Price"></Column>
                <Column field="quantity" header="Total Volume"></Column>
                <Column field="dailyhigh" header="24 Hour High"></Column>
                <Column field="dailylow" header="24 Hour Low"></Column>
            </DataTable>
        </div>
      </div>
      <Toast />
  </div>
</template>

<script>
import { ref } from "vue"
import axios from "axios"
import Toast from 'primevue/toast'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import ProgressSpinner from 'primevue/progressspinner'

export default {
  name: 'TrendingChart',
  components: {
    Toast,
    DataTable,
    Column,
    ProgressSpinner
  },
  async mounted(){
    this.loading = true
    await this.getTrendingData()
    this.loading = false
  },
  setup(){
      const loading = ref(false);
      const trendingData = ref([]);
      const trendingSymbolsData = ref([]);

    return { trendingData, trendingSymbolsData, loading }
  },
  methods: {
    async getTrendingData(){
      this.loading = true
      await axios.get('http://127.0.0.1:3005/get-trending-data')
              .then(async (resp) => {
                const coinData = resp.data.coins

                await this.setTableData(coinData)
              })
              .catch((error) => {
                this.createToast('error', 'Failed', 'Failed to obtain trending data')
                this.loading = false
              })
      this.createToast('success', 'Data Obtained', 'Retrevied trending data')
    },
    createToast(severity, summary, message){
        this.$toast.add({severity: severity, summary:  summary, detail: message, life: 3000})
    },
    async setTableData(data){
      let counter = 0
      
      await data.forEach(async (val) => {
        await axios.get(`http://127.0.0.1:3005/get-coin-data?coin_id=${val.item.id}`).then((resp => {
          this.normalizeChartDataAndSetTableData(resp.data)
          counter++
        })).catch(() => {
          this.createToast('error', 'Failed', 'Failed to obtain coin data')
        })

        if(counter === 7){
          this.$emit('success', this.trendingData)
        }
      })      
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
  },
}
</script>