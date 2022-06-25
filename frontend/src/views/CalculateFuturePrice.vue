<template>
  <div>
    <h1>This will be where we calculate future data and such</h1>
    <Button @click="getCrytoSymbols">Get Data</Button>
    <div v-if="cryptoSymbols.length > 0">
      {{cryptoSymbols}}
    </div>
    <Toast />
  </div>
</template>

<script>
import axios from "axios"
import Button from 'primevue/button'
import Toast from 'primevue/toast'

export default {
  name: 'CalculateFuturePrice',
  components: {
    Button,
    Toast
  },
  data() {
    return {
      cryptoSymbols: []
    }
  },
  computed: {

  },
  methods: {
    async getCrytoSymbols(){
      await axios.get('http://127.0.0.1:3005/get-crypto-symbols')
      .then(resp => {
        this.normalizeData(resp.data.data)
        this.$toast.add({severity:'success', summary: 'Success', detail:'Successfully obtained records', life: 3000})
      })
      .catch(error => 
        this.$toast.add({severity:'error', summary: 'Failed', detail:'Failed to Obtain Cryptocurrency records', life: 3000})
      )
    },
    normalizeData(cryptoData){
      cryptoData.forEach(data => this.cryptoSymbols.push(data))
    }
  }
}
</script>