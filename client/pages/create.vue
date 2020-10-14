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
          v-model="title"
          type="text"
          class="form-control"
          placeholder="Enter event title"
        />
        <input
          v-model="city"
          type="text"
          class="form-control mt-3"
          placeholder="Enter event city"
        />
        <input v-model="date" type="date" class="form-control mt-3" />
        <input v-model="time" type="time" class="form-control mt-3" />
        <input
          v-model="max"
          type="number"
          class="form-control mt-3"
          placeholder="Enter the number of seats"
        />
        <input
          v-model="hostName"
          type="text"
          class="form-control mt-3"
          placeholder="Host name"
        />
        <input
          v-model="hostURL"
          type="url"
          class="form-control mt-3"
          placeholder="Host link"
        />
        <input
          v-model="link"
          type="text"
          class="form-control mt-3"
          placeholder="Event link"
        />
        <input
          v-model="img"
          type="url"
          class="form-control mt-3 mb-3"
          placeholder="Cover image link"
        />
        <select v-model="type" class="form-select">
          <option value="Online event">Online event</option>
          <option value="Physical location">Physical location</option>
        </select>
        <button class="btn btn-primary mt-5" @click="Create">
          Create event
        </button>
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
  data: () => ({
    title: '',
    img: '',
    date: '',
    time: '',
    city: '',
    type: '',
    max: null,
    hostName: '',
    hostURL: '',
    link: '',
  }),
  methods: {
    async Create() {
      const componentState = {
        title: this.title,
        img: this.img,
        date: this.date,
        time: this.time,
        city: this.city,
        type: this.type,
        max: Number(this.max),
        hostName: this.hostName,
        hostURL: this.hostURL,
        link: this.link,
      }
      const create = await this.$axios.post(
        'http://localhost:3333/api/v1/events',
        componentState,
        {
          headers: {
            Authorization: `Bearer ${this.$store.state.Token}`,
          },
        }
      )
      if (!create) {
        // eslint-disable-next-line no-console
        console.log('Create event error.')
      }
    },
  },
})
</script>
