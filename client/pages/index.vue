<template>
  <div>
    <!-- Navbar -->
    <Navbar />
    <!-- Hero Landing -->
    <div class="container mt-5">
      <div class="row py-5">
        <div class="col-sm">
          <h1 class="text-left display-1 font-weight-bold">
            Descubra eventos para todas as coisas que você ama.
          </h1>
          <button type="button" class="btn btn-primary btn-lg mt-5">
            Começar.
          </button>
        </div>
      </div>
    </div>
    <!-- Search Component -->
    <div class="container d-flex justify-content-center mt-4">
      <div class="row">
        <div class="col">
          <div class="card px-4 py-4">
            <Search />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Navbar from '../components/Navbar.vue'
import Search from '../components/Search.vue'

export default Vue.extend({
  components: {
    Navbar,
    Search,
  },
})
</script>
