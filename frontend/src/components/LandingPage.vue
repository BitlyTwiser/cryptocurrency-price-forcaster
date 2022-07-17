<template>
  <div class="hello">
    <h1>Welcome to the Capstone Cryptocurrency project</h1>
    <p>You can use this app to visualize current/previous cryprocurrency data for trending currencies in the market.</p>
    <p>You can also utlize the machine learning elements of this project to attempt to ascertain what the price of selected coins might be in the future!</p>
    <p>The prediction mechanism utilizes an adaptation of Guassian Naive Bayes algorithm to attempt to curate a prediction of the future cost of the specified token.</p>
    <p>The Guassian Naive Bayes algorithm is trained off of the last 30 days of OHLC data that is cleansed and ingested to formulate probabilities of price values and what those price values may be.</p>
    <p>Use the navBar above to visit a selected page to get started!</p>
  </div>
</template>

<script>

export default {
  name: 'LandingPage',
  props: {
    msg: String
  }
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}

button.p-button.p-component {
 align-items: center; 
 background-color: #2196f3; 
 color: #ffffff; 
 display: inline-flex; 
 font-family: -apple-system; 
 text-align: left; 
 margin-left: 10px;
}

</style>
