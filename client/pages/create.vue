<template>
  <div>
    <!-- Navbar -->
    <section>
      <Navbar class="fixed-top" />
    </section>
    <!-- Page content -->
    <section class="mt-5 py-5 bg-light">
      <div class="container">
        <input
          type="text"
          class="form-control"
          placeholder="Enter event title"
        />
        <input type="date" class="form-control mt-3" />
        <input type="time" class="form-control mt-3" />
        <input
          type="number"
          class="form-control mt-3"
          placeholder="Enter the number of seats"
        />
        <input type="text" class="form-control mt-3" placeholder="Host name" />
        <input type="url" class="form-control mt-3" placeholder="Host link" />
        <input type="text" class="form-control mt-3" placeholder="Event link" />
        <input
          type="url"
          class="form-control mt-3 mb-3"
          placeholder="Cover image link"
        />
        <select class="form-select">
          <option value="1">Online event</option>
          <option value="2">Physical location</option>
        </select>
        <button class="btn btn-primary mt-5">Create event</button>
      </div>
    </section>
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
