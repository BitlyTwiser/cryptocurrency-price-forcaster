<template>
  <div class="hello">
    <h1>Welcome to the Capstone Cryptocurrency project</h1>
    <p>You can use this app to visualize current/previous cryprocurrency data for selected assets</p>
    <p>You can also utlize the machine learning elements of this project to attempt to ascertain what the price of selected coins might be in the future.</p>
    <p>Use the navBar above to visit a selecte page or click any of the buttons below to get started</p>
    <Button>button1</Button>
    <Button>button2</Button>
    <Button>button3</Button>
  </div>
</template>

<script>
import Button from 'primevue/button'

export default {
  name: 'LandingPage',
  components: { Button },
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
</style>
