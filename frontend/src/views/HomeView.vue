<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <p><small>Photo courtesy of <a href="https://unsplash.com/">Unsplash</a></small></p>
    <landing-page />
  </div>
</template>

<script>
// @ is an alias to /src
import LandingPage from '@/components/LandingPage.vue'

export default {
  name: 'HomeView',
  components: {
    LandingPage
  }
}
</script>
