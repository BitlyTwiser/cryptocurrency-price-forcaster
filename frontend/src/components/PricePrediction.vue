<template>
  <div>
    <h1>price Prediction</h1>
    <div v-if="loading">
        <Dropdown loading/>
    </div>
    <div v-else>
        <div v-if="cryptoSymbols.length > 0">
            <Dropdown v-model="selectedCryptoSymbol" :options="cryptoSymbols" optionLabel="name" placeholder="Selected your crypto" />
        </div>
        <div v-else>
            No Data Loaded            
        </div>
        <Button @click="getCrytoSymbols">Request Price Prediction</Button>
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
    Dropdown
  },
  setup(){
    const selectedCryptoSymbol = ref();
    const cryptoSymbols = ref([]);
    const loading = ref(false);

    return { selectedCryptoSymbol, cryptoSymbols, loading }
  },
  mounted(){
    this.getCrytoSymbols()
  },
  computed: {

  },
  methods: {
    async getCrytoSymbols(){
      this.loading = true
      await axios.get('http://127.0.0.1:3005/get-crypto-symbols')
      .then(resp => {
        this.normalizeData(resp.data.data)
        this.$toast.add({severity:'success', summary: 'Success', detail:'Successfully obtained records', life: 3000})
        this.loading = false
      })
      .catch(error => {
        this.$toast.add({severity:'error', summary: 'Failed', detail:'Failed to Obtain Cryptocurrency records', life: 3000})
        this.loading = false
      })
    },
    normalizeData(cryptoData){
      cryptoData.forEach(data => this.cryptoSymbols.push(data))
    }
  }
}
</script>