<template>
  <div>
    <!-- Navbar -->
    <section>
      <Navbar />
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
            <div class="card border-light px-4 py-4 shadow mb-5">
              <Search />
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="py-5 bg-light">
      <div class="container">
        <h4 class="font-weight-bold py-5">
          Events near <span class="badge bg-primary">Zurich.</span>
        </h4>
        <div class="d-flex flex-wrap">
          <div
            v-for="event in events"
            :key="event.id"
            class="card border-light shadow w mb-4 ml-4"
          >
            <span class="badge bg-white text-dark bdg">{{ event.type }}.</span>
            <img :src="event.img" class="card-img-top" alt="..." />
            <div class="card-body">
              <h5 class="card-title">{{ event.title }}</h5>
              <small class="card-title">{{ event.date }}</small>
              <div
                class="mt-4 d-flex justify-content-between align-items-center"
              >
                <nuxt-link to="/" class="btn btn-outline-primary"
                  >View event</nuxt-link
                >
                <button
                  v-if="event.current != event.max"
                  class="btn btn-light"
                  disabled
                >
                  {{ event.current }} /
                  <span class="font-weight-bold">{{ event.max }}</span>
                </button>
                <button v-else class="btn btn-light font-weight-bold" disabled>
                  Full
                </button>
              </div>
            </div>
          </div>
        </div>
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

export default Vue.extend({
  components: {
    Navbar,
    Search,
    Footer,
  },
  data: () => ({
    events: [
      {
        id: 1,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 35,
        current: 12,
      },
      {
        id: 2,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 50,
        current: 4,
      },
      {
        id: 3,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 12,
        current: 12,
      },
      {
        id: 4,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 40,
        current: 27,
      },
      {
        id: 5,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 200,
        current: 44,
      },
    ],
  }),
})
</script>

<style scoped>
.shadow {
  -webkit-box-shadow: 3px 3px 5px 6px #ccc; /* Safari 3-4, iOS 4.0.2 - 4.2, Android 2.3+ */
  -moz-box-shadow: 3px 3px 5px 6px #ccc; /* Firefox 3.5 - 3.6 */
  box-shadow: 3px 3px 5px 6px #ccc; /* Opera 10.5, IE 9, Firefox 4+, Chrome 6+, iOS 5 */
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
