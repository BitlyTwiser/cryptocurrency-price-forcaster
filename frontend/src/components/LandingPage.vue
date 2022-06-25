<template>
  <div class="hello">
    <h1>Welcome to the Capstone Cryptocurrency project</h1>
    <p>You can use this app to visualize current/previous cryprocurrency data for trending currencies in the market.</p>
    <p>You can also utlize the machine learning elements of this project to attempt to ascertain what the price of selected coins might be in the future!</p>
    <p>The prediction mechanism utilizes an adaptation of Guassian Naive Bayes algorithm.</p>
    <p>Use the navBar above to visit a selected page</p>
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
