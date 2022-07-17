<template>
  <div>
    <h1>Crypto Price Dashboard</h1>
    <p>Select a currency to continue</p>
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
        <div>
            <Dialog header="Disclaimer" v-model:visible="displayConfirmation" :breakpoints="{'960px': '75vw', '640px': '90vw'}" :style="{width: '350px'}" :modal="true" >
	            <div class="confirmation-content">
                    <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                <span>{{modalContent}}</span>
            </div>
            <template #footer>
                <Button label="Decline" icon="pi pi-times" @click="closeConfirmation" class="p-button-text" autofocus />
                <Button label="Accept" icon="pi pi-check" @click="obtainFuturePricingEstimate" class="p-button-text" autofocus />
            </template>
            </Dialog>
        </div>
        <div v-if="selectedCryptoSymbol && series[0].data.length > 0">
            <p><small>OLHC data for {{selectedCryptoSymbol.name}} (30 Days)</small></p>
            <div class="chart">
                <apexchart type="candlestick" height="350" :options="chartOptions" :series="seriesData"></apexchart>
            </div>

            <div v-if="probablityData.length > 0 && probablityDataLoaded" class="probability-table">
              <p><small>Note: Accuracy score is derrived from the given dataset and accuracy score from Sklearn ML accuracy_analysis algorithm. A score of 1 is effective > 90% accuracy</small></p>
              <DataTable :value="probablityData" responsiveLayout="scroll">
                <template #header>
                    Probability Table
                </template>
                <Column field="current_price" header="Current Price"></Column>
                <Column field="current_probability" header="Probability for Current Price"></Column>
                <Column field="low_price_estimate" header="Low Price Estimation"></Column>
                <Column field="low_probability" header="Probability of Low Price Occurance"></Column>
                <Column field="high_price_estimate" header="High Price Estimation"></Column>
                <Column field="high_probability" header="Probability of High Price Occurance"></Column>
                <Column field="accuracy_score" header="Accuracy Score (Floating Point)"></Column>
            </DataTable>
            </div>
            <div v-if="loadingPrediction">
              <p>Performing calculations and training datasets for probability analysis and price estimation</p>
              <ProgressSpinner style="width:50px;height:50px" strokeWidth="8" fill="var(--surface-ground)" animationDuration=".5s"/>
            </div>
            <div v-else>
            <Button 
                @click="showDisclaimerModal" 
                class="prediction-data" 
                :disabled="buttonDisabled">Request Future Price Prediction</Button>
            </div>
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
import Dialog from 'primevue/dialog'
import { ref } from "vue"
import ProgressSpinner from 'primevue/progressspinner'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

export default {
  name: 'PricePrediction',
  components: {
    Button,
    Toast,
    Dropdown,
    Dialog,
    ProgressSpinner,
    DataTable,
    Column
  },
  setup(){
    const selectedCryptoSymbol = ref(null);
    const cryptoSymbols = ref([]);
    const loading = ref(false);
    const loadingPrediction = ref(false);
    const displayConfirmation = ref(false);
    const probablityDataLoaded = ref(false);
    const probablityData = ref([]);
    const modalContent = ref("This program is not intended to provide finanical advice of any kind. This application is intended to utilize algoriothms to provide the best plausible guess at future prices of cryptocurrency tokens. USE AT YOUR OWN RISK. Accept to continue")
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
    const openConfirmation = () => {
            displayConfirmation.value = true;
        };
    const closeConfirmation = () => {
        displayConfirmation.value = false;
    }

    return { probablityDataLoaded, probablityData, selectedCryptoSymbol, cryptoSymbols, loading, chartOptions, series, modalContent, displayConfirmation, openConfirmation, closeConfirmation, loadingPrediction }
  },
  mounted(){
    this.getCrytoSymbols()
  },
  computed: {
    buttonDisabled(){
        return this.selectedCryptoSymbol.length < 1
    },
    seriesData(){
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
    showDisclaimerModal(){
        this.openConfirmation()
    },
    async obtainFuturePricingEstimate(){
        this.loadingPrediction = true
        this.closeConfirmation()
        const currentPrice = await this.getCurrentPriceOfSelectedCoin()
        
        if(this.selectedCryptoSymbol){
            axios.post(`http://127.0.0.1:3005/get-prediction`, {coin_id: this.selectedCryptoSymbol.id, price: currentPrice})
            .then((resp) => {
              this.setProbabilityTableData(resp.data)
              this.probablityDataLoaded = true
              this.loadingPrediction = false
            })
            .catch((error) => {
              console.log(error)
              this.loadingPrediction = false
              this.createToast('warning', 'Error', 'Error Performing calculations, please try later.')
            })
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
    async getCurrentPriceOfSelectedCoin(){
      let price

      await axios.get(`https://api.coingecko.com/api/v3/simple/price?ids=${this.selectedCryptoSymbol.id}&vs_currencies=usd`)
        .then((resp) => {
          price = resp.data[this.selectedCryptoSymbol.id].usd
        })
        .catch((error) => {
          this.createToast('warning', 'No Data', 'Error getting value price, please try again')
        })

        return price
    },
    setProbabilityTableData(data){
      this.probablityData.push({
        current_price: data.current_price,
        current_probability: data.current_probability,
        high_price_estimate: data.high_price_estimate,
        high_probability: data.high_probability,
        low_price_estimate: data.low_price_estimate,
        low_probability: data.low_probability,
        accuracy_score: data.accuracy_score
      })
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
    .probability-table{
      margin-bottom: 10px;
    }
</style>