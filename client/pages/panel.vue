<template>
  <div>
    <!-- Navbar -->
    <section>
      <Navbar class="fixed-top" />
    </section>
    <!-- Page content -->
    <h1 class="mt-5 py-5 px-5">Panel</h1>
    <!-- Footer -->
    <section>
      <Footer />
    </section>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Navbar from '../components/Navbar.vue'
import Footer from '../components/Footer.vue'

export default Vue.extend({
  components: {
    Navbar,
    Footer,
  },
})
</script>
