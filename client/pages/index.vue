<template>
  <div>
    <Navbar />
    <div class="container mt-5">
      <div class="row py-5">
        <div class="col-sm">
          <h1 class="text-left display-1 font-weight-bold">
            Discover events for all the things you love.
          </h1>
          <button type="button" class="btn btn-primary btn-lg mt-5">
            Join Meetup
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Navbar from '../components/Navbar.vue'

export default Vue.extend({
  components: {
    Navbar,
  },
})
</script>
