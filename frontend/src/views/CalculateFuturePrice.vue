<template>
  <div>
    <h1>This will be where we calculate future data and such</h1>
    <Button @click="getCrytoSymbols">Get Data</Button>
    <div v-if="showCrypoData">
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
      cryptoSymbols: ['asdsd']
    }
  },
  compouted: {

  },
  methods: {
    async getCrytoSymbols(){
      await axios.get('http://127.0.0.1:3005/get-crypto-symbols')
      .then(resp => this.cryptoSymbols = resp)
      .catch(error => console.log(error))
    },
    showCrypoData(){
      console.log('here')
      return this.cryptoSymbols.length > 0
    },
  }
}
</script>