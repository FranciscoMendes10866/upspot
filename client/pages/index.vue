<template>
  <div>
    <!-- Navbar -->
    <section>
      <Navbar />
    </section>
    <!-- Hero Landing -->
    <section class="mt-2">
      <div class="container">
        <div class="row py-5">
          <div class="col-sm">
            <h1 class="text-left display-1 font-weight-bold">
              Descubra eventos para todas as coisas que você ama.
            </h1>
            <button type="button" class="btn btn-outline-primary btn-lg mt-4">
              Começar.
            </button>
          </div>
        </div>
      </div>
    </section>
    <!-- Search Component -->
    <section class="mt-4 py-4">
      <div class="container d-flex justify-content-center">
        <div class="row">
          <div class="col">
            <div class="card border-light px-4 py-4 shadow">
              <Search />
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="mt-4 py-5">
      <div class="container">
        <div class="d-flex justify-content-start">
          <h4 class="font-weight-bold">
            Eventos perto de <span class="badge bg-primary">Coimbra.</span>
          </h4>
        </div>
      </div>
    </section>
    <section class="py-3">
      <div class="container">
        <div class="d-flex flex-wrap">
          <div
            v-for="event in events"
            :key="event.id"
            class="card border-light shadow w mb-4 ml-4"
          >
            <img :src="event.img" class="card-img-top" alt="..." />
            <div class="card-body">
              <h5 class="card-title">{{ event.title }}</h5>
              <small class="card-title">{{ event.date }}</small>
              <nuxt-link to="/" class="btn btn-light mt-3"
                >Ver evento</nuxt-link
              >
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
  data: () => ({
    events: [
      {
        id: 1,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Evento online',
      },
      {
        id: 2,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Evento online',
      },
      {
        id: 3,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Evento online',
      },
      {
        id: 4,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Evento online',
      },
      {
        id: 5,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Evento online',
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
</style>
