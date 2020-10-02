<template>
  <div>
    <!-- Navbar -->
    <section>
      <Navbar />
    </section>
    <!-- Hero Landing -->
    <section class="mt-5">
      <div class="container">
        <div class="row py-5">
          <div class="col-sm">
            <h1 class="text-left display-1 font-weight-bold">
              Descubra eventos para todas as coisas que você ama.
            </h1>
            <button type="button" class="btn btn-primary btn-lg mt-4">
              Começar.
            </button>
          </div>
        </div>
      </div>
    </section>
    <!-- Search Component -->
    <section class="mt-4">
      <div class="container d-flex justify-content-center">
        <div class="row">
          <div class="col">
            <div class="card px-4 py-4 shadow">
              <Search />
            </div>
          </div>
        </div>
      </div>
    </section>
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

<style scoped>
.shadow {
  -webkit-box-shadow: 3px 3px 5px 6px #ccc; /* Safari 3-4, iOS 4.0.2 - 4.2, Android 2.3+ */
  -moz-box-shadow: 3px 3px 5px 6px #ccc; /* Firefox 3.5 - 3.6 */
  box-shadow: 3px 3px 5px 6px #ccc; /* Opera 10.5, IE 9, Firefox 4+, Chrome 6+, iOS 5 */
}
</style>
