<template>
  <div>
    <!-- Navbar -->
    <section>
      <Navbar class="fixed-top" />
    </section>
    <!-- Hero Landing -->
    <section class="py-5 mt-5">
      <div class="container">
        <div class="row py-5">
          <div class="col-sm text-center">
            <h1 class="text-left display-4 font-weight-bold text-center">
              Discover events for all the things you love.
            </h1>
            <div class="mt-5">
              <button type="button" class="btn btn-primary btn-lg mt-4">
                Get Started.
              </button>
              <button type="button" class="btn btn-outline-primary btn-lg mt-4">
                Know more.
              </button>
            </div>
          </div>
        </div>
      </div>
    </section>
    <!-- Search Component -->
    <section class="py-3">
      <div class="container d-flex justify-content-center">
        <div class="row">
          <div class="col">
            <div class="card border-0 px-4 py-4 shadow mb-5">
              <Search />
            </div>
          </div>
        </div>
      </div>
    </section>
    <!-- Card Component -->
    <section class="py-5 bg-light">
      <div class="container">
        <h4 class="font-weight-bold py-5">
          Events near <span class="badge bg-primary">Zurich.</span>
        </h4>
        <Card />
      </div>
    </section>
    <!-- Footer section -->
    <section>
      <Footer />
    </section>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Navbar from '../components/Navbar.vue'
import Search from '../components/Search.vue'
import Footer from '../components/Footer.vue'
import Card from '../components/Card.vue'

export default Vue.extend({
  components: {
    Navbar,
    Search,
    Footer,
    Card,
  },
})
</script>

<style>
.shadow {
  -webkit-box-shadow: 2px 2px 2px 2px #ccc; /* Safari 3-4, iOS 4.0.2 - 4.2, Android 2.3+ */
  -moz-box-shadow: 2px 2px 2px 2px #ccc; /* Firefox 3.5 - 3.6 */
  box-shadow: 2px 2px 2px 2px #ccc; /* Opera 10.5, IE 9, Firefox 4+, Chrome 6+, iOS 5 */
}
.w {
  width: 18rem;
}
.bdg {
  position: absolute;
  padding: 5px;
  top: 0;
  margin-top: 10px;
  margin-left: 10px;
}
</style>
