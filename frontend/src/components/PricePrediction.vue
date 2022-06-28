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
                @change="crypoSelected"
            />
        </div>
        <div v-else>
            No Data Loaded            
        </div>
        <div v-if="selectedCryptoSymbol && series[0].data.length > 0">
            <p><small>OLHC data for {{selectedCryptoSymbol.name}}</small></p>
            <div class="chart">
                <apexchart type="candlestick" height="350" :options="chartOptions" :series="seriesData"></apexchart>
            </div>
        
            <Button 
                @click="obtainFuturePricingEstimate" 
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
import { ref } from "vue"

export default {
  name: 'PricePrediction',
  components: {
    Button,
    Toast,
    Dropdown,
  },
  setup(){
    const selectedCryptoSymbol = ref(null);
    const cryptoSymbols = ref([]);
    const loading = ref(false);
    const chartOptions = ref({
            chart: {
              type: 'candlestick',
              height: 350
            },
            title: {
              text: 'CandleStick Chart',
              align: 'left'
            },
            xaxis: {
              type: 'datetime'
            },
            yaxis: {
              tooltip: {
                enabled: true
              }
            }
          });
    const series =  ref([{data: []}]);

    return { selectedCryptoSymbol, cryptoSymbols, loading, chartOptions, series }
  },
  mounted(){
    this.getCrytoSymbols()
  },
  computed: {
    buttonDisabled(){
        return this.selectedCryptoSymbol.length < 1
    },
    seriesData(){
        debugger
        return this.series
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
        await this.getOLHCDataForSelectedCoin()
    },
    async getOLHCDataForSelectedCoin(){
        await axios.get(`https://api.coingecko.com/api/v3/coins/${this.selectedCryptoSymbol.id}/ohlc?vs_currency=usd&days=30`)
                .then((resp) => {
                    debugger
                    resp.data.forEach((data) => this.series[0].data.push({x: new Date(data[0]), y: data.slice(1)}))  
                })
                .catch((error) => {
                    this.createToast('error', 'Failure', `Failure to get OLHC Data for ${this.selectedCryptoSymbol.name}`)
                })
    },
    async showCurrentDataForSelectedSymbol(){
        const date = new Date()
        
        await axios.get(`https://api.coingecko.com/api/v3/coins/${this.selectedCryptoSymbol.id}/market_chart/range?vs_currency=usd&from=${date.getTime()}&to=${date.setFullYear(date.getFullYear() + 1)}`)
                .then((resp)=>{
                    console.table(resp.data.prices)
                }).catch((error) => {
                    console.debug(error)
                })
    },
    returnDatatableValue(){
        return `${this.selectedCryptoSymbol.name} Data`
    },
    async createDataTableValues(){
        console.log("WTF")
        this.chartOptions.title.text = this.returnDatatableValue()
    },
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