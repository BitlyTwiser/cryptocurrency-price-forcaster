<template>
  <div>
    <h1>Crypto Price Estimation</h1>
    <div v-if="loading">
        <Dropdown loading/>
    </div>
    <div v-else>
        <div v-if="cryptoSymbols.length > 0">
            <Dropdown 
                v-model="selectedCryptoSymbol" 
                :options="cryptoSymbols" 
                optionLabel="name" 
                placeholder="Selected your crypto" 
                class="dropdown"
                :filter="true"
                @change="showCurrentDataForSelectedSymbol"
            />
        </div>
        <div v-else>
            No Data Loaded            
        </div>
        <div v-if="selectedCryptoSymbol">
            <p><small>OLH</small></p>
            <DataTable :value="returnDatatableValue">
                <Column v-for="col of columns" :field="col.field" :header="col.header" :key="col.field"></Column>
            </DataTable>

            <Button 
                @click="crypoSelected" 
                class="prediction-data" 
                :disabled="buttonDisabled">Request Price Prediction</Button>
        </div>

    </div>
    <Toast />
  </div>
</template>

<script>
import axios from "axios"
import Button from 'primevue/button'
import Toast from 'primevue/toast'
import Dropdown from 'primevue/dropdown'
import Column from 'primevue/column'
import DataTable from 'primevue/datatable'
import { ref } from "vue"

export default {
  name: 'PricePrediction',
  components: {
    Button,
    Toast,
    Dropdown,
    DataTable,
    Column
  },
  setup(){
    const selectedCryptoSymbol = ref(null);
    const cryptoSymbols = ref([]);
    const loading = ref(false);
    const columns = ref([]);

    return { selectedCryptoSymbol, cryptoSymbols, loading }
  },
  mounted(){
    this.getCrytoSymbols()
  },
  computed: {
    buttonDisabled(){
        return this.selectedCryptoSymbol.length < 1
    }
  },
  methods: {
    async getCrytoSymbols(){
      this.loading = true
      await axios.get('http://127.0.0.1:3005/get-crypto-symbols')
      .then(resp => {
        this.normalizeData(resp.data)
        this.createToast('success', 'Success', 'Successfully obtained records')
        this.loading = false
      })
      .catch(() => {
        this.createToast('error', 'Failed', 'Failed to Obtain Cryptocurrency records')
        this.loading = false
      })
    },
    normalizeData(cryptoData){
      cryptoData.forEach(data => this.cryptoSymbols.push(data))
    },
    createToast(severity, summary, message){
        this.$toast.add({severity: severity, summary:  summary, detail: message, life: 3000})
    },
    obtainFuturePricingEstimate(){
        if(this.selectedCryptoSymbol){
            console.log(this.selectedCryptoSymbol)
        } else {
            this.createToast('warning', 'No Data', 'A cryptosumbol was not selected, please select an element to predict.')
        }
    },
    async crypoSelected(){
        await this.showCurrentDataForSelectedSymbol()
        await this.createDataTableValues()
    },
    async showCurrentDataForSelectedSymbol(){
        const date = new Date()
        
        await axios.get(`https://api.coingecko.com/api/v3/coins/${this.selectedCryptoSymbol.id}/market_chart/range?vs_currency=usd&from=${date.getTime()}&to=${date.setFullYear(date.getFullYear() + 1)}`)
                .then((resp)=>{
                    debugger
                    console.table(resp.data.prices)
                }).catch((error) => {
                    console.debug(error)
                })
    },
    returnDatatableValue(){
        return `${this.selectedCryptoSymbol} Data`
    },
    async createDataTableValues(){
        console.log('Shalom')
    }
  }
}
</script>

<style>
    .prediction-data {
        margin-top: 10;
    }
    .dropdown{
        margin-bottom: 10px;
    }
</style>