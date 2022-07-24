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
import { onMounted, ref } from "vue"
import ProgressSpinner from 'primevue/progressspinner'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import { useToast } from "primevue/usetoast"

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
    const toast = useToast();
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
    };

    onMounted( () => {
      getCrytoSymbols()
    });

    const getCrytoSymbols = async () => {
      loading.value = true
      await axios.get('http://127.0.0.1:3005/get-crypto-symbols')
      .then(resp => {
        normalizeData(resp.data)
        createToast('success', 'Success', 'Successfully obtained records')
        loading.value = false
      })
      .catch(() => {
        createToast('error', 'Failed', 'Failed to Obtain Cryptocurrency records')
        loading.value = false
      })
    };

    const normalizeData = (cryptoData) => {
      cryptoData.forEach(data => cryptoSymbols.value.push(data))
    };

    const createToast = (severity, summary, message) => {
        toast.add({severity: severity, summary:  summary, detail: message, life: 3000})
    };

    const showDisclaimerModal = () => {
        openConfirmation()
    };

    const obtainFuturePricingEstimate = async () => {
        loadingPrediction.value = true
        closeConfirmation()
        const currentPrice = await getCurrentPriceOfSelectedCoin()
        
        if(selectedCryptoSymbol.value){
            axios.post(`http://127.0.0.1:3005/get-prediction`, {coin_id: selectedCryptoSymbol.value.id, price: currentPrice})
            .then((resp) => {
              setProbabilityTableData(resp.data)
              probablityDataLoaded.value = true
              loadingPrediction.value = false
            })
            .catch((error) => {
              loadingPrediction.value = false
              createToast('warning', 'Error', 'Error Performing calculations, please try later.')
            })
        } else {
            createToast('warning', 'No Data', 'A cryptosumbol was not selected, please select an element to predict.')
        }     
    };

    const crypoSelected = async () => {
        await showCurrentDataForSelectedSymbol()
        await getOLHCDataForSelectedCoin()
    };

    const getOLHCDataForSelectedCoin = async () => {
        await axios.get(`https://api.coingecko.com/api/v3/coins/${selectedCryptoSymbol.value.id}/ohlc?vs_currency=usd&days=30`)
                .then((resp) => {
                    resp.data.forEach((data) => series.value[0].data.push({x: new Date(data[0]), y: data.slice(1)}))  
                })
                .catch((error) => {
                    createToast('error', 'Failure', `Failure to get OLHC Data for ${selectedCryptoSymbol.value.name}`)
                })
    };

    const showCurrentDataForSelectedSymbol = async () => {
        const date = new Date()
        
        await axios.get(`https://api.coingecko.com/api/v3/coins/${selectedCryptoSymbol.value.id}/market_chart/range?vs_currency=usd&from=${date.getTime()}&to=${date.setFullYear(date.getFullYear() + 1)}`)
                .then((resp)=>{
                    console.table(resp.data.prices)
                }).catch((error) => {
                    console.debug(error)
                })
    };

    const returnDatatableValue = () => {
        return `${selectedCryptoSymbol.value.name} Data`
    };

    const getCurrentPriceOfSelectedCoin = async () => {
      let price

      await axios.get(`https://api.coingecko.com/api/v3/simple/price?ids=${selectedCryptoSymbol.value.id}&vs_currencies=usd`)
        .then((resp) => {
          price = resp.data[selectedCryptoSymbol.value.id].usd
        })
        .catch((error) => {
          createToast('warning', 'No Data', 'Error getting value price, please try again')
        })

        return price
    };

    const setProbabilityTableData = (data) => {
      probablityData.value.push({
        current_price: data.current_price,
        current_probability: data.current_probability,
        high_price_estimate: data.high_price_estimate,
        high_probability: data.high_probability,
        low_price_estimate: data.low_price_estimate,
        low_probability: data.low_probability,
        accuracy_score: data.accuracy_score
      })
    };

    return { 
      probablityDataLoaded, 
      probablityData, 
      selectedCryptoSymbol, 
      cryptoSymbols, 
      loading, 
      chartOptions, 
      series, 
      modalContent, 
      displayConfirmation, 
      openConfirmation, 
      closeConfirmation, 
      loadingPrediction, 
      showDisclaimerModal, 
      obtainFuturePricingEstimate,
      returnDatatableValue,
      crypoSelected
    }
  },
  computed: {
    buttonDisabled(){
        return this.selectedCryptoSymbol.length < 1
    },
    seriesData(){
        return this.series
    }
  },
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