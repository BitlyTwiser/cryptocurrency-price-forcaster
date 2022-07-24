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
import { onMounted, ref } from "vue"
import axios from "axios"
import Toast from 'primevue/toast'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import ProgressSpinner from 'primevue/progressspinner'
import { useToast } from "primevue/usetoast"

export default {
  name: 'TrendingChart',
  components: {
    Toast,
    DataTable,
    Column,
    ProgressSpinner
  },
  setup(props, { emit }){
      const loading = ref(false);
      const trendingData = ref([]);
      const trendingSymbolsData = ref([]);
      const toast = useToast();

      onMounted( async () => {
        loading.value = true
        await getTrendingData()
        loading.value = false
      });

      const getTrendingData = async () => {
        loading.value = true
        await axios.get('http://127.0.0.1:3005/get-trending-data')
                .then(async (resp) => {
                  const coinData = resp.data.coins

                  await setTableData(coinData)
                })
                .catch((error) => {
                  createToast('error', 'Failed', 'Failed to obtain trending data')
                  loading.value = false
                })
        createToast('success', 'Data Obtained', 'Retrevied trending data')
    };

    const createToast = (severity, summary, message) => {
        toast.add({severity: severity, summary:  summary, detail: message, life: 3000})
    };

    const setTableData = async (data) => {
      let counter = 0
      
      await data.forEach(async (val) => {
        await axios.get(`http://127.0.0.1:3005/get-coin-data?coin_id=${val.item.id}`).then((resp => {
          normalizeChartDataAndSetTableData(resp.data)
          counter++
        })).catch(() => {
          createToast('error', 'Failed', 'Failed to obtain coin data')
        })

        if(counter === 7){
          emit('success', trendingData.value)
        }
      })      
    };

    const normalizeChartDataAndSetTableData = (data) => {
      trendingData.value.push({
        position: trendingData.value.length,
        id: data.id,
        name: data.name,
        price: data.market_data.current_price.usd,
        quantity: data.market_data.total_supply,
        dailyhigh: data.market_data.high_24h.usd,
        dailylow: data.market_data.low_24h.usd,
      })
    };

    return { trendingData, trendingSymbolsData, loading }
  },
}
</script>