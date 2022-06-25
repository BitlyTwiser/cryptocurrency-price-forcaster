<template>
  <div>
    <h1>This will be where we calculate future data and such</h1>
    <Button @click="getCrytoSymbols">Get Data</Button>
    <div v-if="cryptoSymbols.length > 0">
      {{cryptoSymbols}}
    </div>
    
  </div>
</template>

<script>
import axios from "axios"
import Button from 'primevue/button'

export default {
  name: 'CalculateFuturePrice',
  components: {
    Button
  },
  data() {
    return {
      cryptoSymbols: []
    }
  },
  compouted: {

  },
  methods: {
    async getCrytoSymbols(){
      await axios.get('http://127.0.0.1:3005/get-crypto-symbols')
      .then(resp => {
        this.normalizeData(resp.data.data)
      })
      .catch(error => 
        console.log(error)
      )
    },
    normalizeData(cryptoData){
      cryptoData.forEach(data => this.cryptoSymbols.push(data))
    }
  }
}
</script>